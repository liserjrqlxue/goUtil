package compress

import (
	"archive/zip"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// AddToZip adds a file to the zip archive.
// from chatGPT4 with modify
func AddToZip(zw *zip.Writer, filePath, fileName string) error {
	fileToZip, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	if fileName == "" {
		header.Name = filePath
	} else {
		header.Name = fileName
	}
	header.Method = zip.Deflate

	writer, err := zw.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

func ZipFiles(target string, source ...string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	// Create a new zip writer
	zw := zip.NewWriter(zipfile)
	defer zw.Close()

	// Add files to zip
	for _, fileName := range source {
		err := AddToZip(zw, fileName, fileName)
		if err != nil {
			return err
		}
	}
	return nil
}

func ZipDir(target, source string, filter func(string) bool) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	// Create a new zip writer
	zw := zip.NewWriter(zipfile)
	defer zw.Close()

	err = filepath.WalkDir(source, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !filter(path) {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		name, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		// Create a new file header
		// Add file to zip
		err = AddToZip(zw, path, name)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// ------------------ 纯 Go 实现 ------------------
func ZipPureGo(outputZip string, files []string, rootDirAbs string) error {
	outFile, err := os.Create(outputZip)
	if err != nil {
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	for _, src := range files {
		err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// 计算相对路径（基于 rootDirAbs）
			pathAbs, _ := filepath.Abs(path)
			relPath, err := filepath.Rel(rootDirAbs, pathAbs)
			if err != nil {
				return err
			}
			relPath = filepath.ToSlash(relPath) // zip 内必须用 `/`

			if info.IsDir() {
				// 保留空目录（必须以 `/` 结尾）
				if relPath != "" {
					_, err := zipWriter.Create(relPath + "/")
					if err != nil {
						return err
					}
				}
				return nil
			}

			// 添加文件
			fw, err := zipWriter.Create(relPath)
			if err != nil {
				return err
			}
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = io.Copy(fw, f)
			return err
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// ------------------ 系统调用实现 ------------------
func ZipBySystem(outputZip string, files []string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows: PowerShell Compress-Archive
		args := []string{"Compress-Archive"}
		for _, f := range files {
			args = append(args, "-Path", f)
		}
		args = append(args, "-DestinationPath", outputZip, "-Force")
		cmd = exec.Command("powershell", args...)
	} else {
		// Linux/macOS: zip -r output.zip files...
		args := append([]string{"-r", outputZip}, files...)
		cmd = exec.Command("zip", args...)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
