package test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/simonalong/gole/util"
)

func TestString(t *testing.T) {
	var s util.ISCString = "abcdefg"
	ss := s.Insert(3, "xyz")
	// ss := s.SubStringAfterLast(",")
	t.Logf("%v\n", ss) // abcxyzdefg

	sss := ss.Delete(3, 3)
	t.Logf("%v\n", sss) // abcdefg

}

func TestStringConvert(t *testing.T) {
	originalStr := "dataBaseUser"
	newStr := util.BigCamel(originalStr)
	assert.Equal(t, "DataBaseUser", newStr)
}

func TestMiddleLine(t *testing.T) {
	originalStr := "dataBaseUser"
	newStr := util.MiddleLine(originalStr)
	assert.Equal(t, "data-base-user", newStr)
}

func TestBigCamelToMiddleLine(t *testing.T) {
	originalStr := "DataBaseUser"
	newStr := util.BigCamelToMiddleLine(originalStr)
	assert.Equal(t, "data-base-user", newStr)
}

func TestBigCamelToSmallCamel(t *testing.T) {
	originalStr := "DataBaseUser"
	newStr := util.BigCamelToSmallCamel(originalStr)
	assert.Equal(t, "dataBaseUser", newStr)
}

func TestBigCamelToPostUnder(t *testing.T) {
	originalStr := "DataBaseUser"
	newStr := util.BigCamelToPostUnder(originalStr)
	assert.Equal(t, "data_base_user_", newStr)
}

func TestPostUnder(t *testing.T) {
	originalStr := "dataBaseUser"
	newStr := util.PostUnder(originalStr)
	assert.Equal(t, "data_base_user_", newStr)
}

func TestPrePostUnder(t *testing.T) {
	originalStr := "dataBaseUser"
	newStr := util.PrePostUnder(originalStr)
	assert.Equal(t, "_data_base_user_", newStr)
}

func TestBigCamelToPrePostUnder(t *testing.T) {
	originalStr := "DataBaseUser"
	newStr := util.BigCamelToPrePostUnder(originalStr)
	assert.Equal(t, "_data_base_user_", newStr)
}

func TestPreUnder(t *testing.T) {
	originalStr := "dataBaseUser"
	newStr := util.PreUnder(originalStr)
	assert.Equal(t, "_data_base_user", newStr)
}

func TestBigCamelToPreUnder(t *testing.T) {
	originalStr := "DataBaseUser"
	newStr := util.BigCamelToPreUnder(originalStr)
	assert.Equal(t, "_data_base_user", newStr)
}

func TestBigCamelToUnderLine(t *testing.T) {
	originalStr := "DataBaseUser"
	newStr := util.BigCamelToUnderLine(originalStr)
	assert.Equal(t, "data_base_user", newStr)
}

func TestBigCamelToUpperMiddle(t *testing.T) {
	originalStr := "DataBaseUser"
	newStr := util.BigCamelToUpperMiddle(originalStr)
	assert.Equal(t, "DATA-BASE-USER", newStr)
}

func TestUpperUnderMiddle(t *testing.T) {
	originalStr := "dataBaseUser"
	newStr := util.UpperUnderMiddle(originalStr)
	assert.Equal(t, "DATA-BASE-USER", newStr)
}

func TestUpperUnder(t *testing.T) {
	originalStr := "dataBaseUser"
	newStr := util.UpperUnder(originalStr)
	assert.Equal(t, "DATA_BASE_USER", newStr)
}

func TestBigCamelToUpperUnder(t *testing.T) {
	originalStr := "DataBaseUser"
	newStr := util.BigCamelToUpperUnder(originalStr)
	assert.Equal(t, "DATA_BASE_USER", newStr)
}

func TestMiddleLineToSmallCamel(t *testing.T) {
	originalStr := "data-base-user"
	newStr := util.MiddleLineToSmallCamel(originalStr)
	assert.Equal(t, "dataBaseUser", newStr)
}

func TestMiddleLineToBigCamel(t *testing.T) {
	originalStr := "data-base-user"
	newStr := util.MiddleLineToBigCamel(originalStr)
	assert.Equal(t, "DataBaseUser", newStr)
}

func TestPreFixUnderLine(t *testing.T) {
	originalStr := "dataBaseUser"
	newStr := util.PreFixUnderLine(originalStr, "pre_")
	assert.Equal(t, "pre_data_base_user", newStr)
}

func TestUnderLineToSmallCamel(t *testing.T) {
	originalStr1 := "data_base_user"
	newStr1 := util.UnderLineToSmallCamel(originalStr1)
	assert.Equal(t, "dataBaseUser", newStr1)

	originalStr2 := "_data_base_user"
	newStr2 := util.UnderLineToSmallCamel(originalStr2)
	assert.Equal(t, "dataBaseUser", newStr2)

	originalStr3 := "data_base_user_"
	newStr3 := util.UnderLineToSmallCamel(originalStr3)
	assert.Equal(t, "dataBaseUser", newStr3)
}

func TestPreFixUnderToSmallCamel(t *testing.T) {
	originalStr := "pre_data_base_user"
	newStr := util.PreFixUnderToSmallCamel(originalStr, "pre_")
	assert.Equal(t, "dataBaseUser", newStr)
}

func TestUnderLineToBigCamel(t *testing.T) {
	originalStr1 := "data_base_user"
	newStr1 := util.UnderLineToBigCamel(originalStr1)
	assert.Equal(t, "DataBaseUser", newStr1)

	originalStr2 := "_data_base_user"
	newStr2 := util.UnderLineToBigCamel(originalStr2)
	assert.Equal(t, "DataBaseUser", newStr2)

	originalStr3 := "_data_base_user_"
	newStr3 := util.UnderLineToBigCamel(originalStr3)
	assert.Equal(t, "DataBaseUser", newStr3)

	originalStr4 := "data_base_user_"
	newStr4 := util.UnderLineToBigCamel(originalStr4)
	assert.Equal(t, "DataBaseUser", newStr4)
}

func TestUpperUnderMiddleToSmallCamel(t *testing.T) {
	originalStr := "DATA-BASE-USER"
	newStr := util.UpperUnderMiddleToSmallCamel(originalStr)
	assert.Equal(t, "dataBaseUser", newStr)
}

func TestUpperUnderToSmallCamel(t *testing.T) {
	originalStr := "DATA_BASE_USER"
	newStr := util.UpperUnderToSmallCamel(originalStr)
	assert.Equal(t, "dataBaseUser", newStr)
}

func TestUpperUnderToBigCamel(t *testing.T) {
	originalStr := "DATA_BASE_USER"
	newStr := util.UpperUnderToBigCamel(originalStr)
	assert.Equal(t, "DataBaseUser", newStr)
}

func TestUpperMiddleToBigCamel(t *testing.T) {
	originalStr := "DATA-BASE-USER"
	newStr := util.UpperMiddleToBigCamel(originalStr)
	assert.Equal(t, "DataBaseUser", newStr)
}
