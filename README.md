# http-proxy

Golang reverse HTTP proxy - sometimes useful in development.

Use `UPSTREAM` env var to control upstream endpoint.

```
 -> docker run -it --rm -p 8080:8080 -e UPSTREAM=http://httpbin.org docker.io/michaeldonat/http-proxy
 -> curl localhost:8080/get
{
  "args": {},
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip",
    "Connection": "close",
    "Host": "httpbin.org",
    "User-Agent": "curl/7.54.0"
  },
  "origin": "172.17.0.1, 94.195.82.83",
  "url": "http://httpbin.org/get"
}

```

