package main

import "github.com/latipovsharif/wallet/api"

func main() {
	s := api.Server{}
	if err := s.Run(); err != nil {
		panic(err)
	}
}
