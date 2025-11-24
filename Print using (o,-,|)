package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
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

			if isTopLeft || isTopRight || isBottomLeft || isBottomRight {
				z01.PrintRune('o')
			} else if isTopEdge || isBottomEdge {
				z01.PrintRune('-')
			} else if isLeftEdge || isRightEdge {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadA(5, 3)
	z01.PrintRune('\n')
	QuadA(5, 1)
	z01.PrintRune('\n')
	QuadA(1, 1)
	z01.PrintRune('\n')
	QuadA(1, 5)
}
