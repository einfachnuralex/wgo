package dependencyInjection

import (
	"fmt"
	"io"
)

func main() {
	fmt.Print(Hello("Waldo"))
}

// Hello is very cool and greetes the subject
func Hello(subject string) string {
	if subject == "" {
		subject = "World"
	}
	return "Hello, " + subject
}

func Greet(writer io.Writer, name string) {
	fmt.Fprint(writer, Hello(name))
}
