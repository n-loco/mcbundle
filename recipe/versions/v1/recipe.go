package v1

type Recipe struct {
	Config  *Config  `json:"config"`
	Header  *Header  `json:"header"`
	Modules []Module `json:"modules"`
	Meta    *Meta    `json:"meta"`
}
