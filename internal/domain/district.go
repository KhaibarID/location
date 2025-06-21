package domain

type District struct {
	Code                 string `bson:"code" json:"code"`
	Name                 string `bson:"name" json:"name"`
	PostalCode           string `bson:"postal_code" json:"postal_code"`
	NumberOfSubDistricts string `bson:"number_of_sub_district" json:"-"`
	NumberOfIsland       string `bson:"number_of_island" json:"-"`
	Regency              string `bson:"regency" json:"-"`
	Province             string `bson:"province" json:"-"`
}
