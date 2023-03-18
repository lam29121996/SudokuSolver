package main

func EasySudoku1() *Sudoku {
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

func MediumSudoku2() *Sudoku {
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

func HardSudoku1() *Sudoku {
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

func HardSudoku2() *Sudoku {
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

func HardSudoku3() *Sudoku {
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

func HardSudoku4() *Sudoku {
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

func HardSudoku5() *Sudoku {
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

func ExpertSudoku1() *Sudoku {
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

func ExpertSudoku2() *Sudoku {
	expertSudoku := NewSudoku()

	expertSudoku.Set(0, 0, 7)

	expertSudoku.Set(0, 1, 2)
	expertSudoku.Set(1, 1, 4)
	expertSudoku.Set(2, 1, 9)
	expertSudoku.Set(6, 1, 8)

	expertSudoku.Set(1, 2, 5)
	expertSudoku.Set(6, 2, 3)

	expertSudoku.Set(0, 3, 3)
	expertSudoku.Set(4, 3, 5)

	expertSudoku.Set(4, 4, 9)
	expertSudoku.Set(5, 4, 6)

	expertSudoku.Set(2, 5, 6)
	expertSudoku.Set(4, 5, 8)
	expertSudoku.Set(5, 5, 4)
	expertSudoku.Set(6, 5, 7)

	expertSudoku.Set(3, 6, 6)
	expertSudoku.Set(4, 6, 7)
	expertSudoku.Set(6, 6, 5)

	expertSudoku.Set(7, 7, 4)
	expertSudoku.Set(8, 7, 9)

	expertSudoku.Set(4, 8, 4)
	expertSudoku.Set(5, 8, 9)
	expertSudoku.Set(7, 8, 2)

	return expertSudoku
}
