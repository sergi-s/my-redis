package main

import (
	"bufio"
	"fmt"
)

func writeResp(writer *bufio.Writer, data interface{}) error {
	switch v := data.(type) {
	case string:
		if _, err := writer.WriteString(fmt.Sprintf("+%s\r\n", v)); err != nil {
			return err
		}
	case error:
		if _, err := writer.WriteString(fmt.Sprintf("-%s\r\n", v.Error())); err != nil {
			return err
		}
	case int:
		if _, err := writer.WriteString(fmt.Sprintf(":%d\r\n", v)); err != nil {
			return err
		}
	case []byte:
		if _, err := writer.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)); err != nil {
			return err
		}
	case nil:
		if _, err := writer.WriteString("$-1\r\n"); err != nil {
			return err
		}
	case []interface{}:
		if _, err := writer.WriteString(fmt.Sprintf("*%d\r\n", len(v))); err != nil {
			return err
		}
		for _, item := range v {
			if err := writeResp(writer, item); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
	return writer.Flush()
}
