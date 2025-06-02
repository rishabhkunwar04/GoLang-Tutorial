## logging framework

- Requirement:
1. The logging framework should support different log levels, such as DEBUG, INFO, WARNING, ERROR, and FATAL.
2. It should allow logging messages with a timestamp, log level, and message content.
3. The framework should support multiple output destinations, such as console, file, and database.
4. It should provide a configuration mechanism to set the log level and output destination.
5. The logging framework should be thread-safe to handle concurrent logging from multiple threads.
6. It should be extensible to accommodate new log levels and output destinations in the future.


- schema design:
 ```sql
 CREATE TABLE IF NOT EXISTS logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp TEXT NOT NULL,
    level TEXT NOT NULL,
    message TEXT NOT NULL
);

 ```
```go
package main

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3" // or your preferred DB driver
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}[l]
}

type Appender interface {
	Append(level LogLevel, message string)
}

type ConsoleAppender struct{}

func (c *ConsoleAppender) Append(level LogLevel, message string) {
	fmt.Printf("%s [%s] %s\n", time.Now().Format(time.RFC3339), level, message)
}

type FileAppender struct {
	file *os.File
}

func NewFileAppender(filePath string) (*FileAppender, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileAppender{file: file}, nil
}

func (f *FileAppender) Append(level LogLevel, message string) {
	fmt.Fprintf(f.file, "%s [%s] %s\n", time.Now().Format(time.RFC3339), level, message)
}

type DBAppender struct {
	db *sql.DB
}

func NewDBAppender(dataSourceName string) (*DBAppender, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	_, _ = db.Exec(`CREATE TABLE IF NOT EXISTS logs (timestamp TEXT, level TEXT, message TEXT)`)
	return &DBAppender{db: db}, nil
}

func (d *DBAppender) Append(level LogLevel, message string) {
	timestamp := time.Now().Format(time.RFC3339)
	_, _ = d.db.Exec(`INSERT INTO logs (timestamp, level, message) VALUES (?, ?, ?)`, timestamp, level.String(), message)
}

type Logger struct {
	mu        sync.RWMutex
	level     LogLevel
	appenders []Appender
}

func NewLogger(level LogLevel, appenders ...Appender) *Logger {
	return &Logger{level: level, appenders: appenders}
}

func (l *Logger) log(level LogLevel, msg string) {
	if level < l.level {
		return
	}
	l.mu.RLock()
	defer l.mu.RUnlock()
	for _, appender := range l.appenders {
		appender.Append(level, msg)
	}
}

func (l *Logger) Debug(msg string)   { l.log(DEBUG, msg) }
func (l *Logger) Info(msg string)    { l.log(INFO, msg) }
func (l *Logger) Warning(msg string) { l.log(WARNING, msg) }
func (l *Logger) Error(msg string)   { l.log(ERROR, msg) }
func (l *Logger) Fatal(msg string)   { l.log(FATAL, msg) }

// Client code
func main() {
	console := &ConsoleAppender{}
	fileAppender, err := NewFileAppender("/mnt/data/logging_example/app.log")
	if err != nil {
		panic(err)
	}
	dbAppender, err := NewDBAppender("/mnt/data/logging_example/logs.db")
	if err != nil {
		panic(err)
	}

	logger := NewLogger(INFO, console, fileAppender, dbAppender)
	logger.Info("Application started")
	logger.Debug("This debug message won't show at INFO level")
	logger.Error("An error occurred")
}

```