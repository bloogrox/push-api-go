package pushapi

// Subscriber ...
type Subscriber struct {
	Token    string `json:"token"`
	UserID   string `json:"user_id"`
	SenderID string `json:"sender_id"`
	Subtool  string `json:"subtool"`
}

// Message ...
type Message struct {
	Title       string `json:"title"`
	Body        string `json:"body"`
	Icon        string `json:"icon"`
	ClickAction string `json:"click_action"`
}

// Postback ...
type Postback struct {
	URL   string `json:"url"`
	Event string `json:"event"`
}

// Push ...
type Push struct {
	Subscriber *Subscriber `json:"subscriber"`
	Message    *Message    `json:"message"`
	Postback   []Postback  `json:"postbacks"`
}
