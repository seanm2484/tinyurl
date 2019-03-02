# tinyURL
TinyURL is a url shortening service based on an unpadded base62 encoding.


## Building
You can either build TinyURL from source, or deploy the docker container

### Building from Source
```
go build -o tinyurl
./tinyurl

// go to localhost:7777/
```

### Deploy with Docker
```
/code/tinyurl $ docker build -t tinyurl .
/code/tinyurl $ docker run -d -p 7777:7777 tinyurl

// go to localhost:7777/
```

## Demo Server
If you would like to try tinyURL out without running it yourself, there is an instance running at:
```
http://198.211.96.191:7777/
```
