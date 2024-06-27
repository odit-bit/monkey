This project contains an interpreter for the "Monkey" programming language with GO as host language, see [monkeylang](https://monkeylang.org/)

### QuickStart  
#
#### REPL
```
git clone https://github.com/odit-bit/monkey
cd monkey
```
run the repl
```
go run ./cli/repl 
```

```
>> let multFive = fn(x) {x * 5; };
>> multFive(5)
```
it should return
```
25
```

#### FILE
interpreted from file 
```go
go run ./cli/monkey ./example/monkey
```


### Syntax
Example can be found in `example/` dir

#### statement
```monkey
 let a = 10;
 let b = 3;
```

#### function
```go
let adderTwo = fn (x) {x + 2;}
let minAddertwo = fn(x, adderTwo) {return x - adderTwo(x);}
minAddertwo(10, adderTwo) // call the function
```

#### arithmatic
```go
5 > 3;          // true
2 + 5 < 10 ;    // true
4 == 2 * 2;     // true
3 > 3 * 3;      // false
```

#### built-in type

### string
```go
// concat string
let a = "hello";
let b = "world";
let concatWord = fn (a,b) {return a + " " + b;};
```

### array
```go
let arr = ["hello", "world", 2000 + 24];
```

##### WIP