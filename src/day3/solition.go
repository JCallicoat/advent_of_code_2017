package main

import (
	"fmt"
	"math"
)

/*

You come across an experimental new kind of memory stored on an infinite two-dimensional grid.

Each square on the grid is allocated in a spiral pattern starting at a location marked 1 and then counting up while spiraling outward. For example, the first few squares are allocated like this:

17  16  15  14  13
18   5   4   3  12
19   6   1   2  11
20   7   8   9  10
21  22  23---> ...

While this is very space-efficient (no squares are skipped), requested data must be carried back to square 1 (the location of the only access port for this memory system) by programs that can only move up, down, left, or right. They always take the shortest path: the Manhattan Distance between the location of the data and square 1.

For example:

    Data from square 1 is carried 0 steps, since it's at the access port.
    Data from square 12 is carried 3 steps, such as: down, left, left.
    Data from square 23 is carried only 2 steps: up twice.
    Data from square 1024 must be carried 31 steps.

How many steps are required to carry the data from the square identified in your puzzle input all the way to the access port?

Your puzzle input is 325489.

*/

func part1(input int) int {
	steps := 0    // how many squares have been created from the center
	iter := 1     // how long a given side of the square is
	current := 0  // the last number of the current square
	previous := 0 // the last number of the previous square

	// calculate how many squares we need before we find `input` in the square
	for current < input {
		iter += 2
		previous = current
		current = iter * iter
		steps++
	}

	// number of items in a side of the square that contains `input`
	count := (current - previous) / 4

	// dist is calculated by finding the distance to the center of
	// the side (count/2) of the square containing `input` from
	// the position where `input` is located
	dist := 0
	if input <= previous+count { // `input` in right side of square
		dist = count/2 - ((previous + count) - input)
	} else if input <= previous+(count*2) { // `input` in top side of square
		dist = count/2 - ((previous + (count * 2)) - input)
	} else if input <= previous+(count*3) { // `input` in left side of square
		dist = count/2 - ((previous + (count * 3)) - input)
	} else if input <= previous+(count*4) { // `input` in bottom side of square
		dist = count/2 - ((previous + (count * 4)) - input)
	}

	// the manhatten distance is the sum of the absolute value of dist
	// and the number of squares created from the center
	return steps + int(math.Abs(float64(dist)))
}

func part2(input int) int {
	return 0
}

func main() {
	input := 325489
	fmt.Printf("%v\n", part1(input))
	fmt.Printf("%v\n", part2(input))
}
