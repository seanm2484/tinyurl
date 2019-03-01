package main

import (
	"strings"

	"github.com/v33ps/base62"
)

// gets the longURL
// pass the longURL and mutex to genShortURL and get the return from that
// write the db record in a goroutine
func shortenURL(longURL string) (string, error) {
	m.Lock()
	shortURL := base62.Encode(recordID)
	longURL, err := checkHTTP(longURL)
	if err != nil {
		return "", err
	}
	rec := URLRecord{recordID, longURL, shortURL}
	recordID++
	m.Unlock()
	go writeRecord(db, rec)
	return shortURL, err
}

// make sure we have "http://" in front of the url. If not, put it there
func checkHTTP(longURL string) (string, error) {
	var err error
	if len(longURL) < 7 {
		return "", err
	}
	if strings.Contains(longURL[0:7], "http://") {
		return longURL, nil
	} else if strings.Contains(longURL[0:8], "https://") {
		return longURL, nil
	} else {
		return "https://" + longURL, nil
	}
}
