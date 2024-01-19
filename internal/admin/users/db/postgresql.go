package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"test-backend/internal/admin/users"
	"test-backend/pkg/logging"
)

type repository struct {
	client *pgxpool.Pool
	logger *logging.Logger
}

func NewRepository(client *pgxpool.Pool, logger *logging.Logger) users.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r repository) GetAllData(ctx context.Context, req users.PaginationDTO) (users.DataDTO, error) {

	var (
		dt users.DataDTO
		ns []users.UserDTO
	)

	qC := `select count(*) from users;`
	r.client.QueryRow(ctx, qC).Scan(&dt.Count)

	q := ` select 
				id, name, surname, 
				patronymic, age,
				gender, country,
				probability, created_at
			 from users 
		   order by created_at desc limit $1 offset $2; `

	rows, err := r.client.Query(ctx, q, req.Limit, (req.Page-1)*req.Limit)
	defer rows.Close()

	if err != nil {
		return dt, nil
	}

	for rows.Next() {
		var n users.UserDTO
		errN := rows.Scan(
			&n.Id, &n.Name,
			&n.Surname, &n.Patronymic,
			&n.Age, &n.Gender, &n.Country,
			&n.Probability, &n.CreatedAt,
		)
		if errN != nil {
			r.logger.Error("errN :::", errN)
			continue
		}

		ns = append(ns, n)
	}
	dt.Users = ns

	return dt, nil
}

func (r repository) AddData(ctx context.Context, req users.ReqUser) (int, error) {

	var id int
	q := `INSERT INTO users (name, surname, patronymic) VALUES ($1, $2, $3)  RETURNING id;;`

	err := r.client.QueryRow(ctx, q, req.Name, req.Surname, req.Patronymic).Scan(&id)
	if err != nil && err.Error() != "no rows in result set" {
		return 0, err
	}

	return id, nil
}

func (r repository) UpdateData(ctx context.Context, req users.ReqUser) error {
	var (
		q string
	)

	q = ` update users set name = $1, surname = $2 ,  patronymic = $3 where Id = $4;`

	_, err := r.client.Exec(ctx, q, req.Name, req.Surname, req.Patronymic, req.Id)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}

func (r repository) UpdateAgeData(ctx context.Context, req users.ResUserAge) error {
	var (
		q string
	)

	q = ` update users set age = $1 where Id = $2;`

	_, err := r.client.Exec(ctx, q, req.Age, req.Id)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}

func (r repository) UpdateGenderData(ctx context.Context, req users.ResUserGender) error {
	var (
		q string
	)

	q = ` update users set gender = $1, probability = $2  where Id = $3;`

	_, err := r.client.Exec(ctx, q, req.Gender, req.Probability, req.Id)

	if err != nil && err.Error() != "no rows in result set" {
		fmt.Println("erro", err)
		return err
	}

	return nil
}

func (r repository) UpdateNationalizeData(ctx context.Context, req users.ResUserCountry) error {
	var (
		q string
	)

	q = ` update users set country = $1 where Id = $2;`

	_, err := r.client.Exec(ctx, q, req.Country, req.Id)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}

func (r repository) DeleteData(ctx context.Context, Id int) error {
	var (
		q string
	)

	q = ` delete from users where Id = $1 ;`
	_, err := r.client.Exec(ctx, q, Id)

	if err != nil && err.Error() != "no rows in result set" {
		return err
	}

	return nil
}
