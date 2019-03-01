# Start from golang v1.11.5 base image
FROM golang:1.11.5

# Add Maintainer Info
LABEL maintainer="Sean - v33ps@protonmail.com"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/v33ps/tinyurl

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

RUN go build -o tinyurl

# This container exposes port 8080 to the outside world
EXPOSE 7777

# Run the executable
CMD ["./tinyurl"]
