## Routing for Go REST services 

### Understading Go net/http package 

```go 

package main 

import (
    "io" 
    "net/http" 
    "log"
)


func MyServer(w http.Response){
    io.WriteString(w, "hello world!\n" )
}

func main(){
    http.HandleFunc("/hello", MyServer) 
    log.Fatal(http.ListenAndServe(":8000", nil))
}

``` 

#### ServeMux a basic router in Go 

ServeMux is an HTTP request multiplexer. The HandleFunc is actually a method 
of ServeMux. By creating a new ServeMux, we can handle multiple routes. 
Before that, we can also create our own multiplexer 

A multiplexer just handles the logic of separating routes with a function called ServeHTTP method, it can do the job.

Consider a route as a key in a dictionary (map), then handler as its value. 
the router finds the handler from the route and tries to execute the 
ServeHTTP method


