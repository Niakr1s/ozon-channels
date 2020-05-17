package main

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		for i := 0; i < n; i++ {
			var x1, x2 int

			select {
			case x1 = <-in1:
				x2 = <-in2
			case x2 = <-in2:
				x1 = <-in1
			}

			go func() {
				res1, res2 := make(chan int), make(chan int)

				getResult := func(x int, res chan<- int) {
					res <- f(x)
				}

				go getResult(x1, res1)
				go getResult(x2, res2)

				out <- <-res1 + <-res2
			}()
		}
	}()
}
