package main

const french = "French"
const frenchgreeting = "Bonjour, "
const spanish = "Spanish"
const englishlanggreeting = "Hello, "
const spanishgreeting = "Hola, "

func hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	if language == spanish {
		return spanishgreeting + name
	}

	if language == french {
		return frenchgreeting + name
	}

	return englishlanggreeting + name
}

func main() {
	// hello("Chris")
	// hello("")
	hello("Elodie", "Spanish")
	hello("Rachel", "French")

}
