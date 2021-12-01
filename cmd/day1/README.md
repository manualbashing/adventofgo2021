## Day 1 Lessons learned

[Day 1 - Advent of Code 2021: Sonar Sweeps](https://adventofcode.com/2021/day/1)

I am completely new to golang and this means for me: googling the absolute basics.

## Part 1

I was following this advice and created a module and cmd sublfolder: [Starting a Go Project · Mark Wolfe's Blog](https://www.wolfe.id.au/2020/03/10/starting-a-go-project/) 

I did not yet figure out the meaning of this. Does everyone do this? 🤔

---

I was surprised to see, that the vscode extension for go does not substitute tabs with spaces, as this is my general setting. Apparently google recommends using tabs for indentation... I thought tabs are dead?  Apparently not: [formatting - Indentation in Go: tabs or spaces? - Stack Overflow](https://stackoverflow.com/questions/19094704/indentation-in-go-tabs-or-spaces)

---

The first challange was to read integers from a file. This provided helpful advice:

- [Go by Example: Reading Files](https://gobyexample.com/reading-files).
- [Golang read file line by line to string - golangprograms.com](https://www.golangprograms.com/golang-read-file-line-by-line-to-string.html)

---

And somehow I needed to interate over resulting array in a for loop: [The for-loop in GoLang - GoLang Docs](https://golangdocs.com/for-loop-in-golang)

---

I was struggling a bit with parsing string to int. Following this: [Convert string to integer type in Go? - Stack Overflow](https://stackoverflow.com/questions/4278430/convert-string-to-integer-type-in-go) I was trying to figure out `strconv.ParseInt()` but was struggling with the fact that an error objects is returned and needs to be assigned away somehow.

In the end I followed this [go - Golang, is there a better way read a file of integers into an array? - Stack Overflow](https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array) and used `strconv.Atoi()` , which I felt was more intuitive.

---

Finally I was looking up how a format string works in golang. Pretty straightforward I woud say: [Go string format - formatting strings in Golang (zetcode.com)](https://zetcode.com/golang/string-format/)

I did not quite figure out the difference yet between `fmt.Printf()` and `fmt.Sprintf()` but I put this article to my reading list: [Golang’s fmt.Sprintf() and Printf() Demystified | by Dirk Avery | FAUN Publication](https://faun.pub/golangs-fmt-sprintf-and-printf-demystified-4adf6f9722a2)

### Part 2

[Day 1 - Advent of Code 2021: Part2](https://adventofcode.com/2021/day/1#part2)

I was able to solve part two with all information that I already gathered for part1. My first approach was straigthforward but contained unnecessary calculations:

```go
var current int = data[i] + data[i+1] + data[i+2]
var next int = data[i+1] + data[i+2] + data[i+3] // <- Here
```

As this iterations `next` is next iterations `current`, the calculation of this value became redundant.

So I was refactoring that bit and met something strange:

```go
var next int = data[0] + data[1] + data[2]
for i := 0; i < len(data)-3; i++ {

    var current int = next
    var next int = data[i+1] + data[i+2] + data[i+3]
    if current < next {
        larger++
    }
}
```

This was giving me the wrong result. Using the debugger I was able to see, that `current` was set to the value of `data[0] + data[1] + data[2]` on every single iteration! 

I had the intensions to overwrite the content of the variable `next` with a new value after assigning it to `current` but instead I was apparently creating a new local `next` variable that only lived within the current iteration of the for loop. 🤦‍♂️ See here:  [Golang Variables Scope: What is the Scope of Variables In Go (appdividend.com)](https://appdividend.com/2020/01/29/scope-of-variables-in-golang-go-variables-scope/)

So I stopped declaring a new variable and did it right instead:

```go
var next int = data[0] + data[1] + data[2]
for i := 0; i < len(data)-3; i++ {

    var current int = next // this is OK to be a local var
    next = data[i+1] + data[i+2] + data[i+3] // this not
    if current < next {
        larger++
    }
}
```

