## Reproducing the problem

The sample code is simple. It just requires `bcrypt` and prints out a hash.

Try to run `go run main.go` and `go run -mod=vendor main.go` to feel the time difference.

## Problem

Because `golang.org/x/...` is blocked by the GFW in China, I follow one solution in [this issue](https://github.com/golang/go/issues/28652#issuecomment-443745942) and alias `golang.org/x/crypto@v0.0.0` to `github.com/golang/crypto@latest` 

```bash
$ go mod edit -require=golang.org/x/crypto@v0.0.0
$ go mod edit -replace=golang.org/x/crypto@v0.0.0=github.com/golang/crypto@latest
```

now my `go.mod` looks like this:

```bash
module github.com/levblanc/x-crypto-issue

require golang.org/x/crypto v0.0.0

replace golang.org/x/crypto v0.0.0 => github.com/golang/crypto v0.0.0-20190123085648-057139ce5d2b
```

But when I run `go run main.go`, it took me around 30 seconds to print out the result. EVERYTIME.

Then I created the `vendor` folder and tried the `-mod=vendor` flag with `go run`, and it worked well!

## Reason

I digged into the `bcrypt.go` source code, and found that it was importing `golang.org/x/crypto/blowfish`,

which is not in my pkg path, because I alias crypto to the github mirror:

```bash
/Users/levblanc/go/pkg/mod/github.com/golang/crypto@v0.0.0-20190123085648-057139ce5d2b
```

While in the `vendor` folder, `blowfish` is downloaded.

And I am guessing this might be the reason of the time difference?

I've seen the `go mod init` command adding subdependencies into `go.mod`, and why is this not the case for `blowfish`？