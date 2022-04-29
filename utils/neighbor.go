package utils

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
)

func IsFoundHost(host string, port uint16) bool {
	target := fmt.Sprintf("%s:%d", host, port)
	// 受け取ったhostとportのアドレスに対して指定のプロトコルでアクセスできるか確認
	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		fmt.Printf("%s %v\n", target, err)
		return false
	}
	return true
}

// 192.168.0.10:5000
// 192.168.0.11:5000
// 192.168.0.12:5000

// 192.168.0.10:5001
// 192.168.0.10:5002
// 192.168.0.10:5003

var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func FindNeighbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}
	// 192.168.0.10の192.168.0を取り出す
	prefixHost := m[1]
	// 192.168.0.10の10を取り出す
	lastIp, _ := strconv.Atoi(m[len(m)-1])
	neighbors := make([]string, 0)

	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			// IPアドレスを作成
			guessHost := fmt.Sprintf("%s%d", prefixHost, lastIp+int(ip))
			// IPアドレスにポートを追加
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if guessTarget != address && IsFoundHost(guessHost, port) {
				// 対象のターゲットが自分のアドレスと同じではない、かつ、対象のホストが存在する場合
				neighbors = append(neighbors, guessTarget)
			}
		}
	}
	return neighbors

}

func GetHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "127.0.0.1"
	}
	fmt.Println(hostname)
	address, err := net.LookupHost(hostname)
	if err != nil {
		return "127.0.0.1"
	}
	// fmt.Println(address)
	return address[0]
}
