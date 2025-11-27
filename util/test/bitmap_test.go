package test

import (
	"github.com/magiconair/properties/assert"
	"github.com/simonalong/gole/util"
	"testing"
)

func TestBitMap1(t *testing.T) {
	bitmap := util.NewBitmap(1000)

	assert.Equal(t, true, bitmap.Set(10), "校验：加入10失败")
	assert.Equal(t, true, bitmap.Set(100), "校验：加入100失败")
	assert.Equal(t, true, bitmap.Set(1000), "校验：加入1000失败")

	assert.Equal(t, true, bitmap.Check(10), "校验：判断10存在失败")
	assert.Equal(t, true, bitmap.Check(100), "校验：判断100存在失败")
	assert.Equal(t, true, bitmap.Check(1000), "校验：判断1000存在失败")
}

func TestBitMap2(t *testing.T) {
	bitmap := util.NewBitmap(100)

	assert.Equal(t, true, bitmap.Set(10), "校验：加入10失败")
	assert.Equal(t, true, bitmap.Set(100), "校验：加入100失败")
	assert.Equal(t, false, bitmap.Set(1000), "校验：加入1000失败")

	assert.Equal(t, true, bitmap.Check(10), "校验：判断10存在失败")
	assert.Equal(t, true, bitmap.Check(100), "校验：判断100存在失败")
	assert.Equal(t, false, bitmap.Check(1000), "校验：判断1000存在失败")
}

// 小于64，则存储内容为64，也就是64以内的数据都可以存储
func TestBitMap3(t *testing.T) {
	bitmap := util.NewBitmap(1)

	assert.Equal(t, true, bitmap.Set(10), "校验：加入10失败")
	assert.Equal(t, true, bitmap.Set(63), "校验：加入63失败")
	assert.Equal(t, false, bitmap.Set(64), "校验：加入64失败")
	assert.Equal(t, false, bitmap.Set(100), "校验：加入100失败")
	assert.Equal(t, false, bitmap.Set(1000), "校验：加入1000失败")

	assert.Equal(t, true, bitmap.Check(10), "校验：判断10存在失败")
	assert.Equal(t, true, bitmap.Check(63), "校验：判断63存在失败")
	assert.Equal(t, false, bitmap.Check(64), "校验：判断64存在失败")
	assert.Equal(t, false, bitmap.Check(100), "校验：判断100存在失败")
	assert.Equal(t, false, bitmap.Check(1000), "校验：判断1000存在失败")
}

// 大于64小于128，则存储内容为128，也就是内容总是按照64的倍数存储的
func TestBitMap4(t *testing.T) {
	bitmap := util.NewBitmap(65)

	assert.Equal(t, true, bitmap.Set(10), "校验：加入10失败")
	assert.Equal(t, true, bitmap.Set(64), "校验：加入64失败")
	assert.Equal(t, true, bitmap.Set(127), "校验：加入127失败")
	assert.Equal(t, false, bitmap.Set(128), "校验：加入128失败")
	assert.Equal(t, false, bitmap.Set(1000), "校验：加入1000失败")

	assert.Equal(t, true, bitmap.Check(10), "校验：判断10存在失败")
	assert.Equal(t, true, bitmap.Check(64), "校验：判断64存在失败")
	assert.Equal(t, true, bitmap.Check(127), "校验：判断127存在失败")
	assert.Equal(t, false, bitmap.Check(128), "校验：判断128存在失败")
	assert.Equal(t, false, bitmap.Check(1000), "校验：判断1000存在失败")
}

func TestBitMap5(t *testing.T) {
	bitmap := util.NewBitmap(127)

	assert.Equal(t, true, bitmap.Set(10), "校验：加入10失败")
	assert.Equal(t, true, bitmap.Set(64), "校验：加入64失败")
	assert.Equal(t, true, bitmap.Set(127), "校验：加入127失败")
	assert.Equal(t, false, bitmap.Set(128), "校验：加入128失败")
	assert.Equal(t, false, bitmap.Set(1000), "校验：加入1000失败")

	assert.Equal(t, true, bitmap.Check(10), "校验：判断10存在失败")
	assert.Equal(t, true, bitmap.Check(64), "校验：判断64存在失败")
	assert.Equal(t, true, bitmap.Check(127), "校验：判断127存在失败")
	assert.Equal(t, false, bitmap.Check(128), "校验：判断128存在失败")
	assert.Equal(t, false, bitmap.Check(1000), "校验：判断1000存在失败")
}
