package main

func Merge2Channels(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		computedCount := 0
		in1Stack, in2Stack := make([]int, 0), make([]int, 0)

		for computedCount < n {

			select {
			case x1 := <-in1:
				// log.Printf("Got %d in stack1", x1)
				in1Stack = append(in1Stack, x1)
			case x2 := <-in2:
				// log.Printf("Got %d in stack2", x2)
				in2Stack = append(in2Stack, x2)
			}

			// log.Printf("computing len1=%d, len2=%d", len(in1Stack), len(in2Stack))
			if res, ok := compute(f, &in1Stack, &in2Stack); ok {
				computedCount++
				out <- res
			}
		}
	}()
}

// pops last element of stacks and returns f(last1)+f(last2), true
// othervise false
func compute(f func(int) int, stack1 *[]int, stack2 *[]int) (int, bool) {
	if len(*stack1) == 0 || len(*stack2) == 0 {
		return 0, false
	}

	// 3rd element always true, because of upper if
	x1, _ := pop(stack1)
	x2, _ := pop(stack2)

	return f(x1) + f(x2), true

}

func pop(slice *[]int) (int, bool) {
	if len(*slice) > 0 {
		res := (*slice)[len(*slice)-1]
		*slice = (*slice)[:len(*slice)-1]
		return res, true
	}
	return 0, false
}
