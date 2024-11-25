package models

type Cliente struct {
	ID    int
	Nome  string `json:"nome" binding:"required,min=4"`
	Cpf   string `json:"cpf" binding:"required,len=14"` //mudando para 14 para nao ter problema com a mascara do front
	Senha string `json:"senha" binding:"required,min=6,containsany=!@#$%"`
}

type Produtos struct {
	ID         int
	Nome       string  `json:"nome"`
	Valor      float64 `json:"valor"`
	Quantidade int     `json:"quantidade"`
}
