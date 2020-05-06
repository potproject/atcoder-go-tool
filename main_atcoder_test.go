package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

// メモリ制限
const mle = 1024 * 1000 * 1000 // 1024MB
// 実行時間制限
const tle = 2 * time.Second

// I/O
// input 入力
const input = "4\r\n10\r\n20\r\n30\r\n40"

// expectOutput 期待される出力
const expectOutput = "4 100"

func TestMain(t *testing.T) {
	tleFlag := false
	fmt.Println("/------ input ------/")
	fmt.Println(input)

	chanOutput := make(chan string)
	start := time.Now()
	go func() {
		output, _ := stubIO(input, func() {
			main()
		})
		chanOutput <- output
	}()
	actualOutput := ""
	select {
	case actualOutput = <-chanOutput:
		actualOutput = strings.TrimRight(actualOutput, "\r\n")
		fmt.Println("/------ output ------/")
		fmt.Println(actualOutput)
	case <-time.After(tle):
		tleFlag = true
	}
	fmt.Println("/------ result ------/")
	fmt.Printf("Time: %d ms\r\n", (time.Now().Sub(start)).Nanoseconds()/int64(time.Millisecond))
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("Mem: %.3f MB\r\n", float64(mem.TotalAlloc)/1000/1000)
	if tleFlag {
		fmt.Println("[TLE]")
		t.Error("Time Limit Exceeded")
		return
	}
	if mem.TotalAlloc > mle {
		fmt.Println("[MLE]")
		t.Error("Memory Limit Exceeded")
		return
	}
	if actualOutput != expectOutput {
		fmt.Println("[WA]")
		t.Errorf("Wrong Answer \r\nactual:\r\n%s\r\nexpact:\r\n%s", actualOutput, expectOutput)
		return
	}
	fmt.Println("[AC]")

}

// stubIO stubs Stdin Stdout Stderr in 'fn'.return Stdout and Stderr
func stubIO(inbuf string, fn func()) (string, string) {
	inr, inw, _ := os.Pipe()
	outr, outw, _ := os.Pipe()
	errr, errw, _ := os.Pipe()
	orgStdin := os.Stdin
	orgStdout := os.Stdout
	orgStderr := os.Stderr
	inw.Write([]byte(inbuf))
	inw.Close()
	os.Stdin = inr
	os.Stdout = outw
	os.Stderr = errw
	fn()
	os.Stdin = orgStdin
	os.Stdout = orgStdout
	os.Stderr = orgStderr
	outw.Close()
	outbuf, _ := ioutil.ReadAll(outr)
	errw.Close()
	errbuf, _ := ioutil.ReadAll(errr)

	return string(outbuf), string(errbuf)

}
