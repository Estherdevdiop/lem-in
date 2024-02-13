package lemin

import "strings"

var (
	AllPath []string
	AllRooms []*Room
	
	Paths    [][]string

)

func DFS(path []string, CurrentRoom *Room) {
	path = append(path, CurrentRoom.name)

	if CurrentRoom.name == strings.Split(End, " ")[0] {
		AllPath = append(AllPath, path...)
	}

	CurrentRoom.visit = true
	for _, voisin := range CurrentRoom.adjacent {
		v := GetRoom(voisin)
		if !v.visit {
			DFS(path, v)
		}
	}
	CurrentRoom.visit = false
}

func FindAdj(room_name string, link []string) []string {
	var adjacent []string
	for _, room := range link {
		Voisins := strings.Split(room, "-")
		if room_name == Voisins[0] {
			adjacent = append(adjacent, Voisins[1])

		} else if room_name == Voisins[1] {
			adjacent = append(adjacent, Voisins[0])
		}
	}
	return adjacent
}

func InitRoom() {
	for _, values := range Rooms {
		r := NewRoom(strings.Fields(values)[0])
		r.adjacent = FindAdj(r.name, Link)
		AllRooms = append(AllRooms, r)
	}
}

func GetRoom(name string) *Room {
	for _, v := range AllRooms {
		if v.name == name {
			return v
		}
	}
	return nil
}

func TriePath() {
	start := strings.Split(Start, " ")[0]
	DFS([]string{}, GetRoom(start))
	
	var path []string
	for _, room := range AllPath {
		path = append(path, room)
		if room == strings.Fields(End)[0] {
			Paths = append(Paths, path)
			path = []string{}
		}
	}
}
