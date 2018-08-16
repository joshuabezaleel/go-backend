# 1
Goal: have a working REST APIs with no persistent database.

## Prerequisite knowledge
The basic concepts of APIs and REST can be found on these sites: 
- lkjsdf
- lsdkjf
- sldkfj

## Explanation
Go comes with a lot of useful libraries to get developer to get started. It takes just a couple of lines to make a simple web server as below.
```
package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
```
