package main

import "os"

// io 操作
func main() {
	var buf [16]byte
	_, _ = os.Stdin.Read(buf[:])
	_, _ = os.Stdin.WriteString(string(buf[:]))
}
