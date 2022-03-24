# Go 1.18 Release Demo

Go 1.18 was just recently released and includes significant changes to the language, and implementation of the 
toolchain, runtime, and libraries.

The release maintains the Go 1 "promise of compatability", they expect **MOST** all Go programs to compile and run
as they did before.

## Changes to the language

### Generics

- The syntax for function and type declarations now accepts type parameters.
- Parameterized functions and types can be instantiated by following them with a list of type arguments in square brackets.
- The new token ~ has been added to the set of operators and punctuation.
- The syntax for Interface types now permits the embedding of arbitrary types (not just type names of interfaces) as well as union and ~T type elements. Such interfaces may only be used as type constraints. An interface now defines a set of types as well as a set of methods.
- The new predeclared identifier any is an alias for the empty interface. It may be used instead of interface{}.
- The new predeclared identifier comparable is an interface that denotes the set of all types which can be compared using == or !=. It may only be used as (or embedded in) a type constraint.


### Fuzzing

Includes an implementation of fuzzing as described by [the fuzzing proposal](https://golang.org/issue/44551).

### Updates to the `Go` command

Not going to go over the specifics in this demo but they can be found [here](https://tip.golang.org/doc/go1.18#go-command).