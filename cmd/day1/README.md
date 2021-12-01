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

