package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (eb englishBot) getGreeting() string {
	// custom logic for generating english greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
