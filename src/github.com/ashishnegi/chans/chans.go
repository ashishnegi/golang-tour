package chans

func sum(ns []int, resChan chan int) {
	sum := 0
	for _, n := range ns {
		sum += n
	}
	resChan <- sum
}

func TwiceFastSum(ns []int) int {
	c1 := make(chan int)
	c2 := make(chan int)

	go sum(ns[:len(ns)/2], c1)
	go sum(ns[len(ns)/2:], c2)

	x, y := <-c1, <-c2
	return x + y
}
