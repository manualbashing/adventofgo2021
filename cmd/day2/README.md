## Day 2 Lessons learned

[Day 2 - Advent of Code 2021: Dive!](https://adventofcode.com/2021/day/2)

## Part 1

I figured that the go file is not executable if the package is not declared as `main` but as any other thing, like `day2`. Need to read this: [Packages — An Introduction to Programming in Go | Go Resources (golang-book.com)](https://www.golang-book.com/books/intro/11)

---

This time I also understood the difference between using `=` and `:=` in a variable assignment. (The latter infers the type.) I like that!

---

I had to look up how to split a string on a whitespace character. Easy peasy with help of this: [3 ways to split a string into a slice · YourBasic Go](https://yourbasic.org/golang/split-string-into-slice/)

---

A little bit of googling was also necessary to figure out, how multidimensional arrays are declared in Go: [Go - Multidimensional Arrays in Go (tutorialspoint.com)](https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.htm)

So far I like Go's syntax. It feels very intuitive.

## Part 2

This part called for a more oop approach. The way Go handles this is definitely something I need to get used to. There are no classes, just structs and interfaces with attached methods! [Structs Instead of Classes - Object Oriented Programming in Golang (golangbot.com)](https://golangbot.com/structs-instead-of-classes/)

---

I had to lookup how to actually add a method to a struct: [How to add a method to struct type in Golang? - GeeksforGeeks](https://www.geeksforgeeks.org/how-to-add-a-method-to-struct-type-in-golang/)

---

This also made me try to get a little more familiar with the package system of Go. Especially importing a local module from a subfolder was something I did not get at first: [How to import local packages in go? - Stack Overflow](https://stackoverflow.com/questions/35480623/how-to-import-local-packages-in-go)

The key was to remember, what name I gave the module when I was running `go mod init {name}` the first time. So instead of importing a module like:

```go
p2 "day2/part2" // 👎
```

I had to do it like this instead:

```go
p2 "github.com/manualbashing/adventofgo2021/cmd/day2/part2" // 👍
```

---

Another important takeaway was, that imported modules **must not** have a name that is already used by a function in the same part of the program. Otherwise the compiler won't know, what to do. (In my case the trouble was the function `part2()`)