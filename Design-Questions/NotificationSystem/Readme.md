## Requirement
1. Multi-Channel Support: Email, SMS, Push Notification, In-App.
2. Notification Types: Alert, Reminder, Marketing, Transactional.
3. User Preferences: User can choose preferred channel(s) and types.
4. Templating: Custom templates for different types of notifications.
5. Retry & Failure Handling: Retry failed notifications.
6. Scheduling: Send now or schedule for future.
7. Rate Limiting & Throttling: Avoid spamming users.
8. Real-Time & Asynchronous Delivery: Some need real-time (in-app), others can be async.
9. Extensibility: Easy to plug in new channels like WhatsApp.
10. Analytics & Audit Logs: Track delivery status, failure reason, open rates, etc.

```go
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
	case "Email": return &EmailChannel{}
	case "SMS": return &SMSChannel{}
	default: return nil
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
	limits map[string]time.Time
	mu     sync.Mutex
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


	nm.RegisterObserver(NewAuditLogger())

	nm.SetUserPreference(UserPreference{
		UserID:            "user123",
		PreferredChannels: []string{"Email", "SMS"},
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

```
### Database Schema
```Sql
Database Schema

-- Users
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT,
    email TEXT,
    phone TEXT
);

-- Notification Templates
CREATE TABLE templates (
    id UUID PRIMARY KEY,
    type TEXT,
    channel TEXT,
    content TEXT
);

-- User Preferences
CREATE TABLE user_preferences (
    id UUID PRIMARY KEY,
    user_id UUID,
    channel TEXT,
    notification_type TEXT,
    enabled BOOLEAN
);

-- Notifications Queue
CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    user_id UUID,
    type TEXT,
    channel TEXT,
    template_id UUID,
    data JSONB,
    status TEXT,
    retries INT DEFAULT 0,
    scheduled_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT now()
);

-- Audit Log
CREATE TABLE audit_logs (
    id UUID PRIMARY KEY,
    notification_id UUID,
    status TEXT,
    message TEXT,
    created_at TIMESTAMPTZ DEFAULT now()
);
```

**design pattern we can use here**
```go
🔁 1. Strategy Pattern
Where: Channel interface and its implementations like EmailChannel, SMSChannel, etc.

Why:
To allow sending a notification using different strategies (channels) interchangeably without altering the core logic. You can switch the channel at runtime.


type Channel interface {
    Send(notification Notification) error
}
Benefit: Easily extendable for new channels like WhatsApp or Slack.

📦 2. Factory Pattern
Where: A potential enhancement — creating Channel instances using a factory.

Why:
To encapsulate object creation logic, especially if creating a Channel requires configuration or external dependencies.


func ChannelFactory(name string) Channel {
    switch name {
    case "EMAIL":
        return EmailChannel{}
    case "SMS":
        return SMSChannel{}
    // ...
    }
}
Benefit: Centralized and configurable creation logic.

📥 3. Command Pattern
Where: In dispatching logic for scheduled or retryable tasks.

Why:
Encapsulates notification sending as a command object, which can be queued, executed, retried, or logged independently.


type NotificationCommand struct {
    Notification Notification
}

func (cmd *NotificationCommand) Execute() error {
    return dispatcher.Dispatch(cmd.Notification)
}
Benefit: Useful for background workers, retry queues, scheduling.

🎁 4. Decorator Pattern
Where: Rate limiting, logging, and retry logic around the actual Channel.Send().

Why:
To dynamically wrap additional behavior around Send() without modifying core logic.


type RateLimitedChannel struct {
    Channel
    Limiter RateLimiter
}

func (r *RateLimitedChannel) Send(n Notification) error {
    if !r.Limiter.Allow(n.UserID) {
        return errors.New("rate limit exceeded")
    }
    return r.Channel.Send(n)
}
Benefit: Add behavior like rate limiting, retries, and logging without modifying original classes.

🧱 5. Builder Pattern
Where: While creating complex Notification objects with various optional fields (data, scheduled time, etc.)

Why:
Helps build objects step-by-step when you have many optional fields.


type NotificationBuilder struct {
    n Notification
}

func (b *NotificationBuilder) WithUserID(id string) *NotificationBuilder { ... }
func (b *NotificationBuilder) Build() Notification                       { ... }
Benefit: Cleaner code when creating customizable notifications.

🔌 6. Observer Pattern (Optional)
Where: If you want multiple services to react to notification events — like analytics, delivery reports, etc.

Why:
To notify interested subscribers (observers) when a notification is sent, failed, etc.

Benefit: Loosely coupled event broadcasting.

🛠️ 7. Template Method Pattern
Where: In templating logic.

Why:
To define the skeleton of notification formatting in a base class and let subclasses override certain steps.

✅ Summary Table
Pattern	Role in Design	Benefit
Strategy	Abstracts channels	Easily switch or add channels
Factory	Creates channels or notification instances	Centralized creation, config-ready
Command	Encapsulates dispatch behavior	Supports scheduling, retries, queues
Decorator	Adds logging, retry, rate limit without core change	Adds flexible features
Builder	Builds Notification objects	Clean construction with optional fields
Observer	Publishes delivery/reporting events	Decoupled analytics, audit, metrics
Template Method	Base logic for message formatting	Custom formatting with shared steps


```