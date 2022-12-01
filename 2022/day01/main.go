package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/jim-at-jibba/advent-of-code/cast"
	"github.com/jim-at-jibba/advent-of-code/util"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		util.CopyToClipboard(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	parsed := parseInput(input)
	return parsed[0]
}

func part2(input string) int {
	var total int
	parsed := parseInput(input)
	sliced := parsed[0:3]

	for _, v := range sliced {
		total += v
	}
	return total
}

func parseInput(input string) (ans []int) {
	var tempAnswer int
	for _, line := range strings.Split(input, "\n") {
		var trimmed = strings.TrimSpace(line)
		if len(trimmed) == 0 {
			ans = append(ans, tempAnswer)
			tempAnswer = 0
		} else {
			tempAnswer += cast.ToInt(trimmed)
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	return ans
}
