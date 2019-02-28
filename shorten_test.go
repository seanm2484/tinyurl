package main

import "testing"

func TestShortenURL(t *testing.T) {
	expected := "0"
	actual := shortenURL("www.google.com")

	if actual != expected {
		t.Errorf("testShortenURL bad return: expected %v but got %v", expected,
			actual)
	}
}

func TestCheckHTTPWithHTTP(t *testing.T) {
	expected := "http://www.google.com"
	actual := checkHTTP("http://www.google.com")
	if actual != expected {
		t.Errorf("TestCheckHTTPWithHTTP bad return: expected %v but got %v",
			expected, actual)
	}
}

func TestCheckHTTPWithHTTPS(t *testing.T) {
	expected := "https://www.google.com"
	actual := checkHTTP("https://www.google.com")
	if actual != expected {
		t.Errorf("TestCheckHTTPWithHTTPS bad return: expected %v but got %v",
			expected, actual)
	}
}

func TestCheckHTTPWithNone(t *testing.T) {
	expected := "https://www.google.com"
	actual := checkHTTP("www.google.com")
	if actual != expected {
		t.Errorf("TestCheckHTTPWithNone bad return: expected %v but got %v",
			expected, actual)
	}
}
