# Day 05 If You Give A Seed A Fertilizer
https://adventofcode.com/2023/day/5

Brute forced.
Used goroutines to speed things up.

## Part 1

Originally tried to precalculate all the mappings given the destination-source-length data.  This worked fine for the example but the actual data created way too many mapping entries to check.  Should look at the data before writing solutions.  Determining the mapping path on the fly was simple enough though.  

## Part 2

Same wrong attempt this time around.  Tried to precalculate all the potential seeds which quickly grew larger than memory.  Iterating through the potential seeds instead of saving them was all that needed to be done.  On an older i5 with 6 cores this took 4m12s with no concurrency.  Breaking the calculations into the 10 seed ranges each running on its own goroutine reduced the run time to 1m57s.