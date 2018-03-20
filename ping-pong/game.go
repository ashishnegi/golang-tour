package main

import (
	"fmt"
	"time"
)

// How do we model a stateful code ?
// We want to model a ping pong game, where two players have access to ball.
// Each hit increments the ball hit counter.

type ball struct {
	hits int
}

func player(name string, table chan *ball) {
	for {
		b := <-table
		b.hits++
		fmt.Printf("%s hit the ball..\n", name)
		time.Sleep(500 * time.Millisecond)
		table <- b
	}
}

func main() {
	table := make(chan *ball)
	go player("Ma Long", table)
	go player("Timo Boll", table)
	// comment out this line to see the deadlock.
	table <- new(ball)
	time.Sleep(3 * time.Second)
	fmt.Println("Game ends..")
	<-table
}

// asnegi@asnegi-sawvm MINGW64 ~/Documents/Gopath/src/github.com/ashishnegi/golang-tour/ping-pong (master)
// $ go build

// asnegi@asnegi-sawvm MINGW64 ~/Documents/Gopath/src/github.com/ashishnegi/golang-tour/ping-pong (master)
// $ ls
// game.go  ping-pong.exe*

// asnegi@asnegi-sawvm MINGW64 ~/Documents/Gopath/src/github.com/ashishnegi/golang-tour/ping-pong (master)
// $ ./ping-pong.exe
// Game ends..
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         C:/Users/asnegi/Documents/Gopath/src/github.com/ashishnegi/golang-tour/ping-pong/game.go:33 +0x119

// goroutine 18 [chan receive]:
// main.player(0x4bf95d, 0x7, 0xc04204e0c0)
//         C:/Users/asnegi/Documents/Gopath/src/github.com/ashishnegi/golang-tour/ping-pong/game.go:18 +0x4f
// created by main.main
//         C:/Users/asnegi/Documents/Gopath/src/github.com/ashishnegi/golang-tour/ping-pong/game.go:28 +0x7d

// goroutine 19 [chan receive]:
// main.player(0x4bfdfd, 0x9, 0xc04204e0c0)
//         C:/Users/asnegi/Documents/Gopath/src/github.com/ashishnegi/golang-tour/ping-pong/game.go:18 +0x4f
// created by main.main
//         C:/Users/asnegi/Documents/Gopath/src/github.com/ashishnegi/golang-tour/ping-pong/game.go:29 +0xb4

// asnegi@asnegi
