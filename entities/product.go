package entities

import "fmt"

type Product struct {
	Id     int64  `json:"id"`
	Data   string `json:"data"`
	Prices int64  `json:"prices"`
}

func (product Product) ToString() string {

	return fmt.Sprintf("id: %d\n name: %s\n ", product.Id, product.Data, product.Prices)

}
