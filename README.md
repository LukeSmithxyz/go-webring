# go-webring

A go program for managing a webring.

This is originally a fork of [Amolith](https://secluded.site)'s [go-webring](https://sr.ht/~amolith/go-webring/), which
is used for [Fediring](https://fediring.net/).

``` shell
git clone https://git.sr.ht/~amolith/go-webring
cd go-webring
go build
cp example-index.html index.html
./go-webring
```

The service runs on port 2857, so you can have NginX/etc. proxy_pass to that port for the domain you want to use as a
host.

The `index.html` is the main webpage of your ring. It can be customized to your liking and a list of the sites in the
ring will be addedd in the place of the `{{ . }}`.

## Pages

This program produces four main pages for use as links:

- `/` -- The index. A mainpage in `index.html` where you can customize your webring homepage. This lists all sites in
  the webring automatically.
- `/next` -- Go to the next site in the webring (from whichever site you are linked from).
- `/previous` -- Go to the previous site in the webring.
- `/random` -- Go to a random site in the webring.

If linked to `/next` or `/previous` without a HTTP referer in the webring, they will function like `/random`.

## Custom Usage

``` text
$ ./go-webring -h
Usage of ./go-webring
  -i, --index string     Path to home page template (default "index.html")
  -l, --listen string    Host and port go-webring will listen on (default "127.0.0.1:2857")
  -m, --members string   Path to list of webring members (default "list.txt")
```
