package transform

type Attachments struct {
	Data   []Attachment `json:"data"`
	Paging Paging       `json:"paging"`
}

type Attachment struct {
	ID        string `json:"string"`
	Name      string `json:"name"`
	MimeType  string `json:"mime_type"`
	Size      uint64 `json:"size"`
	FileUrl   string `json:"file_url,omitempty"`
	VideoData Video  `json:"video_data,omitempty"`
}

type Video struct {
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Length     int    `json:"length"`
	VideoType  int    `json:"video_type"`
	PreviewUrl string `json:"preview_url"`
	Url        string `json:"url"`
	Rotation   int    `json:"rotation"`
}

type Paging struct {
	Next    string  `json:"next"`
	Cursors Cursors `json:"cursors"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}
