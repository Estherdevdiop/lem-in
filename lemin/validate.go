package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	Start    string
	End      string
	Rooms    []string
	Link     []string
	NbAnt    int
)

func ReadAndValidate(file string) (result bool) {
	var emptyln []string
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("erreur lors de lecture du fichier")
		
	}
	
	line := strings.Split(string(data), "\r\n")

	for _, v := range line {
		if v != "" {
			if !strings.HasPrefix(v, "#") || (v == "##start" || v == "##end") {
				emptyln = append(emptyln, v)
			}
		}
	}
	NbAnt, err = strconv.Atoi(emptyln[0])
	if err != nil {
		fmt.Println(err)
	}

	if NbAnt <= 0 {
		fmt.Println("Error no valid ant found")
		result = false
		//return
		os.Exit(1)
	}

	p := ""
	for i, v := range emptyln[1:] {
		parts := strings.Fields(v)
		switch v {
		case "##start":
			if p == "" {
				p = "##start"
				Start = emptyln[i+2]
				// fmt.Println(Start)
				//continue
			}

		case "##end":
			if p == "##start" {
				p = "##end"
				End = emptyln[i+2]
				//continue

			} else {
				fmt.Println("no inverse commande")
				os.Exit(1)
			}

		}
		if len(parts) > 3 || len(parts) == 2 {
			fmt.Println("no valide rooms", v)

		}
		if len(parts) == 3 {
			Rooms = append(Rooms, v)

		} else {
			parts = strings.Split(v, "-")
			if len(parts) == 2 {
				Link = append(Link, v)
			}
		}
	}

	return
}

func ValidRoms(Roms []string) bool {
	for _, r := range Roms {
		if strings.HasPrefix(r, "L") || strings.Contains(r, "-") {
			fmt.Println("error")
		}
		parts := strings.Split(r, " ")
		last := parts[1:]

		_, err1 := strconv.Atoi(last[0])
		_, err2 := strconv.Atoi(last[1])
		if err1 != nil || err2 != nil {
			fmt.Println("error")
		}
	}
	return true
}
