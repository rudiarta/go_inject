package service

import "test/rudi/repo"

type Service interface {
	Insert(string) string
}

func NewService(r repo.Repository) Service {
	return Ser{r}
}

type Ser struct {
	r repo.Repository
}

func (s Ser) Insert(data string) string {
	return data + "av" + s.r.Get()
}
