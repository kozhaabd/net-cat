package funcs

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func (server *Server) Name(conn net.Conn, scanner *bufio.Scanner) (string, error) {
	var Username string
	for {
		_, err := fmt.Fprintf(conn, "[ENTER YOUR NAME]:")
		if err != nil {
			return "", err
		}
		if scanner.Scan() {
			if len(strings.TrimSpace(scanner.Text())) == 0 || !server.CheckName(scanner.Text()) || !server.CheckSymbols(scanner.Text()) || !server.ValidString(scanner.Text()) {
				_, err := fmt.Fprintln(conn, "Error: Illegal name, try to choose another one")
				if err != nil {
					return "", err
				}
				continue
			}
			Username = strings.TrimSpace(scanner.Text())
			break
		}
	}
	return Username, nil
}
