package pkg

import (
	"errors"
	"strings"

	"github.com/khaibarid/location/internal/domain"
	"github.com/khaibarid/location/internal/repository"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-10-01 22:19:43
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type LocationUsecase interface {
	GetByCode(code string) (*domain.Location, error)
}

type locUsecase struct {
	districtRepo    repository.DistrictRepository
	subDistrictRepo repository.SubDistrictRepository
	regencyRepo     repository.RegencyRepository
	provinceRepo    repository.ProvinceRepository
}

func NewLocationUsecase(
	districtRepo repository.DistrictRepository,
	subDistrictRepo repository.SubDistrictRepository,
	regencyRepo repository.RegencyRepository,
	provinceRepo repository.ProvinceRepository,
) LocationUsecase {
	return &locUsecase{
		districtRepo:    districtRepo,
		subDistrictRepo: subDistrictRepo,
		regencyRepo:     regencyRepo,
		provinceRepo:    provinceRepo,
	}
}

func (u *locUsecase) GetByCode(code string) (*domain.Location, error) {
	loc := &domain.Location{}
	codes := strings.Split(code, ".")
	if len(codes) < 4 {
		return nil, errors.New("invalid code")
	}

	subDistrict, _ := u.subDistrictRepo.FindByCode(code)
	if subDistrict != nil {
		loc.SubDistrict = subDistrict.Name
		loc.SubDistrictID = subDistrict.Code
		loc.District = subDistrict.District
		loc.Province = subDistrict.Province
		loc.PostCode = subDistrict.PostalCode
	}

	district, _ := u.districtRepo.FindByCode(codes[0] + "." + codes[1] + "." + codes[2])
	if district != nil {
		loc.District = district.Name
		loc.DistrictID = district.Code
	}

	regency, _ := u.regencyRepo.FindByCode(codes[0] + "." + codes[1])
	if regency != nil {
		loc.City = regency.Name
		loc.CityID = regency.Code
	}

	province, _ := u.provinceRepo.FindByCode(codes[0])
	if province != nil {
		loc.Province = province.Name
		loc.ProvinceID = province.Code
	}

	return loc, nil
}
