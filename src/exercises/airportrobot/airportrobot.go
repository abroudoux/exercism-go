package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Gretter interface {
	LanguageName() string
	Greet(string) string
}

func (g Greeter) SayHello(name string) {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}
