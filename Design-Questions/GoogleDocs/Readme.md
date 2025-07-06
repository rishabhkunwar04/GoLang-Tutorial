## Google Doc Design
```go
package main

import (
	"fmt"
)

// --- User Struct ---
type User struct {
	Username string
	UserID   string
}

// --- Permission Enum ---
type Permission int

const (
	READ Permission = iota
	WRITE
	OWNER
)

func (p Permission) String() string {
	return [...]string{"READ", "WRITE", "OWNER"}[p]
}

// --- Document Struct ---
type Document struct {
	Content string
	DocName string
	//// one userId -> many permissions
	// DocMap map[string][]Permission
	DocMap map[string]Permission
}

func NewDocument(user User, content string, docName string) *Document {
	return &Document{
		Content: content,
		DocName: docName,
		// DocMap:  map[string][]Permission{u.UserID: {OWNER}},

		DocMap: map[string]Permission{user.UserID: OWNER},
	}
}

/*
// helper: add perm if not already present
func addPerm(list []Permission, perm Permission) []Permission {
	for _, p := range list {
		if p == perm { return list }
	}
	return append(list, perm)
}
*/

func (d *Document) GrantAccess(user User, perm Permission) {
	d.DocMap[user.UserID] = perm
	//d.DocMap[u.UserID] = addPerm(d.DocMap[u.UserID], perm)
}

func (d *Document) RevokeAccess(user User) {
	delete(d.DocMap, user.UserID)
}

func (d *Document) WriteContent(user User, content string) {
	if perm, ok := d.DocMap[user.UserID]; ok && (perm == OWNER || perm == WRITE) {
		d.Content += content
		fmt.Println("Content written by", user.Username)
	} else {
		fmt.Println("Permission denied to WRITE!!")
	}
}

func (d *Document) ReadContent(user User) {
	if perm, ok := d.DocMap[user.UserID]; ok && (perm == OWNER || perm == READ) {
		fmt.Println("Content read by", user.Username)
		fmt.Println("-- Content --\n" + d.Content)
	} else {
		fmt.Println("Permission denied to READ!!")
	}
}

func (d *Document) DeleteContent(user User) {
	if perm, ok := d.DocMap[user.UserID]; ok && perm == OWNER {
		d.Content = ""
		fmt.Println("Content deleted by", user.Username)
	} else {
		fmt.Println("Permission denied to DELETE!!")
	}
}

// --- Main ---
func main() {
	user1 := User{"Rishabh", "rishabh1"}
	user2 := User{"Sumit", "sumit1"}
	// user3 := User{"Sushant", "sushant1"}

	doc := NewDocument(user1, "initial content !! ", "myDoc")

	fmt.Println(doc.Content)
	doc.WriteContent(user1, "content added by owner")
	fmt.Println(doc.Content)

	doc.GrantAccess(user2, READ)
	doc.RevokeAccess(user2)

	doc.ReadContent(user2) // Should say permission denied
}

```