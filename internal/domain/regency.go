package domain

type Regency struct {
	Code                 string `bson:"code" json:"code"`
	Name                 string `bson:"name" json:"name"`
	PostalCodePrefix     string `bson:"postal_code_prefix" json:"-"`
	PostalCodeRange      string `bson:"postal_code_range" json:"-"`
	NumberOfDistricts    string `bson:"number_of_districts" json:"-"`
	NumberOfSubDistricts string `bson:"number_of_sub_districts" json:"-"`
	NumberOfIsland       string `bson:"number_of_island" json:"-"`
	Province             string `bson:"province" json:"-"`
}
