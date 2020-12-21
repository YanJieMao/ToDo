package reso

// GetUser GET "/user" response object
type GetUser struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Gender   int64  `json:"gender"`
	Age      int64  `json:"age"`
}

// PutUser PUT "/user" response object
type PutUser struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
	Gender   int64  `json:"gender"`
	Age      int64  `json:"age"`
}

// PostUser POST "/user" response object
type PostUser struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
}

// PostLogin POST "/login" response object
type PostLogin struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
	Token    string `json:"token"`
}
