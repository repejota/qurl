# qurl

[![License][License-Image]][License-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]
[![GoDoc](https://godoc.org/github.com/repejota/qurl?status.svg)][GoDoc-Reference-Url]

## Table of contents

* [Introduction](https://github.com/repejota/qurl#introduction)
* [Documentation](https://github.com/repejota/qurl#documentation)
* [Continuous Integration](https://github.com/repejota/qurl#continuous-integration)
  * [Tests](https://github.com/repejota/qurl#license)
  * [Coverage](https://github.com/repejota/qurl#coverage)
* [License](https://github.com/repejota/qurl#license)

## Introduction

Qurl is a drop-in and easy to deploy microservice who exposes an HTTP API you
can use to extract content from any web page as JSON. Any information available
at a given public URL can be extracted using selector queries (check examples
below).

Example:

```javascript
$ curl -s 'http://localhost:8080/q?url=https://example.com&header=Content-Type&selector=meta' | json_pp
{
   "url" : "https://example.com",
   "status" : 200,
   "selectors" : {
      "meta" : [
         {
            "text" : "",
            "attributes" : [
               {
                 "key" : "charset",
                  "value" : "utf-8"
               }
            ]
         },
         {
           "text" : "",
            "attributes" : [
               {
                  "key" : "name",
                  "value" : "viewport"
               },
               {
                 "key" : "content",
                  "value" : "width=device-width, initial-scale=1"
               }
            ]
         }
      ]
   },
   "headers" : {
      "Content-Type" : [
         "text/html; charset=utf-8"
      ]
   }
}
```

## Documentation

Qurl is really powerful and full of posibilities, that's why we have been
creating an entire site that documents all the features you can use.

Plesae refer to the [documentation](https://repejota.github.io/qurl/) website
where you'll gonna find accurated information aboug how to use the API that
provides Qurl but also how to download, build and launch your own insgtance
of the service.

**[Visit the documentation](https://repejota.github.io/qurl/)**

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
[Godoc-Reference-Url]: http://godoc.org/github.com/repejota/qurl