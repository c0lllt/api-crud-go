package models

type Cliente struct {
	ID    int
	Nome  string `json:"nome" binding:"required,min=4"`
	Cpf   string `json:"cpf" binding:"required,len=11"`
	Senha string `json:"senha" binding:"required,min=6,containsany=!@#$%"`
}

type Produtos struct {
	ID      int    `json:"id"`
	Nome    string `json:"nome"`
	Preco   int16  `json:"preco"`
	Estoque int    `json:"estoque"`
}