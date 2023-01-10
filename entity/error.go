package entity

type TwitterError struct {
	Message   string `json:"message"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations"`
	Path       []interface{} `json:"path"`
	Extensions struct {
		Name    string `json:"name"`
		Source  string `json:"source"`
		Code    int    `json:"code"`
		Kind    string `json:"kind"`
		Tracing struct {
			TraceId string `json:"trace_id"`
		} `json:"tracing"`
	} `json:"extensions"`
	Code    int    `json:"code"`
	Kind    string `json:"kind"`
	Name    string `json:"name"`
	Source  string `json:"source"`
	Tracing struct {
		TraceId string `json:"trace_id"`
	} `json:"tracing"`
}
