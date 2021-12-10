## Day 4 Lessons learned

[Day 4 - Advent of Code 2021: Giant Squid](https://adventofcode.com/2021/day/4)

This was all about slices for me! [Go by Example: Slices](https://gobyexample.com/slices)

---

Time to use regex! I learned about a way how to split a string with several delimiters when reading data from a file: [go - How to split a string by multiple delimiters - Stack Overflow](https://stackoverflow.com/questions/39862613/how-to-split-a-string-by-multiple-delimiters)

```go
line := regexp.MustCompile(`[, ]+`).Split(scanner.Text(), -1)
```

---

Found a nice way to skip empty lines when reading data from a file: [[bufio.Scanner\] skip empty lines (google.com)](https://groups.google.com/g/golang-nuts/c/B3YVAZXi6nc?pli=1)

```go
line := bytes.TrimSpace(s.Bytes())
if len(line) == 0 {
    continue
}
```

---

There was a good question on stackoverflow about the different ways to initialize a slice: [arrays - Correct way to initialize empty slice - Stack Overflow](https://stackoverflow.com/questions/29164375/correct-way-to-initialize-empty-slice)

---

I also learned how to properly empty a slice if necessary: [How to best clear a slice: empty vs. nil · YourBasic Go](https://yourbasic.org/golang/clear-slice/)

---

There is no builtin `pop()` in Go but it is easy to do anyways. (Reminds me of Python!)

```go
func main() {
	dwarfs := []string{"Doc", "Grumpy", "Happy", "Sleepy", "Bashful", "Sneezy", "Dopey"}
	// remove first element
	first := dwarfs[0]
	dwarfs = dwarfs[1:]

	fmt.Println(first)
	fmt.Println(dwarfs)
}
```

Source: [Remove first element of slice (shift, pop(0)) (code-maven.com)](https://code-maven.com/slides/golang/remove-first-element-of-slice)

---

Related to this is also the question of how to determine the last element in a slice?

```go
func main() {
    dwarfs := []string{"Doc", "Grumpy", "Happy", "Sleepy", "Bashful", "Sneezy", "Dopey"}
    // remove last element
    lastIndex := len(dwarfs) - 1
    last := dwarfs[lastIndex]
    dwarfs = dwarfs[:lastIndex]

    fmt.Println(last)
    fmt.Println(dwarfs)
}
```

---

The method for checking if a string starts with a certain expression is called `HasPrefix()`... But why? Can't we just all agree that `StartsWith()` is the best name for this? [Go: check if a string starts with a substring #golang (github.com)](https://gist.github.com/flaviocopes/6717857ca7ba7266d41330329eeb5ad0)