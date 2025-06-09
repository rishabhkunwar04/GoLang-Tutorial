// Notification Main - Strategy, Factory, and Observer Patterns Applied
package main

import (
	"fmt"
	"sync"
	"time"
)

// Notification Types
// ------------------
type NotificationType string

const (
	Alert         NotificationType = "ALERT"
	Reminder      NotificationType = "REMINDER"
	Marketing     NotificationType = "MARKETING"
	Transactional NotificationType = "TRANSACTIONAL"
)

// Notification Struct
// -------------------
type Notification struct {
	UserID      string
	Type        NotificationType
	Title       string
	Message     string
	ScheduledAt time.Time
}

// NotificationChannel Interface - Strategy Pattern
// ------------------------------------------------
type NotificationChannel interface {
	Send(notification Notification) error
	ChannelName() string
}

// Concrete Strategy Implementations (Channels)
// --------------------------------------------
type EmailChannel struct{}

func (e *EmailChannel) Send(notification Notification) error {
	fmt.Printf("[Email] Sending to %s: %s - %s\n", notification.UserID, notification.Title, notification.Message)
	return nil
}
func (e *EmailChannel) ChannelName() string { return "Email" }

type SMSChannel struct{}

func (s *SMSChannel) Send(notification Notification) error {
	fmt.Printf("[SMS] Sending to %s: %s - %s\n", notification.UserID, notification.Title, notification.Message)
	return nil
}
func (s *SMSChannel) ChannelName() string { return "SMS" }

// ChannelFactory - Factory Method Pattern
// ---------------------------------------
type ChannelFactory struct{}

func (f *ChannelFactory) CreateChannel(name string) NotificationChannel {
	switch name {
	case "Email":
		return &EmailChannel{}
	case "SMS":
		return &SMSChannel{}
	default:
		return nil
	}
}

// Retry Manager
// -------------
type RetryManager struct {
	maxRetries int
}

func NewRetryManager(maxRetries int) *RetryManager {
	return &RetryManager{maxRetries: maxRetries}
}
func (rm *RetryManager) RetrySend(channel NotificationChannel, notification Notification) {
	for i := 0; i < rm.maxRetries; i++ {
		err := channel.Send(notification)
		if err == nil {
			return
		}
		time.Sleep(time.Duration(i+1) * time.Second)
	}
	fmt.Printf("[Retry Failed] Channel: %s, User: %s\n", channel.ChannelName(), notification.UserID)
}

// Rate Limiter
// ------------
type RateLimiter struct {
	limits   map[string]time.Time
	mu       sync.Mutex
	cooldown time.Duration
}

func NewRateLimiter(cooldown time.Duration) *RateLimiter {
	return &RateLimiter{limits: make(map[string]time.Time), cooldown: cooldown}
}
func (rl *RateLimiter) Allow(userID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	if nextAllowed, exists := rl.limits[userID]; exists && now.Before(nextAllowed) {
		return false
	}
	rl.limits[userID] = now.Add(rl.cooldown)
	return true
}

// Observer Pattern - Observer Interface and Subscriber
// ----------------------------------------------------
type NotificationObserver interface {
	Update(notification Notification, channel string, status string)
}

type AuditLogger struct {
	mu sync.Mutex
}

func NewAuditLogger() *AuditLogger { return &AuditLogger{} }
func (al *AuditLogger) Update(notification Notification, channel string, status string) {
	al.mu.Lock()
	defer al.mu.Unlock()
	fmt.Printf("[AuditLog] User: %s, Channel: %s, Status: %s, Title: %s\n", notification.UserID, channel, status, notification.Title)
}

// User Preference
// ---------------
type UserPreference struct {
	UserID            string
	PreferredChannels []string
	PreferredTypes    []NotificationType
}

// Notification Manager - Publisher with Observer Registration
// ------------------------------------------------------------
type NotificationManager struct {
	Channels     map[string]NotificationChannel
	Preferences  map[string]UserPreference
	Observers    []NotificationObserver
	mu           sync.Mutex
	RateLimiter  *RateLimiter
	RetryMgr     *RetryManager
	ChannelMaker *ChannelFactory
}

func NewNotificationManager() *NotificationManager {
	return &NotificationManager{
		Channels:     make(map[string]NotificationChannel),
		Preferences:  make(map[string]UserPreference),
		Observers:    make([]NotificationObserver, 0),
		RateLimiter:  NewRateLimiter(3 * time.Second),
		RetryMgr:     NewRetryManager(3),
		ChannelMaker: &ChannelFactory{},
	}
}

func (nm *NotificationManager) RegisterObserver(o NotificationObserver) {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	nm.Observers = append(nm.Observers, o)
}

func (nm *NotificationManager) NotifyObservers(notification Notification, channel string, status string) {
	for _, observer := range nm.Observers {
		observer.Update(notification, channel, status)
	}
}

func (nm *NotificationManager) RegisterChannel(channel NotificationChannel) {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	nm.Channels[channel.ChannelName()] = channel
}

func (nm *NotificationManager) SetUserPreference(pref UserPreference) {
	nm.mu.Lock()
	defer nm.mu.Unlock()
	nm.Preferences[pref.UserID] = pref
}

func (nm *NotificationManager) SendNotification(notification Notification) {
	if !nm.RateLimiter.Allow(notification.UserID) {
		fmt.Printf("[RateLimit] Notification skipped for user %s\n", notification.UserID)
		return
	}

	nm.mu.Lock()
	pref, exists := nm.Preferences[notification.UserID]
	nm.mu.Unlock()
	if !exists {
		fmt.Printf("No preferences found for user %s\n", notification.UserID)
		return
	}

	typeAllowed := false
	for _, typ := range pref.PreferredTypes {
		if typ == notification.Type {
			typeAllowed = true
			break
		}
	}
	if !typeAllowed {
		fmt.Printf("Notification type %s not preferred by user %s\n", notification.Type, notification.UserID)
		return
	}

	for _, channelName := range pref.PreferredChannels {
		if channel, ok := nm.Channels[channelName]; ok {
			go func(ch NotificationChannel) {
				err := ch.Send(notification)
				status := "Sent"
				if err != nil {
					nm.RetryMgr.RetrySend(ch, notification)
					status = "Retry Failed"
				}
				nm.NotifyObservers(notification, ch.ChannelName(), status)
			}(channel)
		}
	}
}

func main() {
	nm := NewNotificationManager()

	nm.RegisterChannel(nm.ChannelMaker.CreateChannel("Email"))
	nm.RegisterChannel(nm.ChannelMaker.CreateChannel("SMS"))
	nm.RegisterChannel(nm.ChannelMaker.CreateChannel("Push"))
	nm.RegisterChannel(nm.ChannelMaker.CreateChannel("InApp"))

	nm.RegisterObserver(NewAuditLogger())

	nm.SetUserPreference(UserPreference{
		UserID:            "user123",
		PreferredChannels: []string{"Email", "InApp"},
		PreferredTypes:    []NotificationType{Alert, Transactional},
	})

	notification := Notification{
		UserID:  "user123",
		Type:    Alert,
		Title:   "System Update",
		Message: "A system update will occur tonight at 2 AM.",
	}

	nm.SendNotification(notification)
	time.Sleep(5 * time.Second) // Wait for async goroutines
}
