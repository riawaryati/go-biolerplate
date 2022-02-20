package master

import (
	"fmt"
	"strings"

	dg "github.com/furee/backend/domain/general"
	dm "github.com/furee/backend/domain/master"
	"github.com/furee/backend/infra"
)

type ProvinceRepo struct {
	DBList *infra.DatabaseList
}

func newProvinceRepo(dbList *infra.DatabaseList) ProvinceRepo {
	return ProvinceRepo{
		DBList: dbList,
	}
}

const (
	pqSelectProvince = `
	SELECT
		province_id,
		country_id,
		name
	FROM
		provinces`

	pqCountProvince = `
	SELECT
		COUNT(1) as count
	FROM
		provinces`

	pqWhere = `
	WHERE`

	pqFilterProvinceID = `
		province_id = ?`

	pqFilterCountryID = `
		country_id = ?`

	pqFilterName = `
		lower(name) LIKE ?`

	pqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	pqOrderBy = `
	ORDER BY`
)

type ProvinceRepoItf interface {
	GetByID(provinceID int64) (dm.Province, error)
	GetByCountryID(countryID int64) ([]dm.Province, error)
	GetByName(name string) (dm.Province, error)
	GetListProvince(pagination dg.PaginationData, filter dm.ProvinceFilter) ([]dm.Province, error)
	GetTotalDataProvince(pagination dg.PaginationData, filter dm.ProvinceFilter) (int64, int64, error)
}

func (pr ProvinceRepo) GetByID(provinceID int64) (dm.Province, error) {
	var res dm.Province

	q := fmt.Sprintf("%s%s%s", pqSelectProvince, pqWhere, pqFilterProvinceID)
	query, args, err := pr.DBList.Backend.Read.In(q, provinceID)
	if err != nil {
		return res, err
	}

	query = pr.DBList.Backend.Read.Rebind(query)
	err = pr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (pr ProvinceRepo) GetByCountryID(countryID int64) ([]dm.Province, error) {
	var res []dm.Province

	q := fmt.Sprintf("%s%s%s", pqSelectProvince, pqWhere, pqFilterCountryID)
	query, args, err := pr.DBList.Backend.Read.In(q, countryID)
	if err != nil {
		return res, err
	}

	query = pr.DBList.Backend.Read.Rebind(query)
	err = pr.DBList.Backend.Read.Select(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (pr ProvinceRepo) GetByName(name string) (dm.Province, error) {
	var res dm.Province

	q := fmt.Sprintf("%s%s%s", pqSelectProvince, pqWhere, pqFilterName)
	query, args, err := pr.DBList.Backend.Read.In(q, strings.ToLower(name))
	if err != nil {
		return res, err
	}

	query = pr.DBList.Backend.Read.Rebind(query)
	err = pr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (pr ProvinceRepo) GetListProvince(pagination dg.PaginationData, filter dm.ProvinceFilter) ([]dm.Province, error) {
	var result []dm.Province
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, pqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.CountryID.Valid {
		fl = append(fl, pqFilterCountryID)
		param = append(param, filter.CountryID.Int64)
	}

	q := pqSelectProvince

	if len(fl) > 0 {
		q += pqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + pqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += pqLimitOffset
		param = append(param, pagination.Limit)
		param = append(param, pagination.Offset)
	}

	query, args, err := pr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, err
	}

	query = pr.DBList.Backend.Read.Rebind(query)
	err = pr.DBList.Backend.Read.Select(&result, query, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (pr ProvinceRepo) GetTotalDataProvince(pagination dg.PaginationData, filter dm.ProvinceFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, pqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.CountryID.Valid {
		fl = append(fl, pqFilterCountryID)
		param = append(param, filter.CountryID.Int64)
	}

	q := pqCountProvince

	if len(fl) > 0 {
		q += dqWhere + strings.Join(fl, " AND ")
	}

	query, args, err := pr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, 0, err
	}

	//Run query to get total data
	query = pr.DBList.Backend.Read.Rebind(query)
	err = pr.DBList.Backend.Read.Get(&result, query, args...)
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
