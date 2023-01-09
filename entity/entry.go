package entity

type Entry struct {
	Content struct {
		EntryType string `json:"entryType"`
	} `json:"content"`
}
