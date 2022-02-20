package order

import (
	"github.com/furee/backend/infra"
	"github.com/sirupsen/logrus"
)

type OrderRepo struct {
	Order OrderDataRepoItf
	Item  ItemDataRepoItf
}

func NewMasterRepo(db *infra.DatabaseList, logger *logrus.Logger) OrderRepo {
	return OrderRepo{
		Order: newOrderDataRepo(db),
		Item:  newItemDataRepo(db),
	}
}
