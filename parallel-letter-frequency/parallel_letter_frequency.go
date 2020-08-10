package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in an array of strings and return this data as a FreMap
func ConcurrentFrequency(strings []string) FreqMap {
	m := FreqMap{}
	quit := make(chan int)
	channel := make(chan rune)

	var mux sync.Mutex

	go func(fm FreqMap, ch chan rune) {
		for {
			r, ok := <-ch

			if !ok {
				break
			}
			mux.Lock()
			fm[r]++
			mux.Unlock()
		}
		quit <- 0

	}(m, channel)

	var wg sync.WaitGroup

	for _, s := range strings {
		wg.Add(1)
		go func(text string, ch chan rune) {
			for _, r := range text {
				ch <- r
			}
			wg.Done()
		}(s, channel)
	}

	wg.Wait()
	close(channel)
	<-quit
	return m
}
