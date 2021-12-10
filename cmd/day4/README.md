## Day 4 Lessons learned

[Day 3 - Advent of Code 2021: Giant Squid](https://adventofcode.com/2021/day/4)

Time to use regex! I learned about a way how to split a string with several delimiters when reading data from a file.

```go
line := regexp.MustCompile(`[, ]+`).Split(scanner.Text(), -1)
```

---

