package domain

import "gopkg.in/mgo.v2/bson"

type SubDistrict struct {
	ID         bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Code       string        `bson:"code" json:"code"`
	PostalCode string        `bson:"postal_code" json:"post_code"`
	Name       string        `bson:"name" json:"name"`
	District   string        `bson:"district" json:"district"`
	Regency    string        `bson:"regency" json:"regency"`
	Province   string        `bson:"province" json:"province"`
}

func NewSubDistrict(code, postalCode, name, district, regency, province string) *SubDistrict {
	return &SubDistrict{
		Code:       code,
		PostalCode: postalCode,
		Name:       name,
		District:   district,
		Regency:    regency,
		Province:   province,
	}
}
