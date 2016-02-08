package main

import (
        "net"
        "fmt"
)

func main () {
        txt, err := net.LookupTXT("rs.dns-oarc.net")
        if err != nil {
                fmt.Println(err)
        }
        fmt.Printf("%v", txt)
}
