package main

import (
	"testing"
	"time"
)

func Mul2(i int) int {
	return i * 2
}

func TestMerge2Channels(t *testing.T) {
	in1, in2, out := make(chan int), make(chan int), make(chan int)

	n := 50

	Merge2Channels(Mul2, in1, in2, out, n)

	go func() {
		for i := 0; i < n; i++ {
			<-out
		}
	}()

	half := n / 2

	for i := 0; i < half; i++ {
		in1 <- i
	}

	for i := 0; i < half; i++ {
		in2 <- i
	}

	for i := half; i < n; i++ {
		in2 <- i
		in1 <- i
	}

	<-time.After(time.Millisecond * 100)

	select {
	case <-out:
		t.Fatalf("Out should be empty")
	default:
	}

	select {
	case in1 <- 1:
		t.Fatalf("shouldn't write in in1")
	case in2 <- 1:
		t.Fatalf("shouldn't write in in2")
	default:
	}
}
