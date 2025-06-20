# Requirements
1. The digital wallet should allow users to create an account and manage their personal information.
2. Users should be able to add and remove payment methods, such as credit cards or bank accounts.
3. The digital wallet should support fund transfers between users and to external accounts.
4. The system should handle transaction history and provide a statement of transactions.
5. The digital wallet should support multiple currencies and perform currency conversions.
6. The system should ensure the security of user information and transactions.
7. The digital wallet should handle concurrent transactions and ensure data consistency.
8. The system should be scalable to handle a large number of users and transactions.

```go
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// ---------- ENUMS ----------
type Currency string

const (
	INR Currency = "INR"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// ---------- STRUCTS ----------
type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Accounts []*Account
}

type Account struct {
	ID          string
	UserID      string
	AccountNo   string
	Currency    Currency
	Balance     float64
	Transactions []*Transaction
	mu          sync.Mutex
}

type Transaction struct {
	ID              string
	SourceAccountID string
	DestAccountID   string
	Amount          float64
	Currency        Currency
	Timestamp       time.Time
}

type PaymentMethod interface {
	Process(amount float64) error
	GetDetails() string
}

type CreditCard struct {
	CardNumber string
	Expiry     string
	CVV        string
	HolderName string
}

func (c CreditCard) Process(amount float64) error {
	fmt.Printf("Processed %.2f from credit card %s\n", amount, c.CardNumber)
	return nil
}

func (c CreditCard) GetDetails() string {
	return fmt.Sprintf("CreditCard: %s", c.CardNumber)
}

type BankAccount struct {
	BankName string
	AccountNumber string
	IFSC string
}

func (b BankAccount) Process(amount float64) error {
	fmt.Printf("Processed %.2f from bank account %s\n", amount, b.AccountNumber)
	return nil
}

func (b BankAccount) GetDetails() string {
	return fmt.Sprintf("BankAccount: %s", b.AccountNumber)
}

// ---------- UTILITY ----------
type CurrencyConverter struct{}

func (CurrencyConverter) Convert(amount float64, from, to Currency) float64 {
	if from == to {
		return amount
	}
	rate := map[Currency]map[Currency]float64{
		INR: {USD: 0.012, EUR: 0.011},
		USD: {INR: 82.0, EUR: 0.91},
		EUR: {INR: 90.0, USD: 1.1},
	}
	return amount * rate[from][to]
}

// ---------- SINGLETON ----------
type DigitalWallet struct {
	Users      map[string]*User
	Accounts   map[string]*Account
	Methods    map[string]PaymentMethod
	Transactions []*Transaction
	mu         sync.Mutex
}

var walletInstance *DigitalWallet
var once sync.Once

func GetDigitalWallet() *DigitalWallet {
	once.Do(func() {
		walletInstance = &DigitalWallet{
			Users:      make(map[string]*User),
			Accounts:   make(map[string]*Account),
			Methods:    make(map[string]PaymentMethod),
			Transactions: []*Transaction{},
		}
	})
	return walletInstance
}

func (dw *DigitalWallet) CreateUser(name, email, password string) *User {
	dw.mu.Lock()
	defer dw.mu.Unlock()
	id := generateID()
	user := &User{ID: id, Name: name, Email: email, Password: password}
	dw.Users[id] = user
	return user
}

func (dw *DigitalWallet) CreateAccount(userID string, currency Currency) *Account {
	dw.mu.Lock()
	defer dw.mu.Unlock()
	id := generateID()
	acc := &Account{ID: id, UserID: userID, AccountNo: generateAccountNo(), Currency: currency, Balance: 0}
	dw.Accounts[id] = acc
	dw.Users[userID].Accounts = append(dw.Users[userID].Accounts, acc)
	return acc
}

func (dw *DigitalWallet) AddPaymentMethod(id string, method PaymentMethod) {
	dw.mu.Lock()
	defer dw.mu.Unlock()
	dw.Methods[id] = method
}

func (dw *DigitalWallet) Transfer(fromID, toID string, amount float64, currency Currency) error {
	dw.mu.Lock()
	fromAcc, ok1 := dw.Accounts[fromID]
	toAcc, ok2 := dw.Accounts[toID]
	dw.mu.Unlock()
	if !ok1 || !ok2 {
		return errors.New("invalid account(s)")
	}

	fromAcc.mu.Lock()
	defer fromAcc.mu.Unlock()
	destCurrency := toAcc.Currency
	convertedAmount := CurrencyConverter{}.Convert(amount, currency, destCurrency)

	if fromAcc.Balance < amount {
		return errors.New("insufficient funds")
	}
	fromAcc.Balance -= amount

	toAcc.mu.Lock()
	toAcc.Balance += convertedAmount
	toAcc.mu.Unlock()

	txn := &Transaction{
		ID: generateID(),
		SourceAccountID: fromID,
		DestAccountID: toID,
		Amount: amount,
		Currency: currency,
		Timestamp: time.Now(),
	}
	fromAcc.Transactions = append(fromAcc.Transactions, txn)
	toAcc.Transactions = append(toAcc.Transactions, txn)
	dw.mu.Lock()
	dw.Transactions = append(dw.Transactions, txn)
	dw.mu.Unlock()

	return nil
}

func (dw *DigitalWallet) GetTransactionHistory(accountID string) []*Transaction {
	dw.mu.Lock()
	defer dw.mu.Unlock()
	if acc, ok := dw.Accounts[accountID]; ok {
		return acc.Transactions
	}
	return nil
}

// ---------- HELPERS ----------
func generateID() string {
	return fmt.Sprintf("TXN%d", rand.Intn(1000000))
}

func generateAccountNo() string {
	return "AC" + strconv.Itoa(rand.Intn(99999999))
}

// ---------- DEMO ----------
func main() {
	wallet := GetDigitalWallet()
	user1 := wallet.CreateUser("Rishabh", "rishabh@example.com", "pass123")
	user2 := wallet.CreateUser("Amit", "amit@example.com", "pass456")

	u1Acc := wallet.CreateAccount(user1.ID, INR)
	u2Acc := wallet.CreateAccount(user2.ID, USD)

	u1Acc.Balance = 10000

	wallet.AddPaymentMethod("rishabh_card", CreditCard{"4111-1111-1111-1111", "12/25", "123", "Rishabh"})
	wallet.AddPaymentMethod("amit_bank", BankAccount{"HDFC", "1234567890", "HDFC0001"})

	err := wallet.Transfer(u1Acc.ID, u2Acc.ID, 1000, INR)
	if err != nil {
		fmt.Println("Transfer Failed:", err)
	} else {
		fmt.Println("Transfer Successful")
	}

	for _, txn := range wallet.GetTransactionHistory(u1Acc.ID) {
		fmt.Printf("Txn %s: Sent %.2f %s to %s\n", txn.ID, txn.Amount, txn.Currency, txn.DestAccountID)
	}
}

```