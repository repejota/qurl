---
title: ""
date: 2017-03-02T12:00:00-05:00

---
## Introduction

*Qurl* is a drop-in and easy to deploy microservice who exposes an *HTTP API*
you can use to extract content from any web page as *JSON*. Any information
available at a given public *URL* can be extracted using selector queries 
(check examples below).

Quick example:

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