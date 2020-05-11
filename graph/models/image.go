package models

// Image - Structure of single image
type Image struct {
	ID         string  `json:"id, omitempty" bson:"id, omitempty"`
	Src        string  `json:"src, omitempty"`
	Name       string  `json:"name, omitempty"`
	Ratio      float64 `json:"ratio, omitempty"`
	Border     bool    `json:"border, omitempty"`
	Texture    bool    `json:"texture, omitempty"`
	Background bool    `json:"background, omitempty"`
	Rotate     bool    `json:"rotate, omitempty"`
	Lighting   int     `json:"lighting, omitempty"`
}
