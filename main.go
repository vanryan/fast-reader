package main

import (
	"fast-reader/reader"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name as a commandline parameter")
		return
	}
	fileName := os.Args[1]

	wpm := reader.DEFAULT_WPM
	if len(os.Args) > 2 {
		arg2, err := strconv.Atoi(os.Args[2])
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
		if arg2 > 100 && arg2 < 1000 {
			// wpm is given
			wpm = arg2
		}
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")
	wordDelims := reader.WORD_DELIMS
	wordDelimiterRegex := regexp.MustCompile(fmt.Sprintf("[%s]+", wordDelims))

	sleepMicroBwEachWordSet := time.Duration(reader.GetSleepTimeMicroByWpm(wpm)) * time.Microsecond

	fmt.Print("\033[H\033[2J")
	for _, line := range lines {
		if len(line) == 0 || len(strings.TrimSpace(line)) == 0 {
			time.Sleep(time.Duration(reader.DEFAULT_LINE_BREAK_MILLISECS) * time.Millisecond)
			continue
		}

		// normal lines
		// TODO: deal with prepositions ('to') and articles ('the')
		// also, notations/operators like +->
		words := wordDelimiterRegex.Split(line, -1)
		for _, w := range words {
			fmt.Print(reader.COLOR_DEF_HEADER)
			fmt.Println(reader.DEFAULT_HEADER)
			fmt.Print("\n")

			reader.ProcessChunk(w)

			//fmt.Print("\033[0m")
			//fmt.Print("\033[H") // this prints the new line right after the cursor, but need to clean the cursor
			// fmt.Print("\x1b[0K") // ? not doing anything?
			// fmt.Print("\x1b[1A") // this prints on the cursor, but the cursor is wilding
			// fmt.Print("\033c") // this cleans everything on the terminal, fatal

			time.Sleep(sleepMicroBwEachWordSet)
			fmt.Print("\033[H\033[2J")
		}
	}
	fmt.Print("\033[0m") // set to default
}
