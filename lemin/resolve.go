package lemin

import (
	"fmt"
	"lem-in/utils"
	"strings"
)

var (
	Parcours []*Parcour
	AllAnts  []*Ant
	result   [][]string
)

func NotCrossing(firtroom, tabroute []string) int {
	count := 0
	for _, R := range firtroom {
		if R != tabroute[1] {
			for _, B := range Paths {
				if B[1] == R {
					if !utils.CrossingPath(tabroute, B) {
						count++
						break
					}
				}
			}

		}
	}
	return count
}

func PathsToUse() {
	TriePath()
	SortPath()
	start := strings.Split(Start, " ")[0]
	firstroom := GetRoom(start).adjacent
	for _, r := range firstroom {
		for _, route := range Paths {
			if route[1] == r {
				if NotCrossing(firstroom, route) == len(firstroom)-1 {
					if len(result) == 0 {
						result = append(result, route)
						break
					} else {
						trace := true
						for _, v := range result {
							if utils.CrossingPath(route, v) {
								trace = false
							}
						}
						if trace {
							result = append(result, route)
							break
						}
					}
				}
			}
		}
	}
	if len(result) == 0 {
		result = append(result, Paths[0])
	}
}

func SortPath() {
	for a := 0; a < len(Paths); a++ {
		for b := a + 1; b < len(Paths); b++ {
			if len(Paths[a]) > len(Paths[b]) {
				Paths[a], Paths[b] = Paths[b], Paths[a]
			}
		}
	}
}

func SortRslt() {
	for a := 0; a < len(result); a++ {
		for b := a + 1; b < len(result); b++ {
			if len(result[a]) > len(result[b]) {
				result[a], result[b] = result[b], result[a]
			}
		}
	}
}

func CountStep(Parcour *Parcour) int {
	return len(Parcour.Route) + len(Parcour.Ants)
}

func SetRoute(ant *Ant) {
	var ShortestStep, ShortestPathIndex int
	for i, p := range Parcours {
		Step := CountStep(p)
		if ShortestStep == 0 || Step < ShortestStep {
			ShortestPathIndex = i
			ShortestStep = Step
		}
	}

	//fmt.Println(AllAnts)
	//fmt.Println(Parcours)
	//for _,ant := range AllAnts {
		Parcours[ShortestPathIndex].Ants = append(Parcours[ShortestPathIndex].Ants, ant)
		ant.Route = Parcours[ShortestPathIndex].Route
	//}
}

func InitMove() {
	for _, p := range Parcours {
		p.Ants[0].Movable = true
	}
}

func Finish() bool {
	for _, ant := range AllAnts {
		if !ant.Arrive {
			return false
		}

	}
	return true
}

func LetsMove() {
	PathsToUse()
	InitAnt()
	SortRslt()
	InitParcours()
	for _, ant := range AllAnts {
		SetRoute(ant)
	}
	InitMove()
	for !Finish() {
		for _, p := range Parcours {
			for i, ant := range p.Ants {
				if ant.Position == len(ant.Route) {
					ant.Arrive = true
					continue
				}
				//if ant.Position != len(ant.Route) {
				if ant.Movable {
					//fmt.Println("route: ", ant.Route)
					fmt.Print("L", ant.Num, "-", ant.Route[ant.Position], " ")
					ant.Position++

				} else if p.SetMove == i {
					ant.Movable = true

				}
			}
		}
		//break
		fmt.Println()
		for _, p := range Parcours {
			p.SetMove++
		}
	}
}
