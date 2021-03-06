package main

import "fmt"

const englishHelloPrefix = "hello, "    //常量可以提高程序的性能
const spanish = "Spanish"
const french = "French"
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func hello(name string, language string) string{
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return 
}

func main(){

	fmt.Println(hello("Chris", ""))
}