// +build tools

// This file exists to declare that its containing
// package explicitly depends on the pluginator
// tool (via go:generate directives)
package tools

import (
  	_ "github.com/monopole/hayduke/foo/v2"
)
