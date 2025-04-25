package infrastructure

import (
	"DummyMultifinance/domain/models"
	"DummyMultifinance/domain/repositories"
	"context"
	"database/sql"
)

type mysqlConsumerRepo struct {
	DB *sql.DB
}

func NewMysqlConsumerRepo(db *sql.DB) repositories.ConsumerRepository {
	return &mysqlConsumerRepo{DB: db}
}

// Insert untuk menambahkan data consumer ke dalam database
func (m *mysqlConsumerRepo) Insert(ctx context.Context, consumer *models.Consumers) (*models.Consumers, error) {
	query := `INSERT INTO consumers
        (
			user_id,
			nik,
			full_name,
			legal_name,
			birth_place,
			birth_date,
			salary,
			ktp_photo,
			selfie_photo,
			created_at,
			updated_at
		)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	result, err := m.DB.ExecContext(ctx, query,
		consumer.UserID,
		consumer.NIK,
		consumer.FullName,
		consumer.LegalName,
		consumer.BirthPlace,
		consumer.BirthDate,
		consumer.Salary,
		consumer.KTPPhoto,
		consumer.SelfiePhoto,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	consumer.ID = int(id)
	return consumer, nil
}

// GetByID untuk mengambil data consumer berdasarkan ID
func (m *mysqlConsumerRepo) GetByID(ctx context.Context, id int) (*models.Consumers, error) {
	query := `SELECT 
				id,
				user_id,
				nik,
				full_name,
				legal_name,
				birth_place,
				birth_date,
				salary,
				ktp_photo,
				selfie_photo,
				created_at,
				updated_at
              FROM consumers WHERE id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)

	var consumer models.Consumers
	err := row.Scan(
		&consumer.ID,
		&consumer.UserID,
		&consumer.NIK,
		&consumer.FullName,
		&consumer.LegalName,
		&consumer.BirthPlace,
		&consumer.BirthDate,
		&consumer.Salary,
		&consumer.KTPPhoto,
		&consumer.SelfiePhoto,
		&consumer.CreatedAt,
		&consumer.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}
