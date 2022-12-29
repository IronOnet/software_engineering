package main 

import (
	"fmt" 
	"math/rand" 
	"net/http"
)

// CustomServeMux is a struct which can be a multiplexer  
type CustomeServeMux struct{

}


// This is the function handler to be overriden 
func (p *CustomeServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request){

	if r.URL.Path == "/"{
		giveRandom(w, r) 
		return 
	}
	http.NotFound(w, r) 
	return 
}

func giveRandom(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func main(){
	mux := &CustomeServeMux{} 

	// Adding multiple handlers using ServeMux 
	newMux := http.NewServeMux() 
	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request){
		fmt.Fptintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, rand.Intn(100))
	})

	http.ListenAndServe(":8000", newMux)

	http.ListenAndServe(":8000", mux)
}
