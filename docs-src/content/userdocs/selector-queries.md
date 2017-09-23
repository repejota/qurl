---
title: "Selector Queries"
date: 2017-09-24T00:35:28+02:00
draft: false
weight: 3
---
This is by far, the most powerful and useful feature of *Qurl*. You can use all the power of *CSS selectors* to retrieve any matching content of the target URL. 

Using *Qurl* query selectors you get an automated content scraping API tool of remote URL contents without needing any external dependencies neither any browser or aditional *HTML* parser/analyzer.

Just calling *Qurl* simple API you will be able to do the same than with other more complex less powerful tools and also in less time and resources.

### Basic Selectors

#### Type Selectors

Type selectors selects all elements that match the given HTML node name.

For each matching node it will return its *text* and a list of pairs of *key* and *value* for all the *attibutes* the node has.

```javascript
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

```javascript
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

The *attributes* for each mathing element from a *selector* is also a list as they could be more than one. In the following exaple you can see an example:

```javascript
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

- In the same way than the *Header Queries*, you can append more than one *selector* query parameter to your URL to query more than one selector with a single call to the API.
- And if there is no matching coincidence using a *selector* we won't add it to the response.

#### Class Selectors

### Combinators

combinators

### Pseudo-classes

pseudo-classes

### Pseudo-elements

pseudo-elements

- The Mozilla Developer Network website maintains accurated documentation about the standarized [CSS Selectors](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Selectors) and how to use them to match content.
