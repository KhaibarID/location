package repository

import (
	"github.com/khaibarid/location/configs"
	"github.com/khaibarid/location/internal/domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SubDistrictRepository interface {
	All(district, q string, page, size int) ([]*domain.SubDistrict, error)
	FindByCode(code string) (*domain.SubDistrict, error)
	Save(d *domain.SubDistrict) error
}

type subDistrictServiceService struct {
	col *mgo.Collection
}

func (p subDistrictServiceService) All(district, q string, page, size int) ([]*domain.SubDistrict, error) {
	var items []*domain.SubDistrict
	err := p.col.Find(bson.M{
		"name": bson.RegEx{Pattern: q + ".*", Options: "i"},
		"code": bson.RegEx{Pattern: district + ".", Options: "i"},
	}).
		Skip((page - 1) * size).Limit(size).
		All(&items)
	if len(items) < 1 {
		return []*domain.SubDistrict{}, err
	}
	return items, err
}

func (p subDistrictServiceService) FindByCode(code string) (*domain.SubDistrict, error) {
	var subDist *domain.SubDistrict
	err := p.col.Find(bson.M{"code": code}).One(&subDist)
	return subDist, err
}

func (p subDistrictServiceService) Save(d *domain.SubDistrict) error {
	err := p.col.Insert(d)
	return err
}

func NewSubDistrictRepo(conn *configs.LocationDB) SubDistrictRepository {
	return &subDistrictServiceService{conn.Database.C("sub_districts")}
}
