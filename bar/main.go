package main


import (
	"fmt"

	"github.com/monopole/hayduke/api/v2/prefixer"
)

func main() {
 fmt.Println(prefixer.Prefix("bar", "bar"))
}
