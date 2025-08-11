## Google Doc Design Basics
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


## Advanced Google Docs

**Features Required:**

- Real-time Collaboration: Multiple users should be able to collaborate and edit a document simultaneously in real-time.

- Document Synchronization: Changes made by one user should be immediately reflected in the document for all other users.

- User Presence: Users should be able to see the presence of other users currently viewing or editing the document.

- Cursor Position Tracking: Users should be able to see the cursor positions of other users in real-time.

- Collaborative Text Editing: Users should be able to add, delete, and modify text in the document collaboratively.

- Version History: The system should maintain a version history of the document, allowing users to revert to previous versions if needed.
Document Sharing: Users should be able to share documents with other users, granting them appropriate access permissions.

- Access Controls: The system should provide access control mechanisms to manage user permissions for viewing and editing documents.

```java
Design Patterns Involved or Used:

Observer Pattern: The Observer pattern can be used to notify users about changes or updates in the document, such as new edits or cursor positions.

State Pattern: The State pattern can be used to manage the different states of a document, such as viewing, editing, or locked states, and handle state transitions.

Proxy Pattern: The Proxy pattern can be used to manage access controls and permissions for documents, providing a level of indirection and encapsulation for user operations.

Command Pattern: The Command pattern can be used to encapsulate document editing operations as commands, allowing for easy undo/redo functionality and collaboration tracking.

Code: Classes Implementation Based on Patterns Mentioned Above

// Observer interface
interface DocumentObserver {
    void update(Document document);
}

// CursorPositionObserver class
class CursorPositionObserver implements DocumentObserver {
    private User user;

    public CursorPositionObserver(User user) {
        this.user = user;
    }

    @Override
    public void update(Document document) {
        // Update cursor position for the user
        System.out.println("Cursor position updated for user " + user.getUsername() + ": " + document.getCursorPosition(user));
    }
}

// Document class
class Document {
    private String content;
    private List<User> collaborators;
    private List<String> versionHistory;
    private Map<User, Integer> cursorPositions;
    private List<DocumentObserver> observers;

    public Document() {
        this.content = "";
        this.collaborators = new ArrayList<>();
        this.versionHistory = new ArrayList<>();
        this.cursorPositions = new HashMap<>();
        this.observers = new ArrayList<>();
    }

    public void editDocument(String newText) {
        // Perform document editing operations
        content = newText;

        // Notify observers about the document update
        notifyObservers();
    }

    public void addCollaborator(User user) {
        collaborators.add(user);
        cursorPositions.put(user, 0);

        // Add cursor position observer for the new collaborator
        CursorPositionObserver observer = new CursorPositionObserver(user);
        observers.add(observer);
    }

    public void removeCollaborator(User user) {
        collaborators.remove(user);
        cursorPositions.remove(user);

        // Remove cursor position observer for the collaborator
        DocumentObserver observer = new CursorPositionObserver(user);
        observers.remove(observer);
    }

    public void updateCursorPosition(User user, int newPosition) {
        cursorPositions.put(user, newPosition);
    }

    public int getCursorPosition(User user) {
        return cursorPositions.getOrDefault(user, 0);
    }

    public void subscribeObserver(DocumentObserver observer) {
        observers.add(observer);
    }

    public void unsubscribeObserver(DocumentObserver observer) {
        observers.remove(observer);
    }

    private void notifyObservers() {
        for (DocumentObserver observer : observers) {
            observer.update(this);
        }
    }

    // Other operations related to document synchronization, version history, and access controls
}

// CollaborativeEditor class
class CollaborativeEditor {
    private Document document;

    public CollaborativeEditor(Document document) {
        this.document = document;
    }

    public void addCollaborator(User user) {
        document.addCollaborator(user);
    }

    public void removeCollaborator(User user) {
        document.removeCollaborator(user);
    }

    public void editDocument(String newText) {
        document.editDocument(newText);
    }

    // Other operations related to real-time collaboration, presence tracking, and cursor position tracking
}

// Main Class
public class RealTimeCollaborativeEditor {
    public static void main(String[] args) {
        // Create users
        User user1 = new User("user1");
        User user2 = new User("user2");

        // Create a document
        Document document = new Document();

        // Create a collaborative editor
        CollaborativeEditor collaborativeEditor = new CollaborativeEditor(document);

        // Add collaborators to the document
        collaborativeEditor.addCollaborator(user1);
        collaborativeEditor.addCollaborator(user2);

        // Edit the document
        collaborativeEditor.editDocument("This is the edited document content.");

        // Update cursor position for user1
        document.updateCursorPosition(user1, 10);

        // Update cursor position for user2
        document.updateCursorPosition(user2, 15);
    }
}

```