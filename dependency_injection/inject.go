package dependency_injection

import (
	"fmt"
	"io"
	"os"
)

func greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}

func printWorld() {
	greet(os.Stdout, "World")
}
