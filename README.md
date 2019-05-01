# IPFS2Png render

[![Build Status](https://travis-ci.org/multiplio/cast-render.svg?branch=master)](https://travis-ci.org/multiplio/cast-render)
[![Go Report Card](https://goreportcard.com/badge/github.com/multiplio/cast-render)](https://goreportcard.com/report/github.com/multiplio/cast-render)

Render posts from IPFS into html, linkable to social media sites.

## Environment (with defaults)
```
NAME=render
ADDRESS=

DPI=
FONTFILE=fonts/NotoSans-Regular.ttf

POST_TEMPLATE=templates/post.mustache

IPFS_ADDRESS=
ROOT_URL=
```

## Routes
```
get /ready => for kubernetes readiness probe

get /post/<hash> => formatted html of post (determined by POST_TEMPLATE)

get /post/<hash>/image => rendered post image
```

## Dependencies

Expects IPFS blocks in the following format:
```
{
  content:     "Hello World!",
  description: "Template message.",
  fontSize:    36,
  spacing:     1.5,
}
```

