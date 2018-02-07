# Golang
FROM golang:1.9.3
WORKDIR /go/src/github.com/akshaysuresh4793
RUN go get -u "github.com/akshaysuresh4793/zapposapi" && cd zapposapi && go get -v && go build main.go restaurant.go menu.go menuitem.go location.go db.go json.go error.go response.go cache.go && cp main /usr/bin/server