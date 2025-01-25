package comfortablefuncs

type Message struct {
	ID   int    `json:"ID"`
	Text string `json:"text"`
}

type Responce struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}
