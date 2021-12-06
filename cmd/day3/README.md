## Day 3 Lessons learned

[Day 3 - Advent of Code 2021: Dive!](https://adventofcode.com/2021/day/3)

## Part 1

`strconv.ParseInt` has that nice feature, that you can specify the base. That means I can easily convert a string representation of a binary number to int.

```go
strconv.ParseInt(gammaRate, 2, 0)
```

Unfortunately it will return a `int64` type. I have not yet figured out, how to return a simple `int` like `strconv.Atoi()` does.

## Part 2

There are no while loops in Go. And that is actually a clever thing.

```go
for matchFound := true; matchFound; matchFound = len(matches) == 1 {
		matches = append(matches, "foo")
}
```

Using a for loop this way seems more flexible anyways.

---
 
 There is an interesting way how functions are defined to returen multiple values:

 ```go
 func function() (string, string) {

     return "foo","bar"
 }

 value1,value2 := function()
 ```
 
 ---

I learned a bit about type conversions:

- `"0"[0]` will be of type `byte`, so there is no use to compare `"1001"[0] == "1"`
- There is no conversion from `bool` to `int`: https://github.com/golang/go/issues/9367