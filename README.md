# IPFS2Png render+serve

[![Build Status](https://travis-ci.org/tekwrks/renderer.svg?branch=master)](https://travis-ci.org/tekwrks/renderer)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ftekwrks%2Frenderer.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Ftekwrks%2Frenderer?ref=badge_shield)
[![Go Report Card](https://goreportcard.com/badge/github.com/tekwrks/renderer)](https://goreportcard.com/report/github.com/tekwrks/renderer)

To build and run:
```
make && make run
```

```
At <IPFS_HASH>:

{
  content:     "Hello World!",
  description: "description",
  fontSize:    22,
  spacing:     1.5,
}
```

Navigate to: [http://localhost:3000/post/IPFS_HASH](http://localhost:3000/post/<IPFS_HASH>) and you should see 'Hello World!' in big friendly letters.

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ftekwrks%2Frenderer.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Ftekwrks%2Frenderer?ref=badge_large)
