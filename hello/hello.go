package hello

import "fmt"

func main() {
	fmt.Println(Hello("Jakob"))
}

func Hello(name string) string {
	return "Hello " + name + "!"
}
