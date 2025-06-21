package domain

type Province struct {
	No                    int    `bson:"-" json:"-"`
	Code                  string `bson:"code" json:"code"`
	Name                  string `bson:"name" json:"name"`
	PostalCodePrefix      string `bson:"postal_code_prefix" json:"-"`
	PostalCodePrefixRange string `bson:"postal_code_prefix_range" json:"-"`
	PostalCodeRange       string `bson:"postal_code_range" json:"-"`
	NumberOfRegencies     string `bson:"number_of_regencies" json:"-"`
	NumberOfCities        string `bson:"number_of_cities" json:"-"`
	NumberOfDistricts     string `bson:"number_of_districts" json:"-"`
	NumberOfSubDistrict   string `bson:"number_of_sub_districts" json:"-"`
	NumberOfIsland        string `bson:"number_of_island" json:"-"`
}
