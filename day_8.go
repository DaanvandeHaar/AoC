package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main(){
	grid, err := loadTreeGrid()
	if err != nil {
		return
	}

	count, err := visibleTrees(grid)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(count)

	highest, coords, err := scenicScore(grid)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(highest, coords)


}

func loadTreeGrid() ([][]int, error) {

	f, err := os.ReadFile("trees.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(string(f)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	treeGrid := make([][]int, len(lines))

	for i, line := range lines {
		treeGrid[i] = make([]int, len(line))
		for j, char := range line {
			treeGrid[i][j] = int(char)
		}
	}
	return treeGrid, nil
}

func visibleTrees(trees [][]int) (int, error)  {

	visible := make([][]int, len(trees))
	for i := range visible {
		visible[i] = make([]int, len(trees[0]))
	}

	// rows
	for i := 0; i < len(trees); i++ {
		highest := -1
		for j := 0; j < len(trees[i]); j++ {
			if trees[i][j] > highest {
				visible[i][j] = 1
				highest = trees[i][j]
				continue
			}
		}
		highest = -1
		for j := len(trees) - 1; j > 0 ; j-- {
			if trees[i][j] > highest {
				visible[i][j] = 1
				highest = trees[i][j]
				continue
			}
		}
	}

	// cols
	for j := 0; j < len(trees); j++ {
		highest := -1
		for i := 0; i < len(trees[j]); i++ {
			if trees[i][j] > highest {
				visible[i][j] = 1
				highest = trees[i][j]
				continue
			}
		}
		highest = -1
		for i := len(trees) - 1; i > 0 ; i-- {
			if trees[i][j] > highest {
				visible[i][j] = 1
				highest = trees[i][j]
				continue
			}
		}
	}

	count := 0
	for i := 0; i < len(visible) ; i++ {
		for j := 0; j < len(visible[i]); j++ {
			if visible[i][j] == 1 {count++}
		}
	}
	return count, nil
}

func scenicScore(trees [][]int) (int, string, error) {
	var highest int
	var highestCoords string

	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {
			var coords = trees[i][j]
			score, count := 1, 0

			for r := j + 1; r < len(trees[0]); r++ {
				count++
				if trees[i][r] >= coords {
					break
				}
			}
			score *= count
			count = 0

			for l := j - 1; l >= 0; l-- {
				count++
				if trees[i][l] >= coords {
					break
				}
			}
			score *= count
			count = 0

			for d := i + 1; d < len(trees); d++ {
				count++
				if trees[d][j] >= coords {
					break
				}
			}
			score *= count
			count = 0

			for u := i - 1; u >= 0; u-- {
				count++
				if trees[u][j] >= coords {
					break
				}
			}
			score *= count


			if score > highest {
				highest, highestCoords = score, fmt.Sprintf("[%v],[%v]", i, j)

			}
		}
	}
	return highest, highestCoords, nil
}
