package main

import "fmt"

type ISportsFactory interface{
  makeShoe() IShoe
  makeShirt() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error){
  if brand == "addidas"{
    return &Addias{}
  }

  if brand == "nike"{
    return &Nike{}
  }

  return nil, fmt.Errof("wrong brand type passed")
}
