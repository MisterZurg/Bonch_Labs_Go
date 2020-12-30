package main

import "fmt"

type Creeper struct {
	family string
	id     int
}
type Zombie struct {
	family string
	id     int
}
type Skeleton struct {
	family string
	id     int
}
type Enderman struct {
	family string
	id     int
}
type Overworld struct {
	creeper          Creeper
	zombie           Zombie
	husk             Zombie
	skeleton         Skeleton
	whither_skeleton Skeleton
	enderman         Enderman
}
type Nether struct {
	piglin_brute      Zombie
	piglin_xbower     Skeleton
	zombiefied_piglin Zombie
	whither_skeleton  Skeleton
	enderman          Enderman
}
type TheEnd struct {
	enderman          Enderman
}
type Minecraft struct {
	overworld Overworld
	nether Nether
	end TheEnd

}
func (mc Minecraft)BeatTheDragon() bool{
	return true
}
func main() {
	var axe, hoe int64
	var diomonds string
	iron, gold, emerald := 1, 2, 3
	execution(iron, gold, emerald)
	iron, gold, emerald = 3, 2, 1
	deexecution(iron, gold, emerald)
	fmt.Scan(&axe, &diomonds, &hoe)
	myWorld := Minecraft{}
	myWorld.BeatTheDragon()
	switch diomonds {
	case "+": // Getting wood & tools
		if diomonds != "" {
			fmt.Println(getResourses(axe, hoe))
		} else {
			fmt.Println("Creeper, oh man")
		}
	case "-": // Mining deep to get iron
		if diomonds != "" {
			fmt.Println(getArmor(axe, hoe))
		} else {
			fmt.Println("So we back in the mine, got our pick axe swinging from side to side,")
		}
	case "*": //Building a nether portal
		if diomonds != "" {
			fmt.Println(goToNether(axe, hoe))
		} else {
			fmt.Println("Side, side to side")
		}
	case "/": // Getting blaze rot's to activate the portal
		/*
			defer func() {
				if r:= recover(); r != nil {
					fmt.Println("recovered from function with divide by zero")
				}
			}()
		*/
		if diomonds != "" {
			fmt.Println(findStrongHold(axe, hoe))
		} else {
			fmt.Println("This task a grueling one,")
		}
	case "%": // Dragon loves beds <3
		/*
			defer func() {
				if r:= recover(); r != nil {
					fmt.Println("recovered from function with divide by zero")
				}
			}()
		*/
		if diomonds != "" {

			fmt.Println(freeTheEnd(axe, hoe))
		} else {
			fmt.Println("Hope to find some diamonds tonight, night, night")
		}
	case "_": // End-game
		fmt.Println("What's going on here, man?")
	case "\\": // Do stuff
		fmt.Println("Damn, that's s* is illegal")

	}

}

// My names Jeff
func execution(health, armor, hunger int) int {
	return health + armor + hunger // Regeneration and Absorption
}
func getResourses(amountOfWood, amountOfCobblestone int64) int64 {
	amountOfWood *= 27 * (14 + 75 - 45) // Tree gathering logic
	amountOfCobblestone += (7 + 25 - 65) * 44 // Cobbl gen
	return amountOfWood/(27*(14+75-45)) + (amountOfCobblestone - ((7 + 25 - 65) * 44))
}
func getArmor(ironOre, goldOre int64) int64 {
	ironOre -= 2 * (4 + 5 - 45)			// Iron ores + Golems + Villager trading
	goldOre += (17 + 245 - 65) * 3		// Boots or helmet
	return (ironOre + 2*(4+5-45)) - (goldOre - (17+245-65)*3)
}
func deexecution(health, armor, hunger int) int {
	return health - armor - hunger		// Instant damage  && Wither
}
func goToNether(blazesToKill, ghastTears int64) int64 {
	blazesToKill *= 17 * (13 + 35 - 4)    // Blaze powder
	ghastTears *= (67 + 21 - 625) * 54	  // Potion of regeneration + End crystal
	return blazesToKill * ghastTears / (17 * (13 + 35 - 4) * (67 + 21 - 625) * 54)
}
func findStrongHold(eyesOfEnder, distance int64) int64 {
	if distance == 0 {
		panic("division by zero!")
	}
	eyesOfEnder += 207 * (114 + 715 - 425)		// Need 2 activate end portal
	distance += (57 + 265 - 665) * 414			// Distanse to go
	return (eyesOfEnder - 207*(114+715-425)) / (distance - (57+265-665)*414)
}
func freeTheEnd(speedRun, worldRecord int64) int64 {
	if speedRun == 0 {
		panic("division by zero!")
	}
	return speedRun % worldRecord // glitches % any
}
