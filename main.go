package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/alteamc/minequery/ping"
	"github.com/korovkin/limiter"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkServer(ip string, ipNum int, file *os.File) {
	resp, err := ping.Ping(ip, 25565)
	if (err == nil) {
		if motd, ok := resp.Description.(map[string]interface{}); ok {
			file.WriteString(fmt.Sprintf(`{"ip":"%s","motd":"%s","online":%d}` + "\n", ip, motd["text"], resp.Players.Online))
		}
	}
}

func main() {
	fmt.Println("beginning scan...")
	limit := limiter.NewConcurrencyLimiter(10000) // set this to the number of concurrent requests you want to make
	f, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkErr(err)
	// read ips.txt and split by newline
	content, err := os.ReadFile("ips.txt")
	
	checkErr(err)
	ips := strings.Split(string(content), "\n")



	// iterate over ips
	for index, chunk := range ips {
		limit.Execute(func() {
			checkServer(chunk, index, f)
		})
	}
	defer limit.WaitAndClose()
	fmt.Println("scan complete")
}
