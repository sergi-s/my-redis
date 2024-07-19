package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func readResp(reader *bufio.Reader) (interface{}, error) {
	prefix, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	switch prefix {
	case '+':
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return strings.TrimSuffix(line, "\r\n"), nil
	case '-':
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return fmt.Errorf(strings.TrimSuffix(line, "\r\n")), nil
	case ':':
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		return strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
	case '$':
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		length, err := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
		if err != nil {
			return nil, err
		}
		if length < 0 {
			return nil, nil
		}
		data := make([]byte, length+2)
		if _, err := io.ReadFull(reader, data); err != nil {
			return nil, err
		}
		return string(data[:length]), nil
	case '*':
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		length, err := strconv.Atoi(strings.TrimSuffix(line, "\r\n"))
		if err != nil {
			return nil, err
		}
		result := make([]interface{}, length)
		for i := 0; i < length; i++ {
			result[i], err = readResp(reader)
			if err != nil {
				return nil, err
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected prefix: %c", prefix)
	}
}
