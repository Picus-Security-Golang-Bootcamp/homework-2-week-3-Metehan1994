# Channels

Channels provide conduits by storing and transmitting data between goroutines. Its working principle is based on Go proverbs which are " Do not communicate by sharing memory; instead, share memory by communicating." They regulate the behavior of goroutines which work concurrently and line up the processes of sending and receiving data. Therefore, they allow these processes to realize on the basis of first in first out (FIFO) concept. In order to understand channels, this article can be useful with its content about channel types, usage and with its coding examples.

***

<p align="center">
  <img width="460" height="300" src="https://www.educative.io/api/edpresso/shot/4668828718989312/image/4671705944424448">
</p>

<p align="center">
   Photo by https://www.educative.io/edpresso/what-are-channels-in-golang
</p>

***

## Channel Types

Channels are categorized into two types which are unbuffered and buffered channels. They are distinguished with respect to their capacities in appearance. However, they significantly differs with their functionalities during the data transmission among goroutines.

```Go
unBuffchan := make(chan string) // only channel data type specified
Buffchan:= make(chan string, 5) // channel data type & buffer size specified
```

### Unbuffered Channels

As it is discussed, they are defined without capacity. Such a definition results in that data exchange only occurs when both goroutines sending and receiving data are ready for the processing. It means that a goroutine (sender) can only send data when another goroutine (receiver) waits to take the same data or vice a versa. Otherwise, gouroutines trying to transmit data will be blocked. So, unbuffered channels provides synchronization since it requires send and receive sequences.

### Buffered Channels

They are created with the capacity declaration in channel definition. These channels can store up to the given capacity. So, it makes goroutines more flexible for data send and receive. The goroutine blockage occurs in the cases of empty or full channels. The logic is basic. Receiver goroutine cannot take anything from an empty channel whereas sender goroutine cannot feed a full channel with more data.

## Examples

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

For buffered channel, the data can be stored in buffer even if there is no receiver. So, buffer creates a space for the data even if it is tranmitted to the receiver. However, if number of data exceeds to its buffer size, then same deadlock problem emerges as it is shown in below example.

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

I hope this article can create a conceptual fundament for further investigations about channels and go routines building blocks for the concurrency in Go.

## References

* https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
* https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
* https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html