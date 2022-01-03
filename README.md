# Go

## Creation of Go

Go was created amongst a group of engineers within Google.

## Go Characteristics

- Fast compilation
- Fully compiled
- Strongly typed
- Concurrent by default (the idea that an application would have to work with multiple threads was a default and thus designed with that in mind)
- Garbage collected language
- Simplicity as a core value

## What is Go Good At?

- Developing web services and web applications: Go was designed with concurrency in mind whereas other applications typically use third party libraries to assist with this task.
- Task automation: Go's syntax is almost as light as other scripting languages.
- GUI/Thick-client: There's an experimental library that provides all basic UI elements required to build apps that can run on different platforms.
- Machine learning: Syntax is light enough and characteristics are similar to Python.

## Go playground

A link to the go playground can be found [here](https://go.dev/play/). This is an online editor that supports development of simple applications. Good resource for learning.

## Hello World Walk Through

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```

As you can see above in the code example it begins with `package main`, every Go source code is going to start with a package statement. A package statement is a directory inside your Go workspace that contains on or more Go source files.

The next thing we see is an import block which in this case is: `import ("fmt")`. Import blocks are used to import functionality from other libraries into the workspace we are working with.

The entry point of the application is the main function `main()`. This is the point at which are application will start. Every executable program in Go will contain a package main statement and main function.

Another key thing with Go is the curly braces need to be declared within the same line as the function statement. You cannot have the curly braces begin on a newline after declaring a function. For instance, this would cause an error:
```
func main()
{
	fmt.Println("Hello, world")
}
```

## Installing Go

To install Go on your machine click [here](https://go.dev/dl/). This link will take you to the download section to install all necessary packages and libraries on your system.
