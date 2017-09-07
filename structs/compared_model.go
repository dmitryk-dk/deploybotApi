package structs

type ComparedObject struct {
	Meta    Meta          `json:"meta"`
	Entries []interface{} `json:"entries,omitempty"`
}
