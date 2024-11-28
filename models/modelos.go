package models

type Cliente struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome" binding:"required,min=4"`
	Cpf   string `json:"cpf" binding:"required,len=14"`
	Senha string `json:"senha" binding:"required,min=6,containsany=!@#$%"`
}

type Produtos struct {
	ID        int     `json:"id"`
	Nome      string  `json:"nome"`
	Descricao string  `json:"descricao"`
	Valor     float64 `json:"valor"`
}
