package util

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Zip 压缩
// zipFilePath：示例：./example/
// extractToPath：示例：example.zip
func Zip(sourceDir, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历文件夹
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 创建ZIP文件中的文件头部
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 调整文件路径，使其适应ZIP格式
		header.Name = filepath.ToSlash(path[len(sourceDir):])

		// 如果是目录，则设置目录头部
		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 如果是文件，则写入文件内容
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return zipWriter.Close()
}

// UnZip 函数解压 ZIP 文件到指定目录
// zipFilePath：示例：example.zip
// extractToPath：示例：example
func UnZip(zipFilePath, extractToPath string) error {
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer r.Close()

	err = os.MkdirAll(extractToPath, 0755)
	if err != nil {
		return err
	}

	// 遍历 ZIP 文件中的所有文件
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		// 创建文件路径
		fPath := filepath.Join(extractToPath, f.Name)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(fPath, f.Mode())
		} else {
			_ = os.MkdirAll(filepath.Dir(fPath), f.Mode())
			f, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
