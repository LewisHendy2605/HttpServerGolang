package main

func main() {
	println("hello")
}

/*
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
				println("Context done called")
				return
			default:
				buffer := make([]byte, 1024)

				bytes_read, err := conn.Read(buffer)
				if err == io.EOF {
					println("EOF")
					return
				} else if err != nil {
					log.Fatalf("error reading from conn: %v", err)
				}

				// Grab whats been read
				buffer = buffer[:bytes_read]

				fmt.Printf("Read Loop: %s\n", string(buffer))

				// Look for new line
				if i := bytes.IndexByte(buffer, '\n'); i != -1 {
					// Grab up to new line
					str += string(buffer[:i])

					fmt.Printf("Found new line: %s\n", str)

					// Pass out line
					out <- strings.Trim(str, parser.CR)

					// Reset
					buffer = buffer[i+1:]
					str = ""

					fmt.Printf("Reset Buffer: %s\n", string(buffer))
				} else {
					fmt.Printf("Stringing up: %s\n", string(buffer))
					// Store in string
					str += string(buffer)
				}
			}
		}
	}()

	return out
}
*/
