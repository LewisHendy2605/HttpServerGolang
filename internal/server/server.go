package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
)

// Starts tcp server
func StartServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error accepting connection: %v", err)
		}

		for lines := range ReadLines(conn, context.Background()) {
			fmt.Println(lines)
		}
	}
}

// Returns lines from a reader/closer
func ReadLines(conn io.ReadCloser, ctx context.Context) <-chan string {
	out := make(chan string)

	go func() {
		defer conn.Close()
		defer close(out)

		str := ""
		for {
			select {
			case <-ctx.Done():
				return
			default:
				buffer := make([]byte, 8)

				bytes_read, err := conn.Read(buffer)
				if err == io.EOF {
					return
				} else if err != nil {
					log.Fatalf("error reading from conn: %v", err)
				}

				// Grab whats been read
				buffer = buffer[:bytes_read]

				// Look for new line
				if i := bytes.IndexByte(buffer, '\n'); i != -1 {
					// Grab up to new line
					str += string(buffer[:i])

					// Pass out line
					out <- str

					// Reset
					buffer = buffer[i+1:]
					str = ""
				} else {
					// Store in string
					str += string(buffer)
				}
			}

			if len(str) != 0 {
				out <- str
			}
		}
	}()

	return out
}
