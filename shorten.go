package main

import (
	"fmt"

	"github.com/v33ps/base62"
)

// gets the longURL
// pass the longURL and mutex to genShortURL and get the return from that
// write the db record in a goroutine
func shortenURL(longURL string) string {
	fmt.Println(recordID)
	m.Lock()
	shortURL := base62.Encode(recordID)
	rec := URLRecord{recordID, longURL, shortURL}
	recordID++
	m.Unlock()
	go writeRecord(db, rec)
	return shortURL
}

/*
// mutex lock recordID and use it to create the shortURL
// create a URLRecord object that stores the id longURL and shortURL
// unlock the mutex and return the URLRecord object
func genShortURL(longURL string, m *sync.Mutex) URLRecord {
	m.Lock()
	shortURL := base62.Encode(recordID)
	rec := URLRecord{recordID, longURL, shortURL}
	recordID++
	m.Unlock()
	return rec
}*/
