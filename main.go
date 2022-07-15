package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"github.com/alteamc/minequery/ping"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func chunks(xs []string, chunkSize int) [][]string {
    if len(xs) == 0 {
        return nil
    }
    divided := make([][]string, (len(xs)+chunkSize-1)/chunkSize)
    prev := 0
    i := 0
    till := len(xs) - chunkSize
    for prev < till {
        next := prev + chunkSize
        divided[i] = xs[prev:next]
        prev = next
        i++
    }
    divided[i] = xs[prev:]
    return divided
}

func checkServer(ip string, ipNum int, file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := ping.Ping(ip, 25565)
	if (err == nil) {
		if motd, ok := resp.Description.(map[string]interface{}); ok {
			file.WriteString(fmt.Sprintf(`{"ip":"%s","motd":"%s","online":%d}` + "\n", ip, motd["text"], resp.Players.Online))
		}
	}
}

func main() {
	fmt.Println("beginning scan...")
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(err)
	// read ips.txt and split by newline
	content, err := os.ReadFile("ips.txt")
	
	checkErr(err)
	ips := strings.Split(string(content), "\n")

	numCPU := runtime.NumCPU()
	chunkSize := (len(ips) + numCPU - 1) / numCPU

	ipsChunk := chunks(ips, chunkSize)

	// iterate over ips
	for index, chunk := range ipsChunk {
		// create a new waitgroup
		var wg sync.WaitGroup
		// print chunk number starting
		fmt.Println("chunk", index, "starting")
		// add one to the waitgroup
		fmt.Println("chunk length: ", len(chunk))
		wg.Add(len(chunk))
		// iterate over chunk
		for ipNum, ip := range chunk {
			// create a goroutine for each ip
			go checkServer(ip, ipNum, f, &wg)
		}
		// wait for the goroutine to finish
		wg.Wait()
		// print chunk number ending
		fmt.Println("chunk", index, "ending")
	}

	fmt.Println("scan complete")
}
