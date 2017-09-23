# qurl

[![License][License-Image]][License-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]

## Table of contents

* [Introduction](https://github.com/repejota/qurl#introduction)
* [User Documentation](https://github.com/repejota/qurl#user-documentation)
	* [Basic Usage](https://github.com/repejota/qurl#basic-usage)
	* [Header Queries](https://github.com/repejota/qurl#header-queries)
	* [Selector Queries](https://github.com/repejota/qurl#selector-queries)
* [Developer Documentation](https://github.com/repejota/qurl#developer-documentation)
* [Continuous Integration](https://github.com/repejota/qurl#continuous-integration)
  * [Tests](https://github.com/repejota/qurl#license)
  * [Coverage](https://github.com/repejota/qurl#coverage)
* [License](https://github.com/repejota/qurl#license)

## Introduction

Qurl is a drop-in and easy to deploy microservice who exposes an HTTP API you can use to extract content from any web page as JSON. Any information available at a given public URL can be extracted using selector queries (check examples below).

## User Documentation

User documentation

### Basic Usage

The most simple query only returns the HTTP Status code response when fetching url contents.

Example:

```
$ curl -s http://qurl.io/q?url=https://example.com | json_pp 
{
   "url" : "https://www.example.com",
   "status" : 200
}
```

> We use `-s` curl flag to enable silent or quiet mode. So curl won't show progress meter or error messages. It will still output the resulting data.

> Also the result is piped to `json_pp` command to pretty print the JSON data.

> See man pages [curl(1)](http://www.manpagez.com/man/1/curl/) and [json_pp(1)](http://www.manpagez.com/man/1/json_pp/) for more information.

### Header Queries

header queries

Example:

```
$ curl -s 'http://localhost:8080/q?url=https://github.com&header=Content-Type' | json_pp
{
   "url" : "https://github.com",
   "status" : 200,
   "headers" : {
      "Content-Type" : [
         "text/html; charset=utf-8"
      ]
   }
}
```

with multiple values

```
$ curl -s 'http://localhost:8080/q?url=https://github.com&header=Set-Cookie' | json_pp
{
   "url" : "https://github.com",
   "status" : 200,
   "headers" : {
      "Set-Cookie" : [
         "_octo=GH1.1.1872640429.1506145235; domain=.github.com; path=/; expires=Mon, 23 Sep 2019 05:40:35 -0000",
         "logged_in=no; domain=.github.com; path=/; expires=Wed, 23 Sep 2037 05:40:35 -0000; secure; HttpOnly",
         "_gh_sess=eyJzZXNzaW9uX2lkIjoiNGU3ZmEwMGUzN2RkOTFkYmYyYjFhM2RmODA3YTc4M2QiLCJsYXN0X3JlYWRfZnJvbV9yZXBsaWNhcyI6MTUwNjE0NTIzNTY2NywiX2NzcmZfdG9rZW4iOiJWMXFCYUlsT1h3YXZTYTErVWJyNCsvRnFFcE5zNVdxUGczUzVKSlo1bTZnPSJ9--8dd729a38c9647489eb95438071a0a5bf083edb6; path=/; secure; HttpOnly"
      ]
   }
}
```

more than one header on the same query

```
$ curl -s 'http://localhost:8080/q?url=https://github.com&header=Date&header=Cache-Control' | json_pp
{
   "url" : "https://github.com",
   "status" : 200,
   "headers" : {
      "Cache-Control" : [
         "no-cache"
      ],
      "Date" : [
         "Sat, 23 Sep 2017 05:47:30 GMT"
      ]
   }
}
 
```

if the header is not found in the response it is also added to the result but with null value

```
$ curl -s 'http://localhost:8080/q?url=https://github.com&header=foobar' | json_pp
{
   "url" : "https://github.com",
   "status" : 200,
   "headers" : {
      "foobar" : null
   }
}
```

### Selector Queries

selector queries

## Developer Documentation

Developer documentation

## Continuous integration

### Tests

* Develop: [![CircleCI](https://circleci.com/gh/repejota/qurl/tree/develop.svg?style=svg)](https://circleci.com/gh/repejota/qurl/tree/develop)
* Master: [![CircleCI](https://circleci.com/gh/repejota/qurl/tree/master.svg?style=svg)](https://circleci.com/gh/repejota/qurl/tree/master)

### Coverage

* Develop: [![Coverage Status](https://coveralls.io/repos/github/repejota/qurl/badge.svg?branch=develop)](https://coveralls.io/github/repejota/qurl?branch=develop)
* Master: [![Coverage Status](https://coveralls.io/repos/github/repejota/qurl/badge.svg?branch=master)](https://coveralls.io/github/repejota/qurl?branch=master)


## License

(The MIT License)

Copyright (c) 2017 qurl Authors.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to
deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/badge/License-MIT-blue.svg
[ReportCard-Url]: http://goreportcard.com/report/repejota/qurl
[ReportCard-Image]: http://goreportcard.com/badge/github.com/repejota/qurl