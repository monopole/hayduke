package main


import (
	"fmt"

	"github.com/monopole/hayduke/api/prefixer"
)

func main() {
 fmt.Println("Hey there " + prefixer.Prefix("bob"))
}
