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
	BuyLink    string  `json:"buyLink, omitempty"`
	Price      float64 `json:"price, omitempty"`
}

type UpdateArtworkInput struct {
	ID         string   `json:"id"`
	Src        string   `json:"src"`
	Name       string   `json:"name"`
	Ratio      float64  `json:"ratio"`
	Border     bool     `json:"border"`
	Texture    bool     `json:"texture"`
	Background bool     `json:"background"`
	Rotate     bool     `json:"rotate"`
	Lighting   int      `json:"lighting"`
	Username   string   `json:"username"`
	Price      *float64 `json:"price"`
	BuyLink    *string  `json:"buyLink"`
}

type AddArtworkInput struct {
	ID         string  `json:"id"`
	Src        string  `json:"src"`
	Name       string  `json:"name"`
	Ratio      float64 `json:"ratio"`
	Border     bool    `json:"border"`
	Texture    bool    `json:"texture"`
	Background bool    `json:"background"`
	Rotate     bool    `json:"rotate"`
	Lighting   int     `json:"lighting"`
	Username   string  `json:"username"`
	BuyLink    string  `json:"buyLink"`
	Price      float64 `json:"price"`
}

type DeleteArtworkInput struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}
