package entity

type Entry struct {
	EntryId   string `json:"entryId"`
	SortIndex string `json:"sortIndex"`
	Content   struct {
		EntryType string `json:"entryType"`
	} `json:"content"`
}
