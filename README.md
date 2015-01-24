# gok

A command line bookmark in Go Lang. It store the URL into a BoltDB and
index into Bleve Search

# Install

```
$ go install ...
```

# Using

### Add an URL

```
$ go a url
```

the URL will be fetched, the title and body is then indexed.

### Search

```
$ go s keyword
```

### List all item

```
$ go l
```

# TODO

Lots of thing

# Should I use it

TL'DR: No, you should not.

This is created to learn some more Go and Bolt and Bleve search only.

