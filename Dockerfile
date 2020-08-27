# using latest golang docker image
FROM golang as builder

# copy code
WORKDIR /go/src/the-posit-engine
COPY . .

# build the code
RUN chmod +x ./scripts/build.sh
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./scripts/build.sh

# run tests
RUN go test ./...

# using latest alpine docker image
FROM alpine

# copy the binary file
COPY --from=builder /go/src/the-posit-engine/bin/the-posit-engine .

# run the server
ENTRYPOINT ["./the-posit-engine"]