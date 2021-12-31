<!--
SPDX-FileCopyrightText: 2021 Amolith <amolith@secluded.site>

SPDX-License-Identifier: CC0-1.0
-->

# go-webring
Simple webring implementation

After [installing Go](https://go.dev/dl/) and [pandoc](https://pandoc.org/) …
``` shell
git clone https://git.sr.ht/~amolith/go-webring
cd go-webring
pandoc -s index.md -o index.html
go build
./go-webring
xdg-open http://127.0.0.1:2857
```

## Usage

``` text
$ ./go-webring -h
Usage of ./go-webring
  -i, --index string     Path to home page template (default "index.html")
  -l, --listen string    Host and port go-webring will listen on (default "127.0.0.1:2857")
  -m, --members string   Path to list of webring members (default "list.txt")
```

See the included `list.txt` and `index.md` for examples of a webring setup. To
run `go-webring` with those examples, first generate `index.html` from
`index.md` like so:

``` shell
$ pandoc -s index.md -o index.html
```

In *your* home page template, place the string `{{ . }}` wherever you want the
table of members inserted. It's a regular HTML table so you can style it with
CSS in the template's `<head>` or in a separate `.css` file.

Given the example `list.txt`, the following is what `go-webring` would insert,
so it's up to you to add the starting `<table>` and `<th>` tags as in the
example `index.md`.

``` html
<tr>
  <td>member1</td>
  <td><a href="https://example.com">example.com</a><td>
</tr>
<tr>
  <td>member2</td>
  <td><a href="https://sometilde.com/~member2">sometilde.com/~member2</a><td>
</tr>
```

Note that you don't *have* to create your webring home page in Markdown, that's
just what I found easiest. As long as your HTML template includes `{{ . }}`, the
table will show up.
