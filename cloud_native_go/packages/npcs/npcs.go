package npcs

import (
	"fmt" 
	"math"
)


func (loc Location) String() string{
	return fmt.Sprintf("(%f, %f, %f)", loc.X, loc.Y, loc.Z)
}

// EuclidianDistance returns the distance between two in-game 
// locations
func (loc Location) EuclidianDistance(target Location) float64{
	return Math.sqrt(
		(loc.X-target.X) * (loc.X-target.X) + 
		(loc.Y-target.Y) * (loc.Y-target.Y) + 
		(loc.Z-target.Z) * (loc.Z-target.Z)
	)
}


// DistanceTo returns the distance between two in-game characters 
func (npc NonPlayerCharacter) DistanceTo(target NonPlayerCharacter) float64{
	return npc.Loc.EuclidianDistance(target.Loc)
}

func (npc NonPlayerCharacter) String() string{
	return fmt.Sprintf("%s %s", npc.Name, npc.Loc)
}


