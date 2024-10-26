package letter

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
	l := len(texts)
	freqChan := make(chan FreqMap, l)
	for _, t := range texts {
		go countFrequency(freqChan, t)
	}

	fm := FreqMap{}
	for range texts {
		for r, c := range <-freqChan {
			fm[r] += c
		}
	}

	return fm
}

func countFrequency(ch chan<- FreqMap, text string) {
	ch <- Frequency(text)
}
