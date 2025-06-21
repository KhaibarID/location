package domain

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

/**
 * Created by Muhammad Muflih Kholidin
 * at 2020-10-01 22:19:43
 * https://github.com/mmuflih
 * muflic.24@gmail.com
 **/

type Location struct {
	MID           bson.ObjectId `bson:"_id,omitempty" json:"-"`
	ID            int64         `bson:"id" json:"id"`
	ProvinceID    string        `bson:"province_id" json:"province_id"`
	Province      string        `bson:"province" json:"province"`
	CityID        string        `bson:"city_id" json:"city_id"`
	City          string        `bson:"city" json:"city"`
	DistrictID    string        `bson:"district_id" json:"district_id"`
	District      string        `bson:"district" json:"district"`
	SubDistrictID string        `bson:"sub_district_id" json:"sub_district_id"`
	SubDistrict   string        `bson:"sub_district" json:"sub_district"`
	PostCode      string        `bson:"post_code" json:"post_code"`
}

type DetailAddress struct {
	PostCode        string `json:"post_code"`
	SubDistrictID   string `json:"sub_district_id"`
	Province        string `json:"province"`
	CodeProvince    string `json:"code_province"`
	City            string `json:"city"`
	CodeCity        string `json:"code_city"`
	District        string `json:"district"`
	CodeDistrict    string `json:"code_district"`
	SubDistrict     string `json:"sub_district"`
	CodeSubDistrict string `json:"code_sub_district"`
}

func (l *Location) GetAddress() string {
	if l == nil {
		return ""
	}
	return fmt.Sprintf(
		"%s, %s, %s, %s, %s",
		l.SubDistrict,
		l.District,
		l.City,
		l.Province,
		l.PostCode,
	)
}
