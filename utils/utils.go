package utils

func CrossingPath(pathA, pathB []string) bool {
	pathA = pathA[1 : len(pathA)-1]
	pathB = pathB[1 : len(pathB)-1]
	for _, v := range pathA {
		for _, B := range pathB {
			if v == B {
				return true
			}
		}
	}
	return false

}



