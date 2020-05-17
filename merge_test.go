package main

import (
	"math/rand"
	"testing"
	"time"
)

func Mul2(i int) int {
	<-time.After(time.Millisecond * time.Duration(rand.Intn(15)))
	return i * 2
}

func TestMerge2Channels(t *testing.T) {
	in1, in2, out := make(chan int), make(chan int), make(chan int)

	n := 50

	Merge2Channels(Mul2, in1, in2, out, n)

	go func(t *testing.T) {

	}(t)

	go func() {
		for i := 0; i < n; i++ {
			<-time.After(time.Millisecond * 1)
			in1 <- i
		}
	}()

	go func() {
		for i := 0; i < n; i++ {
			<-time.After(time.Millisecond * 1)
			in2 <- i
		}
	}()

	for i := 0; i < n; i++ {
		if got, expected := <-out, Mul2(i)+Mul2(i); got != expected {
			t.Errorf("Expected: %d, got %d", expected, got)
		}
	}
}
