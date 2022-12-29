### Introduction 

Types of web services:  
    - SOAP 
    - UDL 
    - WSDL 
    - REST

A SOAP request usually consists of theses basic components: 
     - Envelope 
     - Header 
     - Body

Characteristics of REST based web services 

    - Client based architecture 
        - This architecture is most essential for the modern web to 
        communicate over HTTP.

    - Statelessness: 
        - This is the most important characteristic of a REST service. 
        A REST HTTP request consists of all the data needed by the server
        to understand and give back the response. Once a request is served 
        the server does'nt remember if the request arrived after a while. 
        So the operation will be a stateless one 

    - Cacheable: In order to scale an application well, we need to cache
    content and deliver it as a response. if the cache is not valid, it is 
    our responsibility to bust it. REST services should be properly cached 
    for scaling

    - Scripts on demand: H

    - Multiple layered system: The REST API can be served from multiple
    servers.

    - Representation of resources: The REST API provides a uniform interface
    to talk to. It uses a Uniform Resource Identifier (URI) to map resources (data). It has also has the advantage of requesting a specific data format 
    as a response.

* The internet MEDIA type (MIME type) can tell the server that the requested resource is of that particular type.
**
    - Implementational freedom: REST is just a mechanism to define your web 
    services. it is an architectural style that can be implemented in multiple 
    ways.


    - SPA reduces the bandwidth and improves performance

Three reasons to use GO for your RESTful webservice: 
    - To scale your API to a wider audience 
    - To enable your developers to build robust systems 
    - To invest in the future viability of your projects

```go 
package romanNumerals 

var Numerals = map[int]string{
    10: "X", 
    9: "IX", 
    8: "VIII", 
    7: "VII", 
    6: "VI", 
    5: "V", 
    4: "IV", 
    3: "III", 
    2: "II", 
    1: "I"
}
```


```go 

package main 

import (
    "fmt" 
    "github.com/narenaryan/romanNumerals" 
    "html" 
    "net/http" 
    "strconv" 
    "strings"
    "time"
)

func main(){
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        urlPathElements := string.Split(r.URL.Path, "/") 
        // if request is GET with correct syntax 
        if urlPathElements[1] == "roman_number"{
            number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
            if number == 0 || number > 10{
                // If resource is not in the list, send not found status 
                w.WriteHeader(http.StatusNotFound) 
                w.Write([]byte("404 - Not Found"))
            } else{
                fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
            }
        } else{
            // for all other requests, tell that client send a bad 
            // request 
            w.WriteHeader(http.StatusBadRequest)
            w.Write([]byte("400 - Bad Request"))
        }
    })


    // Create a server and run it on port 8080 
    s := &http.Server{
        Addr: "8080", 
        ReadTimeout: 10 * time.Second, 
        WriteTimeout: 10 * time.Second, 
        MaxHeaderBytes: 1 << 20,
    }

    s.ListenAndServe()
}

```