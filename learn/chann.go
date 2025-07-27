package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func sendMessage(ch chan string, msg string) {
	ch <- msg
}
func bufferedChannel() {

	for i := 0; i < 10; i++ {
		fmt.Println()
		res := fmt.Sprintf("the %d th generate", i)
		fmt.Println("res" + res)
		var usr1Chan chan string = make(chan string)

		var usr2Chan chan string = make(chan string)
		go sendMessage(usr1Chan, "hello ")
		go sendMessage(usr1Chan, "world")

		go sendMessage(usr2Chan, "永不")
		go sendMessage(usr2Chan, "言弃")

		fmt.Print(<-usr1Chan)
		fmt.Print(<-usr1Chan)
		fmt.Println()
		fmt.Print(<-usr2Chan)
		fmt.Print(<-usr2Chan)
	}

}

func selectChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	go sendMessage(ch1, "hi")
	go sendMessage(ch2, "hi")
	go sendMessage(ch3, "hi3")
	select {
	case msg1 := <-ch1:
		fmt.Println("Channel 1 " + msg1)
	case <-ch2:
		fmt.Println("Channel 2")
	case <-ch3:
		fmt.Println("Channel 3")
	}
}

func rangeChanUse() {
	ch1 := make(chan string)
	for i := 1; i < 10; i++ {
		go sendMessage(ch1, strconv.Itoa(i))
		//if i == 9 {
		//	close(ch1)
		//}
	}

	for v := range ch1 {
		fmt.Printf("the send index : %s\n", v)
	}
}

func channel2Return() {
	ch1 := make(chan string)
	go func() {
		ch1 <- "first"
		fmt.Println("数据已发送")
		close(ch1)

	}()

	time.Sleep(10 * time.Second)
	go func() {
		v1, ok := <-ch1
		fmt.Printf("exits: %t, value is %s", ok, v1)
	}()

}

func commonProblemsWithGoRoutine() {
	num := 10
	go func() {
		num = num * 2
		fmt.Printf("fisrt * 2")
	}()

	go func() {
		num = num * 2
		fmt.Printf("second * 2")
	}()

	time.Sleep(time.Second)
	// the answer maybe 10
	fmt.Println("number is :" + strconv.Itoa(num))
}

type Provider struct {
	mutex sync.Mutex
	value int
	wait  *sync.WaitGroup
}

func (ticktes *Provider) sell() bool {
	//fmt.Printf("current value is %d\n", ticktes.value)
	defer ticktes.wait.Done()
	ticktes.mutex.Lock()
	defer ticktes.mutex.Unlock()

	if ticktes.value > 0 {
		time.Sleep(time.Microsecond)
		ticktes.value = ticktes.value - 1
		fmt.Printf("value is %d\n", ticktes.value)
		//fmt.Println("you got a ticket !")
		return true
	}

	return false
}
func mutex() {

}
func main() {

	wg := &sync.WaitGroup{}
	wg.Add(11)
	solder := &Provider{value: 5, wait: wg}

	for i := 10; i >= 0; i-- {
		go solder.sell()
	}
	//fmt.Println("waitingg....")
	wg.Wait()
	fmt.Printf("last vlaue is %d\n", solder.value)

}
