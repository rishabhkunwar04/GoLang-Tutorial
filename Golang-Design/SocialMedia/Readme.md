
## Social Media
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type NotificationType string

const (
	FriendRequest         NotificationType = "FriendRequest"
	FriendRequestAccepted NotificationType = "FriendRequestAccepted"
	Like                  NotificationType = "Like"
	Comment               NotificationType = "Comment"
	Mention               NotificationType = "Mention"
)

type User struct {
	ID          int
	Name        string
	Email       string
	Password    string
	ProfilePic  string
	Bio         string
	Interests   []string
	Friends     map[int]bool
	Posts       []int
	Notifications []Notification
}

type Post struct {
	ID        int
	UserID    int
	Content   string
	Images    []string
	Videos    []string
	Timestamp time.Time
	Likes     map[int]bool
	Comments  []Comment
}

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	Content   string
	Timestamp time.Time
}

type Notification struct {
	ID        int
	UserID    int
	Type      NotificationType
	Content   string
	Timestamp time.Time
}

type SocialNetworkingService struct {
	Users   map[int]*User
	Posts   map[int]*Post
	Mutex   sync.Mutex
	NextUID int
	NextPID int
	NextCID int
	NextNID int
}

var instance *SocialNetworkingService
var once sync.Once

func GetSNSInstance() *SocialNetworkingService {
	once.Do(func() {
		instance = &SocialNetworkingService{
			Users: make(map[int]*User),
			Posts: make(map[int]*Post),
			Mutex: sync.Mutex{},
		}
	})
	return instance
}

func (sns *SocialNetworkingService) RegisterUser(name, email, password string) *User {
	sns.Mutex.Lock()
	defer sns.Mutex.Unlock()
	user := &User{
		ID:       sns.NextUID,
		Name:     name,
		Email:    email,
		Password: password,
		Friends:  make(map[int]bool),
	}
	sns.Users[user.ID] = user
	sns.NextUID++
	return user
}

func (sns *SocialNetworkingService) CreatePost(userID int, content string, images, videos []string) *Post {
	sns.Mutex.Lock()
	defer sns.Mutex.Unlock()
	post := &Post{
		ID:        sns.NextPID,
		UserID:    userID,
		Content:   content,
		Images:    images,
		Videos:    videos,
		Timestamp: time.Now(),
		Likes:     make(map[int]bool),
	}
	sns.Posts[post.ID] = post
	sns.Users[userID].Posts = append(sns.Users[userID].Posts, post.ID)
	sns.NextPID++
	return post
}

func (sns *SocialNetworkingService) SendFriendRequest(senderID, receiverID int) {
	notification := Notification{
		ID:        sns.NextNID,
		UserID:    receiverID,
		Type:      FriendRequest,
		Content:   fmt.Sprintf("%s sent you a friend request.", sns.Users[senderID].Name),
		Timestamp: time.Now(),
	}
	sns.Users[receiverID].Notifications = append(sns.Users[receiverID].Notifications, notification)
	sns.NextNID++
}

func (sns *SocialNetworkingService) AcceptFriendRequest(userID, friendID int) {
	sns.Users[userID].Friends[friendID] = true
	sns.Users[friendID].Friends[userID] = true
	notification := Notification{
		ID:        sns.NextNID,
		UserID:    friendID,
		Type:      FriendRequestAccepted,
		Content:   fmt.Sprintf("%s accepted your friend request.", sns.Users[userID].Name),
		Timestamp: time.Now(),
	}
	sns.Users[friendID].Notifications = append(sns.Users[friendID].Notifications, notification)
	sns.NextNID++
}

func (sns *SocialNetworkingService) LikePost(userID, postID int) {
	post := sns.Posts[postID]
	post.Likes[userID] = true
	owner := sns.Users[post.UserID]
	notification := Notification{
		ID:        sns.NextNID,
		UserID:    owner.ID,
		Type:      Like,
		Content:   fmt.Sprintf("%s liked your post.", sns.Users[userID].Name),
		Timestamp: time.Now(),
	}
	owner.Notifications = append(owner.Notifications, notification)
	sns.NextNID++
}

func (sns *SocialNetworkingService) CommentOnPost(userID, postID int, content string) {
	comment := Comment{
		ID:        sns.NextCID,
		UserID:    userID,
		PostID:    postID,
		Content:   content,
		Timestamp: time.Now(),
	}
	sns.Posts[postID].Comments = append(sns.Posts[postID].Comments, comment)
	sns.NextCID++
}

func (sns *SocialNetworkingService) GetNewsFeed(userID int) []Post {
	var feed []Post
	for friendID := range sns.Users[userID].Friends {
		for _, postID := range sns.Users[friendID].Posts {
			feed = append(feed, *sns.Posts[postID])
		}
	}
	for _, postID := range sns.Users[userID].Posts {
		feed = append(feed, *sns.Posts[postID])
	}
	return feed // Simplified, no sort for now
}

func main() {
	sns := GetSNSInstance()
	u1 := sns.RegisterUser("Alice", "alice@example.com", "pass123")
	u2 := sns.RegisterUser("Bob", "bob@example.com", "pass456")

	sns.SendFriendRequest(u1.ID, u2.ID)
	sns.AcceptFriendRequest(u2.ID, u1.ID)

	p1 := sns.CreatePost(u1.ID, "Hello from Alice!", nil, nil)
	sns.LikePost(u2.ID, p1.ID)
	sns.CommentOnPost(u2.ID, p1.ID, "Nice post!")

	feed := sns.GetNewsFeed(u1.ID)
	for _, post := range feed {
		fmt.Println("Post:", post.Content)
	}
}