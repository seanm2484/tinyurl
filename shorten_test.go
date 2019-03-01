package main

import "testing"

func TestShortenURL(t *testing.T) {
	expected := "0"
	actual, err := shortenURL("www.google.com")
	if err != nil {
		t.Errorf("TestShortenURL error: %v", err)
	}

	if actual != expected {
		t.Errorf("testShortenURL bad return: expected %v but got %v", expected,
			actual)
	}
}

func TestCheckHTTPWithHTTP(t *testing.T) {
	expected := "http://www.google.com"
	actual, err := checkHTTP("http://www.google.com")
	if err != nil {
		t.Errorf("TestCheckHTTPWithHTTP error: %v", err)
	}
	if actual != expected {
		t.Errorf("TestCheckHTTPWithHTTP bad return: expected %v but got %v",
			expected, actual)
	}
}

func TestCheckHTTPWithHTTPS(t *testing.T) {
	expected := "https://www.google.com"
	actual, err := checkHTTP("https://www.google.com")
	if err != nil {
		t.Errorf("TestCheckHTTPWithHTTPS error: %v", err)
	}
	if actual != expected {
		t.Errorf("TestCheckHTTPWithHTTPS bad return: expected %v but got %v",
			expected, actual)
	}
}

func TestCheckHTTPWithNone(t *testing.T) {
	expected := "https://www.google.com"
	actual, err := checkHTTP("www.google.com")
	if err != nil {
		t.Errorf("TestCheckHTTPWithNone error: %v", err)
	}
	if actual != expected {
		t.Errorf("TestCheckHTTPWithNone bad return: expected %v but got %v",
			expected, actual)
	}
}
