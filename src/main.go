package airportrobot

import "fmt"

type Gretter interface {
	LanguageName() string
	Greet(string) string
}

func (g Greeter) SayHello(name string) {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}
