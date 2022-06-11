package hello

const german = "de"
const englishHelloPrefix = "Hello"
const germanHelloPrefix = "Hallo"
const frenchHelloPrefix = "Bonjour"
const french = "fr"

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(lang) + " " + name + "!"
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case german:
		prefix = germanHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
