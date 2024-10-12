package compress

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
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
