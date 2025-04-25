package repositories

import "DummyMultifinance/domain/models"

type ConsumerRepository interface {
	GetByID(id int) (*models.Consumer, error)
	Save(consumer *models.Consumer) error
}
