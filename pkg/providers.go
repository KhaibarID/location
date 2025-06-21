package pkg

import (
	"github.com/khaibarid/location/internal/repository"
	"go.uber.org/dig"
)

/**
* Created by Muhammad Muflih Kholidin
* at 2020-10-01 22:19:43
* https://github.com/mmuflih
* muflic.24@gmail.com
**/

func BuildProvider(c *dig.Container) *dig.Container {
	var uc []interface{}

	uc = append(uc, repository.NewSubDistrictRepo)
	uc = append(uc, repository.NewDistrictRepo)
	uc = append(uc, repository.NewRegencyRepo)
	uc = append(uc, repository.NewProvinceRepo)

	uc = append(uc, NewLocationUsecase)

	for _, usecase := range uc {
		if err := c.Provide(usecase); err != nil {
			panic(err)
		}
	}
	return c
}
