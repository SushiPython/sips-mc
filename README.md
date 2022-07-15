# sips-mc
Sushi's Internet Protocol Scanner (for Minecraft)

## Info
This is an ip scanner I wrote in go to find the status of minecraft servers given a large list of IPs. 
It divides the IPs into chunks and concurrently scans the servers and prints the output to a file.
- This should not be used on your home wifi because it may make your ISP angry.
- This project is not designed for scanning the entire internet (it will be too slow). For that, try masscan.

## Usage
- Create a file "ips.txt" and "output.txt".
- "ips.txt" should be a list of IPs seperated by newlines.
- "output.txt" should be blank because that is where the program will write to.

Note: this does not fully work yet and is not tested on all operating systems.
