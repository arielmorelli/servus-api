package route

import (
	"fmt"
	"testing"
)

// AsRouteName
func TestAsRouteName(t *testing.T) {
	name := AsRouteName("/a/b/c/")
	if name != "a/b/c" {
		t.Fatalf(fmt.Sprint("Expects: a/b/c got ", name))
	}
}

// isMapSubset
func TestIsMapSubset(t *testing.T) {
	base := map[string]string{"a": "a", "b": "b"}
	// is subset
	if !isMapSubset(base, map[string]string{"a": "a"}) {
		t.Fatalf("Is a subset, wrong return")
	}

	if !isMapSubset(base, map[string]string{"b": "b"}) {
		t.Fatalf("Is a subset, wrong return")
	}

	if !isMapSubset(base, map[string]string{"a": "a", "b": "b"}) {
		t.Fatalf("Is a subset, wrong return")
	}

	// not subset
	if isMapSubset(base, map[string]string{"a": "a", "b": "b", "c": "c"}) {
		t.Fatalf("Is not a subset, wrong return")
	}

	if isMapSubset(base, map[string]string{"c": "c"}) {
		t.Fatalf("Is not a subset, wrong return")
	}
}
