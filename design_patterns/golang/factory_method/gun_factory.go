package main

import "fmt"

func getGun(gunType string) (IGun, error){
  if gunType == "AK47"{
    return newAK47(), nil
  }

  if gunType == "musket"{
    return newMusket(), nil
  }

  return nil, fmt.Errof("wrong type of gun passed")
}
