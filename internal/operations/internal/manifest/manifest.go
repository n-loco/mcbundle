package manifest

type Manifest struct {
	FormatVersion uint8        `json:"format_version"`
	Header        Header       `json:"header"`
	Modules       []Module     `json:"modules"`
	Dependencies  []Dependency `json:"dependencies,omitempty,omitzero"`
	Meta          Meta         `json:"meta,omitempty,omitzero"`
}
