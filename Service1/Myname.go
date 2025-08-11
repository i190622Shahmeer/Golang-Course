package Name

// In one directory, you can have multiple files but "ONE Package Name".
import (
	"fmt"
)

func Myname(message string) {
	fmt.Println(message)
}
func Myname2() string {
	// fmt.Println("Hello from Myname2 function!")
	return "Hello from Myname2 function!"
}
