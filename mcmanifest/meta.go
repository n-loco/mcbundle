package mcmanifest

type Meta struct {
	Authors []string `json:"authors,omitempty,omitzero"`
	License string   `json:"license,omitempty,omitzero"`
}

func (meta *Meta) IsZero() bool {
	return (len(meta.Authors) != 0) && (len(meta.License) != 0)
}
