package main

// It's impossible to implement the classic factory Method pattern in Go
// due to lack of OOP features such as classes and inheritance. However
// we can still implement basic version of the pattenr, the Simple Factory

// In this example we're going to build various types of weapons using
// a factory struct.

// First, we create the IGun interface, which defines all methods a gun should
// have. There is a gun struct type that implements the IGun interface.
// Two concrete guns ak47 and musket, both embed gun struct and indirectly
// implement all iGun methods.

// The gunFactory struct serves as a factory, which creates

import "fmt "

func main(){
  ak47, _ := getGun("ak47")
  musket, _ := getGun("musket")

  printDetails(ak47)
  printDetails(musket)
}

func printDetails(g IGun){
  fmt.Printf("Gun: %s", g.getName())
  fmt.Println()
  fmt.Printf("Power: %d", g.getPower())
  fmt.Println()
}
