---
title: "Header Queries"
description: "With Qurl you can fetch any header the server is returning on
each HTTP Request."
date: 2017-09-24T00:35:13+02:00
draft: false
weight: 2
---
Get any HTTP headers defined in the response, just adding a `header` parameter 
with the name of the header you want to get as value to the URL query.

```javascript
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

- Notice that the resulting value is not an string but a list of strings. This
  is because a single header can have multiple values in the same response.

In this case we will also return a list of values for the response header so
you'll be able to iterate over them.

```javascript
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

It is possible also to query more than one response header at once. Just append
as much `header` parameters with their names as you neeed to the query URL.

```javascript
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

If the header you query is not present in the response, we will still add it to
the result and its value will be `null`.

```javascript
$ curl -s 'http://localhost:8080/q?url=https://github.com&header=foobar' | json_pp
{
   "url" : "https://github.com",
   "status" : 200,
   "headers" : {
      "foobar" : null
   }
}
```

Finally if you query for an specific header more than once, the result will be
only populated once so bandwitch will be saved and the response data will be
simpler.

```javascript
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

- Mozila Developer Network website offers good and accurated documentation
  about what [HTTP Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)
  are their use. Their documentation is also available in multiple languages.
- For an updated list of standarizeed HTTP Headers, you can check the
  [RFC4229](https://tools.ietf.org/html/rfc4229) and the
  [IANA Message Headers](https://www.iana.org/assignments/message-headers/message-headers.xhtml)
  document.
- Wikipedia also maintains a list of standard and non-standard
  [Response HTTP Headers](https://en.wikipedia.org/wiki/List_of_HTTP_header_fields#Response_fields)
  in case you want more information.