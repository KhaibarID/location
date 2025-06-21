package repository

import (
	"github.com/khaibarid/location/configs"
	"github.com/khaibarid/location/internal/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProvinceRepository interface {
	All(q string, page, size int) ([]*domain.Province, error)
	FindByCode(code string) (*domain.Province, error)
}

type provinceServiceService struct {
	col *mgo.Collection
}

func (p provinceServiceService) All(q string, page, size int) ([]*domain.Province, error) {
	var items []*domain.Province
	err := p.col.Find(bson.M{"name": bson.RegEx{Pattern: q + ".*", Options: "i"}}).
		Skip((page - 1) * size).Limit(size).
		All(&items)
	if len(items) < 1 {
		return []*domain.Province{}, err
	}
	return items, err
}

func (p provinceServiceService) FindByCode(code string) (*domain.Province, error) {
	var province *domain.Province
	err := p.col.Find(bson.M{"code": code}).One(&province)
	return province, err
}

func NewProvinceRepo(conn *configs.LocationDB) ProvinceRepository {
	return &provinceServiceService{conn.Database.C("provinces")}
}
