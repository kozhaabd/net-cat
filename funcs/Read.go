package funcs

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func Read(conn net.Conn, str string) {
	file, err := os.Open(str)
	if err != nil {
		fmt.Fprintln(conn, "Welcome to TCP-Chat!")
		fmt.Println("Error: Can not read Welcome.txt")
	}
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Fprint(conn, line)
	}
}
