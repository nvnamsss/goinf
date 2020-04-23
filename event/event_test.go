package event

import (
	"fmt"
	"testing"
)

func meomeo() {
	fmt.Println("meo meo")
}
func TestAsyncEvent(t *testing.T) {
	var onTrigger *Event = NewEvent()
	onTrigger.Subscribe(meomeo)

	onTrigger.Invoke()
	onTrigger.Subscribe(func() {
		fmt.Println("Hi mom")
	})

	onTrigger.Invoke()

	onTrigger.Unsubscribe(meomeo)
	onTrigger.Invoke()
}
