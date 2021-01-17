package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("%s <games> <first> <second>\n", os.Args[0])
	}
	games, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	first, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	second, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("# Playing %d games, until %d followed by %d\n", games, first, second)

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	var gameSum float64

	distribution := make(map[int]int)
	pipCount := make(map[int]int)
	for pips := 1; pips <= 6; pips++ {
		pipCount[pips] = 0
	}

	for i := 0; i < games; i++ {
		// new game
		rolls := 0
		previous := 0

		for {
			n := rand.Intn(6) + 1
			pipCount[n]++
			rolls++
			if n == second && previous == first {
				break
			}
			previous = n
		}

		distribution[rolls]++
		gameSum += float64(rolls)
	}

	fmt.Printf("# Mean fee %.02f\n", gameSum/float64(games))

	for pips := 1; pips <= 6; pips++ {
		fmt.Printf("# roll %d\t%d\n", pips, pipCount[pips])
	}

	var rolls []int
	for roll, _ := range distribution {
		rolls = append(rolls, roll)
	}

	sort.Sort(sort.IntSlice(rolls))

	fg := float64(games)

	for _, roll := range rolls {
		count := distribution[roll]
		proportion := float64(count) / fg
		fmt.Printf("%d\t%d\t%f\n", roll, count, proportion)
	}
}
