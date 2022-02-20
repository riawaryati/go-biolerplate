package order

import (
	"database/sql"
	"fmt"
	"strings"

	du "github.com/furee/backend/domain/order"
	"github.com/furee/backend/infra"
)

type OrderDataRepo struct {
	DBList *infra.DatabaseList
}

func newOrderDataRepo(dbList *infra.DatabaseList) OrderDataRepo {
	return OrderDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectOrder = `
	SELECT
		order_id,
		customer_name,
		ordered_at
	FROM
		orders`

	uqInsertOrder = `
	INSERT INTO orders (
		customer_name,
		ordered_at
	) VALUES (
		?, ?
	)
	RETURNING order_id`

	uqUpdateOrder = `
	UPDATE 
		orders
	SET
		`

	uqDeleteOrder = `
	DELETE FROM 
			orders `

	uqWhere = `
	WHERE`

	uqFilterOrderID = `
		order_id = ?`

	uqFilterCustomerName = `
		customer_name = ?`

	uqFilterOrderedAt = `
		ordered_at = ?`
)

type OrderDataRepoItf interface {
	GetByID(orderID int64) (*du.Order, error)
	GetList() ([]du.Order, error)
	DeleteByID(tx *sql.Tx, orderID int64) error
	InsertOrder(tx *sql.Tx, data du.Order) (int64, error)
	UpdateOrder(tx *sql.Tx, data du.Order) error
}

func (ur OrderDataRepo) GetByID(orderID int64) (*du.Order, error) {
	var res du.Order

	q := fmt.Sprintf("%s%s%s", uqSelectOrder, uqWhere, uqFilterOrderID)
	query, args, err := ur.DBList.Backend.Read.In(q, orderID)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res.OrderID == 0 {
		return nil, nil
	}

	return &res, nil
}

func (ur OrderDataRepo) GetList() ([]du.Order, error) {
	var res []du.Order

	// q := fmt.Sprintf("%s", uqSelectOrder)
	query, args, err := ur.DBList.Backend.Read.In(uqSelectOrder)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Select(&res, query, args...)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return res, err
	}

	return res, nil
}

func (ur OrderDataRepo) InsertOrder(tx *sql.Tx, data du.Order) (int64, error) {
	param := make([]interface{}, 0)

	param = append(param, strings.Title(strings.ToLower(data.CustomerName)))
	param = append(param, data.OrderedAt)

	// orderedAt, err := time.Parse(time.RFC3339, request.OrderedAt)

	// var orderedAt *time.Time
	// if err == nill {
	// 	orderedAt = &data.OrderedAt.Time
	// }
	// param = append(param, orderedAt)

	query, args, err := ur.DBList.Backend.Write.In(uqInsertOrder, param...)
	if err != nil {
		return 0, err
	}

	query = ur.DBList.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = ur.DBList.Backend.Write.QueryRow(query, args...)
	} else {
		res = tx.QueryRow(query, args...)
	}

	if err != nil {
		return 0, err
	}

	err = res.Err()
	if err != nil {
		return 0, err
	}

	var orderID int64
	err = res.Scan(&orderID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func (ur OrderDataRepo) UpdateOrder(tx *sql.Tx, data du.Order) error {
	var err error

	q := fmt.Sprintf("%s %s, %s %s %s", uqUpdateOrder, uqFilterCustomerName, uqFilterOrderedAt, uqWhere, uqFilterOrderID)
	query, args, err := ur.DBList.Backend.Read.In(q, data.CustomerName, data.OrderedAt, data.OrderID)
	if err != nil {
		return err
	}

	query = ur.DBList.Backend.Write.Rebind(query)
	_, err = ur.DBList.Backend.Write.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (ur OrderDataRepo) DeleteByID(tx *sql.Tx, orderID int64) error {
	var err error

	q := fmt.Sprintf("%s %s %s", uqDeleteOrder, uqWhere, uqFilterOrderID)

	query, args, err := ur.DBList.Backend.Read.In(q, orderID)
	if err != nil {
		return err
	}

	query = ur.DBList.Backend.Write.Rebind(query)
	_, err = ur.DBList.Backend.Write.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
