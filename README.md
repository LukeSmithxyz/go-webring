<!--
SPDX-FileCopyrightText: 2021 Amolith <amolith@secluded.site>

SPDX-License-Identifier: CC0-1.0
-->

# go-webring
[![REUSE status][reuse-shield]][reuse]
[![Donate with fosspay][fosspay-shield]][fosspay]
![Time spent on project][wakapi-shield]

Simple webring implementation created for the [Fediring](https://fediring.net/)

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

This webring implementation handles four paths:
- **Root:** returns the home page template replacing the string "`{{ . }}`" with
  an HTML table of ring members
- **Next:** returns a 302 redirect pointing to the next site in the list
- **Previous:** returns a 302 redirect pointing to the previous site in the list
- **Random:** returns a 302 redirect pointing to a random site in the list

The **next** and **previous** paths require a `?host=` parameter containing a
URL-encoded URI of the site being visited. For example, if Sam is a member of a
webring on `ring.com` and her site is `sometilde.com/~sam`, she will need the
following links on her page for directing visitors to the next/previous ring
members.

- `https://ring.com/next?host=sometilde.com%2F~sam`
- `https://ring.com/previous?host=sometilde.com%2F~sam`

### With provided examples

See the included `list.txt` and `index.md` for examples of a webring setup. To
run `go-webring` with those examples, first install [pandoc](https://pandoc.org)
then generate `index.html` from `index.md` like so:

``` shell
$ pandoc -s index.md -o index.html
```

Next, you'll need to [install Go](https://go.dev/dl) and build the project.

``` shell
$ go build
```

After that, simply execute the binary then open
[localhost:2857](http://localhost:2857) in your browser.

``` shell
$ ./go-webring
```

### With custom files

To run your own webring, you'll first need a template homepage. This should be
any HTML file with the string "`{{ . }}`" placed wherever you want the table of
members inserted. This table is plain HTML so you can style it with CSS in the
template's `<head>` or in a separate `.css` file.

Pandoc produces very pleasing (in my opinion) standalone HTML pages; if you just
want something simple, I would recommend modifying the included `index.md` and
generating your homepage as in section above.

Next, you'll need a text file containing a list of members. On each line should
be the member's unique identifer (such as their username) followed by a single
space followed by their site's URI omitting the scheme. For example, if a user
is `bob` and his site is `https://bobssite.com`, his line would look like the
following.

``` text
bob bobssite.com
```

If the user was `sam` and her site was `https://sometilde/~sam`, her line would
look like this:

``` text
sam sometilde/~sam
```

With those two members in the text file, the HTML inserted into the home page
will be the following.

``` html
<tr>
  <td>bob</td>
  <td><a href="https://bobssite.com">bobssite.com</a><td>
</tr>
<tr>
  <td>sam</td>
  <td><a href="https://sometilde.com/~sam">sometilde.com/~sam</a><td>
</tr>
```

Assuming this webring is on `ring.com`, Bob will need to have the following
links on his page.

- `https://ring.com/next?host=bobssite.com`
- `https://ring.com/previous?host=bobssite.com`

Because Sam has a forward slash in her URI, she'll need to percent-encode it so
browsers interpret the parameter correctly.

- `https://ring.com/next?host=sometilde.com%2F~sam`
- `https://ring.com/previous?host=sometilde.com%2F~sam`


## Questions & Contributions
Questions, comments, and patches can always be sent to my public inbox, but I'm
also in my IRC channel/XMPP room pretty much 24/7. However, I might not see
messages right away because I'm working on something else (or sleeping) so
please stick around!

If you're wanting to introduce a new feature and I don't feel like it fits with
this project's goal, I encourage you to fork the repo and make whatever changes
you like!

- Email: [~amolith/public-inbox@lists.sr.ht][email]
- IRC: [irc.nixnet.services/#secluded][irc]
- XMPP: [secluded@muc.secluded.site][xmpp]

*If you haven't used mailing lists before, please take a look at [SourceHut's
documentation](https://man.sr.ht/lists.sr.ht/), especially the etiquette
section.*

[reuse]: https://api.reuse.software/info/git.sr.ht/~amolith/go-webring
[reuse-shield]: https://shields.io/reuse/compliance/git.sr.ht/~amolith/go-webring

[fosspay]: https://secluded.site/donate/
[fosspay-shield]: https://shields.io/badge/donate-fosspay-yellow

[wakapi-shield]: https://img.shields.io/endpoint?url=https://waka.secluded.site/api/compat/shields/v1/amolith/project:go-webring&color=blue&label=time%20spent

[email]: mailto:~amolith/public-inbox@lists.sr.ht
[irc]: irc://irc.nixnet.services/#secluded
[xmpp]: xmpp:secluded@muc.secluded.site?join
