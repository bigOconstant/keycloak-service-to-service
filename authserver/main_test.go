package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetToken(t *testing.T) {

	var h http.Header

	h = make(map[string][]string)
	h["Authorization"] = append(h["Authorization"], "foo")

	result, _ := getToken(h)
	if result != "foo" {
		t.Error("token incorrect")
	} else {
		fmt.Println("token:", result)
	}
}
