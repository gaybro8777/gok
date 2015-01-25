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

# TODO

Lots of thing

* Import from Pocket
* Export to JSON
* Trigger to Pocket

# Should I use it

TL'DR: No, you should not.

This is created to learn some more Go and Bolt and Bleve search only.

