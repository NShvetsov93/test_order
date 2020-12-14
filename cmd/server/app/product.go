package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Product_id int32
	Quantity   int32
}

func GetProduct(w http.ResponseWriter, r *http.Request) (*Product, error) {
	product := &Product{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, product)
	if err != nil {
		return nil, err
	}
	if err = product.validate(); err != nil {
		return nil, err
	}
	return product, nil

}

func (p *Product) validate() error {
	if p.Product_id == 0 {
		return errors.New("product_id required")
	}
	if p.Quantity == 0 {
		return errors.New("quantity required")
	}
	return nil
}
