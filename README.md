# FAST-READER
Command line reader based on the Spritz fast reading method.

## What? Why?
Essentially, we want to boost our reading speed from maybe 100-200 wpm (depending on the material) to maybe 250-300 wpm by "forcing" our eyes and brain to process words and remove the need for our eyes to move along the lines of a text. With a fixated gaze, one is supposed to be able to read faster and less prone to be distracted.

# How to use it?
- `go build main.go`
- set up your terminal --> font size, window size, etc.
- `./main <filename>` or `./main <filename> <wpm>`

## Workflow
1. Read commandline input for the text to process.
2. Break down sentences into words and show up in the terminal. (The original Spritz method involves coloring the `center of gravity` in the word)
3. Maintain a clear screen.

## Implementation
1. Read file line by line
2. Print line by line, and only one line
3. Print given number of words with the correct wpm

## References
- Spritz
    - https://spritz.com/#features 
    - https://theconversation.com/spritz-and-other-speed-reading-apps-prose-and-cons-24467 
- ANSI escape code
    - https://en.wikipedia.org/wiki/ANSI_escape_code
- web app
    - https://tools-unite.com/tools/speed-reading-app 
- similar projects
    - https://github.com/radiofreejohn/cfastread 