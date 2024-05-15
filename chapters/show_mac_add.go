package main

import (
    "fmt"
    "net"
)

func main() {
    interfaces, err := net.Interfaces()
    if err != nil {
        fmt.Print(err)
        return
    }

    for _, i := range interfaces {
        if i.HardwareAddr != nil {
            fmt.Println(i.Name, i.HardwareAddr)
        }
    }
}