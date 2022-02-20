package master

import (
	"fmt"
	"strings"

	dg "github.com/furee/backend/domain/general"
	dm "github.com/furee/backend/domain/master"
	"github.com/furee/backend/infra"
)

type CountryRepo struct {
	DBList *infra.DatabaseList
}

func newCountryRepo(dbList *infra.DatabaseList) CountryRepo {
	return CountryRepo{
		DBList: dbList,
	}
}

const (
	prqSelectCountry = `
	SELECT
		country_id,
		name
	FROM
		countries`

	prqCountCountry = `
	SELECT
		COUNT(1) as count
	FROM
		countries`

	prqWhere = `
	WHERE`

	prqFilterCountryID = `
		country_id = ?`
	prqFilterName = `
		lower(name) LIKE ?`

	prqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	prqOrderBy = `
	ORDER BY`
)

type CountryRepoItf interface {
	GetByID(countryID int64) (dm.Country, error)
	GetByName(countryName string) (dm.Country, error)
	GetListCountry(pagination dg.PaginationData, filter dm.CountryFilter) ([]dm.Country, error)
	GetTotalDataCountry(pagination dg.PaginationData, filter dm.CountryFilter) (int64, int64, error)
}

func (cr CountryRepo) GetByID(countryID int64) (dm.Country, error) {
	var res dm.Country

	q := fmt.Sprintf("%s%s%s", prqSelectCountry, prqWhere, prqFilterCountryID)
	query, args, err := cr.DBList.Backend.Read.In(q, countryID)
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

func (cr CountryRepo) GetByName(countryName string) (dm.Country, error) {
	var res dm.Country

	q := fmt.Sprintf("%s%s%s", prqSelectCountry, prqWhere, prqFilterName)
	query, args, err := cr.DBList.Backend.Read.In(q, strings.ToLower(countryName))
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

func (cr CountryRepo) GetListCountry(pagination dg.PaginationData, filter dm.CountryFilter) ([]dm.Country, error) {
	var result []dm.Country
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, prqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	q := prqSelectCountry

	if len(fl) > 0 {
		q += prqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + prqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += prqLimitOffset
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

func (cr CountryRepo) GetTotalDataCountry(pagination dg.PaginationData, filter dm.CountryFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, prqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	q := prqCountCountry

	if len(fl) > 0 {
		q += prqWhere + strings.Join(fl, " AND ")
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
