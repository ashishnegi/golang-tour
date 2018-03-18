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

// Note: Only the sender should close a channel, never the receiver.
// Sending on a closed channel will cause a panic.

// Another note: Channels aren't like files; you don't usually need to close them.
// Closing is only necessary when the receiver must be told there are no more
// values coming, such as to terminate a range loop.

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case.
// It chooses one at random if multiple are ready.

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		Walk(t1, c1)
		close(c1)
	}()
	go func() {
		Walk(t2, c2)
		close(c2)
	}()

	for {
		x, more1 := <-c1
		y, more2 := <-c2
		if more1 == false && more2 == false {
			return true
		} else if more1 != more2 {
			return false
		} else if x != y {
			return false
		}
	}
}
