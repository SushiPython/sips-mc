package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/alteamc/minequery/ping"
	"sync"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkServer(ip string, file *os.File, wg *sync.WaitGroup) {
	resp, err := ping.Ping(ip, 25565)
	if (err == nil) {
		if motd, ok := resp.Description.(map[string]interface{}); ok {
			file.WriteString(fmt.Sprintf(`{"ip":"%s","motd":"%s","online":%d}` + "\n", ip, motd["text"], resp.Players.Online))
		}
	}
	wg.Done()
}

func main() {
	fmt.Println("beginning scan...")
	var wg sync.WaitGroup
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(err)
	// read ips.txt and split by newline
	content, err := os.ReadFile("ips.txt")
	
	checkErr(err)
	ips := strings.Split(string(content), "\n")



	// iterate over ips
	for _, ip := range ips {
		wg.Add(1)
		go checkServer(ip, f, &wg)
	}
	wg.Wait()
	fmt.Println("scan complete")
}
