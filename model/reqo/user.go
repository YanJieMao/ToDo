package reqo

// GetUser GET "/user" request object
type GetUser struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
}

// PutUser PUT "/user" request object
type PutUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Gender   int64  `json:"gender"`
	Age      int64  `json:"age"`
}

// PostUser POST "/user" request object
type PostUser struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Gender   int64  `json:"gender"`
	Age      int64  `json:"age"`
}

// PostLogin POST "/login" request object
type PostLogin struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}
