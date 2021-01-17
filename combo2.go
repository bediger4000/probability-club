package main

import (
	"flag"
	"fmt"
	"os"
	"probability-club/runenumber"
	"sort"
)

func main() {
	first := flag.String("f", "5", "first of 2 rolls to stop")
	second := flag.String("s", "6", "second of 2 rolls to stop")
	n := flag.Int("n", 4, "max number of rolls")

	flag.Parse()
	fmt.Printf("%d rolls, %s followed by %s to win\n", *n, *first, *second)

	if *n < 2 {
		fmt.Fprintf(os.Stderr, "Must have at least 2 rolls\n")
		return
	}

	var d6 = []rune{'1', '2', '3', '4', '5', '6'}

	var num runenumber.Number
	for i := 0; i < *n; i++ {
		num = append(num, runenumber.NewDigit(d6))
	}

	done := false
	var combo []rune

	f := rune((*first)[0])
	s := rune((*second)[0])

	var total int
	wins := make(map[int]int)

	secondToLast := *n - 2

	for !done {
		combo, done = num.Next()
		total++
		rolls := -1
		for i := 0; i <= secondToLast; i++ {
			if combo[i] == f && combo[i+1] == s {
				rolls = i + 2
				break
			}
		}
		if rolls > 0 {
			wins[rolls]++
		}
	}

	fmt.Printf("%d max rolls, %c:%c game ending rolls\n", *n, f, s)
	fmt.Printf("%d total combinations\n", total)
	fmt.Printf("Roll count distribution\n")

	var rolls []int
	for roll, _ := range wins {
		rolls = append(rolls, roll)
	}

	sort.Sort(sort.IntSlice(rolls))

	fg := float64(total)

	for _, roll := range rolls {
		count := wins[roll]
		proportion := float64(count) / fg
		fmt.Printf("%d\t%d\t%f\n", roll, count, proportion)
	}
}
