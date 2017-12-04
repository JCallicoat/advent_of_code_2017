package main

import (
	"bufio"
	"check"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

The spreadsheet consists of rows of apparently-random numbers. To make sure the recovery process is on the right track, they need you to calculate the spreadsheet's checksum. For each row, determine the difference between the largest value and the smallest value; the checksum is the sum of all of these differences.

For example, given the following spreadsheet:

5 1 9 5
7 5 3
2 4 6 8

    The first row's largest and smallest values are 9 and 1, and their difference is 8.
    The second row's largest and smallest values are 7 and 3, and their difference is 4.
    The third row's difference is 6.

In this example, the spreadsheet's checksum would be 8 + 4 + 6 = 18.

*/
func part1(rows []string) int {
	large := 1
	small := 1
	for idx, char := range rows {
		num, _ := strconv.Atoi(char)
		if idx == 0 {
			small = num
		}
		if num > large {
			large = num
		}
		if num < small {
			small = num
		}
	}
	return large - small
}

/*

"Based on what we're seeing, it looks like all the User wanted is some information about the evenly divisible values in the spreadsheet. Unfortunately, none of us are equipped for that kind of calculation - most of us specialize in bitwise operations."

It sounds like the goal is to find the only two numbers in each row where one evenly divides the other - that is, where the result of the division operation is a whole number. They would like you to find those numbers on each line, divide them, and add up each line's result.

For example, given the following spreadsheet:

5 9 2 8
9 4 7 3
3 8 6 5

    In the first row, the only two numbers that evenly divide are 8 and 2; the result of this division is 4.
    In the second row, the two numbers are 9 and 3; the result is 3.
    In the third row, the result is 2.

In this example, the sum of the results would be 4 + 3 + 2 = 9.

*/

func part2(rows []string) int {
	for idx, last := range rows {
		last, err := strconv.Atoi(last)
		check.Check(err)
		for _, next := range rows[idx+1:] {
			next, err := strconv.Atoi(next)
			check.Check(err)
			if last%next == 0 {
				return last / next
			} else if next%last == 0 {
				return next / last
			}
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	defer file.Close()
	check.Check(err)

	part1sum := 0
	part2sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cells := strings.Split(scanner.Text(), "\t")
		part1sum += part1(cells)
		part2sum += part2(cells)
	}
	check.Check(scanner.Err())

	fmt.Printf("%v\n%v\n", part1sum, part2sum)
}
