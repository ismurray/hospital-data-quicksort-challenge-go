# Technical Challenge: Sorting Hospital Data
## Premise:
With a language you've never used before (Golang), sort the .csv data from `Unplanned Hospital Visits - Hospital.csv` (available from https://data.medicare.gov/data/hospital-compare#) by the `Score` column, ascending.

This Golang CLI program should load the .csv data, sort the data by the Score column, and then print out the records in ascending order to stdout. You only need to be able to load this particular CSV and sort by that column, but bonus if you can make it more flexible. The program should not use any existing Golang sorting APIs, and you should implement either quicksort or bubblesort.
