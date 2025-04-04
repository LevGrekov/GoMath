package matrix

import (
	"errors"
	"math/rand"
)

type Matrix struct {
	data [][]float64
	rows int
	cols int
}

func New(rows, cols int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}
	return &Matrix{data: data, rows: rows, cols: cols}
}

func (m *Matrix) Rows() int {
	return m.rows
}

func (m *Matrix) Cols() int {
	return m.cols
}

func (m *Matrix) Set(row, col int, value float64) error {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		return errors.New("index out of range")
	}
	m.data[row][col] = value
	return nil
}

func (m *Matrix) Get(row, col int) (float64, error) {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		return 0, errors.New("index out of range")
	}
	return m.data[row][col], nil
}

func (m *Matrix) Multiply(other *Matrix) (*Matrix, error) {
	if m.cols != other.rows {
		return nil, errors.New("incompatible matrix dimensions")
	}

	result := New(m.rows, other.cols)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < other.cols; j++ {
			sum := 0.0
			for k := 0; k < m.cols; k++ {
				sum += m.data[i][k] * other.data[k][j]
			}
			result.data[i][j] = sum
		}
	}

	return result, nil
}

func GenerateRandom(rows, cols int) *Matrix {
	m := New(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m.data[i][j] = rand.Float64()*10 - 5
		}
	}
	return m
}

func (m *Matrix) Equal(other *Matrix) bool {
	if m.rows != other.rows || m.cols != other.cols {
		return false
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			if m.data[i][j] != other.data[i][j] {
				return false
			}
		}
	}

	return true
}
