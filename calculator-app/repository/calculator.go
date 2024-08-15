package repository

import (
	"calculator-app/model"
	"errors"
)

type Calculator struct{}

func (c *Calculator) Tambah(op model.Operation) float32 {
	return op.Value1 + op.Value2
}

func (c *Calculator) Kurangi(op model.Operation) float32 {
	return op.Value1 - op.Value2
}

func (c *Calculator) Perkalian(op model.Operation) float32 {
	return op.Value1 * op.Value2
}

func (c *Calculator) Pembagaian(op model.Operation) (float32, error) {
	if op.Value2 == 0 {
		return 0, errors.New("tidak bisa membagi nol")
	}
	return op.Value1 / op.Value2, nil
}
