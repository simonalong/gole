package file

import (
	"github.com/simonalong/gole/util"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		return os.IsExist(err)
	}
	return true
}

func DirectoryExists(dirPath string) bool {
	if s, err := os.Stat(dirPath); err != nil {
		return false
	} else {
		return s.IsDir()
	}
}

func ExtractFilePath(filePath string) string {
	idx := strings.LastIndex(filePath, string(os.PathSeparator))
	return filePath[:idx]
}

func ExtractFileName(filePath string) string {
	idx := strings.LastIndex(filePath, string(os.PathSeparator))
	return filePath[idx+1:]
}

func ExtractFileExt(filePath string) string {
	idx := strings.LastIndex(filePath, ".")
	if idx != -1 {
		return filePath[idx+1:]
	}
	return ""
}

func ChangeFileExt(filePath string, ext string) string {
	mext := ExtractFileExt(filePath)
	if mext == "" {
		return filePath + "." + ext
	} else {
		return filePath[:len(filePath)-len(mext)] + ext
	}
}

func MkDirs(path string) bool {
	if !DirectoryExists(path) {
		return os.MkdirAll(path, os.ModePerm) == nil
	} else {
		return false
	}
}

func DeleteDirs(path string) error {
	return os.RemoveAll(path)
}

func DeleteFile(filePath string) error {
	return os.Remove(filePath)
}

func ReadFile(filePath string) string {
	var ret = ""
	if b, err := os.ReadFile(filePath); err == nil {
		ret = string(b)
	}
	return ret
}

func ReadFileBytes(filePath string) []byte {
	var ret []byte
	if b, err := os.ReadFile(filePath); err == nil {
		ret = b
	}
	return ret
}

func ReadFileLines(filePath string) []string {
	var ret []string
	if b, err := os.ReadFile(filePath); err == nil {
		ret = strings.Split(string(b), "\n")
	}
	return ret
}

func WriteFile(filePath string, text string) bool {
	return WriteFileBytes(filePath, []byte(text))
}

func WriteFileBytes(filePath string, data []byte) bool {
	p0 := ExtractFilePath(filePath)
	if !DirectoryExists(p0) {
		MkDirs(p0)
	}
	if FileExists(filePath) {
		DeleteFile(filePath)
	}
	if fl, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return false
	} else {
		_, err := fl.Write(data)
		_ = fl.Close()
		return err == nil
	}
}

func GetFile(filePath string) (*os.File, error) {
	p0 := ExtractFilePath(filePath)
	if !DirectoryExists(p0) {
		MkDirs(p0)
	}

	if fl, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644); err != nil {
		return nil, err
	} else {
		return fl, nil
	}
}

func AppendFile(filePath string, text string) bool {
	return AppendFileBytes(filePath, []byte(text))
}

func AppendFileBytes(filePath string, data []byte) bool {
	p0 := ExtractFilePath(filePath)
	if !DirectoryExists(p0) {
		MkDirs(p0)
	}
	if fl, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return false
	} else {
		_, err := fl.Write(data)
		_ = fl.Close()
		return err == nil
	}
}

func CopyFile(srcFilePath string, destFilePath string) error {
	p0 := ExtractFilePath(destFilePath)
	if !DirectoryExists(p0) {
		MkDirs(p0)
	}
	src, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}
	defer func(src *os.File) { _ = src.Close() }(src)
	dst, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer func(dst *os.File) { _ = dst.Close() }(dst)
	_, err = io.Copy(dst, src)
	return err
}

func CopyDirs(srcDirPath string, destDirPath string) error {
	err := filepath.Walk(srcDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			path = filepath.Clean(path)
			srcDirPath = filepath.Clean(srcDirPath)
			if srcDirPath == path {
				return nil
			}

			destDirPath = filepath.Clean(destDirPath)
			dir := getNewDstDir(srcDirPath, destDirPath, path)
			if !DirectoryExists(dir) {
				MkDirs(dir)
			}
		} else {
			dstFilePath := getNewDstDir(srcDirPath, destDirPath, path)
			err := CopyFile(path, dstFilePath)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func CopyDirsExcept(srcDirPath, destDirPath, exceptSrcDirPath string) error {
	srcDirPath, _ = filepath.Abs(srcDirPath)
	destDirPath, _ = filepath.Abs(destDirPath)
	exceptSrcDirPath, _ = filepath.Abs(exceptSrcDirPath)
	err := filepath.Walk(srcDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			path = filepath.Clean(path)
			srcDirPath = filepath.Clean(srcDirPath)
			if path == srcDirPath {
				return nil
			}

			exceptSrcDirPath = filepath.Clean(exceptSrcDirPath)
			if path == exceptSrcDirPath {
				return nil
			}

			destDirPath = filepath.Clean(destDirPath)
			dir := getNewDstDir(srcDirPath, destDirPath, path)
			if !DirectoryExists(dir) {
				MkDirs(dir)
			}
		} else {
			except := filepath.Clean(exceptSrcDirPath)
			if filepath.Dir(path) == except {
				return nil
			}

			dstFilePath := getNewDstDir(srcDirPath, destDirPath, path)
			err := CopyFile(path, dstFilePath)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func getNewDstDir(srcDirPath, destDirPath, path string) string {
	index := strings.Index(path, srcDirPath)
	if index == -1 {
		return path
	}
	return filepath.Join(destDirPath, path[len(srcDirPath):])
}

func RenameFile(srcFilePath string, destFilePath string) bool {
	p0 := ExtractFilePath(destFilePath)
	if !DirectoryExists(p0) {
		MkDirs(p0)
	}
	return os.Rename(srcFilePath, destFilePath) == nil
}

func CreateFile(filePath string) bool {
	if FileExists(filePath) {
		return true
	}

	p0 := ExtractFilePath(filePath)
	if !DirectoryExists(p0) {
		MkDirs(p0)
	}

	if _, err := os.OpenFile(filePath, os.O_CREATE, 0644); err != nil {
		return false
	} else {
		return true
	}
}

func Child(filePath string) ([]os.DirEntry, error) {
	return os.ReadDir(filePath)
}

// Size 返回文件/目录的大小
func Size(filePath string) int64 {
	if !DirectoryExists(filePath) {
		fi, err := os.Stat(filePath)
		if err == nil {
			return fi.Size()
		}
		return 0
	} else {
		var size int64
		err := filepath.Walk(filePath, func(_ string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				size += info.Size()
			}
			return err
		})
		if err != nil {
			return 0
		}
		return size
	}
}

func Sizes(filePaths ...string) int64 {
	var count int64
	for _, filePath := range filePaths {
		count += Size(filePath)
	}
	return count
}

func SizeList(filePaths []string) int64 {
	var count int64
	for _, filePath := range filePaths {
		count += Size(filePath)
	}
	return count
}

// SizeFormat 返回文件/目录的可读大小
func SizeFormat(filePath string) string {
	return util.FormatSize(Size(filePath))
}
