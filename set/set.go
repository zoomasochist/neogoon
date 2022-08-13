package set

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Set struct {
	Urls      []string `toml:"urls"`
	Texts     []string `toml:"texts"`
	Prompts   []string `toml:"prompts"`
	Filenames []string `toml:"filenames"`

	AllTexts []string

	Animated []string
	Audio    []string
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

	s.AllTexts = append(s.Texts, s.Prompts...)

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

	if err = Walk(&s.Audio, filepath.Join(outDir, "audio")); err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}

func DecompressZip(zipPath, out string) error {
	r, _ := zip.OpenReader(zipPath)
	defer r.Close()

	for _, f := range r.File {
		filePath := filepath.Join(out, f.Name)
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}
		defer destinationFile.Close()

		zippedFile, err := f.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		if _, err := io.Copy(destinationFile, zippedFile); err != nil {
			return err
		}
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
