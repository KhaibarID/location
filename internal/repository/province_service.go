package repository

import (
	"github.com/khaibarid/location/configs"
	"github.com/khaibarid/location/internal/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProvinceRepository interface {
	All(q string, page, size int) (error, []*domain.Province)
	FindByCode(code string) (*domain.Province, error)
}

type provinceServiceService struct {
	col *mgo.Collection
}

func (p provinceServiceService) All(q string, page, size int) (error, []*domain.Province) {
	var items []*domain.Province
	err := p.col.Find(bson.M{"name": bson.RegEx{Pattern: q + ".*", Options: "i"}}).
		Skip((page - 1) * size).Limit(size).
		All(&items)
	if len(items) < 1 {
		return err, []*domain.Province{}
	}
	return err, items
}

func (p provinceServiceService) FindByCode(code string) (*domain.Province, error) {
	var province *domain.Province
	err := p.col.Find(bson.M{"code": code}).One(&province)
	return province, err
}

func NewProvinceRepo(conn *configs.MongoDB) ProvinceRepository {
	return &provinceServiceService{conn.LOC.C("provinces")}
}
