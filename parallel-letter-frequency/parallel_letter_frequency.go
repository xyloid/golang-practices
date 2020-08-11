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
	resultChannel := make(chan FreqMap, 10)

	var mux sync.Mutex

	go func(fm FreqMap, ch chan FreqMap) {

		for mp := range ch {
			mux.Lock()
			for k, v := range mp {
				m[k] += v
			}
			mux.Unlock()
		}

		quit <- 0

	}(m, resultChannel)

	var wg sync.WaitGroup

	for _, s := range strings {
		wg.Add(1)
		go func(text string, ch chan FreqMap) {
			localFreq := FreqMap{}
			for _, r := range text {
				localFreq[r]++
			}

			ch <- localFreq
			wg.Done()
		}(s, resultChannel)
	}

	wg.Wait()
	close(resultChannel)
	<-quit
	return m
}
