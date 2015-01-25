# gok

A command line bookmark in Go Lang. It store the URL into a BoltDB and
index into Bleve Search.

Pocket is cool but it's too slow and doesn't have a good search feature.

# Build status

[![wercker status](https://app.wercker.com/status/bf910e4c4a30dc6c9293f037163dcdec/m "wercker status")](https://app.wercker.com/project/bykey/bf910e4c4a30dc6c9293f037163dcdec)

# Install

```
$ go install ...
```

# Using

### Add an URL

```
$ gok a url
```

the URL will be fetched, the title and body is then indexed.

### Search

```
$ gok s keyword
```

### List all item

```
$ gok l
```

# GetPocket import

GetPocket used OAuth2 for authentication. It requires a step where we
have to run the authentication in a GUI and authorize the access on
pocket website. To do that, we implement a small HTTP in `gok` which do
authentication, get the access token and push it to other co-routine to
retrieve and index the URL

```
$ gok import pocket --consumer_key=yourconsumerkeyhere
```

# TODO

Lots of thing

* Import from Pocket
* Export to JSON
* Trigger to Pocket

# Should I use it

TL'DR: No, you should not.

This is created to learn some more Go and Bolt and Bleve search only.

