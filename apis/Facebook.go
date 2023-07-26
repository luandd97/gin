package apis

// API Get FaceBook Pages

type FacebookPageResponse struct {
	Data []FacebookPage `json:"data"`
}

type FacebookPage struct {
	Name        string  `json:"name"`
	ID          string  `json:"id"`
	AccessToken string  `json:"access_token"`
	Picture     Picture `json:"picture"`
}

type Picture struct {
	Data Data `json:"data"`
}

type Data struct {
	Height       int    `json:"height"`
	IsSilhouette bool   `json:"is_silhouette"`
	Url          string `json:"url"`
	Width        int    `json:"width"`
}

// API Get FaceBook Page Conversation
type FacebookPageConversation struct {
	Data []Conversation `json:"data"`
}

type Conversation struct {
	ID           string  `json:"id"`
	Link         string  `json:"link"`
	Snippet      string  `json:"snippet"`
	UnreadCount  uint    `json:"unread_count"`
	UpdatedTime  string  `json:"updated_time"`
	Senders      Senders `json:"senders"`
	MessageCount uint64  `json:"message_count"`
}

type Senders struct {
	Data []SenderData `json:"data"`
}

type SenderData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
}

// API Get FaceBook Page Conversation Message ( Convo Detail)

type FBConvoMessages struct {
	Data   []Messages `json:"data"`
	Paging Paging     `json:"paging"`
}

type Messages struct {
	ID          string `json:"id"`
	CreatedTime string `json:"created_time"`
}

type Paging struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

// API Get FaceBook Page Conversation Message Deatail

type Message struct {
	ID          string `json:"id"`
	CreatedTime string `json:"created_time"`
	Message     string `json:"message"`
	From        From   `json:"from"`
	To          To     `json:"to"`
}

type From struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
}

type To struct {
	Data []UserInfo `json:"data"`
}

// Global
type ID struct {
	ID string `json:"id"`
}

type Name struct {
	Name string `json:"name"`
}

type Email struct {
	Email string `json:"email"`
}

type UserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    string `json:"id"`
}
	