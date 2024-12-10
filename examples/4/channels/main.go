package main

func process(f float64) {
	// do something
}

func processData(ch <-chan float64) {
	for c := range ch {
		process(c)
	}
}

func main() {
	ch := make(chan float64)
	go processData(ch)

	for i := 0; i < 10; i++ {
		ch <- float64(i)
	}
	close(ch)
}
