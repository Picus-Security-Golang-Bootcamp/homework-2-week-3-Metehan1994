# Goroutines & Channels

The phenomenon of "concurrency" which holds a significant place in Go means the ability of the different parts to be executed independently. It is generally confused with the term of "parallelism" ( [Concurrency is not Parallelism - Rob Pike](https://www.youtube.com/watch?v=oV9rvDllKEg) ). However, we can say that concurrency (a structure) enables the parallelism (an execution type), but it does not limit a process to it. Two things can be realized in a discontinuous and subsequent way. Let's think about a person writing on two papers. He/she writes something on first paper, and then, performs some work on the second paper, and later, continues to write on first paper. This type of execution is still a concurrent operation. It is clearly exemplified in the figure given below:

***

<p align="center">
  <img width="460" height="200" src=https://images.squarespace-cdn.com/content/v1/562ea223e4b0e0b9dab0b930/1567230852223-R1D5NTFFUWKX98UNW479/concurrency_vs_parallelism.png?format=750w>
</p>

<p align="center">
   Photo by https://engineering.appfolio.com/appfolio-engineering/2019/9/13/benchmarking-fibers-threads-and-processes
</p>

***

So, we can interpret that paralellism requires more than one processor while concurrency can be performed in one processor. In Go language, concurrency is naturally supported with channels & goroutines. In order to understand goroutines & channels, this article can be useful with their brief info & coding examples.

## Goroutines

Goroutines are originated from the functions, and they are actually functions. Main function is a good example of a goroutine. It has a lifetime starting with the code execution and closing after the execution of last line in the code. In addition, we can create other goroutines in the main function which are normally independent of main function and each other. Main function and some goroutines in the main function are started with the code execution and they work separately. Whenever they finish their duties, they are closed and do not affect execution of each other. So, a question can arise here. What happens if main function finishes earlier than other goroutines ? As it is discussed, main function does not wait other goroutines which are indepedent from main function and each other.

To specify a function as a goroutine, we write "go" keyword before the function name. Now, we will see their usage with a coding example.

**Goroutines without main function time delay**

```Go
package main

import (
	"fmt"
)

func declaredFunc(numOfLoops int) {
	for i := 0; i < numOfLoops; i++ {
		fmt.Println("declared function working", i)
	}
}

func main() {
	go func() {
		fmt.Println("anonymous function")
	}()

	go declaredFunc(3)
	fmt.Println("done")
}
```

```
Output:

done
```

In the example, there are two goroutines except for main function. One of them forms from an anonymous function whereas other one is a declared and called function. Only "done" is printed in output which is the last line of the code. It reveals the issue discussed above. Main function closes earlier than goroutines. It executes all the line before the goroutines complete their duties. So, their output is not observed. Some time delay for main function can be introduced in above example to see the result of goroutines.

**Goroutines with main function time delay**

```Go
package main

import (
	"fmt"
	"time"
)

func declaredFunc(numOfLoops int) {
	for i := 0; i < numOfLoops; i++ {
		fmt.Println("declared function working", i)
	}
}

func main() {
	go func() {
		fmt.Println("anonymous function")
	}()

	go declaredFunc(3)
	<-time.After(time.Second) //Line added
	fmt.Println("done")
}
```

```
Output:

anonymous function
declared function working 0
declared function working 1
declared function working 2
done
```

## Channels

Channels provide conduits by storing and transmitting data between goroutines. Its working principle is based on Go proverbs which are " Do not communicate by sharing memory; instead, share memory by communicating." They regulate the behavior of goroutines which work concurrently and line up the processes of sending and receiving data. Therefore, they allow these processes to realize on the basis of first in first out (FIFO) concept.

***

<p align="center">
  <img width="460" height="300" src="https://www.educative.io/api/edpresso/shot/4668828718989312/image/4671705944424448">
</p>

<p align="center">
   Photo by https://www.educative.io/edpresso/what-are-channels-in-golang
</p>

***

### Channel Types

Channels are categorized into two types which are unbuffered and buffered channels. They are distinguished with respect to their capacities in appearance. However, they significantly differs with their functionalities during the data transmission among goroutines.

```Go
unBuffchan := make(chan string) // only channel data type specified
Buffchan:= make(chan string, 5) // channel data type & buffer size specified
```

### Unbuffered Channels

They are defined without capacity. Such a definition results in that data exchange only occurs when both goroutines sending and receiving data are ready for the processing. It means that a goroutine (sender) can only send data when another goroutine (receiver) waits to take the same data or vice a versa. Otherwise, gouroutines trying to transmit data will be blocked. So, unbuffered channels provides synchronization since it requires send and receive sequences.

### Buffered Channels

They are created with the capacity declaration in channel definition. These channels can store up to the given capacity. So, it makes goroutines more flexible for data send and receive. The goroutine blockage occurs in the cases of empty or full channels. The logic is basic. Receiver goroutine cannot take anything from an empty channel whereas sender goroutine cannot feed a full channel with more data.

### Examples

Now, let's compare buffered channels with unbuffered channels in two examples. We can comprehend what the buffer means through first example.

#### Example 1

**Unbuffered Channels**

```Go
package main

import "fmt"

func main() {
	unBufferedChannel := make(chan string)
	unBufferedChannel <- "Hello"
	fmt.Println(<-unBufferedChannel)
}
```

```
Output:

fatal error: all goroutines are asleep - deadlock!

```

**Buffered Channels**

```Go
package main

import "fmt"

func main() {
	BufferedChannel := make(chan string, 1)
	BufferedChannel <- "Hello"
	fmt.Println(<-BufferedChannel)
}
```

```
Output:

Hello

```

For unbuffered channel, deadlock problem has been encountered since there is no receiver goroutine for data. As it is declared in unbuffered channel section, only way for data transmission of these channels is to have ready go routines for both sending and receiving data.

For buffered channel, the data can be stored in buffer even if there is no receiver. So, buffer creates a space for the data even if it is transmitted to the receiver. However, if number of data exceeds to its buffer size, then same deadlock problem emerges as it is shown in below example.

***Buffered Channel with exceeded capacity**

```Go
package main

import "fmt"

func main() {
	BufferedChannel := make(chan string, 1)
	BufferedChannel <- "Hello"
	BufferedChannel <- "World"
	fmt.Println(<-BufferedChannel)
}
```

```
Output:
fatal error: all goroutines are asleep - deadlock!
```

#### Example 2

**Unbuffered Channel**
```Go
package main

import (
	"fmt"
	"time"
)

func Send(channel chan int) {
	channel <- 1
	fmt.Println("1st data sent")
	channel <- 2
	fmt.Println("2nd data sent")
	channel <- 3
	fmt.Println("3rd data sent")
}

func Receive(channel chan int) {
	read1 := <-channel
	fmt.Println("Read: ", read1)
	read2 := <-channel
	fmt.Println("Read: ", read2)
	read3 := <-channel
	fmt.Println("Read: ", read3)
}

func main() {
	unbufferedChannel := make(chan int)
	go Send(unbufferedChannel)
	time.Sleep(time.Second * 1)
	go Receive(unbufferedChannel)
	time.Sleep(time.Second * 5)
}

```

```
Output:

1st data sent
Read:  1
Read:  2
2nd data sent
3rd data sent
Read:  3
```

**Buffered Channel**

```Go
package main

import (
	"fmt"
	"time"
)

func Send(channel chan int) {
	channel <- 1
	fmt.Println("1st data sent")
	channel <- 2
	fmt.Println("2nd data sent")
	channel <- 3
	fmt.Println("3rd data sent")
}

func Receive(channel chan int) {
	read1 := <-channel
	fmt.Println("Read: ", read1)
	read2 := <-channel
	fmt.Println("Read: ", read2)
	read3 := <-channel
	fmt.Println("Read: ", read3)
}

func main() {
	bufferedChannel := make(chan int, 3)
	go Send(bufferedChannel)
	time.Sleep(time.Second * 1)
	go Receive(bufferedChannel)
	time.Sleep(time.Second * 5)
}
```

```
Output:

1st data sent
2nd data sent
3rd data sent
Read:  1
Read:  2
Read:  3
```

As it can be seen, these two examples are same, only difference is channel type. In func main(), there are two goroutines whose one sends data while other one takes it. For unbuffered channels, data is transmitted one by one when both goroutines are ready. However, for buffered channels, three slots are created for three works (number of data stored and received). So, all data is fastly stored before they are received through receiver goroutine.

Nevertheless, buffered channels does not guarantee to prevent goroutine blockages. When buffer size is lower than the work and/or sending or receiving data faster process than other one, the blockage will be a still problem due to empty or full channel.

Final example combines the channel and goroutines. A generator goroutine creates a bundle of data transmitted by a channel to a receiver goroutine. FIFO is guaranteed in the process.

```Go
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	//Generator
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Printf("%d stored in channel\n", i)
		}
		close(c)
	}()
	//Receiver
	go func() {
		for v := range c {
			fmt.Printf("%d transmitted to receiver\n", v)
		}
	}()

	time.Sleep(time.Second * 2)
}
```

```
Output:

0 transmitted to receiver
0 stored in channel
1 stored in channel
1 transmitted to receiver
2 transmitted to receiver
2 stored in channel
3 stored in channel
3 transmitted to receiver
4 transmitted to receiver
4 stored in channel
```

I hope this article can create a conceptual fundament for further investigations about channels and goroutines building blocks for the concurrency in Go.

## References

* https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
* https://www.yakuter.com/go-dilinde-concurrency/
* https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html