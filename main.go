package main

import (
	"fmt"

	"test/rudi/repo"
	"test/rudi/service"

	"github.com/karlkfi/inject"
)

func main() {
	var (
		r repo.Repository
		s service.Service
	)

	graph := inject.NewGraph()

	graph.Define(&r, inject.NewProvider(repo.NewRepository))
	graph.Resolve(&r)

	graph.Define(&s, inject.NewProvider(service.NewService, &r))
	graph.Resolve(&s)

	fmt.Println(r.Insert("Aa"))
	fmt.Println(s.Insert("rudi"))
}
