package usecase

import (
	"errors"
	"strings"

	"github.com/khaibarid/location/internal/domain"
	"github.com/khaibarid/location/internal/repository"
)

type LocationUsecase struct {
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
) *LocationUsecase {
	return &LocationUsecase{
		districtRepo:    districtRepo,
		subDistrictRepo: subDistrictRepo,
		regencyRepo:     regencyRepo,
		provinceRepo:    provinceRepo,
	}
}

func (u *LocationUsecase) GetByCode(code string) (*domain.Location, error) {
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
	}

	district, _ := u.districtRepo.FindByCode(codes[2])
	if district != nil {
		loc.District = district.Name
		loc.DistrictID = district.Code
	}

	regency, _ := u.regencyRepo.FindByCode(codes[1])
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
