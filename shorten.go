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

func checkHTTP(longURL string) (string, error) {
	// if we aren't given a protocol, default them to https.
	if strings.Contains(longURL, "://") != true {
		return "https://"+longURL, nil
	} else {
		return longURL, nil
	}
}
