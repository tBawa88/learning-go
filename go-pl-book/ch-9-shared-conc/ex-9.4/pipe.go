package main

func pipeline(numStages int) (in chan int, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < numStages; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for value := range in {
				out <- value
			}
			close(out)
		}(in, out)
	}
	return first, out
}
