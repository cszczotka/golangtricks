# Basic concepts

###  Goroutine 
A goroutine is a lightweight thread managed by the Go runtime. Using the `go` keyword before any named or anonymous function invocation, the function executes in a separate goroutine, providing concurrent execution within the same process.

### Channel

Channels are used to communicate between different goroutines by sending and receiving shared data of almost any type. Channels handle locking or shared resources internally, blocking sends or receives as needed, depending on whether the channel is buffered or unbuffered.

### Select

### Context
The `context` package is used to manage execution contexts across parent and child functions or goroutines. It provides ways to handle timeouts, cancellations, and deadlines, enabling graceful shutdowns in various scenarios.
### Sync



# References

https://medium.com/@miantalhaakram/unlocking-the-power-of-concurrency-in-go-a-practical-deep-dive-e714a921252a