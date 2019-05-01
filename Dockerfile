# Builder
FROM golang:1.12 as builder

WORKDIR /cast-render

# Force the go compiler to use modules
ENV GO111MODULE=on
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the source code
COPY . .
# Compile the project
RUN GOPATH=/go GOOS=linux CGO_ENABLED=0 go install -a -installsuffix cgo .

# Deploy
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app/
COPY --from=builder /cast-render/fonts ./fonts
COPY --from=builder /cast-render/templates ./templates
COPY --from=builder /go/bin/cast-render .

# defaults
ENV FONTFILE=fonts/NotoSans-Regular.ttf
ENV POST_TEMPLATE=templates/post.mustache

EXPOSE 3000
ENTRYPOINT ["/app/cast-render"]
