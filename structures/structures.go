package structures

type NumerologyQuery struct {
	Word string `json:"Word"`
}

//TODO fix the client to pass real json so that this can be a uint32, not a string with later explicit conversion
type TexterologyQuery struct {
	Number string `json:"Number"`
	Length string `json:"Length"`
}
