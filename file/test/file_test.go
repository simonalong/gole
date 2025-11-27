package test

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/simonalong/gole/util"
	"os"
	"path/filepath"
	"testing"

	"github.com/simonalong/gole/file"
)

func TestFile(t *testing.T) {
	// file.WriteFile("./sample.txt", "aaa")
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "sample.txt")

	file.AppendFile(path, "ccc")
	file.DeleteFile("sample.txt")
}

func TestExtract(t *testing.T) {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "sample.txt")
	p0 := file.ExtractFilePath(path)
	t.Logf("p0: %s", p0)
	n0 := file.ExtractFileName(path)
	t.Logf("n0: %s", n0)
	e0 := file.ExtractFileExt(path)
	t.Logf("e0: %s", e0)
	c0 := file.ChangeFileExt(path, "xyz")
	t.Logf("c0: %s", c0)
}

func TestCreatFile(t *testing.T) {
	file.CreateFile("./file/test.txt")
	file.CreateFile("./test2.txt")

	file.DeleteFile("./test2.txt")
	file.DeleteDirs("./file/")
}

func TestChild(t *testing.T) {
	f, _ := file.Child("../")
	for i := range f {
		t.Logf("file_name: %s", f[i].Name())
	}
}

func TestFileSize(t *testing.T) {
	assert.Equal(t, util.ToInt64(40), file.Size("./assert_file_size.txt"))
}

func TestFileFormatSize(t *testing.T) {
	assert.Equal(t, "40.00B", file.SizeFormat("./assert_file_size.txt"))
}

func TestFileCopy(t *testing.T) {
	file.DeleteFile("./demo2/src.txt")
	file.CopyFile("./demo/src.txt", "./demo2/src.txt")
}

func TestCopyDirs1(t *testing.T) {
	err := file.CopyDirs("./srcDir/", "./dstDir/")
	if err != nil {
		t.Fatal(err)
	}
}

// 测试拷贝绝对目录
func TestCopyDirs2(t *testing.T) {
	err := file.CopyDirs("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/", "/Users/zhouzhenyong/project/seatak/gole/file/test/dstDir/")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCopyDirsExcept1(t *testing.T) {
	err := file.CopyDirsExcept("./srcDir/", "./dstDir/", "./srcDir/demo2/")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCopyDirsExcept2(t *testing.T) {
	err := file.CopyDirsExcept("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/",
		"/Users/zhouzhenyong/project/seatak/gole/file/test/dstDir/",
		"/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo2/")
	if err != nil {
		t.Fatal(err)
	}
}

func TestFileWalk(t *testing.T) {
	filepath.Walk("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		fmt.Println(filepath.Clean(path))
		fmt.Println(info.Name())
		fmt.Println(info.IsDir())
		fmt.Println("---")
		return nil
	})
}

func TestFilepath(t *testing.T) {
	//Clean =  srcDir/demo1/text.txt
	fmt.Println("Clean = ", filepath.Clean("./srcDir/demo1/text.txt"))
	///Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.txt <nil>
	fmt.Println(filepath.Abs("./srcDir/demo1/text.txt"))
	//Ext =  .txt
	fmt.Println("Ext = ", filepath.Ext("./srcDir/demo1/text.txt"))
	//Base =  text.txt
	fmt.Println("Base = ", filepath.Base("./srcDir/demo1/text.txt"))
	//Dir =  srcDir/demo1
	fmt.Println("Dir = ", filepath.Dir("./srcDir/demo1/text.txt"))
	//FromSlash =  ./srcDir/demo1/text.txt
	fmt.Println("FromSlash = ", filepath.FromSlash("./srcDir/demo1/text.txt"))
	//IsLocal =  true
	fmt.Println("IsLocal = ", filepath.IsLocal("./srcDir/demo1/text.txt"))

	// ----
	fmt.Println("----")
	//Clean =  /Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx
	fmt.Println("Clean = ", filepath.Clean("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx"))
	///Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx <nil>
	fmt.Println(filepath.Abs("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx"))
	//Ext =  .tx
	fmt.Println("Ext = ", filepath.Ext("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx"))
	//Base =  text.tx
	fmt.Println("Base = ", filepath.Base("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx"))
	//Dir =  /Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1
	fmt.Println("Dir = ", filepath.Dir("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx"))
	//FromSlash =  /Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx
	fmt.Println("FromSlash = ", filepath.FromSlash("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx"))
	//IsLocal =  false
	fmt.Println("IsLocal = ", filepath.IsLocal("/Users/zhouzhenyong/project/seatak/gole/file/test/srcDir/demo1/text.tx"))
}
