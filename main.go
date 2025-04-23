package main

import (
    "log"
    "net"
)

func main() {
    // Roon 使用的 Multicast 訊息
    addr, err := net.ResolveUDPAddr("udp", "239.255.90.90:9003")
    if err != nil {
        log.Fatalf("無法解析 multicast 位址: %v", err)
    }

    conn, err := net.ListenMulticastUDP("udp", nil, addr)
    if err != nil {
        log.Fatalf("無法監聽 multicast: %v", err)
    }
    defer conn.Close()

    buf := make([]byte, 2048)
    for {
        n, src, err := conn.ReadFromUDP(buf)
        if err != nil {
            log.Printf("讀取失敗: %v", err)
            continue
        }

        log.Printf("收到來自 %v 的封包 (%d bytes)", src, n)

        // 廣播到 255.255.255.255:9003
        bcastAddr, _ := net.ResolveUDPAddr("udp", "255.255.255.255:9003")
        outConn, _ := net.DialUDP("udp", nil, bcastAddr)
        outConn.Write(buf[:n])
        outConn.Close()
    }
}
