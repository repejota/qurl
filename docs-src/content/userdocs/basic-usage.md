---
title: "Basic Usage"
date: 2017-09-24T00:27:22+02:00
draft: false
weight: 1
---
The most simple query only returns the HTTP Status code response when fetching
url contents.

```javascript
$ curl -s http://qurl.io/q?url=https://example.com | json_pp
{
   "url" : "https://www.example.com",
   "status" : 200
}
```

- We use `-s` curl flag to enable silent or quiet mode. So curl won't show
progress meter or error messages. It will still output the resulting data.
- Also the result is piped to `json_pp` command to pretty print the JSON data.
- See man pages [curl(1)](http://www.manpagez.com/man/1/curl/) and
[json_pp(1)](http://www.manpagez.com/man/1/json_pp/) for more information.
