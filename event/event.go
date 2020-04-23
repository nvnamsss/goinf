package event

import (
	"fmt"
	"reflect"
	"runtime"
)

//Event presentation a state whichWW is call the subscriber everytime it's trigger
type Event struct {
	Async     bool
	callbacks map[string]func()
}

//Subscribe the callback function
func (e *Event) Subscribe(callback func()) {
	id := runtime.FuncForPC(reflect.ValueOf(callback).Pointer()).Name()
	fmt.Println(id)
	e.callbacks[id] = callback
}

//Unsubscribe the callback function if it been subscribed
func (e *Event) Unsubscribe(callback func()) {
	id := runtime.FuncForPC(reflect.ValueOf(callback).Pointer()).Name()
	delete(e.callbacks, id)
}

//Trigger all subscribed callbacks
//If trigger type is async, all subscribed callback will run by goroutine
//Other whise it will be sequentially runned
func (e *Event) Invoke() {
	for _, item := range e.callbacks {
		if e.Async {
			go item()
		} else {
			item()
		}
	}
}

func NewEvent() *Event {
	e := new(Event)
	e.Async = true
	e.callbacks = make(map[string]func())
	return e
}
