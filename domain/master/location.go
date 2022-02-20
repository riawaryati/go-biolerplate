package master

import (
	"gopkg.in/guregu/null.v4"
)

type Country struct {
	ID   int64  `json:"id" db:"country_id"`
	Name string `json:"name" db:"name"`
}

type CountryFilter struct {
	Name null.String
}

type Province struct {
	ID        int64  `json:"id" db:"province_id"`
	CountryID int64  `json:"country_id" db:"country_id"`
	Name      string `json:"name" db:"name"`
}

type ProvinceFilter struct {
	CountryID null.Int
	Name      null.String
}

type City struct {
	ID         int64  `json:"id" db:"city_id"`
	ProvinceID int64  `json:"province_id" db:"province_id"`
	Name       string `json:"name" db:"name"`
}

type CityFilter struct {
	ProvinceID null.Int
	Name       null.String
}

type District struct {
	ID       int64       `json:"id" db:"district_id"`
	CityID   int64       `json:"city_id" db:"city_id"`
	Name     string      `json:"name" db:"name"`
	Metadata null.String `json:"metadata" db:"metadata"`
}

type DistrictMetadata struct {
	JNE        *DistrictJNEMetadata `json:"jne"`
	JETExpress *DistrictJETMetadata `json:"jet"`
}

type DistrictJNEMetadata struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type DistrictJETMetadata struct {
	Code string `json:"code"`
}

type DistrictFilter struct {
	CityID null.Int
	Name   null.String
}

type SubDistrict struct {
	ID         int64  `json:"id" db:"sub_district_id"`
	DistrictID int64  `json:"district_id" db:"district_id"`
	Name       string `json:"name" db:"name"`
}

type SubDistrictFilter struct {
	DistrictID null.Int
	Name       null.String
}
