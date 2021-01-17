package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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

	fmt.Printf("Playing %d games, until %d followed by %d\n", games, first, second)

	rand.Seed(time.Now().UnixNano() + int64(os.Getpid()))

	pipCount := make(map[int]int)
	for pips := 1; pips <= 6; pips++ {
		pipCount[pips] = 0
	}

	var totalRolls float64

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

		totalRolls += float64(rolls)
	}

	fmt.Printf("Mean fee %.02f\n", totalRolls/float64(games))

	rollCount := 0
	fmt.Printf("Check pip count distribution\n")
	for pips := 1; pips <= 6; pips++ {
		fmt.Printf("roll a %d\t%d times\n", pips, pipCount[pips])
		rollCount += pipCount[pips]
	}
	fmt.Printf("%d total rolls\n", rollCount)

}
