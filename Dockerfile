FROM golang:alpine
COPY . /app
WORKDIR /app
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go build -o GoVueApp
CMD ./GoVueApp -dev