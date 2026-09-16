package main

import (
	"fmt"
)

const (
	greetingStart        = "Hello "
	greetingStartSpanish = "Hola "
	greetingStartFrench  = "Bonjour "

	defaultName        = "world"
	defaultSpanishName = "mundo"
	defaultFrenchName  = "le monde"

	spanishLang = "spanish"
	frenchLang  = "french"
)

func chooseName(name, lang string) string {
	if len(name) == 0 {
		switch lang {
		case spanishLang:
			name = defaultSpanishName
		case frenchLang:
			name = defaultFrenchName
		default:
			name = defaultName
		}
	}
	return name
}

func fullGreeting(name, lang string) string {
	switch lang {
	case spanishLang:
		return greetingStartSpanish + name
	case frenchLang:
		return greetingStartFrench + name
	default:
		return greetingStart + name
	}
}

func Hello(name, lang string) string {
	name = chooseName(name, lang)
	return fullGreeting(name, lang)
}

func main() {
	fmt.Println(Hello("", ""))
}
