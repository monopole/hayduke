package main


import (
	"fmt"

	"github.com/monopole/hayduke/api/prefixer"
	v2p "github.com/monopole/hayduke/api/v2/prefixer"
)

func main() {
 fmt.Println("Hey there " + prefixer.Prefix("bob") + v2p.Prefix("bob", "cheese"))
}
