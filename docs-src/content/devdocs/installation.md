---
title: "Installation"
description: "How to build Qurl from sources, install and execute it on any host."
date: 2017-09-24T00:28:50+02:00
draft: false
weight: 2
---
To install *Qurl* the most easy way is to use the golang tooling and execute
the following:

```bash
$ go get install github.com/repejota/qurl
```

If you have already installed a previous version of *Qurl*, you can upgrade it
using the *-u* flag like this:

```bash
$ go get -u install github.com/repejota/qurl
```

## Building from sources

If you want to get the latest version or execute the development version of
*Qurl* you can build it from sources:

```bash
$ git clone https://github.com/repejota/qurl
Cloning into 'qurl'...
remote: Counting objects: 1018, done.
remote: Compressing objects: 100% (77/77), done.
remote: Total 1018 (delta 41), reused 86 (delta 22), pack-reused 903
Receiving objects: 100% (1018/1018), 1.16 MiB | 214.00 KiB/s, done.
Resolving deltas: 100% (539/539), done.
$ cd qurl
$ make install
```

* As you can see we use `make` as a tool to build *Qurl*. Please refer to our
  [Makefile](https://github.com/repejota/qurl/blob/master/Makefile) to see all
  available tasks.

#### Documentation

This documentation is a *Hugo* based static website. You will find the sources
under the `docs-src` folder.

To execute the development server for the documentation just use:

```bash
$ make docs-serve
```