package repository

import (
	"github.com/khaibarid/location/configs"
	"github.com/khaibarid/location/internal/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegencyRepository interface {
	All(provinceID, q string, page, size int) ([]*domain.Regency, error)
	FindByCode(code string) (*domain.Regency, error)
}

type regencyServiceService struct {
	col *mgo.Collection
}

func (p regencyServiceService) All(provinceID, q string, page, size int) ([]*domain.Regency, error) {
	var items []*domain.Regency
	err := p.col.Find(
		bson.M{
			"name": bson.RegEx{Pattern: q + ".*", Options: "i"},
			"code": bson.RegEx{Pattern: provinceID + ".", Options: "i"},
		}).
		Skip((page - 1) * size).Limit(size).
		All(&items)
	if len(items) < 1 {
		return []*domain.Regency{}, err
	}
	return items, err
}

func (p regencyServiceService) FindByCode(code string) (*domain.Regency, error) {
	var regency *domain.Regency
	err := p.col.Find(bson.M{"code": code}).One(&regency)
	return regency, err
}

func NewRegencyRepo(conn *configs.LocationDB) RegencyRepository {
	return &regencyServiceService{conn.Database.C("regency")}
}
