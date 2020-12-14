package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

//Storage ...
type Storage struct {
	db *pgxpool.Pool
}

//NewStorage ...
func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{
		db: db,
	}
}

//InsertProduct ...
func (s *Storage) InsertProduct(ctx context.Context, productId int32, quantity int32) error {
	var q int32
	rows, err := s.db.Query(ctx, "select quantity from products_q where product_id=$1 ", productId)
	if err != nil {
		return err
	}

	for rows.Next() {
		err := rows.Scan(&q)
		if err != nil {
			return err
		}
	}
	if q != 0 {
		_, err = s.db.Exec(ctx, "update products_q set quantity = $2 where product_id=$1", productId, quantity+q)
		if err != nil {
			return err
		}
	} else {
		_, err = s.db.Exec(ctx, "insert into products_q (product_id,quantity) values ($1,$2)", productId, quantity)
		if err != nil {
			return err
		}
	}
	return nil
}

//GetProduct ...
func (s *Storage) GetProduct(ctx context.Context, productId int32, quantity int32) (error, error) {
	var oq int32
	var res int32

	rows, err := s.db.Query(ctx, "SELECT \"check_limit\"();")
	if err != nil {
		return err, nil
	}
	for rows.Next() {
		err := rows.Scan(&oq)
		if err != nil {
			return err, nil
		}
		if res == 0 {
			return nil, errors.New("Too many requests")
		}
	}

	//Логика
	rows, err = s.db.Query(ctx, "SELECT \"get_product\"($1,$2);", productId, quantity)
	//Сброс счетчика
	_, errReset := s.db.Exec(ctx, "SELECT \"reset_limit\"();")
	if errReset != nil {
		return errReset, nil
	}

	if err != nil {
		return err, nil
	}
	for rows.Next() {
		err = rows.Scan(&res)
		if err != nil {
			return err, nil
		}
		if res == 0 {
			return errors.New("Такого количества нет на складе!"), nil
		}
	}

	return nil, nil
}

func (s *Storage) SelectProduct(ctx context.Context, productId int32) (int32, error) {
	rows, selErr := s.db.Query(ctx, "SELECT quantity from products_q where product_id = $1", productId)
	if selErr != nil {
		return 0, errors.Wrap(selErr, "Не удалось сделать выборку")
	}
	var res int32
	for rows.Next() {
		err := rows.Scan(&res)
		if err != nil {
			return 0, errors.Wrap(err, "Не удалось прочитать результат")
		}
	}

	return res, nil
}
