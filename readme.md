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
```
go run ./cli/monkey ./example/monkey
```


### Syntax
Example can be found in `example/` dir

#### statement
```
 let a = 10;
 let b = 3;
```
#### arithmatic
```
5 > 3;          // true
2 + 5 < 10 ;    // true
4 == 2 * 2;     // true
3 > 3 * 3;      // false
```

#### built-in type

##### WIP