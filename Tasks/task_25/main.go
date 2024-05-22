package main

import (
	"fmt"
	"runtime"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func Sleep(milliseconds int) {
	start := currentMillis()
	for {
		if currentMillis()-start >= int64(milliseconds) {
			break
		}
		runtime.Gosched() // Позволяет другим горутинам выполняться
	}
}

func currentMillis() int64 {
	return time.Now().UnixNano() / 1e6
}

func main() {
	fmt.Println("Start sleep")
	start := time.Now()

	Sleep(2000)            // ввод времени задержки в мс
	sleep(2 * time.Second) // ввод времени в с
	fmt.Println(time.Since(start))
	fmt.Println("End sleep")
}
