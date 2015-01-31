gzip
-----

A simple `net.http.Handler` that adds gzip compression to your `http.Handler`s

Pretty basic, but avoids so typing...


#usage
```
package main

import (
    "net/http"
    "log"
    "github.com/julien/gzip"
)

func handler(func w.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello there"))
}

func main() {
    http.HandleFunc("/", GZip(handler))
    log.Fatalf(http.ListenAndServe(":8000", nil))
}
```

