# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.13

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/ferdingler/go-hammer/cli

# Install dependencies
WORKDIR /go/src/github.com/ferdingler/go-hammer/cli
RUN go get
RUN go install

# Run the CLI by default when the container starts.
ENTRYPOINT ["/go/bin/cli"]
CMD [ "--help" ]