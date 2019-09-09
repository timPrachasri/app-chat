#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------

#Builder stage
FROM golang:1.11-alpine3.7 as builder
# Update OS package and install Git
RUN apk update && apk add git && apk add build-base
# Set working directory
WORKDIR /app
# Get Golang dependency management tool
RUN go get github.com/golang/dep/cmd/dep
# Install Fresh for local development
RUN go get github.com/pilu/fresh
# Install go tool for convert go test output to junit xml
RUN go get -u github.com/jstemmer/go-junit-report &&\
    go get github.com/axw/gocov/... &&\
    go get github.com/AlekSi/gocov-xml &&\
    go get -u github.com/golang/lint/golint
# Install wait-for
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /usr/local/bin/wait-for &&\
    chmod +x /usr/local/bin/wait-for

# Copy Go dependency file
COPY go.mod .
COPY go.sum .

# run go mod download
RUN go mod download

# Copy Go source code
COPY . .

ENV PORT=880

EXPOSE 880
# Set Docker's entry point commands
RUN cd app/ && go build

# ------------------------------------------------------------------------------
# Deployment image
# ------------------------------------------------------------------------------

#App stage
FROM golang:1.11-alpine3.7
# Set working directory
WORKDIR /go/src/github.com/TimothyPrachasri/my-simple-app-chat/server
#Get artifact from buiber stage
COPY --from=builder /usr/local/bin/wait-for /usr/local/bin/wait-for
COPY --from=builder /go/src/github.com/TimothyPrachasri/my-simple-app-chat/server/app app/
COPY --from=builder /go/src/github.com/TimothyPrachasri/my-simple-app-chat/server/vendor vendor
# Set Docker's entry point commands

ENV PORT=880

EXPOSE 880

CMD cd app/ && ./app;
