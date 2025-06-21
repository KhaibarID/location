package repository

import (
	"github.com/khaibarid/location/configs"
	"github.com/khaibarid/location/internal/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RegencyRepository interface {
	All(provinceID, q string, page, size int) (error, []*domain.Regency)
	FindByCode(code string) (*domain.Regency, error)
}

type regencyServiceService struct {
	col *mgo.Collection
}

func (p regencyServiceService) All(provinceID, q string, page, size int) (error, []*domain.Regency) {
	var items []*domain.Regency
	err := p.col.Find(
		bson.M{
			"name": bson.RegEx{Pattern: q + ".*", Options: "i"},
			"code": bson.RegEx{Pattern: provinceID + ".", Options: "i"},
		}).
		Skip((page - 1) * size).Limit(size).
		All(&items)
	if len(items) < 1 {
		return err, []*domain.Regency{}
	}
	return err, items
}

func (p regencyServiceService) FindByCode(code string) (*domain.Regency, error) {
	var regency *domain.Regency
	err := p.col.Find(bson.M{"code": code}).One(&regency)
	return regency, err
}

func NewRegencyRepo(conn *configs.MongoDB) RegencyRepository {
	return &regencyServiceService{conn.LOC.C("regency")}
}
