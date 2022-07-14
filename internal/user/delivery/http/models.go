package http

// naming: Inp or Request
type loginInp struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type loginOut struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type registerInp struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type registerOut struct {
	Message string `json:"message"`
}
