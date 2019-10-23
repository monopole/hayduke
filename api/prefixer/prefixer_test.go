package prefixer_test

import (
	"testing"

	. "github.com/monopole/hayduke/api/v2/prefixer"
)

func TestPrefix(t *testing.T) {
  other := Prefix("eat", "salad")
  if other != "eat-kale-salad" {
    t.Errorf("unexpected other: %s", other)
  }
}

