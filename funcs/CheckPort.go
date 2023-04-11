package funcs

func CheckPort(port string) bool {
	for _, v := range port {
		if len(port) == 4 {
			if v >= '0' && v <= '9' {
				return true
			}
		}
	}
	return false
}

func (server *Server) CheckName(name string) bool {
	for _, user := range server.Users {
		if user == name {
			return false
		}
	}
	return true
}

func (server *Server) CheckSymbols(name string) bool {
	for _, v := range name {
		if (v >= 'a' && v <= 'z') || (v >= 'Z' && v <= 'A') {
			return true
		}
	}
	return false
}
