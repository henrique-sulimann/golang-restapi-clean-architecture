package entity

import (
	"errors"
	"time"
)

type User struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

// Criando o Construtor
func NewUser() *User {
	return &User{}
}

// Criando um método para o meu struct(struc é como se fosse uma classe) chamado IsValid
func (g *User) IsValid() error {
	if g.Name == "" {
		return errors.New("Preencha o nome do jogo")
	}
	if g.Password == "" {
		return errors.New("Preencha a senha do usuário")
	}
	return nil
}
