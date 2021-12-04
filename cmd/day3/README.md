## Day 3 Lessons learned

[Day 3 - Advent of Code 2021: Dive!](https://adventofcode.com/2021/day/3)

## Part 1

`strconv.ParseInt` has that nice feature, that you can specify the base. That means I can easily convert a string representation of a binary number to int.

```go
strconv.ParseInt(gammaRate, 2, 0)
```

Unfortunately it will return a `int64` type. I have not yet figured out, how to return a simple `int` like `strconv.Atoi()` does.