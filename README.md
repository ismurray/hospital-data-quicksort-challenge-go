## Technical Challenge: Sorting Hospital Data with a New Language
### Context:
The following is a technical challenge given by an interviewer for a Back End Developer position. With the primary goal of the challenge being to quickly pick up a completely new language and implement some simple functionality/algorithm with it.

### Premise:
With a language you've never used before (Golang), sort the .csv data from `Unplanned Hospital Visits - Hospital.csv` (available from https://data.medicare.gov/data/hospital-compare#) by the `Score` column, ascending.

This Golang CLI program should load the .csv data, sort the data by the Score column, and then print out the records in ascending order to stdout. You only need to be able to load this particular CSV and sort by that column, but bonus if you can make it more flexible. The program should not use any existing Golang sorting APIs, and you should implement either quicksort or bubblesort.

### Installation and Use
1. Clone the repo
2. Install Go (easily installed with homebrew on Mac using `brew install go`)
3. run `go run sort.go` in the terminal (note: this may take several moments to run)
4. The sorted results will be printed to the terminal, and also saved in a local log file (`sorted-results.txt`)

### Implementation Thoughts/Notes:
It was really interesting to work with Golang for the first time! It's quite different from most of the languages I've worked with in the past, but it seems to have some powerful functionality.

I wasn't sure whether the interviewer wanted rows without available score data to be included in the results, so I decided to push those to the bottom of the results, similar to the behavior you would get if you used the Excel-native sort functionality on this file.

Also, the output was a bit much to copy-paste from the terminal s requested (670,107 rows!), so instead I have the script both print to terminal as requested and write to a logfile (sorted-results.txt). This has an added benefit of making it easier to share/read the results, and it would be easier to modify this program to output the results in .csv format for further processing if that was desired.

### Next Steps:
I was running short on time, so for now it just does the sort on the `Score` column, as opposed to adding in some flexibility to the interface.

With some more time, I would have liked to first add flexibility to which columns you can sort by and whether to sort it ascending/descending. The current implementation is somewhat coupled to the specific typing of the `Score` column, so in addition to interface changes, this would require some type-flexible refactoring.

A further stretch goal would be to allow this to work on any CSV dataset. Both of these feature extension would require more flexible data type handling and error handling.

Another thing I'd like to learn more about going forward is the goroutine/parallelism functionality of Go. This seems to be one of the big strengths of this language, and I imagine I could get this program to run much faster by making use of it.
