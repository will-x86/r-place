package models

type Canvas struct {
	X   int    `json:"x"`
	Y   int    `json:"y"`
	Hex string `json:"hex,omitempty"`
}
