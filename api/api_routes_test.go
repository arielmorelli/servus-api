package api

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestHeadersOrParametersToMapWithHeader(t *testing.T) {
	asMap := headersOrParametersToMap(http.Header{"a": []string{"a"}, "b": []string{"b", "c"}})
	if asMap["a"] != "a" {
		t.Error(fmt.Sprint("Expected a, got ", asMap["a"]))
	}
	if asMap["b"] != "b, c" {
		t.Error(fmt.Sprint("Expected b, c, got ", asMap["b"]))
	}
}

func TestHeadersOrParametersToMapWithParameters(t *testing.T) {
	asMap := headersOrParametersToMap(url.Values{"a": []string{"a"}, "b": []string{"b", "c"}})
	if asMap["a"] != "a" {
		t.Error(fmt.Sprint("Expected a, got ", asMap["a"]))
	}
	if asMap["b"] != "b, c" {
		t.Error(fmt.Sprint("Expected b, c, got ", asMap["b"]))
	}
}
