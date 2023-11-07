# Pointers

A variable is a piece of storage containing a value. Variables created by
declaration are identified by name, such as x, but many variables are identified
only by expressions like x[i] or x.f. All these expressions read the value
of a varialbe, except when they appear on the left hand side of an assignment

==> If a variable is declared var x int, the expression &x ("address of x")
yields a pointer to an integer variable, that is  a value of type *int
which is pronounced "pointer to int"
*

```go

x := 1
p := &x // p of type *int, points to x
*p = 2 // equivalent to x = 2
fmt.Println(x) // 2

```

```go

var p = f()

func f() *int{
  v := 1
  return &v
}

```

==> Because a pointer contains the address of a variable, passing a pointer
argument to a function makes it possible for the function to update the variable
that was indirectly passed.

```go

func incr(p *int) int{
  *p++ // increments what p points to; does not change p
  return *p
}

```


## Lifetime of variables

The lifetime of a variable is the interval of time during which it exists as
the program executes. The lifetime of a package level variable is the entire
execution of the program.

```go

for t := 0; t < cycles*2*math.Pi; t += res{
  x := math.Sin(t)
  y := math.Sin(t * freq + phase)
  img.SetColorIndex(size+int(x*size+0.5), size + int(y * size + 0.5),
blackIndex)
}

```

## Type Declarations

The type of a variable or expression defines the characteristics of the values
it may tak on such as their size (number of bits or number of elements, perhaps)
how thy are represented internally

```go

import "fmt"

type Celsius float64
type Farenheit float64

const (
  AbsoluteZeroC Celsius = -273.15
  FreeZingC Celsius = 0
  BoilingC Celsius = 100
)

func CToF(c Celsius) Farenheit { return Farenheit(c*9/5 + 32)}
func FToC(c Celsius) Celsius { return Celsius((f - 32) * 5/9)}

```

### Packages and Files

Packages in Go serve the same purposes as libraries or modules in other languages,
supporting modularity, encapsulation, separate compilation, and reuse
