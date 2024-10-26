package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	l := 0
	for _, text := range texts {
		l += len(text)
	}
	freqChan := make(chan rune, l)
	wg := sync.WaitGroup{}
	wg.Add(len(texts))
	for _, t := range texts {
		go countFrequency(&wg, freqChan, t)
	}

	go func() {
		wg.Wait()
		close(freqChan)
	}()

	fm := FreqMap{}
	for c := range freqChan {
		fm[c]++
	}

	return fm
}

func countFrequency(wg *sync.WaitGroup, ch chan<- rune, text string) {
	defer wg.Done()
	for _, c := range text {
		ch <- c
	}
}
