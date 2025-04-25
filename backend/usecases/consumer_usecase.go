package usecases

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
)

type ConsumerUseCase struct {
	ConsumerRepo repositories.ConsumerRepository
}

func (uc *ConsumerUseCase) GetConsumerByID(id int) (*models.Consumer, error) {
	return uc.ConsumerRepo.GetByID(id)
}

func (uc *ConsumerUseCase) CreateConsumer(consumer *models.Consumer) error {
	return uc.ConsumerRepo.Save(consumer)
}
