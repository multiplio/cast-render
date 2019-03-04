# IPFS2Png render+serve

[![Build Status](https://travis-ci.org/tekwrks/renderer.svg?branch=master)](https://travis-ci.org/tekwrks/renderer)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ftekwrks%2Frenderer.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Ftekwrks%2Frenderer?ref=badge_shield)
[![Go Report Card](https://goreportcard.com/badge/github.com/tekwrks/renderer)](https://goreportcard.com/report/github.com/tekwrks/renderer)

Render posts from IPFS into html, linkable to social media sites.

## Environment (with defaults)
```
NAME=
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

get /post/<hash>/image => renderer post image
```

## Dependencies

Expects IPFS blocks in the following format:
```
At <IPFS_HASH>:

{
  content:     "Hello World!",
  description: "description",
  fontSize:    22,
  spacing:     1.5,
}
```

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ftekwrks%2Frenderer.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Ftekwrks%2Frenderer?ref=badge_large)
