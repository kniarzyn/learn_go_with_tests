package main

import "fmt"

func Hello() string {
	return "Hello world"
}

func HelloTo(name string, lang ...string) string {
	greatings := map[string]string{
		"english": "Hello, ",
		"polish":  "Witaj, ",
		"spanish": "Hola, ",
		"french":  "Bonjour, ",
	}
	if name == "" {
		return Hello()
	}
	if lang == nil {
		return "Hello, " + name
	}
	return greatings[lang[0]] + name
}

func main() {
	fmt.Println(Hello())
}
