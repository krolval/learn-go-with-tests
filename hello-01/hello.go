package main

import "fmt"

const helloPrefixBr = "oi... "
const helloPrefixSp = "Hola... "
const helloPrefixFr = "Bonjour... "
const spanish = "Spanish"
const french = "French"

func hello(name string, args ...string) string {
	if name == "" {
		name = "People"
	}
	prefix := helloPrefixBr
	if args != nil {
		language := args[0]
		prefix = greetingPrefix(language)

	}
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSp
	case french:
		prefix = helloPrefixFr

	}
	return prefix
}

func main() {
	fmt.Println(hello("Test"))
}
