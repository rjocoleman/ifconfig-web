package main

import (
  "net"
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  addrs, _ := net.InterfaceAddrs()

  for _, a := range addrs {
    if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
      fmt.Fprintln(w, ipnet.IP.String())
    }
  }
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8000", nil)
}
