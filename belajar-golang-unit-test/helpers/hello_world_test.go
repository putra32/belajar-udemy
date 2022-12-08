package helpers

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expecred string
	}{
		{
			name:     "Putra",
			request:  "Putra",
			expecred: "Hello Putra",
		},
		{
			name:     "Boby",
			request:  "Boby",
			expecred: "Hello Boby",
		},
		{
			name:     "Lubis",
			request:  "Lubis",
			expecred: "Hello Lubis",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expecred, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Boby", func(t *testing.T) {
		result := HelloWorld("Boby")
		assert.Equal(t, "Hello Boby", result, "result must be 'Hello Boby'")
	})
	t.Run("Putra", func(t *testing.T) {
		result := HelloWorld("Putra")
		assert.Equal(t, "Hello Putra", result, "result must be 'Hello Putra'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("cannot run on lunix")
	}
	result := HelloWorld("Putra")
	assert.Equal(t, "Hello Putra", result, "result must be 'Hello Putra'")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Putra")
	assert.Equal(t, "Hello Putra", result, "result must be 'Hello Putra'")
}
