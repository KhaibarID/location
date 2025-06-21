package repository

import (
	"github.com/khaibarid/location/configs"
	"github.com/khaibarid/location/internal/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DistrictRepository interface {
	All(provinceCode, q string, page, size int) ([]*domain.District, error)
	FindByCode(code string) (*domain.District, error)
}

type districtServiceService struct {
	col *mgo.Collection
}

func (p districtServiceService) All(provinceCode, q string, page, size int) ([]*domain.District, error) {
	var items []*domain.District
	err := p.col.Find(bson.M{
		"name": bson.RegEx{Pattern: q + ".*", Options: "i"},
		"code": bson.RegEx{Pattern: provinceCode + ".", Options: "i"},
	}).
		Skip((page - 1) * size).Limit(size).
		All(&items)
	if len(items) < 1 {
		return []*domain.District{}, err
	}
	return items, err
}

func (p districtServiceService) FindByCode(code string) (*domain.District, error) {
	dist := new(domain.District)
	err := p.col.Find(bson.M{"code": code}).One(&dist)
	return dist, err
}

func NewDistrictRepo(conn *configs.MongoDB) DistrictRepository {
	return &districtServiceService{conn.LOC.C("districts")}
}
