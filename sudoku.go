package main

import (
	"fmt"
	"os"
)

type Node struct {
	Val        int     `bson:"val,omitempty" json:"val"`
	IsPossible [9]bool `bson:"isPossible,omitempty" json:"isPossible"`
}

type Sudoku [9][9]Node

func NewSudoku() *Sudoku {
	mySudoku := &Sudoku{}

	myNode := Node{Val: -1}
	for i := range myNode.IsPossible {
		myNode.IsPossible[i] = true
	}

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			mySudoku[x][y] = myNode
		}
	}

	return mySudoku
}

func (s *Sudoku) Set(targetX, targetY, value int) error {
	// 1. Check possible
	if !s[targetX][targetY].IsPossible[value-1] {
		return fmt.Errorf("not possible")
	}

	// 2. Set value
	s[targetX][targetY].Val = value

	// 3. Update nodes in same row, column & block
	for i := 0; i < 9; i++ {
		s[i][targetY].IsPossible[value-1] = false
		s[targetX][i].IsPossible[value-1] = false
	}

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if targetX/3 == x/3 && targetY/3 == y/3 {
				s[x][y].IsPossible[value-1] = false
			}
		}
	}

	fmt.Printf("Set x: %d, y: %d to %d\n", targetX, targetY, value)

	return nil
}

func (s *Sudoku) PrettyPrint() {
	str := ""
	red, reset := "\033[31m", "\033[0m"
	separateLine := "|-------|-------|-------|\n"

	for y := 0; y < 9; y++ {
		if y%3 == 0 {
			str += separateLine
		}

		for x := 0; x < 9; x++ {
			if x%3 == 0 {
				str += "| "
			}

			if 1 <= s[x][y].Val && s[x][y].Val <= 9 {
				str += red + fmt.Sprintf("%d ", s[x][y].Val) + reset
			} else {
				str += ". "
			}
		}

		str += "|\n"
	}

	str += separateLine

	fmt.Println(str)
}

func (s *Sudoku) OutputAsSVG(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	startX, startY := 50, 50
	width, height := 900, 900
	nodeSize := 90
	canvas := NewCanvas(file, width, height)

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			x, y := startX+column*nodeSize, startY+row*nodeSize

			canvas.Rect(x, y, nodeSize, nodeSize, "stroke:black; stroke-width:2; fill:white")

			if 1 <= s[row][column].Val && s[row][column].Val <= 9 {
				canvas.Text(x+nodeSize/2, y+nodeSize/2, fmt.Sprintf("%d", s[row][column].Val), "text-anchor:middle; dominant-baseline:central; font-family:Helvetica; font-size:48px; fill:black")
			} else {
				for i := range s[row][column].IsPossible {
					x2, y2 := x+(nodeSize/3)*(i%3), y+(nodeSize/3)*(i/3)

					if s[row][column].IsPossible[i] {
						canvas.Text(x2+nodeSize/6, y2+nodeSize/6, fmt.Sprintf("%d", i+1), "text-anchor:middle; dominant-baseline:central; font-family:Helvetica; font-size:24px; fill:black")
					}
				}
			}
		}
	}

	// Wider squares around 3x3 blocks
	for blockRow := 0; blockRow < 3; blockRow++ {
		for blockColumn := 0; blockColumn < 3; blockColumn++ {
			canvas.Rect(startX+blockColumn*nodeSize*3, startY+blockRow*nodeSize*3, nodeSize*3, nodeSize*3, "stroke:black; stroke-width:5; fill-opacity:0.0")
		}
	}

	canvas.End()

	return nil
}

func (s *Sudoku) SolveByNode() error {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			// 1. Skip if the node is filled
			if 1 <= s[x][y].Val && s[x][y].Val <= 9 {
				continue
			}

			// 2. Count how many possibilities of the node
			count := 0
			lastPossibleValue := -1

			for i := 0; i < 9; i++ {
				if s[x][y].IsPossible[i] {
					count++
					lastPossibleValue = i + 1
				}

				if count > 1 {
					break
				}
			}

			// 3. If only one possible value, set the node
			if count == 1 {
				fmt.Println("Solved by node!")
				s.Set(x, y, lastPossibleValue)
				return nil
			}
		}
	}

	return fmt.Errorf("no solution")
}

func (s *Sudoku) SolveByRow() error {
	for value := 1; value <= 9; value++ {
		for y := 0; y < 9; y++ {
			count := 0
			lastX := -1

			for x := 0; x < 9; x++ {
				// Skip if the node is filled
				if 1 <= s[x][y].Val && s[x][y].Val <= 9 {
					continue
				}

				if s[x][y].IsPossible[value-1] {
					count++
					lastX = x
				}

				if count > 1 {
					break
				}
			}

			if count == 1 {
				fmt.Println("Solved by row!")
				s.Set(lastX, y, value)
				return nil
			}
		}
	}

	return fmt.Errorf("no solution")
}

func (s *Sudoku) SolveByColumn() error {
	for value := 1; value <= 9; value++ {
		for x := 0; x < 9; x++ {
			count := 0
			lastY := -1

			for y := 0; y < 9; y++ {
				// Skip if the node is filled
				if 1 <= s[x][y].Val && s[x][y].Val <= 9 {
					continue
				}

				if s[x][y].IsPossible[value-1] {
					count++
					lastY = y
				}

				if count > 1 {
					break
				}
			}

			if count == 1 {
				fmt.Println("Solved by column!")
				s.Set(x, lastY, value)
				return nil
			}
		}
	}

	return fmt.Errorf("no solution")
}

func (s *Sudoku) SolveByBlock() error {
	for value := 1; value <= 9; value++ {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				count := 0
				lastX, lastY := -1, -1

				for x := i * 3; x < i*3+3; x++ {
					for y := j * 3; y < j*3+3; y++ {
						// Skip if the node is filled
						if 1 <= s[x][y].Val && s[x][y].Val <= 9 {
							continue
						}

						if s[x][y].IsPossible[value-1] {
							count++
							lastX, lastY = x, y
						}

						if count > 1 {
							break
						}
					}

					if count > 1 {
						break
					}
				}

				if count == 1 {
					fmt.Println("Solved by block!")
					s.Set(lastX, lastY, value)
					return nil
				}
			}
		}
	}

	return fmt.Errorf("no solution")
}

func (s *Sudoku) SolveByUncertainNodes() error {

	return nil
}
