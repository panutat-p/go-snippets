# Channel

## Declare

```go
var ch = make(chan int)
```

```go
var ch = make(chan int, 1)
```

## Write

⚠️ Deadlock
```go
ch := make(chan string)

ch <- `{"apple": 15, "banana": 8}` // block

msg := <-ch
fmt.Println(msg)
```

```go
ch := make(chan string)

go func() {
    ch <- `{"apple": 15, "banana": 8}`
}()

msg := <-ch
fmt.Println(msg)
```

## Read

```go
ch := make(chan string, 5)
for i := 1; i <= 5; i++ {
    ch <- fmt.Sprintf("Message %d", i)
}

close(ch)

for message := range ch {
    fmt.Println(message)
}
```

> ⚠️ In Go, when you read from a closed channel, you will receive the zero value for the channel's type.
> For a channel of strings, this would be an empty string `""`.
> For a channel of integers, this would be `0`.

```go
ch := make(chan string, 5)
for i := 1; i <= 5; i++ {
    ch <- fmt.Sprintf("Message %d", i)
}

close(ch)

for {
    select {
    case msg, ok := <- ch:
        if !ok {
            // ⚠️ channel has been closed
            return
        }
        fmt.Println(msg)
    }
}
```

## Preemptive vs Cooperative
> Preemption refers to the ability of the Go runtime scheduler to interrupt a running goroutine and switch the execution to a different goroutine.

> Go routines are managed by the Go runtime's scheduler, this scheduler is not preemptive in the traditional sense.
> Go's scheduler is cooperative, meaning it only checks whether it should switch goroutines at certain points in the program,
> such as function calls, channel operations, and blocking system calls.

> Starting from Go 1.14, the runtime has introduced asynchronous preemption to prevent long-running goroutines from monopolizing the processor, making the scheduler more "preemptive".

> Unlike Java, threads are managed by the operating system's thread scheduler, which uses a preemptive scheduling algorithm.
> a preemptive scheduler will forcibly interrupt a running task after a fixed time slice.
> this scheduler can forcibly interrupt a running thread to give another thread a chance to run.

## FIFO queue

```go
func main() {
    var (
        ch = make(chan int, 10)
        wg sync.WaitGroup
    )
    
    wg.Add(2)

    go func ()  {
        for i := 1; i < 11; i++ {
            ch <- i
        }
        fmt.Println("positive producer complete")
        wg.Done()
    }()

    go func ()  {
        for i := -1; i > -11; i-- {
            ch <- i
        }
        fmt.Println("negative producer complete")
        wg.Done()
    }()

    go func() {
        wg.Wait()
        fmt.Println("✅ done")
        close(ch)
    }()

    // Read from the channel until it's closed
    for v := range ch {
        fmt.Println(v)
    }
}
```
