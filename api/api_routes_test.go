package api

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestHeadersToMap(t *testing.T) {
	asMap := headersToMap(http.Header{"a": []string{"a"}, "b": []string{"b", "c"}})
	if asMap["a"] != "a" {
		t.Error(fmt.Sprint("Expected a, got ", asMap["a"]))
	}
	if asMap["b"] != "b, c" {
		t.Error(fmt.Sprint("Expected b, c, got ", asMap["b"]))
	}
}

func TestParamsToMap(t *testing.T) {
	asMap := paramsToMap(url.Values{"a": []string{"a"}, "b": []string{"b", "c"}})
	if asMap["a"] != "a" {
		t.Error(fmt.Sprint("Expected a, got ", asMap["a"]))
	}
	if asMap["b"] != "b, c" {
		t.Error(fmt.Sprint("Expected b, c, got ", asMap["b"]))
	}
}
