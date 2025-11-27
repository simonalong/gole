package util

const bitsPerWord = 64 // 每个uint64可以存储的位数

// Bitmap 表示一个位图
type Bitmap struct {
	data []uint64
}

// NewBitmap 创建一个新的位图，初始化指定大小的位图
// 内容总是按照64的倍数存储的
// 小于64，则存储内容为64
// 大于64小于128，则存储内容为128
// 以此类推
func NewBitmap(size uint) *Bitmap {
	wordCount := (size + bitsPerWord - 1) / bitsPerWord // 计算需要多少个uint64
	return &Bitmap{
		data: make([]uint64, wordCount),
	}
}

// Size 当前bitmap可以存储多少个数据
func (b *Bitmap) Size() int64 {
	return int64(len(b.data) * bitsPerWord)
}

// Set 设置指定位置的位为1
func (b *Bitmap) Set(index uint64) bool {
	wordIndex := index / bitsPerWord // 计算位所在的uint64的索引
	bitIndex := index % bitsPerWord  // 计算位在uint64中的位置
	if wordIndex >= uint64(len(b.data)) {
		// 当前index值超过了bitmap最大容量
		return false
	}
	b.data[wordIndex] |= 1 << bitIndex // 设置位
	return true
}

// Clear 清除指定位置的位，将其设置为0
func (b *Bitmap) Clear(index uint64) bool {
	wordIndex := index / bitsPerWord
	bitIndex := index % bitsPerWord
	if wordIndex >= uint64(len(b.data)) {
		// 当前index值超过了bitmap最大容量
		return false
	}
	b.data[wordIndex] &^= 1 << bitIndex // 清除位
	return true
}

// Check 检查指定位置的位是否被设置
func (b *Bitmap) Check(index uint64) bool {
	wordIndex := index / bitsPerWord
	bitIndex := index % bitsPerWord
	if wordIndex >= uint64(len(b.data)) {
		// 当前index值超过了bitmap最大容量
		return false
	}
	return (b.data[wordIndex] & (1 << bitIndex)) != 0 // 检查位是否被设置
}
