package order

import (
	"database/sql"
	"fmt"

	du "github.com/furee/backend/domain/order"
	"github.com/furee/backend/infra"
)

type ItemDataRepo struct {
	DBList *infra.DatabaseList
}

func newItemDataRepo(dbList *infra.DatabaseList) ItemDataRepo {
	return ItemDataRepo{
		DBList: dbList,
	}
}

const (
	uqSelectItem = `
	SELECT
		item_id,
		order_id,
		item_code,
		description,
		quantity
	FROM
		items`

	uqInsertItem = `
	INSERT INTO items (
		order_id,
		item_code,
		description,
		quantity
	) VALUES (
		?, ?, ?, ?
	)
	RETURNING item_id`

	uqUpdateItem = `
	UPDATE 
		items
	SET
		`

	uqDeleteItem = `
	DELETE FROM 
			items `

	uqFilterItemID = `
		item_id = ?`

	uqFilterItemCode = `
		item_code = ?`

	uqFilterDescription = `
		description = ?`

	uqFilterQuantity = `
		quantity = ?`
)

type ItemDataRepoItf interface {
	GetByID(itemID int64) (*du.Item, error)
	GetByIDAndOrderID(itemID int64, orderID int64) (*du.Item, error)
	GetList() ([]du.Item, error)
	GetListByOrderID(orderID int64) ([]du.Item, error)
	DeleteByOrderID(tx *sql.Tx, orderID int64) error
	InsertItem(tx *sql.Tx, data du.Item) (int64, error)
	UpdateItem(tx *sql.Tx, data du.Item) error
}

func (ur ItemDataRepo) GetByID(itemID int64) (*du.Item, error) {
	var res du.Item

	q := fmt.Sprintf("%s%s%s", uqSelectItem, uqWhere, uqFilterItemID)
	query, args, err := ur.DBList.Backend.Read.In(q, itemID)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res.ItemID == 0 {
		return nil, nil
	}

	return &res, nil
}

func (ur ItemDataRepo) GetByIDAndOrderID(itemID int64, orderID int64) (*du.Item, error) {
	var res du.Item

	q := fmt.Sprintf("%s %s %s AND %s", uqSelectItem, uqWhere, uqFilterItemID, uqFilterOrderID)
	query, args, err := ur.DBList.Backend.Read.In(q, itemID, orderID)
	if err != nil {
		return nil, err
	}

	query = ur.DBList.Backend.Read.Rebind(query)
	err = ur.DBList.Backend.Read.Get(&res, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res.ItemID == 0 {
		return nil, nil
	}

	return &res, nil
}

func (ur ItemDataRepo) GetList() ([]du.Item, error) {
	var res []du.Item

	q := fmt.Sprintf("%s%s%s", uqSelectItem, uqWhere, uqFilterItemID)
	query, args, err := ur.DBList.Backend.Read.In(q)
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

func (ur ItemDataRepo) GetListByOrderID(orderID int64) ([]du.Item, error) {
	var res []du.Item

	q := fmt.Sprintf("%s %s %s", uqSelectItem, uqWhere, uqFilterOrderID)
	query, args, err := ur.DBList.Backend.Read.In(q, orderID)
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

func (ur ItemDataRepo) InsertItem(tx *sql.Tx, data du.Item) (int64, error) {
	param := make([]interface{}, 0)

	param = append(param, data.OrderID)
	param = append(param, data.ItemCode)
	param = append(param, data.Description)
	param = append(param, data.Quantity)

	// itemedAt, err := time.Parse(time.RFC3339, request.ItemedAt)

	// var itemedAt *time.Time
	// if err == nill {
	// 	itemedAt = &data.ItemedAt.Time
	// }
	// param = append(param, itemedAt)

	query, args, err := ur.DBList.Backend.Write.In(uqInsertItem, param...)
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

	var itemID int64
	err = res.Scan(&itemID)
	if err != nil {
		return 0, err
	}

	return itemID, nil
}

func (ur ItemDataRepo) UpdateItem(tx *sql.Tx, data du.Item) error {
	var err error

	q := fmt.Sprintf("%s %s, %s, %s %s %s", uqUpdateItem, uqFilterItemCode, uqFilterDescription, uqFilterQuantity, uqWhere, uqFilterItemID)
	query, args, err := ur.DBList.Backend.Read.In(q, data.ItemCode, data.Description, data.Quantity, data.ItemID)
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

func (ur ItemDataRepo) DeleteByOrderID(tx *sql.Tx, orderID int64) error {
	var err error

	q := fmt.Sprintf("%s %s %s", uqDeleteItem, uqWhere, uqFilterOrderID)

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
