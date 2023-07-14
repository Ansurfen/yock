// Copyright 2023 The Yock Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package yockc

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/ansurfen/yock/util"
)

func Tar(src, dst string) error {
	fw, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fw.Close()

	gw := gzip.NewWriter(fw)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()
	root := filepath.Base(src)
	err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if src == path {
			return nil
		}
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		hdr.Name = root + "/" + strings.TrimPrefix(strings.ReplaceAll(relPath, "\\", "/"), "/")
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if info.Mode()&fs.ModeType == 0 {
			return nil
		}
		fr, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fr.Close()
		_, err = io.Copy(tw, fr)
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

func Zip(src, dst string) error {
	err := util.Mkdirs(filepath.Base(dst))
	if err != nil {
		return err
	}
	archive, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer archive.Close()

	zw := zip.NewWriter(archive)
	dst = strings.TrimSuffix(src, string(filepath.Separator))
	err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate
		header.Name, err = filepath.Rel(filepath.Dir(dst), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += string(filepath.Separator)
		}
		headerWriter, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		fp, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fp.Close()
		_, err = io.Copy(headerWriter, fp)
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

func Untar(src, dst string) error {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

RET:
	for {
		header, err := tarReader.Next()
		switch err {
		case nil:
			targetPath := strings.ReplaceAll(filepath.Join(dst, header.Name), "\\", "/")
			if header.Typeflag == tar.TypeDir {
				if err := util.Mkdirs(targetPath); err != nil {
					return err
				}
			} else if header.Typeflag == tar.TypeReg {
				if err := util.Mkdirs(filepath.Dir(targetPath)); err != nil {
					return err
				}
				fp, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY, header.FileInfo().Mode())
				if err != nil {
					return err
				}
				defer fp.Close()
				_, err = io.Copy(fp, tarReader)
				if err != nil {
					return err
				}
			} else {
				return util.ErrInvalidFile
			}
		case io.EOF:
			break RET
		default:
			return err
		}
	}
	return nil
}

func Unzip(src, dst string) error {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		filePath := filepath.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			if err := util.Mkdirs(filePath); err != nil {
				return err
			}
			continue
		}
		err := util.Mkdirs(filepath.Dir(filePath))
		if err != nil {
			return err
		}
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		w, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer w.Close()

		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
	}
	return nil
}
