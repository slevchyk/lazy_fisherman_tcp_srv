package models

type BoardVersion struct {
	Id          int64    `json:"id"`
	PcbType     string `json:"pcb_type"`
	Cpu         string `json:"cpu"`
	Description string `json:"description"`
}
