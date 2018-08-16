# 1
### Goal: have a working REST APIs with no persistent database.

## Prerequisite
You can install Postman to send request to APIs here.

The basic concepts of APIs and REST can be found on these sites: 
- lkjsdf
- lsdkjf
- sldkfj

## Explanation
Go comes with a lot of useful packages to get developer to get started. It takes just a couple of lines to make a simple web server as below.
```
package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello world!")
}
```
So what we just did was telling that when there is a http request to the "/" endpoint, we hand it to the handler function, which irrespective of the request coming, write the "Hello world" back as a response. And then we serve it at a specific port, combined with log's Fatal function to log whenever there is something wrong happened.

Pretty neat, right? Now let's get to the exciting part.
