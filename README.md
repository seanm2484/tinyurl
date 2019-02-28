# tinyurl
make your big urls tiny 

## Known Issues
* you need to supply the `http://`/`https://` otherwise it treats it as a relative path which doesn't work.
* you can hammer the submit endpoint and crush the sqlitedb, which panics the service
* it's uggo
* probably a bunch of other stuff since I haven't really tested or wrote unit tests along the way


## build
```
go build -o tinyurl src/*.go 
./tinyurl 

// go to localhost:8080/

```

