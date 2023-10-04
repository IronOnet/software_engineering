package main

import "fmt"

func main(){
  addidasFactory, _ := GetSportsFactory("addidas")
  nikeFactory , _ := GetSportsFactory("nike")

  nikeShoe := nikeFactory.makeShoe()
  nikeShirt := nikeFactory.makeShirt()

  addidasShoe := addidasFactory.makeShoe()
  addidasShirt := addidasFactory.makeShirt()

  printShoeDetails(addidasShoe)
  printShoeDetails(nikeShoe)

  printShirtDetails(addidasShirt)
  printShirtDetails(nikeShirt)
}

func printShoeDetails(s IShoe){
  fmt.Printf("Logo: %s", s.getLogo())
  fmt.Println()
  fmt.Printf("Size: %d", s.getSize())
  fmt.Println()
}

func printShirtDetails(s IShirt){
  fmt.Printf("Logo: %s", s.getLogo())
  fmt.Println()
  fmt.Printf("Size: %d", s.getSize())
  fmt.Println() 
}
