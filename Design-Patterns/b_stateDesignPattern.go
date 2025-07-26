package main

/*
import "fmt"

// ----------- State Interface ------------
type State interface {
	InsertMoney(v *VendingMachine)
	ProvideChangeMoney(v *VendingMachine)
	DispenseProduct(v *VendingMachine)
}

// ----------- NoMoneyState -------------
type NoMoneyState struct{}

func (n *NoMoneyState) InsertMoney(v *VendingMachine) {
	fmt.Println("Inserting money to purchase product")
	v.SetState(&HasMoneyState{})
}

func (n *NoMoneyState) ProvideChangeMoney(v *VendingMachine) {
	fmt.Println("No money to return")
}

func (n *NoMoneyState) DispenseProduct(v *VendingMachine) {
	fmt.Println("No product to return! Insert money to purchase the item")
}

// ----------- HasMoneyState -------------
type HasMoneyState struct{}

func (h *HasMoneyState) InsertMoney(v *VendingMachine) {
	fmt.Println("Money already inserted")
}

func (h *HasMoneyState) ProvideChangeMoney(v *VendingMachine) {
	fmt.Println("Giving change money left after buying product")
	v.SetState(&NoMoneyState{})
}

func (h *HasMoneyState) DispenseProduct(v *VendingMachine) {
	fmt.Println("Please collect product")
	v.SetState(&NoMoneyState{})
}

// ----------- VendingMachine -------------
type VendingMachine struct {
	currentState State
}

func NewVendingMachine() *VendingMachine {
	return &VendingMachine{
		currentState: &NoMoneyState{},
	}
}

func (v *VendingMachine) SetState(s State) {
	v.currentState = s
}

func (v *VendingMachine) InsertMoney() {
	v.currentState.InsertMoney(v)
}

func (v *VendingMachine) ProvideChangeMoney() {
	v.currentState.ProvideChangeMoney(v)
}

func (v *VendingMachine) DispenseProduct() {
	v.currentState.DispenseProduct(v)
}

// ----------- Main -------------
func main() {
	vendingMachine := NewVendingMachine()

	vendingMachine.InsertMoney()
	vendingMachine.DispenseProduct()

	fmt.Println("------------------")

	vendingMachine.InsertMoney()
	vendingMachine.InsertMoney()
	vendingMachine.ProvideChangeMoney()
	vendingMachine.DispenseProduct()

	fmt.Println("------------------")

	vendingMachine.InsertMoney()
	vendingMachine.DispenseProduct()
	vendingMachine.ProvideChangeMoney()
}

*/
