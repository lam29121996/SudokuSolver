package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
type Trainer struct {
	Name string
	Age  int
	City string
}

func mustConnectDB(url string) *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
*/

func main() {
	var mySudoku *Sudoku

	fmt.Printf("Please select test example: ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "Easy1":
			mySudoku = EasySudoku1()
		case "Medium1":
			mySudoku = MediumSudoku2()
		case "Hard1":
			mySudoku = HardSudoku1()
		case "Hard2":
			mySudoku = HardSudoku2()
		case "Hard3":
			mySudoku = HardSudoku3()
		case "Hard4":
			mySudoku = HardSudoku4()
		case "Hard5":
			mySudoku = HardSudoku5()
		case "Expert1":
			mySudoku = ExpertSudoku1()
		case "Expert2":
			mySudoku = ExpertSudoku2()
		default:
			fmt.Printf("No such test example. Please type 'Easy1' / 'Medium1' / 'Hard1' / 'Expert1' / etc.\n")
			continue
		}

		break
	}

	fmt.Printf("\nYou have selected %s\n", scanner.Text())

	mySudoku.PrettyPrint()

	/*
		for mySudoku.SolveByNode() == nil {
			mySudoku.PrettyPrint()
		}

		for mySudoku.SolveByRow() == nil {
			mySudoku.PrettyPrint()
		}

		for mySudoku.SolveByColumn() == nil {
			mySudoku.PrettyPrint()
		}

		for mySudoku.SolveByBlock() == nil {
			mySudoku.PrettyPrint()
		}
	*/

	for mySudoku.SolveByNode() == nil || mySudoku.SolveByRow() == nil || mySudoku.SolveByColumn() == nil || mySudoku.SolveByBlock() == nil {
		mySudoku.PrettyPrint()
	}

	mySudoku.OutputAsSVG("1.svg")

	/*
		client := mustConnectDB("mongodb://localhost:27017")
		collection := client.Database("sudokuDB").Collection("sudoku")

		s := Sudoku{}
		fmt.Println(s)

		sudokuSamples := []interface{}{
			Sudoku{},
			Sudoku{},
		}

		insertResult, err := collection.InsertMany(context.TODO(), sudokuSamples)
		if err != nil {
			panic(err)
		}

		fmt.Println(insertResult.InsertedIDs...)
	*/

	/*
		insertResult, err := collection.InsertOne(context.TODO(), bson.M{"tags": "Customer"})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	*/

	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-termCh
}
