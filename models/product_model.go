package models

import (
	"database/sql"

	"github.com/keveaux/go_CRUD_application/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productmodel ProductModel) FindAll() (product []entities.Product, err error) {

	rows, err := productmodel.Db.Query("select * from example")

	if err != nil {
		return nil, err
	} else {
		var products []entities.Product

		for rows.Next() {
			var id int64
			var data string
			var prices int64

			err2 := rows.Scan(&id, &data, &prices)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:     id,
					Data:   data,
					Prices: prices,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productmodel ProductModel) Search(keyword string) (product []entities.Product, err error) {

	rows, err := productmodel.Db.Query("select * from example where Data like ?", "%"+keyword+"%")

	if err != nil {
		return nil, err
	} else {
		var products []entities.Product

		for rows.Next() {
			var id int64
			var data string
			var prices int64

			err2 := rows.Scan(&id, &data, &prices)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:     id,
					Data:   data,
					Prices: prices,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productmodel ProductModel) SearchPrices(min, max float64) (product []entities.Product, err error) {

	rows, err := productmodel.Db.Query("select * from example where prices >=? and prices<=?", min, max)

	if err != nil {
		return nil, err
	} else {
		var products []entities.Product

		for rows.Next() {
			var id int64
			var data string
			var prices int64

			err2 := rows.Scan(&id, &data, &prices)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:     id,
					Data:   data,
					Prices: prices,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productmodel ProductModel) Create(product *entities.Product) (err error) {

	result, err := productmodel.Db.Exec("insert into example(data,prices) values(?,?)", product.Data, product.Prices)

	if err != nil {
		return err
	} else {
		product.Id, _ = result.LastInsertId()
		return nil
	}
}

func (productmodel ProductModel) Update(product *entities.Product) (int64, error) {

	result, err := productmodel.Db.Exec("update example set data=? , prices=? where id=?", product.Data, product.Prices, product.Id)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (productmodel ProductModel) Delete(id int64) (int64, error) {

	result, err := productmodel.Db.Exec("delete from example where id=?", id)

	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
