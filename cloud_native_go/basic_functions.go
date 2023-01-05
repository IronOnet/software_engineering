package main 

import (
	"math/rand"
	"time"
)

func dieRoll(size int) int{
	rand.Seed(time.Now().UnixNano())
	return rand.Int(size) + 1
}


type dieRollFunc func(int) int


func fakeDieRoll(size int) int{
	return 42
}

func getDieRolls() []dieRollFunc{
	return []dieRollFunc{
		fakeDieRoll, 
		dieRoll
	}
}