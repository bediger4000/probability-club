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

	var totalRolls float64

	for i := 0; i < games; i++ {
		// new game
		rolls := 0
		previous := 0

		for {
			n := rand.Intn(6) + 1
			rolls++
			if n == second && previous == first {
				break
			}
			previous = n
		}

		totalRolls += float64(rolls)
	}

	fmt.Printf("Mean fee %.02f\n", totalRolls/float64(games))
}
