package entity

import (
	"errors"
	"time"
)

type Game struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"createdat"`
	UpdatedAt   time.Time `json:"updatedat"`
}

// Criando o Construtor
func NewGame() *Game {
	return &Game{}
}

// Criando um método para o meu struct(struc é como se fosse uma classe) chamado IsValid
func (g *Game) IsValid() error {
	if g.Name == "" {
		return errors.New("Preencha o nome do jogo")
	}
	if g.Description == "" {
		return errors.New("Preencha a descrição do jogo")
	}
	return nil
}
