package master

import (
	"fmt"
	"strings"

	dg "github.com/furee/backend/domain/general"
	dm "github.com/furee/backend/domain/master"
	"github.com/furee/backend/infra"
)

type SubDistrictRepo struct {
	DBList *infra.DatabaseList
}

func newSubDistrictRepo(dbList *infra.DatabaseList) SubDistrictRepo {
	return SubDistrictRepo{
		DBList: dbList,
	}
}

const (
	sdqSelectSubDistrict = `
	SELECT
		sub_district_id,
		district_id,
		name
	FROM
		sub_districts`

	sdqCountSubDistrict = `
	SELECT
		COUNT(1) as count
	FROM
		sub_districts`

	sdqWhere = `
	WHERE`

	sdqFilterSubDistrictID = `
		sub_district_id = ?`

	sdqFilterDistrictID = `
		district_id = ?`

	sdqFilterName = `
		lower(name) LIKE ?`

	sdqLimitOffset = `
	LIMIT ?
	OFFSET ?`

	sdqOrderBy = `
	ORDER BY`
)

type SubDistrictRepoItf interface {
	GetByID(subDistrictID int64) (dm.SubDistrict, error)
	GetByDistrictID(districtID int64) ([]dm.SubDistrict, error)
	GetByName(name string) (dm.SubDistrict, error)
	GetListSubDistrict(pagination dg.PaginationData, filter dm.SubDistrictFilter) ([]dm.SubDistrict, error)
	GetTotalDataSubDistrict(pagination dg.PaginationData, filter dm.SubDistrictFilter) (int64, int64, error)
}

func (sdr SubDistrictRepo) GetByID(subDistrictID int64) (dm.SubDistrict, error) {
	var res dm.SubDistrict

	q := fmt.Sprintf("%s%s%s", sdqSelectSubDistrict, sdqWhere, sdqFilterSubDistrictID)
	query, args, err := sdr.DBList.Backend.Read.In(q, subDistrictID)
	if err != nil {
		return res, err
	}

	query = sdr.DBList.Backend.Read.Rebind(query)
	err = sdr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (sdr SubDistrictRepo) GetByDistrictID(districtID int64) ([]dm.SubDistrict, error) {
	var res []dm.SubDistrict

	q := fmt.Sprintf("%s%s%s", sdqSelectSubDistrict, sdqWhere, sdqFilterDistrictID)
	query, args, err := sdr.DBList.Backend.Read.In(q, districtID)
	if err != nil {
		return res, err
	}

	query = sdr.DBList.Backend.Read.Rebind(query)
	err = sdr.DBList.Backend.Read.Select(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (sdr SubDistrictRepo) GetByName(name string) (dm.SubDistrict, error) {
	var res dm.SubDistrict

	q := fmt.Sprintf("%s%s%s", sdqSelectSubDistrict, sdqWhere, sdqFilterName)
	query, args, err := sdr.DBList.Backend.Read.In(q, strings.ToLower(name))
	if err != nil {
		return res, err
	}

	query = sdr.DBList.Backend.Read.Rebind(query)
	err = sdr.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (sdr SubDistrictRepo) GetListSubDistrict(pagination dg.PaginationData, filter dm.SubDistrictFilter) ([]dm.SubDistrict, error) {
	var result []dm.SubDistrict
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, sdqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.DistrictID.Valid {
		fl = append(fl, sdqFilterDistrictID)
		param = append(param, filter.DistrictID.Int64)
	}

	q := sdqSelectSubDistrict

	if len(fl) > 0 {
		q += sdqWhere + strings.Join(fl, " AND ")
	}

	// Add orderby value.
	q += " " + sdqOrderBy + " " + pagination.OrderBy.String + " " + strings.ToLower(pagination.Sort)

	if !pagination.IsGetAll {
		// Add limit & page to param.
		q += sdqLimitOffset
		param = append(param, pagination.Limit)
		param = append(param, pagination.Offset)
	}

	query, args, err := sdr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, err
	}

	query = sdr.DBList.Backend.Read.Rebind(query)
	err = sdr.DBList.Backend.Read.Select(&result, query, args...)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (sdr SubDistrictRepo) GetTotalDataSubDistrict(pagination dg.PaginationData, filter dm.SubDistrictFilter) (int64, int64, error) {
	var result int64
	param := make([]interface{}, 0)
	var fl []string

	if filter.Name.Valid {
		fl = append(fl, sdqFilterName)
		param = append(param, "%"+strings.Title(strings.ToLower(filter.Name.String))+"%")
	}

	if filter.DistrictID.Valid {
		fl = append(fl, sdqFilterDistrictID)
		param = append(param, filter.DistrictID.Int64)
	}

	q := sdqCountSubDistrict

	if len(fl) > 0 {
		q += sdqWhere + strings.Join(fl, " AND ")
	}

	query, args, err := sdr.DBList.Backend.Read.In(q, param...)
	if err != nil {
		return result, 0, err
	}

	//Run query to get total data
	query = sdr.DBList.Backend.Read.Rebind(query)
	err = sdr.DBList.Backend.Read.Get(&result, query, args...)
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
