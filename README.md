# Writing an Interpreter
Reference: "Writing An Interpreter In Go" by Thorsten Ball

# The Monkey Language & Interpreter
The interpreter will parse and evaluate a made up language called Monkey
- The features of Monkey include:
	- C-like syntax
	- variable bindings
	- integers and booleans
	- arithmetic expressions
	- built-in functions
	- first-class and higher-order functions
	- closures
	- a string data structure
	- an array data structure
	- a hash data structure
- The interpreter will be a *tree-walking* interpreter
	- The interpreter parses the source code in order to build an abstract syntax tree (AST) out of it and then evaluates the tree
## Monkey Syntax
Binding values to names in Monkey can be done like this:
```c
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

Binding an array of integers to a name looks like:
```c
let myArray = [1, 2, 3, 4, 5];
```

The syntax for a hash looks like:
```c
let myHash = {"name": "Monkey", "version": 1};
```

Accessing elements in arrays and hashes is done with index expressions:
```c
myArray[0]     // => 1
myHash["name"] // => "Monkey"
```

Functions in Monkey are just values, like integers or strings. `let` statements can be used to bind functions to names, e.g.,
```c
let add = fn(a, b) { return a + b; };
```

The language supports *implicit* return values so the code below is equivalent to the code block above:
```c
let add = fn(a, b) { a + b; };
```

Calling a function works by calling the variable assigned in its `let` statement:
```c
add(1, 2);
```

Monkey supports higher-order functions that can take other functions as arguments. For example:
```c
let twice = fn(f, x) {
	return f(f(x));
};

let addTwo = fn(x) {
	return x + 2;
};

twice(addTwo, 2); // => 6
```

## Components of the Interpreter
- The interpreter will tokenize and parse Monkey source code in a read-evaluate-print loop (REPL), in order to built an abstract syntax tree and evaluate the tree.
- The major components of the interpreter include:
	- the lexer
	- the parser
	- the Abstract Syntax Tree (AST)
	- the internal object system
	- the evaluator