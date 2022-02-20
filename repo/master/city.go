package master

import (
	"fmt"
	"strings"

	dg "github.com/furee/backend/domain/general"
	dm "github.com/furee/backend/domain/master"
	"github.com/furee/backend/infra"
)

type CityRepo struct {
	DBList *infra.DatabaseList
}

func newCityRepo(dbList *infra.DatabaseList) CityRepo {
	return CityRepo{
		DBList: dbList,
	}
}

const (
	cqSelectCity = `
	SELECT
		city_id,
		province_id,
		name
	FROM
		cities`

	cqCountCity = `
	SELECT
		COUNT(1) as count
	FROM
		cities`

	cqWhere = `
	WHERE`

	cqFilterCityID = `
		city_id = ?`

	cqFilterProvinceID = `
		province_id = ?`

	cqFilterName = `
		lower(name) LIKE ?`

	cqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	cqOrderBy = `
	ORDER BY`
)

type CityRepoItf interface {
	GetByID(cityID int64) (dm.City, error)
	GetByProvinceID(provinceID int64) ([]dm.City, error)
	GetByName(name string) (dm.City, error)
	GetListCity(pagination dg.PaginationData, filter dm.CityFilter) ([]dm.City, error)
	GetTotalDataCity(pagination dg.PaginationData, filter dm.CityFilter) (int64, int64, error)
}

func (cr CityRepo) GetByID(cityID int64) (dm.City, error) {
	var res dm.City

	q := fmt.Sprintf("%s%s%s", cqSelectCity, cqWhere, cqFilterCityID)
	query, args, err := cr.DBList.Backend.Read.In(q, cityID)
	if err != nil {
		return res, err
	}

	query = cr.DBList.Backend.Read.Rebind(query)
	err = cr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (cr CityRepo) GetByProvinceID(provinceID int64) ([]dm.City, error) {
	var res []dm.City

	q := fmt.Sprintf("%s%s%s", cqSelectCity, cqWhere, cqFilterProvinceID)
	query, args, err := cr.DBList.Backend.Read.In(q, provinceID)
	if err != nil {
		return res, err
	}

	query = cr.DBList.Backend.Read.Rebind(query)
	err = cr.DBList.Backend.Read.Select(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (cr CityRepo) GetByName(name string) (dm.City, error) {
	var res dm.City

	q := fmt.Sprintf("%s%s%s", cqSelectCity, cqWhere, cqFilterName)
	query, args, err := cr.DBList.Backend.Read.In(q, strings.ToLower(name))
	if err != nil {
		return res, err
	}

	query = cr.DBList.Backend.Read.Rebind(query)
	err = cr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (cr CityRepo) GetListCity(pagination dg.PaginationData, filter dm.CityFilter) ([]dm.City, error) {
	var result []dm.City
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, cqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.ProvinceID.Valid {
		fl = append(fl, cqFilterProvinceID)
		param = append(param, filter.ProvinceID.Int64)
	}

	q := cqSelectCity

	if len(fl) > 0 {
		q += cqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + cqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += cqLimitOffset
		param = append(param, pagination.Limit)
		param = append(param, pagination.Offset)
	}

	query, args, err := cr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, err
	}

	query = cr.DBList.Backend.Read.Rebind(query)
	err = cr.DBList.Backend.Read.Select(&result, query, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (cr CityRepo) GetTotalDataCity(pagination dg.PaginationData, filter dm.CityFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, cqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.ProvinceID.Valid {
		fl = append(fl, cqFilterProvinceID)
		param = append(param, filter.ProvinceID.Int64)
	}

	q := cqCountCity

	if len(fl) > 0 {
		q += dqWhere + strings.Join(fl, " AND ")
	}

	query, args, err := cr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, 0, err
	}

	//Run query to get total data
	query = cr.DBList.Backend.Read.Rebind(query)
	err = cr.DBList.Backend.Read.Get(&result, query, args...)
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
