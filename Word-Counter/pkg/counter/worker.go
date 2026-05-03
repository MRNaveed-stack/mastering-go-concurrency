package counter

import (
	"bufio"
	"os"
	"strings"
	"sync"
)

func FileWorker(id int, paths <-chan string, stats *ResultStats, wg *sync.WaitGroup) {
	defer wg.Done()
	for path := range paths {
		file, err := os.Open(path)
		if err != nil {
			continue
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			words := strings.Fields(line)
			for _, w := range words {
				word := strings.ToLower(strings.Trim(w, ",.!?:;\""))
				stats.Update(word, 1)
			}
		}
		file.Close()
	}
}
