# Instructions

Count the frequency of letters in texts using parallel computation.

Parallelism is about doing things in parallel that can also be done sequentially. A common example is counting the frequency of letters. Create a function that returns the total frequency of each letter in a list of texts and that employs parallelism.

# Concurrency vs Parallelism

Go supports concurrency via "goroutines" which are started with the go keyword. It is a simple, lightweight and elegant way to provide concurrency support and is one of the greatest strengths of the language.

You may notice that while this exercise is called Parallel letter frequency you don't see the term "Parallel" used very often in Go. Gophers prefer to use the term Concurrent to describe the management of multiple independent goroutines ("processes" or "threads" in other language contexts). Although these terms are often used interchangeably Gophers like to be technically correct and use "concurrent" when discussing the seemingly simultaneous executions of goroutines.

While we can plan for our programs to run in parallel, and at times they may appear to run in parallel, without strict knowledge of the execution context of our code all we can guarantee is that processes will run concurrently. In other words they may be executing sequentially faster than we can distinguish but not strictly simultaneously.

For more take a look at The Go Blog's post: Concurrency is not parallelism.

# Solution

- [Code](solution.go)