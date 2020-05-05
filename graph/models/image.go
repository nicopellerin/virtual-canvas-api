package models

// Images
type Image struct {
	ID         string  `json:"id, omitempty" bson:"id, omitempty"`
	Src        string  `json:"src, omitempty"`
	Name       string  `json:"name, omitempty"`
	Ratio      float32 `json:"ratio, omitempty"`
	Border     bool    `json:"border, omitempty"`
	Texture    bool    `json:"texture, omitempty"`
	Background bool    `json:"background, omitempty"`
	Rotate     bool    `json:"rotate, omitempty"`
	Lighting   string  `json:"lighting, omitempty"`
}
