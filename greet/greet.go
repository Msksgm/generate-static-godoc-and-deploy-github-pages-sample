package greet

import "fmt"

// Greet は引数に対して挨拶する関数
func Greet(name string) string {
	return fmt.Sprintf("Hello %s !", name)
}
