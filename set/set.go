package set

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Downloader struct {
	Enabled      bool     `toml:"enabled"`
	Booru        string   `toml:"booru"`
	Tags         []string `toml:"tags"`
	MinimumScore int      `toml:"minimum-score"`
}

type Set struct {
	Urls       []string   `toml:"urls"`
	Texts      []string   `toml:"texts"`
	Downloader Downloader `toml:"downloader"`
	Filenames  []string   `toml:"filenames"`

	Animated []string
	Images   []string
	Videos   []string
}

func Load(s *Set, filePath string) error {
	cachePath, err := os.UserCacheDir()
	if err != nil {
		return err
	}

	outDir := filepath.Join(cachePath, "neogoon/set")
	if err = os.RemoveAll(outDir); err != nil {
		return err
	}

	if err = DecompressZip(filePath, outDir); err != nil {
		return err
	}

	setConfig := filepath.Join(outDir, "set.toml")
	_, err = toml.DecodeFile(setConfig, s)
	if err != nil {
		return err
	}

	if err = Walk(&s.Animated, filepath.Join(outDir, "animated")); err != nil &&
		!os.IsNotExist(err) {
		return err
	}

	if err = Walk(&s.Images, filepath.Join(outDir, "images")); err != nil && !os.IsNotExist(err) {
		return err
	}

	if err = Walk(&s.Videos, filepath.Join(outDir, "videos")); err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}

func DecompressZip(filePath, out string) error {
	r, _ := zip.OpenReader(filePath)
	defer r.Close()

	for _, f := range r.File {
		filePath := filepath.Join(out, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, 0644)
			continue
		}

		outWriter, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer outWriter.Close()

		archiveFile, _ := f.Open()
		defer archiveFile.Close()

		io.Copy(outWriter, archiveFile)
	}

	return nil
}

func Walk(slice *[]string, dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			*slice = append(*slice, path)
		}

		return nil
	})
}
