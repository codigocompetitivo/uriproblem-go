package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

func main() {
	var totalFarm int
	fmt.Scanln(&totalFarm)
	sheepsFarmSlice := make([]int64, totalFarm)
	farmVisited := make([]int, totalFarm)
	var totalSheepNotStolen int64
	for i := 0; i < totalFarm; i++ {
		scanf("%d", &sheepsFarmSlice[i])
		totalSheepNotStolen += sheepsFarmSlice[i]
	}

	for i := 0; i >= 0 && i < totalFarm; {
		lado := sheepsFarmSlice[i] % 2
		if sheepsFarmSlice[i] > 0 {
			sheepsFarmSlice[i] -= 1
			farmVisited[i] = 1
			totalSheepNotStolen -= 1
		}

		if lado == 0 {
			i -= 1
		} else {
			i += 1
		}
	}

	totalFarmVisited := 0
	for _, visited := range farmVisited {
		totalFarmVisited += visited
	}

	fmt.Printf("%d %d\n", totalFarmVisited, totalSheepNotStolen)
}
