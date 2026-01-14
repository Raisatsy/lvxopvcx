package main

import "fmt"

type notifier interface {
	notify()
}

type email struct {
	name string
}

type sms struct {
	address string
	count   int
}

func (s *sms) notify() {
	fmt.Println(s.address)
	s.count++
}

func (e *email) notify() {
	fmt.Println(e.name)
}

func sendNotification(n notifier) {
	n.notify()
}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	email := email{name: "abdul"}
	sms := sms{address: "123"}

	sendNotification(&email)
	sendNotification(&sms)
	sendNotification(&sms)
	sendNotification(&sms)
	sendNotification(&sms)
	sendNotification(&sms)
	sendNotification(&sms)
	sendNotification(&sms)
	fmt.Println(sms.count)
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>

}
