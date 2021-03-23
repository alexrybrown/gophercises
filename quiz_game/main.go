package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvFlag := flag.String(
		"csv",
		"problems.csv",
		"Provide path to csv with quiz questions and answers. Defaults to problems.csv.")
	tFlag := flag.Int(
		"t",
		30,
		"Provide amount of time for the quiz. Defaults to 30 seconds.")

	flag.Parse()

	file, err := os.Open(*csvFlag)
	if err != nil {
		log.Fatal("Bad file name: ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Problems reading file: ", err)
	}

	correctAns := 0
	totalQ := 0
	for _, v := range lines {
		q := v[0]
		a := v[1]
		fmt.Println(q)

		buf := bufio.NewReader(os.Stdin)
		ans, err := buf.ReadBytes('\n')
		if err != nil {
			log.Fatal("Problem reading input: ", err)
		}
		if strings.TrimSpace(string(ans)) == a {
			correctAns++
		}
		totalQ++
	}

	fmt.Println(fmt.Sprintf("%v/%v correct.", correctAns, totalQ))
}
