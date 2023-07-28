package greetings
import "fmt"

func greetings(name string) string {
	message:= fmt.Sprintf("Hello, %v, consider yourself great!", name);
	return message; 
}
