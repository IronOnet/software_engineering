package main 

import (
	"fmt" 
	"sync" 

)

var lock = &sync.Mutex{} 

type Single struct{

}

var singletonInstance *Single 
func getInstance() *Single{
	if singletonInstance == nil{
		lock.Lock() 
		defer lock.Unlock() 
		if singletonInstance == nil {
			fmt.Println("Creating single instance now.") 
			singletonInstance = &Single{}
		} else{
			fmt.Println("single instance already created")
		}
	} else{
		fmt.Println("singleton instance already creeated")
	}

	return singletonInstance
}