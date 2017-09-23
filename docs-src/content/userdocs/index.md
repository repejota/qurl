+++
title = "QURL User Documentation"
category = "documentation"
+++
# Introduction

Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim.

## Basic Usage

The most simple query only returns the HTTP Status code response when fetching url contents.

Example:

```bash
$ curl -s http://qurl.io/q?url=https://example.com | json_pp 
{
   "url" : "https://www.example.com",
   "status" : 200
}
```

> We use `-s` curl flag to enable silent or quiet mode. So curl won't show progress meter or error messages. It will still output the resulting data.

> Also the result is piped to `json_pp` command to pretty print the JSON data.

> See man pages [curl(1)](http://www.manpagez.com/man/1/curl/) and [json_pp(1)](http://www.manpagez.com/man/1/json_pp/) for more information.

## Header Queries

Get any HTTP headers defined in the response, just adding a `header` parameter with the name of the header you want to get as value to the URL query.

Example:

```bash
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

> Notice that the resulting value is not an string but a list of strings. This is because a single header can have multiple values in the same response. 

In this case we will also return a list of values for the response header so you'll be able to iterate over them.

Example:

```bash
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

It is possible also to query more than one response header at once. Just append as much `header` parameters with their names as you neeed to the query URL.

Example:

```bash
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

If the header you query is not present in the response, we will still add it to the result and its value will be `null`.

Example:

```bash
$ curl -s 'http://localhost:8080/q?url=https://github.com&header=foobar' | json_pp
{
   "url" : "https://github.com",
   "status" : 200,
   "headers" : {
      "foobar" : null
   }
}
```

Finally if you query for an specific header more than once, the result will be only populated once so bandwitch will be saved and the response data will be simpler.

Example:

```bash
$ curl -s 'http://localhost:8080/q?url=https://github.com&header=Date&header=Date' | json_pp
{
   "url" : "https://github.com",
   "status" : 200,
   "headers" : {
      "Date" : [
         "Sat, 23 Sep 2017 06:03:56 GMT"
      ]
   }
}
```

> Mozila Developer Network website offers good and accurated documentation about what [HTTP Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers) are their use. Their documentation is also available in multiple languages.

> For an updated list of standarizeed HTTP Headers, you can check the [RFC4229](https://tools.ietf.org/html/rfc4229) and the [IANA Message Headers](https://www.iana.org/assignments/message-headers/message-headers.xhtml) document. 

> Wikipedia also maintains a list of standard and non-standard [Response HTTP Headers](https://en.wikipedia.org/wiki/List_of_HTTP_header_fields#Response_fields) in case you want more information.

## Selector Queries

This is by far, the most powerful and useful feature of *Qurl*. You can use all the power of *CSS selectors* to retrieve any matching content of the target URL. 

Using *Qurl* query selectors you get an automated content scraping API tool of remote URL contents without needing any external dependencies neither any browser or aditional *HTML* parser/analyzer.

Just calling *Qurl* simple API you will be able to do the same than with other more complex less powerful tools and also in less time and resources.

### Basic Selectors

#### Type Selectors

Type selectors selects all elements that match the given HTML node name.

For each matching node it will return its *text* and a list of pairs of *key* and *value* for all the *attibutes* the node has.

Example:

```bash
$ curl -s 'http://localhost:8080/q?url=https://example.com&selector=title' | json_pp
{
   "url" : "https://example.com",
   "status" : 200,
   "selectors" : {
      "title" : [
         {
            "text" : "Example Domain",
            "attributes" : null
         }
      ]
   }
}
```

As you can see for each *selector* it will return a list of matching elements, in the following example we are getting more than one element for a single selector.


Example:

```bash
$ curl -s 'http://localhost:8080/q?url=https://example.com&selector=p' | json_pp
{
   "url" : "https://example.com",
   "status" : 200,
   "selectors" : {
      "p" : [
         {
            "text" : "This domain is established to be used for illustrative examples in documents. You may use this\n    domain in examples without prior coordination or asking for permission.",
            "attributes" : null
         },
         {
            "text" : "More information...",
            "attributes" : null
         }
      ]
   }
}
```

The *attributes* for each mathing element from a *selector* is also a list as they could be more than one. In the following exaple you can see an example  

Example:

```bash
$ curl -s 'http://localhost:8080/q?url=https://example.com&selector=meta' | json_pp
{
   "url" : "https://example.com",
   "selectors" : {
      "meta" : [
         {
            "text" : "",
            "attributes" : [
               {
                  "value" : "utf-8",
                  "key" : "charset"
               }
            ]
         },
         {
            "text" : "",
            "attributes" : [
               {
                  "key" : "http-equiv",
                  "value" : "Content-type"
               },
               {
                  "value" : "text/html; charset=utf-8",
                  "key" : "content"
               }
            ]
         },
         {
            "attributes" : [
               {
                  "key" : "name",
                  "value" : "viewport"
               },
               {
                  "value" : "width=device-width, initial-scale=1",
                  "key" : "content"
               }
            ],
            "text" : ""
         }
      ]
   },
   "status" : 200
}
```

> In the same way than the *Header Queries*, you can append more than one *selector* query parameter to your URL to query more than one selector with a single call to the API.

> And if there is no matching coincidence using a *selector* we won't add it to the response.

#### Class Selectors

### Combinators

combinators

### Pseudo-classes

pseudo-classes

### Pseudo-elements

pseudo-elements

> The Mozilla Developer Network website maintains accurated documentation about the standarized [CSS Selectors](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Selectors) and how to use them to match content.
