package main

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			isFirstRow := row == 1
			isLastRow := row == y
			isFirstCol := col == 1
			isLastCol := col == x

			isTopLeft := isFirstRow && isFirstCol
			isTopRight := isFirstRow && isLastCol
			isBottomLeft := isLastRow && isFirstCol
			isBottomRight := isLastRow && isLastCol

			isTopEdge := isFirstRow && col > 1 && col < x
			isBottomEdge := isLastRow && col > 1 && col < x
			isLeftEdge := isFirstCol && row > 1 && row < y
			isRightEdge := isLastCol && row > 1 && row < y

			if isTopLeft {
				z01.PrintRune('A')
			} else if isTopRight {
				z01.PrintRune('C')
			} else if isBottomLeft {
				z01.PrintRune('C')
			} else if isBottomRight {
				z01.PrintRune('A')
			} else if isTopEdge || isBottomEdge {
				z01.PrintRune('B')
			} else if isLeftEdge || isRightEdge {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadE(5, 3)
	z01.PrintRune('\n')
	QuadE(5, 1)
	z01.PrintRune('\n')
	QuadE(1, 1)
	z01.PrintRune('\n')
	QuadE(1, 5)
}
