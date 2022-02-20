package master

import (
	"fmt"
	"strings"

	dg "github.com/furee/backend/domain/general"
	dm "github.com/furee/backend/domain/master"
	"github.com/furee/backend/infra"
)

type DistrictRepo struct {
	DBList *infra.DatabaseList
}

func newDistrictRepo(dbList *infra.DatabaseList) DistrictRepo {
	return DistrictRepo{
		DBList: dbList,
	}
}

const (
	dqSelectDistrict = `
	SELECT
		district_id,
		city_id,
		name,
		metadata
	FROM
		districts`

	dqCountDistrict = `
	SELECT
		COUNT(1) as count
	FROM
		districts`

	dqWhere = `
	WHERE`

	dqFilterDistrictID = `
		district_id = ?`

	dqFilterCityID = `
		city_id = ?`

	dqFilterName = `
		lower(name) LIKE ?`

	dqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	dqOrderBy = `
	ORDER BY`
)

type DistrictRepoItf interface {
	GetByID(districtID int64) (dm.District, error)
	GetByCityID(districtID int64) ([]dm.District, error)
	GetByName(name string) (dm.District, error)
	GetListDistrict(pagination dg.PaginationData, filter dm.DistrictFilter) ([]dm.District, error)
	GetTotalDataDistrict(pagination dg.PaginationData, filter dm.DistrictFilter) (int64, int64, error)
}

func (dr DistrictRepo) GetByID(districtID int64) (dm.District, error) {
	var res dm.District

	q := fmt.Sprintf("%s%s%s", dqSelectDistrict, dqWhere, dqFilterDistrictID)
	query, args, err := dr.DBList.Backend.Read.In(q, districtID)
	if err != nil {
		return res, err
	}

	query = dr.DBList.Backend.Read.Rebind(query)
	err = dr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (dr DistrictRepo) GetByCityID(districtID int64) ([]dm.District, error) {
	var res []dm.District

	q := fmt.Sprintf("%s%s%s", dqSelectDistrict, dqWhere, dqFilterCityID)
	query, args, err := dr.DBList.Backend.Read.In(q, districtID)
	if err != nil {
		return res, err
	}

	query = dr.DBList.Backend.Read.Rebind(query)
	err = dr.DBList.Backend.Read.Select(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (dr DistrictRepo) GetByName(name string) (dm.District, error) {
	var res dm.District

	q := fmt.Sprintf("%s%s%s", dqSelectDistrict, dqWhere, dqFilterName)
	query, args, err := dr.DBList.Backend.Read.In(q, strings.ToLower(name))
	if err != nil {
		return res, err
	}

	query = dr.DBList.Backend.Read.Rebind(query)
	err = dr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (dr DistrictRepo) GetListDistrict(pagination dg.PaginationData, filter dm.DistrictFilter) ([]dm.District, error) {
	var result []dm.District
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, dqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.CityID.Valid {
		fl = append(fl, dqFilterCityID)
		param = append(param, filter.CityID.Int64)
	}

	q := dqSelectDistrict

	if len(fl) > 0 {
		q += dqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + dqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += dqLimitOffset
		param = append(param, pagination.Limit)
		param = append(param, pagination.Offset)
	}

	query, args, err := dr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, err
	}

	query = dr.DBList.Backend.Read.Rebind(query)
	err = dr.DBList.Backend.Read.Select(&result, query, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (dr DistrictRepo) GetTotalDataDistrict(pagination dg.PaginationData, filter dm.DistrictFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, dqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.CityID.Valid {
		fl = append(fl, dqFilterCityID)
		param = append(param, filter.CityID.Int64)
	}

	q := dqCountDistrict

	if len(fl) > 0 {
		q += dqWhere + strings.Join(fl, " AND ")
	}

	query, args, err := dr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, 0, err
	}

	//Run query to get total data
	query = dr.DBList.Backend.Read.Rebind(query)
	err = dr.DBList.Backend.Read.Get(&result, query, args...)
	if err != nil {
		return result, 0, err
	}

	//Calculate Total Page
	totalPage := result / int64(pagination.Limit)
	if result%int64(pagination.Limit) > 0 {
		totalPage++
	}

	return result, totalPage, nil
}
