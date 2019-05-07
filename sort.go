package main

import (
    "fmt"
    "math/rand"
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
  arr := []int{2, 1, 42, 13, 99}
	fmt.Println(qsort(arr))
}
