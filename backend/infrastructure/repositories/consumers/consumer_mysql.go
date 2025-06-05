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

func (m *mysqlConsumerRepo) Insert(ctx context.Context, consumer *models.Consumers) (*models.Consumers, error) {
	query := `INSERT INTO consumers
        (user_id, nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo, selfie_photo, created_at, updated_at)
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

func (m *mysqlConsumerRepo) GetAllData(ctx context.Context) (*models.Consumers, error) {
	query := `SELECT id, user_id, nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo, selfie_photo, created_at, updated_at
              FROM consumers`

	row := m.DB.QueryRowContext(ctx, query)

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

func (m *mysqlConsumerRepo) GetByID(ctx context.Context, id int) (*models.Consumers, error) {
	query := `SELECT id, user_id, nik, full_name, legal_name, birth_place, birth_date, salary, ktp_photo, selfie_photo, created_at, updated_at
              FROM limits WHERE id = ?`

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

func (m *mysqlConsumerRepo) GetConsumerLimit(ctx context.Context, consumer_id int) ([]models.ConsumersLimit, error) {
	query := `SELECT 
			consumers.id,
			tenors.id,
			tenors.name,
			limits.amount
			FROM limits
			LEFT JOIN tenors ON tenors.id = limits.tenor_id
			LEFT JOIN consumers ON consumers.id = limits.consumer_id
			WHERE consumers.id = ?`

	// Gunakan QueryContext karena kita mengharapkan lebih dari satu hasil
	rows, err := m.DB.QueryContext(ctx, query, consumer_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Inisialisasi limits sebagai slice kosong
	var limits []models.ConsumersLimit

	// Loop untuk menambahkan setiap hasil ke dalam slice
	for rows.Next() {
		var limit models.ConsumersLimit
		err := rows.Scan(
			&limit.ConsumerID,
			&limit.TenorID,
			&limit.Tenor,
			&limit.Amount,
		)
		if err != nil {
			return nil, err
		}
		// Tambahkan limit ke dalam slice
		limits = append(limits, limit)
	}

	// Periksa apakah ada error pada saat iterasi
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return limits, nil
}
