package main 

import (
	"fmt" 
	"github.com/irononet/packages/npcs"
)

func main(){
	mob := npcs.NonPlayerCharacter{Name: "Alfred"} 
	fmt.Println(mob)

	hortense := npc.NonPlayerCharacter{Name: "Hortense", 
									Loc: npcs.Location{X: 10.0, Y: 10.0, Z: 10.0}} 
	fmt.Println(hortense)

	fmt.Println("Alfred is %s units of distance from Hortense.\n", 
	mob.DistanceTo(hortense))

}

