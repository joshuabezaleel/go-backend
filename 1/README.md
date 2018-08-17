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

Try to run the program with the command ```go run main.go``` and hover to your browser with the address of ```localhost:8080``` and you will be greeted by the program.
Pretty neat, right? Now let's get to the exciting part.

First, we'll need to get a Go package from Gorilla toolkit called mux which we can get by the command
```go get github.com/gorilla/mux```
and which we will use to direct an incoming request to a specific routes. You can say it as a yellow pages for the HTTP requests.

Let's say that we're going to make an API for an e-commerce, and first we'll start with its products. The endpoints that we'll probably need are:
- ```/products``` (GET) --> retrieve all of the products listed
- ```/product/{id}``` (GET) --> retrieve a product on a specific id
- ```/product/{id}``` (POST) --> create a new product
- ```/product/{id}``` (DELETE) --> delete a product


