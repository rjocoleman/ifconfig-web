# ifconfig-web

A very simple webserver written in go to output all non-loopback IP addresses.
Like a simple ifconfig over HTTP. Listens on port 8000.


## Usage

`go run ./ifconfig.go`

or compile it for whatever platform/arch you wish and run it there.


## Attribution

IP address detection: https://github.com/mccoyst/myip
