package repository

import (
	"day-one-golang/model"
	"errors"
	"fmt"
)

type ProductRepository struct {
	product model.Product
}

func NewProductRepository(id int, namaProduct string) *ProductRepository {
	return &ProductRepository{
		product: model.Product{Id: id, NamaProduct: namaProduct, Jumlah: 0},
	}
}

func (repo *ProductRepository) Tambah(jumlah int) {
	repo.product.Jumlah += jumlah
	fmt.Printf("jumlah product '%s' sekrang: %d\n ", repo.product.NamaProduct, repo.product.Jumlah)
}

func (repo *ProductRepository) Kurang(jumlah int) error {
	if jumlah > repo.product.Jumlah {
		return errors.New("jumlah yang dikurangi melebihi jumlah produk yang ada")
	}
	repo.product.Jumlah -= jumlah
	fmt.Printf("jumlah product '%s' sekrang: %d\n ", repo.product.NamaProduct, repo.product.Jumlah)
	return nil
}

func (repo *ProductRepository) Tampilkan() {
	fmt.Printf("ID Product: %d, Nama product: %s, jumlah %d\n", repo.product.Id, repo.product.NamaProduct, repo.product.Jumlah)
}
