package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	// Initialize AOF
	if err := initAOF("appendonly.aof"); err != nil {
		fmt.Println("Error initializing AOF:", err)
		return
	}
	defer aofFile.Close()

	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error setting up the server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server is running on port 6379...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		data, err := readResp(reader)
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}

		response := handleCommand(data)
		if err := writeResp(writer, response); err != nil {
			fmt.Println("Error writing response:", err)
			return
		}
	}
}

func handleCommand(data interface{}) interface{} {
	if cmd, ok := data.([]interface{}); ok && len(cmd) > 0 {
		switch strings.ToLower(cmd[0].(string)) {
		case "ping":
			if len(cmd) > 1 {
				return cmd[1]
			}
			return "PONG"
		case "echo":
			if len(cmd) > 1 {
				return cmd[1]
			}
			return nil
		}

		// Append the command to AOF
		appendToAOF(cmd)
	}
	return fmt.Errorf("unknown command")
}
