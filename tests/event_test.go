package tests

import (
	`testing`
	
	`github.com/chaodoing/boot/task`
)

type tw struct {
	Name string
	Age  int
}

func TestEvent(t *testing.T) {
	var event = task.NewEvent()
	event.AddEventListener("mysql:connect", func(data tw, boot bool) {
		t.Log("mysql:connect", data, boot)
	})
	event.Trigger("mysql:connect", tw{Name: "何烨霖", Age: 18}, true)
}
