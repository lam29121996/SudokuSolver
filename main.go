package main

import "fmt"

type Node struct {
	val        int
	isPossible [9]bool
}

type Sudoku [9][9]Node

func NewSudoku() Sudoku {
	mySudoku := Sudoku{}

	myNode := Node{val: -1}
	for i := range myNode.isPossible {
		myNode.isPossible[i] = true
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
	if !s[targetX][targetY].isPossible[value-1] {
		return fmt.Errorf("not possible")
	}

	// 2. Set value
	s[targetX][targetY].val = value

	// 3. Update nodes in same row, column & block
	for i := 0; i < 9; i++ {
		s[i][targetY].isPossible[value-1] = false
		s[targetX][i].isPossible[value-1] = false
	}

	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if targetX/3 == x/3 && targetY/3 == y/3 {
				s[x][y].isPossible[value-1] = false
			}
		}
	}

	fmt.Printf("Set x: %d, y: %d to %d\n", targetX, targetY, value)

	return nil
}

func (s *Sudoku) PrettyPrint() {
	str := ""

	for y := 0; y < 9; y++ {
		if y%3 == 0 {
			str += "|-------|-------|-------|\n"
		}

		for x := 0; x < 9; x++ {
			if x%3 == 0 {
				str += "| "
			}

			if 1 <= s[x][y].val && s[x][y].val <= 9 {
				str += fmt.Sprintf("%d ", s[x][y].val)
			} else {
				str += ". "
			}
		}

		str += "|\n"
	}

	str += "|-------|-------|-------|\n"

	fmt.Println(str)
}

func (s *Sudoku) SolveByNode() error {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			// 1. Skip if the node is filled
			if 1 <= s[x][y].val && s[x][y].val <= 9 {
				continue
			}

			// 2. Count how many possibilities of the node
			count := 0
			lastPossibleValue := -1

			for i := 0; i < 9; i++ {
				if s[x][y].isPossible[i] {
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
				if 1 <= s[x][y].val && s[x][y].val <= 9 {
					continue
				}

				if s[x][y].isPossible[value-1] {
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
				if 1 <= s[x][y].val && s[x][y].val <= 9 {
					continue
				}

				if s[x][y].isPossible[value-1] {
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
						if 1 <= s[x][y].val && s[x][y].val <= 9 {
							continue
						}

						if s[x][y].isPossible[value-1] {
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

func EasySudoku() Sudoku {
	easySudoku := NewSudoku()

	easySudoku.Set(3, 0, 7)
	easySudoku.Set(4, 0, 4)
	easySudoku.Set(8, 0, 6)

	easySudoku.Set(0, 1, 4)
	easySudoku.Set(2, 1, 6)
	easySudoku.Set(3, 1, 8)
	easySudoku.Set(6, 1, 5)
	easySudoku.Set(8, 1, 7)

	easySudoku.Set(0, 2, 7)
	easySudoku.Set(4, 2, 9)
	easySudoku.Set(8, 2, 4)

	easySudoku.Set(1, 3, 3)
	easySudoku.Set(3, 3, 9)
	easySudoku.Set(4, 3, 8)
	easySudoku.Set(5, 3, 4)
	easySudoku.Set(6, 3, 7)

	easySudoku.Set(0, 4, 8)
	easySudoku.Set(1, 4, 2)
	easySudoku.Set(3, 4, 6)
	easySudoku.Set(4, 4, 1)
	easySudoku.Set(5, 4, 3)
	easySudoku.Set(6, 4, 4)
	easySudoku.Set(8, 4, 9)

	easySudoku.Set(1, 5, 4)
	easySudoku.Set(6, 5, 3)

	easySudoku.Set(1, 6, 6)
	easySudoku.Set(2, 6, 2)
	easySudoku.Set(3, 6, 3)
	easySudoku.Set(4, 6, 7)
	easySudoku.Set(8, 6, 5)

	easySudoku.Set(2, 7, 5)
	easySudoku.Set(3, 7, 4)
	easySudoku.Set(5, 7, 9)

	easySudoku.Set(1, 8, 7)
	easySudoku.Set(4, 8, 6)
	easySudoku.Set(5, 8, 1)
	easySudoku.Set(6, 8, 2)
	easySudoku.Set(8, 8, 8)

	return easySudoku
}

func MediumSudoku() Sudoku {
	mediumSudoku := NewSudoku()

	mediumSudoku.Set(1, 0, 1)
	mediumSudoku.Set(4, 0, 7)
	mediumSudoku.Set(5, 0, 2)
	mediumSudoku.Set(8, 0, 4)

	mediumSudoku.Set(1, 1, 9)
	mediumSudoku.Set(2, 1, 6)
	mediumSudoku.Set(4, 1, 3)
	mediumSudoku.Set(6, 1, 2)
	mediumSudoku.Set(7, 1, 5)
	mediumSudoku.Set(8, 1, 7)

	mediumSudoku.Set(2, 2, 8)
	mediumSudoku.Set(5, 2, 9)
	mediumSudoku.Set(8, 2, 3)

	mediumSudoku.Set(3, 3, 4)
	mediumSudoku.Set(4, 3, 1)

	mediumSudoku.Set(3, 4, 7)
	mediumSudoku.Set(7, 4, 2)
	mediumSudoku.Set(8, 4, 9)

	mediumSudoku.Set(0, 5, 7)
	mediumSudoku.Set(4, 5, 2)
	mediumSudoku.Set(6, 5, 5)

	mediumSudoku.Set(0, 6, 9)
	mediumSudoku.Set(1, 6, 6)

	mediumSudoku.Set(3, 7, 3)
	mediumSudoku.Set(6, 7, 7)
	mediumSudoku.Set(7, 7, 4)

	mediumSudoku.Set(2, 8, 7)
	mediumSudoku.Set(4, 8, 8)
	mediumSudoku.Set(5, 8, 4)
	mediumSudoku.Set(6, 8, 1)

	return mediumSudoku
}

func HardSudoku1() Sudoku {
	hardSudoku := NewSudoku()

	hardSudoku.Set(3, 0, 6)

	hardSudoku.Set(5, 1, 7)
	hardSudoku.Set(8, 1, 1)

	hardSudoku.Set(0, 2, 5)
	hardSudoku.Set(1, 2, 1)
	hardSudoku.Set(2, 2, 3)
	hardSudoku.Set(3, 2, 9)
	hardSudoku.Set(6, 2, 6)

	hardSudoku.Set(1, 3, 6)

	hardSudoku.Set(1, 4, 8)
	hardSudoku.Set(5, 4, 5)
	hardSudoku.Set(7, 4, 9)

	hardSudoku.Set(5, 5, 6)
	hardSudoku.Set(6, 5, 2)
	hardSudoku.Set(8, 5, 4)

	hardSudoku.Set(0, 6, 9)
	hardSudoku.Set(4, 6, 7)
	hardSudoku.Set(5, 6, 8)
	hardSudoku.Set(7, 6, 1)

	hardSudoku.Set(0, 7, 1)
	hardSudoku.Set(1, 7, 5)
	hardSudoku.Set(4, 7, 2)

	hardSudoku.Set(0, 8, 3)
	hardSudoku.Set(6, 8, 9)
	hardSudoku.Set(7, 8, 2)

	return hardSudoku
}

func HardSudoku2() Sudoku {
	hardSudoku := NewSudoku()

	hardSudoku.Set(2, 0, 8)
	hardSudoku.Set(5, 0, 3)

	hardSudoku.Set(1, 1, 4)
	hardSudoku.Set(3, 1, 6)
	hardSudoku.Set(4, 1, 8)

	hardSudoku.Set(0, 2, 1)
	hardSudoku.Set(2, 2, 7)
	hardSudoku.Set(3, 2, 2)
	hardSudoku.Set(4, 2, 5)
	hardSudoku.Set(6, 2, 8)

	hardSudoku.Set(0, 3, 9)
	hardSudoku.Set(2, 3, 6)
	hardSudoku.Set(3, 3, 3)

	hardSudoku.Set(0, 4, 8)
	hardSudoku.Set(1, 4, 5)
	hardSudoku.Set(4, 4, 2)
	hardSudoku.Set(7, 4, 3)

	hardSudoku.Set(5, 5, 1)
	hardSudoku.Set(6, 5, 5)

	hardSudoku.Set(0, 6, 7)
	hardSudoku.Set(4, 6, 6)
	hardSudoku.Set(6, 6, 9)

	hardSudoku.Set(3, 7, 4)
	hardSudoku.Set(5, 7, 9)
	hardSudoku.Set(8, 7, 2)

	hardSudoku.Set(0, 8, 2)
	hardSudoku.Set(2, 8, 9)
	hardSudoku.Set(8, 8, 4)

	return hardSudoku
}

func HardSudoku3() Sudoku {
	hardSudoku := NewSudoku()

	hardSudoku.Set(1, 0, 8)
	hardSudoku.Set(2, 0, 5)
	hardSudoku.Set(5, 0, 2)
	hardSudoku.Set(7, 0, 9)

	hardSudoku.Set(7, 1, 2)
	hardSudoku.Set(8, 1, 6)

	hardSudoku.Set(3, 2, 7)

	hardSudoku.Set(3, 3, 4)
	hardSudoku.Set(4, 3, 2)
	hardSudoku.Set(5, 3, 7)

	hardSudoku.Set(0, 4, 3)
	hardSudoku.Set(4, 4, 1)
	hardSudoku.Set(5, 4, 5)

	hardSudoku.Set(1, 5, 9)
	hardSudoku.Set(8, 5, 5)

	hardSudoku.Set(0, 6, 9)
	hardSudoku.Set(1, 6, 2)
	hardSudoku.Set(2, 6, 6)
	hardSudoku.Set(3, 6, 8)
	hardSudoku.Set(6, 6, 3)

	hardSudoku.Set(4, 7, 4)
	hardSudoku.Set(6, 7, 7)

	hardSudoku.Set(0, 8, 4)
	hardSudoku.Set(2, 8, 3)
	hardSudoku.Set(8, 8, 9)

	return hardSudoku
}

func HardSudoku4() Sudoku {
	hardSudoku := NewSudoku()

	hardSudoku.Set(2, 0, 5)
	hardSudoku.Set(4, 0, 3)
	hardSudoku.Set(8, 0, 9)

	hardSudoku.Set(5, 1, 8)
	hardSudoku.Set(7, 1, 6)

	hardSudoku.Set(3, 2, 7)

	hardSudoku.Set(0, 3, 5)
	hardSudoku.Set(4, 3, 2)
	hardSudoku.Set(6, 3, 9)
	hardSudoku.Set(7, 3, 1)

	hardSudoku.Set(0, 4, 3)
	hardSudoku.Set(3, 4, 9)
	hardSudoku.Set(4, 4, 1)
	hardSudoku.Set(6, 4, 6)

	hardSudoku.Set(1, 5, 9)
	hardSudoku.Set(5, 5, 3)
	hardSudoku.Set(6, 5, 2)

	hardSudoku.Set(2, 6, 1)
	hardSudoku.Set(4, 6, 4)
	hardSudoku.Set(5, 6, 9)

	hardSudoku.Set(0, 7, 9)
	hardSudoku.Set(2, 7, 6)
	hardSudoku.Set(3, 7, 8)
	hardSudoku.Set(5, 7, 1)
	hardSudoku.Set(8, 7, 5)

	hardSudoku.Set(1, 8, 7)
	hardSudoku.Set(3, 8, 2)
	hardSudoku.Set(8, 8, 1)

	return hardSudoku
}

func HardSudoku5() Sudoku {
	hardSudoku := NewSudoku()

	hardSudoku.Set(8, 0, 9)

	hardSudoku.Set(5, 1, 7)
	hardSudoku.Set(7, 1, 1)

	hardSudoku.Set(0, 2, 7)
	hardSudoku.Set(1, 2, 6)
	hardSudoku.Set(3, 2, 9)
	hardSudoku.Set(6, 2, 3)
	hardSudoku.Set(8, 2, 8)

	hardSudoku.Set(2, 3, 1)
	hardSudoku.Set(3, 3, 6)
	hardSudoku.Set(6, 3, 4)
	hardSudoku.Set(7, 3, 3)

	hardSudoku.Set(8, 4, 6)

	hardSudoku.Set(1, 5, 5)
	hardSudoku.Set(4, 5, 7)
	hardSudoku.Set(7, 5, 8)

	hardSudoku.Set(2, 6, 3)
	hardSudoku.Set(5, 6, 1)
	hardSudoku.Set(7, 6, 2)

	hardSudoku.Set(0, 7, 9)
	hardSudoku.Set(1, 7, 1)
	hardSudoku.Set(5, 7, 3)

	hardSudoku.Set(5, 8, 5)
	hardSudoku.Set(6, 8, 1)
	hardSudoku.Set(7, 8, 9)

	return hardSudoku
}

func ExpertSudoku() Sudoku {
	expertSudoku := NewSudoku()

	expertSudoku.Set(2, 0, 1)
	expertSudoku.Set(4, 0, 6)
	expertSudoku.Set(6, 0, 9)

	expertSudoku.Set(1, 1, 4)
	expertSudoku.Set(7, 1, 5)

	expertSudoku.Set(0, 2, 8)
	expertSudoku.Set(4, 2, 5)
	expertSudoku.Set(5, 2, 9)
	expertSudoku.Set(8, 2, 3)

	expertSudoku.Set(6, 3, 8)

	expertSudoku.Set(4, 4, 3)
	expertSudoku.Set(6, 4, 7)

	expertSudoku.Set(0, 5, 7)
	expertSudoku.Set(8, 5, 6)

	expertSudoku.Set(0, 6, 9)
	expertSudoku.Set(1, 6, 6)
	expertSudoku.Set(3, 6, 1)

	expertSudoku.Set(7, 7, 2)

	expertSudoku.Set(0, 8, 3)
	expertSudoku.Set(3, 8, 2)
	expertSudoku.Set(5, 8, 8)
	expertSudoku.Set(6, 8, 5)

	return expertSudoku
}

func main() {
	mySudoku := ExpertSudoku()

	mySudoku.PrettyPrint()
	
	/*
		for mySudoku.SolveByNode() == nil {
			mySudoku.FirstPrettyPrint()
		}

		for mySudoku.SolveByRow() == nil {
			mySudoku.FirstPrettyPrint()
		}

		for mySudoku.SolveByColumn() == nil {
			mySudoku.FirstPrettyPrint()
		}

		for mySudoku.SolveByBlock() == nil {
			mySudoku.FirstPrettyPrint()
		}
	*/

	for mySudoku.SolveByNode() == nil || mySudoku.SolveByRow() == nil || mySudoku.SolveByColumn() == nil || mySudoku.SolveByBlock() == nil {
		mySudoku.PrettyPrint()
	}
}
