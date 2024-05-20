package main

import (
	"errors"
	"net/http"
	"testing"
)

type testCase struct {
	name          string
	headers       http.Header
	expectedAPIKey string
	expectedErr    error
}