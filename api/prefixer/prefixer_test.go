package prefixer_test

import (
	"testing"

	. "github.com/monopole/hayduke/api/prefixer"
)

func TestPrefix(t *testing.T) {
  bob := "bob"
  other := Prefix(bob)
  if other != "hayduke-bob" {
    t.Errorf("unexpected other: %s", other)
  }
}

