package main

import (
	"fmt"
	"os"
)

type Ghost struct {
	currentBand Band
	amountPrev  int
}

type Band struct {
	members []int
	number  int
}

var ghosts []Ghost
var bands []Band

func createNewBand(x, y int) Band {
	if ghosts[x].currentBand.number == 0 && ghosts[y].currentBand.number == 0 {
		var newBand Band
		newBand.number = len(bands)
		newBand.members = append(newBand.members, x, y)
		bands = append(bands, newBand)

		ghosts[x].currentBand = newBand
		ghosts[x].amountPrev++
		ghosts[y].currentBand = newBand
		ghosts[y].amountPrev++

		return newBand
	}
	if ghosts[x].currentBand.number == 0 {
		var newBand Band
		newBand.number = len(bands)
		newBand.members = append(ghosts[y].currentBand.members, x)
		bands = append(bands, newBand)

		ghosts[x].currentBand = newBand
		ghosts[x].amountPrev++
		oldBand := bands[ghosts[y].currentBand.number]
		for _, member := range oldBand.members {
			ghosts[member].currentBand = newBand
			ghosts[member].amountPrev++
		}
		return newBand
	}
	if ghosts[y].currentBand.number == 0 {
		var newBand Band
		newBand.number = len(bands)
		newBand.members = append(ghosts[x].currentBand.members, y)
		bands = append(bands, newBand)

		ghosts[y].currentBand = newBand
		ghosts[y].amountPrev++
		oldBand := bands[ghosts[x].currentBand.number]
		for _, member := range oldBand.members {
			ghosts[member].currentBand = newBand
			ghosts[member].amountPrev++
		}
		return newBand
	}
	if ghosts[x].currentBand.number != 0 && ghosts[y].currentBand.number != 0 {
		var newBand Band
		newBand.number = len(bands)
		newBand.members = append(ghosts[x].currentBand.members, ghosts[y].currentBand.members...)
		bands = append(bands, newBand)

		oldBandX := bands[ghosts[x].currentBand.number]
		for _, member := range oldBandX.members {
			ghosts[member].currentBand = newBand
			ghosts[member].amountPrev++
		}
		oldBandY := bands[ghosts[y].currentBand.number]
		for _, member := range oldBandY.members {
			ghosts[member].currentBand = newBand
			ghosts[member].amountPrev++
		}
		return newBand
	}
	return Band{}
}

func checkBands(x, y int) bool {
	return ghosts[x].currentBand.number == ghosts[y].currentBand.number && ghosts[y].currentBand.number != 0
}

func main() {
	var n, m int

	fmt.Fscan(os.Stdin, &n)
	fmt.Fscan(os.Stdin, &m)

	bands = append(bands, Band{number: 0})

	for i := 0; i < n; i++ {
		ghosts = append(ghosts, Ghost{currentBand: bands[0], amountPrev: 1})
	}
	var questionNum int

	for i := 0; i < m; i++ {
		fmt.Fscan(os.Stdin, &questionNum)
		if questionNum == 1 {
			var x, y int
			fmt.Fscan(os.Stdin, &x)
			fmt.Fscan(os.Stdin, &y)
			createNewBand(x, y)
		}
		if questionNum == 2 {
			var x, y int
			fmt.Fscan(os.Stdin, &x)
			fmt.Fscan(os.Stdin, &y)
			if checkBands(x, y) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
		if questionNum == 3 {
			var x int
			fmt.Println(ghosts[x].amountPrev)
		}

	}

}
