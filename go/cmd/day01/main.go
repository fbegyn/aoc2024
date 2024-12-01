package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/fbegyn/aoc2024/go/helpers"
)

func main() {
	file := os.Args[1]
	input := helpers.InputToLines(file)
	left, right := make([]string, len(input)), make([]string, len(input))
	for ind, line := range input {
		spl := strings.Fields(line)
		left[ind] = spl[0]
		right[ind] = spl[1]
	}
	sort.Strings(left)
	sort.Strings(right)
	// Part 1
	sum := 0
	for i := range left {
		dist := helpers.Atoi(left[i]) - helpers.Atoi(right[i])
		if dist < 0 {
			dist *= -1
		}
		sum += dist
	}
	fmt.Printf("The answer for part 1 is: %d\n", sum)
	// Part 2
	sum2 := 0
	tracker := map[int]int{}
	for _, l := range left {
		li := helpers.Atoi(l)
		counter := 0
		if val, ok := tracker[li]; ok {
			sum2 += val
			fmt.Println(val)
			continue
		}
		t := true
		i := 0
		for t && i < len(right){
			r := right[i]
			ri := helpers.Atoi(r)
			switch {
			case li < ri:
				score := li * counter
				tracker[li] = score
				sum2 += score
				t = false
			case li == ri:
				counter++
			}
			i++
		}
	}
	fmt.Printf("The answer for part 2 is: %d\n", sum2)
}
