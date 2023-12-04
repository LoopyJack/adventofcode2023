# Day 04 Scratchcards
https://adventofcode.com/2023/day/4

Experimented with a couple different methodologies to get an idea of performance.  First, I tinkered with how matches were found.  And then, the difference between a slice of structs and a slice of pointers to structs.

## Part 1

straightforward

## Part 2

With the various ways of finding matches, at first, I was finding matches on a per request basis just to get some kind of baseline performance.  Part 2 took about 20 full seconds to run this way.  After precalculating the matches slice and taking the length, it took about 25 milliseconds.  After that, I precalculated the length, and that didn't have any noticeable effect on performance.  Len() is a pretty cheap operation.
Switching from a slice of structs to a slice of pointers to structs shaved roughly a ms off.
