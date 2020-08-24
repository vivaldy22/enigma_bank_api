package validation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStruct struct {
	s string
	i int
}

func TestValidateInputNotEmpty(t *testing.T) {
	t.Run("", func(t *testing.T) {
		i := 1
		// mockDataValid := mockStruct{s: "string", i: 1, p: &i}
		mockDataNotValid := []mockStruct{
			{
				s: "",
				i: i,
			},
			{
				s: "string",
				i: 0,
			},
			{
				s: "string",
				i: i,
			},
		}
		// err := ValidateInputNotEmpty(mockDataValid.s, mockDataValid.i, mockDataValid)
		// assert.Nil(t, err)
		for _, mock := range mockDataNotValid {
			err := ValidateInputNotEmpty(mock.s, mock.i)
			assert.NotNil(t, err)
		}
	})
}

func TestValidateStructNotEmpty(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// mockDataNotEmpty := mockStruct{
		// 	s: "string",
		// 	i: 1,
		// 	e: errors.New("error"),
		// }
		// mockDataEmpty := []mockStruct{
		// 	{
		// 		s: "",
		// 	},
		// 	{
		// 		i: 0,
		// 	},
		// 	{
		// 		e: nil,
		// 	},
		// }
		// err := ValidateStructNotEmpty(mockDataNotEmpty)
		// assert.Nil(t, err)
		// for _, mock := range mockDataEmpty {
		// 	err = ValidateStructNotEmpty(mock)
		// 	assert.NotNil(t, err)
		// }
	})
}

type mockUUID struct {
	uuid           string
	expectedResult error
}

func TestValidateUUID(t *testing.T) {
	t.Run("", func(t *testing.T) {
		uuid := []mockUUID{
			{
				uuid:           "",
				expectedResult: errors.New("value '' is not a valid UUID"),
			},
			{
				uuid:           "a1a6f66e-bff4-11ea-8207-1c3e84e5e81",
				expectedResult: errors.New("value 'a1a6f66e-bff4-11ea-8207-1c3e84e5e81' is not a valid UUID"),
			},
			{
				uuid:           "a1a6f66e-bff4-11ea-8207-1c3e84e5e8191",
				expectedResult: errors.New("value 'a1a6f66e-bff4-11ea-8207-1c3e84e5e8191' is not a valid UUID"),
			},
			{
				uuid:           "a1a6f66e-bff4-11ea-8207-1c3e84e5e819",
				expectedResult: nil,
			},
		}
		for _, mock := range uuid {
			err := ValidateUUID(mock.uuid)
			assert.Equal(t, mock.expectedResult, err)
		}
	})
}
