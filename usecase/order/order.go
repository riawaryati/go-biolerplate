package order

import (
	"errors"
	"time"

	"github.com/furee/backend/domain/general"
	du "github.com/furee/backend/domain/order"
	"github.com/furee/backend/infra"
	"github.com/furee/backend/repo"
	ru "github.com/furee/backend/repo/order"
	"github.com/sirupsen/logrus"
)

type OrderDataUsecaseItf interface {
	GetList() ([]du.Order, error)
	GetByID(orderID int64) (*du.Order, error)
	DeleteByID(orderID int64) (bool, error)
	CreateOrder(data du.OrderRequest) (int64, error)
	UpdateOrder(data du.OrderRequest) (bool, error)
}

type OrderDataUsecase struct {
	Repo     ru.OrderDataRepoItf
	RepoItem ru.ItemDataRepoItf
	DBList   *infra.DatabaseList
	Conf     *general.SectionService
	Log      *logrus.Logger
}

func newOrderDataUsecase(r repo.Repo, conf *general.SectionService, logger *logrus.Logger, dbList *infra.DatabaseList) OrderDataUsecase {
	return OrderDataUsecase{
		Repo:     r.Order.Order,
		RepoItem: r.Order.Item,
		Conf:     conf,
		Log:      logger,
		DBList:   dbList,
	}
}

func (uu OrderDataUsecase) GetList() ([]du.Order, error) {
	orders, err := uu.Repo.GetList()
	if err != nil {
		// uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist order")
		return nil, err
	}

	retOrders := []du.Order{}
	for _, order := range orders {
		items, err := uu.RepoItem.GetListByOrderID(order.OrderID)

		if err != nil {
			return orders, err
		}

		order.Items = items
		retOrders = append(retOrders, order)
	}

	return retOrders, nil
}

func (uu OrderDataUsecase) GetByID(orderID int64) (*du.Order, error) {
	order, err := uu.Repo.GetByID(orderID)
	if err != nil {
		// uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist order")
		return nil, err
	}

	if order == nil {
		return nil, errors.New("order data not found")
	}

	items, err := uu.RepoItem.GetListByOrderID(orderID)
	if err != nil {
		// uu.Log.WithField("request", utils.StructToString(data)).WithError(err).Errorf("fail to checking is exist order")
		return order, err
	}

	if items != nil {
		order.Items = items
	}

	return order, nil
}

func (uu OrderDataUsecase) DeleteByID(orderID int64) (bool, error) {
	tx, err := uu.DBList.Backend.Write.Begin()
	if err != nil {
		return false, err
	}

	err = uu.Repo.DeleteByID(tx, orderID)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	err = uu.RepoItem.DeleteByOrderID(tx, orderID)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}

func (uu OrderDataUsecase) UpdateOrder(data du.OrderRequest) (bool, error) {
	tx, err := uu.DBList.Backend.Write.Begin()
	if err != nil {
		return false, err
	}

	orderedAt, err := time.Parse(time.RFC3339, data.OrderedAt)
	if err != nil {
		return false, err
	}

	order := du.Order{OrderID: data.OrderID, CustomerName: data.CustomerName, OrderedAt: orderedAt}

	err = uu.Repo.UpdateOrder(tx, order)
	if err != nil {
		tx.Rollback()
		return false, err
	}

	for _, item := range data.Items {
		oldItem, err := uu.RepoItem.GetByIDAndOrderID(item.ItemID, data.OrderID)

		if err != nil {
			tx.Rollback()
			return false, err
		}

		if oldItem != nil {
			err := uu.RepoItem.UpdateItem(tx, item)

			if err != nil {
				tx.Rollback()
				return false, err
			}
		} else {
			item.OrderID = data.OrderID
			_, err := uu.RepoItem.InsertItem(tx, item)

			if err != nil {
				tx.Rollback()
				return false, err
			}
		}
	}

	tx.Commit()
	return true, nil
}

func (uu OrderDataUsecase) CreateOrder(data du.OrderRequest) (int64, error) {
	tx, err := uu.DBList.Backend.Write.Begin()
	if err != nil {
		return 0, err
	}

	orderedAt, err := time.Parse(time.RFC3339, data.OrderedAt)
	if err != nil {
		return 0, err
	}

	order := du.Order{CustomerName: data.CustomerName, OrderedAt: orderedAt}

	orderID, err := uu.Repo.InsertOrder(tx, order)
	if err != nil {
		tx.Rollback()
		return 0, errors.New("failed to insert order")
	}

	for _, item := range data.Items {
		item.OrderID = orderID
		_, err := uu.RepoItem.InsertItem(tx, item)

		if err != nil {
			tx.Rollback()
			return 0, errors.New("failed to insert item")
		}
	}

	tx.Commit()
	return orderID, nil
}
