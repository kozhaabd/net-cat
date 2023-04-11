package funcs

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func (server *Server) HandleConn(conn net.Conn) {
	// go clientWriter(conn, ch)
	Read(conn, "welcome.txt")
	scanner := bufio.NewScanner(conn)
	username, err := server.Name(conn, scanner)
	if err != nil {
		fmt.Println("error in the username", err)
		return
	}
	if len(server.Users) > 9 {
		_, err := fmt.Fprintln(conn, "chat is full")
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	server.Lock()
	server.Users[conn] = username
	server.Unlock()

	Read(conn, "history.txt")
	Sprintf := fmt.Sprintf("[%s]:[%s]", time.Now().Format("2006-01-02 12:01:02"), username)
	_, err = fmt.Fprint(conn, Sprintf)
	if err != nil {
		log.Println(err)
	}
	join := username + " has joined our chat"
	server.Msg <- Message{Text: join, Name: username}
	for scanner.Scan() {
		if !server.ValidString(scanner.Text()) {
			_, err = fmt.Fprint(conn, "Unacceptable Text\n")
			if err != nil {
				log.Println(err)
			}
		}
		_, err = fmt.Fprint(conn, Sprintf)
		if err != nil {
			log.Fatal(err)
			return
		}

		text := fmt.Sprintf("[%s]:[%s]:%s", time.Now().Format("2006-01-02 15:04:05"), username, scanner.Text())
		// send the message

		server.Msg <- Message{Text: text, Name: username}
		if len(scanner.Text()) > 0 {
			server.History.WriteString(text + "\n")
		}

	}
	left := username + " has left our chat"
	server.Msg <- Message{Text: left, Name: username}
	server.Lock()
	delete(server.Users, conn)
	server.Unlock()
}

// func clientWriter(conn net.Conn, ch <-chan string) {
// 	for msg := range ch {
// 		fmt.Fprintln(conn, msg) // Примечание: игнорируем ошибки сети
// 	}
// }

// messages <- username + " подключился"
// entering <- ch
// input := bufio.NewScanner(conn)
// for input.Scan() {
// 	messages <- username + ": " + input.Text()
// }
// // Примечание: игнорируем потенциальные ошибки input.Err()
// leaving <- ch
// messages <- username + " отключился"
// conn.Close()
