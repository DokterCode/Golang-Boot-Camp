package models

// Bentuk data mahasiswa
type Mahasiswa struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	NIM  string `json:"nim"`
}
