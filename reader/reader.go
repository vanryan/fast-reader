package reader

import (
	"fmt"
	"strings"
)

const (
	COLOR_DEF_HEADER             string = "\033[35m"
	DEFAULT_HEADER               string = "<<<<<<<<<<<<<<<<<<<<<<<< FAST READER >>>>>>>>>>>>>>>>>>>>>>>>"
	COLOR_TEXT                   string = "\033[32m"
	DEFAULT_WPM                  int    = 300
	WORD_DELIMS                  string = ",;: \t"
	EXPECTD_CHUNK_LEN_LIMIT      int    = 60
	COLOR_MID                    string = "\033[93m"
	DEFAULT_LINE_BREAK_MILLISECS int    = 1500
)

func GetSleepTimeMicroByWpm(wpm int) int {
	// get the interval b/w each word given wpm
	return int(60.0 / float64(wpm) * 1000 * 1000)
}

func ProcessChunk(chunk string) {
	// the chunk can be one word or a couple of words
	l := len(chunk)
	if l == 0 {
		return
	}
	// get the length of word and print spaces to balance the word
	// if the chunk length limit is 60, then make sure 30th is the middle in the line on the screen
	preSpaces := EXPECTD_CHUNK_LEN_LIMIT/2 - l/2
	fmt.Print(strings.Repeat(" ", preSpaces))

	fmt.Print(COLOR_TEXT)
	if l < 3 {
		fmt.Print(chunk)
		return
	}
	// find the middle character and color it
	// to do this, need to split the chunk into 3 parts
	//fmt.Print(chunk)
	mid := l / 2
	p1 := chunk[:mid]
	fmt.Print(p1)

	p2 := chunk[mid]
	p3 := chunk[mid+1:]

	fmt.Print(COLOR_MID)
	fmt.Print(string(p2))
	fmt.Print(COLOR_TEXT)
	fmt.Print(p3)

}
