### STAGE 1: Build ###
# Building the binary of the App
#
#
# We label our stage as 'builder'
FROM golang:1.20 AS builder

ENV PATH=*/go/bin:${PATH}
ENV CGO_ENABLED=0
ENV GO1111MODULE=on

RUN mkdir /go/src/ViniciusDJM
WORKDIR /go/src/ViniciusDJM

# Downloads all the dependencies in advance (could be left out, but it's more clear this way) 
COPY go.mod go.sum ./
RUN go mod download

# Copy all the Code and stuff to compile everything
ADD . .

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -a -installsuffix cgo -o process_search_app .
#
#
# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest


WORKDIR /app


# Copy the generated binary from builder image to execution image
COPY --from=builder /go/src/ViniciusDJM/process_search_app /bin/process_search_app

# Expose port 8000 to the outside world
EXPOSE 8000

# Run the binary program produced by `go build`
ENTRYPOINT ["/bin/process_search_app"]

CMD ["/bin/process_search_app"]
