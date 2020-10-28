package services

import (
	"fmt"
	"github.com/ponsonio/avanticaChallenge/pkg/api/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldBeValidQuadrant(t *testing.T) {

	assert := assert.New(t)

	var tests = []struct {
		input    int
		expected error
	}{
		{1, nil},
		{2, nil},
		{3, nil},
		{4, nil},
		{5, fmt.Errorf("Not a valid quadrant number %v", 5)},
	}

	for _, test := range tests {
		q := models.Quadrant{
			Number: test.input,
		}
		assert.Equal(isValidQuadrant(&q), test.expected)
	}
}
