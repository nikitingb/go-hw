package greeter

import "fmt"

// Greet prints a greeting message
func Greet(name string) string {
    return fmt.Sprintf("Hello, %s! Welcome to Go modules.", name)
}