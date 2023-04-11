package funcs

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

func (server *Server) Broadcaster() {
	for {
		select {
		case msg := <-server.Msg:
			server.Lock()
			for conn, name := range server.Users {
				if name != msg.Name {
					_, err := fmt.Fprintln(conn, "\n"+msg.Text)
					if err != nil {
						log.Println(err)
					}
				} else {
					continue
				}
				_, err := fmt.Fprint(conn, fmt.Sprintf("[%s]:[%s]", time.Now().Format("2000-01-02 12:01:02"), name))
				if err != nil {
					log.Println(err)
				}
			}
			server.Unlock()
		case left := <-server.Msg:
			server.Lock()
			for conn, name := range server.Users {
				if name != left.Name {
					_, err := fmt.Fprintln(conn, "\n"+left.Text)
					if err != nil {
						log.Println(err)
					}
				} else {
					continue
				}
				_, err := fmt.Fprint(conn, fmt.Sprintf("[%s]:[%s]:", time.Now().Format("2006-01-02 15:04:05"), name))
				if err != nil {
					log.Println(err)
				}
			}
			server.Unlock()
		case join := <-server.Msg:
			server.Lock()
			for conn, name := range server.Users {
				if name != join.Name {
					_, err := fmt.Fprintln(conn, "\n"+join.Text)
					if err != nil {
						log.Println(err)
					}
				} else {
					continue
				}
				_, err := fmt.Fprint(conn, fmt.Sprintf("[%s]:[%s]:", time.Now().Format("2006-01-02 15:04:05"), name))
				if err != nil {
					log.Println(err)
				}
			}
			server.Unlock()
		}
	}
}

func (server *Server) ValidString(s string) bool {
	s = strings.TrimSuffix(s, "\n")
	rxmsg := regexp.MustCompile("^[\u0400-\u04FF\u0020-\u007F]+$")
	if !rxmsg.MatchString(s) {
		return false
	}
	return true
}

// clients := make(map[client]bool) // Все подключенные клиенты
// for {
// 	select {
// 	case msg := <-server.Msg:
// 		// Широковещательно евходящее сообщение во все
// 		// каналы исходящих сообщений для клиентов,
// 		for conn, name := range server.Users {
// 			if name != msg.Name {
// 				_, err := fmt.Fprintln(conn, "\n"+msg.Text)
// 				if err != nil {
// 					log.Println(err)
// 				}
// 			} else {
// 				continue
// 			}
// 			_, err := fmt.Fprint(conn, fmt.Sprintf("[%s]:[%s]:", time.Now().Format("2006-01-02 15:04:05"), name))
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		}
// 	case cli := <-entering:
// 		clients[cli] = true
// 	case cli := <-leaving:
// 		delete(clients, cli)
// 		close(cli)
// 	}
// }
