package main

/*

✅ 1. Lazy Initialization (Not Thread-Safe)

package singleton

type Singleton struct{}

var instance *Singleton

func GetInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}
✅ 2. Eager Initialization

package singleton

type Singleton struct{}

var instance = &Singleton{} // created at the time of package loading

func GetInstance() *Singleton {
	return instance
}
✅ 3. Thread-Safe Lazy Initialization (Double-Checked Locking)

package singleton

import (
	"sync"
)

type Singleton struct{}

var (
	instance *Singleton
	lock     sync.Mutex
)

func GetInstance() *Singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}
✅ 4. Enum Singleton Equivalent in Go
Go doesn’t have enums with behavior like Java, but you can achieve the same effect using a constant struct instance.

package singleton

type enumSingleton struct{}

var EnumInstance = &enumSingleton{}

func (e *enumSingleton) DoSomething() {
	// logic here
}


*/
