package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "math/rand"
    "os"
    "reflect"
)

func qsort(s []int) []int {
  if len(s) < 2 { return s }

  leftIndex := 0
  rightIndex := len(s) - 1

  // Pick a pivot
  pivotIndex := rand.Int() % len(s)

  // Move the pivot to the right
  s[pivotIndex], s[rightIndex] = s[rightIndex], s[pivotIndex]

  // Pile elements smaller than the pivot on the left
  for i := range s {
    if s[i] < s[rightIndex] {
      s[i], s[leftIndex] = s[leftIndex], s[i]
      leftIndex++
    }
  }

  // Place the pivot after the last smaller element
  s[leftIndex], s[rightIndex] = s[rightIndex], s[leftIndex]

  // Go down the rabbit hole
  qsort(s[:leftIndex])
  qsort(s[leftIndex + 1:])

  return s
}

func main() {
  rows := readOrders("Unplanned_Hospital_Visits_-_Hospital.csv")
  fmt.Println(rows[0])
  fmt.Println(rows[1][12])
  fmt.Println(reflect.TypeOf(rows[1][12]))
  fmt.Println(rows[3][12])
  fmt.Println(reflect.TypeOf(rows[3][12]))
  arr := []int{2, 1, 42, 13, 99}
	fmt.Println(qsort(arr))
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
