# comp590-finalexam

**Author**: Madison Roberts 
**PID**: 730460151

## Argument Against Deadlock 
The combination of **buffered channels**, clear synchronization with sync.WaitGroup, and no circular dependencies in communication guarantees that my Go program won't experience deadlock. 
- **Buffered channels**: Both inCh and outCh are buffered channels with a capacity of 5. This means that the producers and consumers can send and receive messages without immediately blocking. 
- **Synchronization**: The sync.WaitGroup ensures that all goroutines finish their work before the program terminates (prevents premature exit so all tasks will be completed properly).
- **No circular dependencies**: Each goroutine has a clear responsibility. The channels pass data between goroutines, but there are no circular dependencies where two goroutines are waiting on each other.

## Handling Fan-out and Fan-in in Elixir
In Elixir, you would use the **actor model** with processes and mailboxes to handle fan-out and fan-in. A central dispatcher process forwards messages from the producers to the consumers. Each consumer acts independently and consumes data concurrently from its own mailbox.

## How to Run
To run the Go program, use the following command:
```bash
go run pipeline.go