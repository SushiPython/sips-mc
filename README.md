# sips-mc

sips-mc is a concurrent Minecraft server scanner written in Go. It takes a list of IP addresses, checks whether each one is running a Minecraft server, and writes the results to an output file.

## What it does

* Scans large IP lists for Minecraft server status
* Uses goroutines to run checks concurrently
* Splits IPs into chunks for faster scanning
* Writes discovered server results to `output.txt`
* Built as a simple CLI-style networking tool

## Tech Stack

* Go
* TCP networking
* File I/O
* Concurrency

## How to use

Create an `ips.txt` file with one IP address per line:

```txt
1.2.3.4
5.6.7.8
```

Then run:

```bash
go run main.go
```

Results are written to:

```txt
output.txt
```
