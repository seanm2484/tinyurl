package main

import "testing"

func TestAll(t *testing.T) {
	const dbpath = "foo.db"

	db := initDb(dbpath)
	defer db.Close()
	createTable(db)

	items := URLRecord{1, "www.google.com/abc", "asdf"}
	writeRecord(db, items)

	readItems := readItem(db, "")
	t.Log(readItems)

}
