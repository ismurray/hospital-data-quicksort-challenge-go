package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strconv"
    "strings"
)

func main() {
  rows := readCSV("Unplanned_Hospital_Visits_-_Hospital.csv")
  printRows(rows[0], qsort(rows[1:]))
}

func readCSV(name string) [][]string {

	f, err := os.Open(name)

	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", name, err.Error())
	}

  // Close the file when finished
	defer f.Close()

	r := csv.NewReader(f)

  // read the whole file at once, may need to be changed for larger datasets
	rows, err := r.ReadAll()

	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}

	return rows
}

func qsort(s [][]string) [][]string {
  const (
    scoreIndex = 12
    // largest possible float64
    maxFloat64 float64 = 1.797693134862315708145274237317043567981e+308
  )

  if len(s) < 2 { return s }

  leftIndex := 0
  rightIndex := len(s) - 1

  // Pick a random pivot
  pivotIndex := rand.Int() % rightIndex


  // Move the pivot to the right
  s[pivotIndex], s[rightIndex] = s[rightIndex], s[pivotIndex]

  // Pile elements smaller than the pivot on the left
  for i := range s {
    // capture the score value as a float, setting any non-convertible values
    // (ex: "Not Available") to the largest value possible, so they end up at the bottom
    currentScore, err := strconv.ParseFloat(s[i][scoreIndex], 64)
		if err != nil {
			currentScore = maxFloat64
    }

    rightScore, err := strconv.ParseFloat(s[rightIndex][scoreIndex], 64)
		if err != nil {
			rightScore = maxFloat64 // largest possible float32
    }

    if currentScore < rightScore {
      s[i], s[leftIndex] = s[leftIndex], s[i]
      leftIndex++
    }
  }

  // Place the pivot after the last smaller element
  s[leftIndex], s[rightIndex] = s[rightIndex], s[leftIndex]

  // recursively sort the rest
  qsort(s[:leftIndex])
  qsort(s[leftIndex + 1:])

  return s
}

func printRows(colNames []string, s [][]string) {
  // If the file doesn't exist, create it, or append to the file
  f, err := os.OpenFile("sorted-results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
      log.Fatal(err)
  }

  colNamesString := strings.Join(colNames, ",") + "\n"
  if _, err := f.Write([]byte(colNamesString)); err != nil {
      log.Fatal(err)
  }
  fmt.Println(colNamesString)

  for row := range s {
    rowString := strings.Join(s[row], ",") + "\n"
    if _, err := f.Write([]byte(rowString)); err != nil {
        log.Fatal(err)
    }
    fmt.Println(rowString)
  }

  if err := f.Close(); err != nil {
      log.Fatal(err)
  }

}
