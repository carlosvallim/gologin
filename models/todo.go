package models

//Todo - estrutura do todo
type Todo struct {
	ID      int    `json:"id" db:"id"`
	Text    string `json:"text" db:"text"`
	Done    bool   `json:"done" db:"done"`
	Usuario *Usuario
}
