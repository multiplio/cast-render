FROM golang:1.7.3 as builder
WORKDIR /go/src/github.com/WhoMeNope/notatweet
RUN go get -d -v golang.org/x/image/font \
 && go get -d -v golang.org/x/image/math/fixed \
 && go get -d -v github.com/golang/freetype \
 && go get -d -v github.com/golang/freetype/truetype \
 && go get -d -v github.com/valyala/fasthttp
COPY notatweet ./notatweet
COPY textrender ./textrender
RUN GOPATH=/go GOOS=linux CGO_ENABLED=0 go install -a -installsuffix cgo ./textrender
RUN GOPATH=/go GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o app ./notatweet

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/WhoMeNope/notatweet/notatweet/fonts ./fonts
COPY --from=builder /go/src/github.com/WhoMeNope/notatweet/app .
EXPOSE 3000
# ENTRYPOINT ["./app"]
CMD ["./app"]
