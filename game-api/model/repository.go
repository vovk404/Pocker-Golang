package model

type Repository interface {
	GetOne(id int) (*Game, error)
	Update(game *Game) error
	DeleteByID(id int) error
	SaveGame(game *Game) (int, error)
}