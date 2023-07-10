package model

type Item struct {
	Id    int `bun:",pk,autoincrement"`
	Name  string
	Price string
}
