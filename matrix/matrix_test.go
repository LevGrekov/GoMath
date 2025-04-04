package matrix

import (
	"testing"
)

func TestNewMatrix(t *testing.T) {
	rows, cols := 3, 4
	m := New(rows, cols)

	if m.Rows() != rows || m.Cols() != cols {
		t.Errorf("Expected matrix size %dx%d, got %dx%d", rows, cols, m.Rows(), m.Cols())
	}
}

func TestSetAndGet(t *testing.T) {
	m := New(2, 2)
	val := 3.14
	m.Set(0, 1, val)

	got, err := m.Get(0, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if got != val {
		t.Errorf("Expected %f, got %f", val, got)
	}

	// Test out of bounds
	_, err = m.Get(2, 0)
	if err == nil {
		t.Error("Expected error for out of bounds access")
	}
}

func TestMatrixMultiplication(t *testing.T) {
	a := New(2, 3)
	b := New(3, 2)

	// Матрица A
	// 1 2 3
	// 4 5 6
	a.Set(0, 0, 1)
	a.Set(0, 1, 2)
	a.Set(0, 2, 3)
	a.Set(1, 0, 4)
	a.Set(1, 1, 5)
	a.Set(1, 2, 6)

	// Матрица B
	// 7  8
	// 9  10
	// 11 12
	b.Set(0, 0, 7)
	b.Set(0, 1, 8)
	b.Set(1, 0, 9)
	b.Set(1, 1, 10)
	b.Set(2, 0, 11)
	b.Set(2, 1, 12)

	// Ожидаемый результат
	// 58  64
	// 139 154
	expected := New(2, 2)
	expected.Set(0, 0, 58)
	expected.Set(0, 1, 64)
	expected.Set(1, 0, 139)
	expected.Set(1, 1, 154)

	result, err := a.Multiply(b)
	if err != nil {
		t.Fatalf("Multiplication failed: %v", err)
	}

	if !result.Equal(expected) {
		t.Errorf("Multiplication result incorrect")
	}
}

func TestInvalidMultiplication(t *testing.T) {
	a := New(2, 3)
	b := New(2, 3) // Несовместимые размеры

	_, err := a.Multiply(b)
	if err == nil {
		t.Error("Expected error for incompatible matrix dimensions")
	}
}

func TestRandomMatrixGeneration(t *testing.T) {
	rows, cols := 5, 5
	m := GenerateRandom(rows, cols)

	if m.Rows() != rows || m.Cols() != cols {
		t.Errorf("Expected random matrix size %dx%d, got %dx%d", rows, cols, m.Rows(), m.Cols())
	}

	// Проверим, что все элементы заполнены (не нули)
	hasNonZero := false
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val, _ := m.Get(i, j)
			if val != 0 {
				hasNonZero = true
				break
			}
		}
		if hasNonZero {
			break
		}
	}

	if !hasNonZero {
		t.Error("Random matrix contains only zeros")
	}
}

func BenchmarkMatrixMultiplication(b *testing.B) {
	// Создаем две большие матрицы для тестирования производительности
	rows, cols := 100, 100
	m1 := GenerateRandom(rows, cols)
	m2 := GenerateRandom(rows, cols)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m1.Multiply(m2)
	}
}
