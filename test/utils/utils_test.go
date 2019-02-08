package utils

import (
	"awesomeProject/src/utils"
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"syscall"
	"testing"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, utils.Abs(-3), utils.Abs(3))
	assert.Equal(t, utils.Abs(5), utils.Abs(-5))
	assert.Equal(t, 5, utils.Abs(-5))
	assert.Equal(t, utils.Abs(0), utils.Abs(-0))
}

func TestTernary(t *testing.T) {
	res := utils.Ternary(false, 1, 0)
	assert.Equal(t, res, 0)
	res = utils.Ternary(true, 1, 0)
	assert.Equal(t, res, 1)
}

func TestReadFile(t *testing.T) {
	testString := "hello world"
	f := func(str string) ([]byte, error) {
		if str != "" {
			arr, err := syscall.ByteSliceFromString(testString)
			return arr, err
		} else {
			return nil, errors.New("empty filename error")
		}
	}
	content, err := utils.ReadFile("random", f)
	assert.Equal(t, strings.TrimSuffix(string(*content), "\x00"), testString)
	assert.Nil(t, err)
	content, err = utils.ReadFile("", f)
	assert.NotNil(t, err)
}