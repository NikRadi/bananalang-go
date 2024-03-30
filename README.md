# Banana-lang
Banana is a high-level, dynamic language that is compiled into Banana byte code and interpreted by the Banana virtual machine (BVM).  
The purpose of this language is for me to learn Go.

![Banana image](img/banana-lang.webp)

### Declarations

#### Variable declaration
```
var pi = 3.1415;
var name = "Earth";
var people = [
    { "age": 30 },
    { "age": 34 },
    { "age": 29 },
];
```

#### Function declaration
```
func Add(a, b) {
    return a + b;
}
```

### Statements

#### If statement
```
if condition1 {
    // code
}
else if condition2 {
    // more code
}
else {
    // even more code
}
```

#### For-loop statement
```
for var i = 0; i < 10; i += 1 { /* code */ }

// Infinite loop
for {
    // break;
}
```

### Examples

#### Hello Banana
```
import "Console"

func Main() {
    Print("Hello Banana");
    return 0;
}
```

#### Fibonacci with compile time execution
```
import "Console"

func Fibonacci(n) {
    if n == 0 {
        return 0;
    }

    if n == 1 || n == 2 {
        return 1;
    }

    return Fibonacci(n - 1) + Fibonacci(n - 2);
}

func Main() {
    // The '@' is executed at compile-time.
    var my_value = @Fibonacci(10)
    return 0;
}
```

