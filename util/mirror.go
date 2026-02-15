package util

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	inv1       = "NoDisplay=true"
	inv2       = "Hidden=true"
	nameKey    = "Name="
	extDesktop = ".desktop"
)

type Icons []IconFile

type IconFile struct {
	Name       string
	OriginName string
	Data       *os.File
}

func (cfg *AppConfig) GetIcons(index int) (Icons, error) {
	if index < 0 || index >= len(cfg.SearchDirs) {
		return nil, os.ErrNotExist
	}

	searchDir := cfg.SearchDirs[index].Path

	entries, err := os.ReadDir(searchDir)
	if err != nil {
		return nil, err
	}

	var icons Icons

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != extDesktop {
			continue
		}

		fullPath := filepath.Join(searchDir, entry.Name())

		file, err := os.Open(fullPath)
		if err != nil {
			continue
		}

		if FileNotValid(file) {
			file.Close()
			continue
		}

		name, ok := IconNameFromFile(file)
		if !ok {
			name = IconNameFromNameFile(entry.Name())
		}

		icons = append(icons, IconFile{
			Name:       name,
			OriginName: entry.Name(),
			Data:       file,
		})
	}

	return icons, nil
}

func (icons *Icons) Close() error {
	var firstErr error

	for _, icon := range *icons {
		if icon.Data == nil {
			continue
		}

		if err := icon.Data.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}

	return firstErr
}

func FileNotValid(file *os.File) bool {
	defer file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if line == inv1 || line == inv2 {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		return true
	}

	return false
}

func IconNameFromFile(file *os.File) (string, bool) {
	defer file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)

	inDesktop := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if !inDesktop {
			if line == "[Desktop Entry]" {
				inDesktop = true
			}
			continue
		}

		if value, ok := strings.CutPrefix(line, "Name="); ok {
			return value, true
		}
	}
	
	return "", false
}

func IconNameFromNameFile(filename string) string {
	base := strings.TrimSuffix(filename, filepath.Ext(filename))
	if i := strings.LastIndex(base, "."); i != -1 {
		return base[i+1:]
	}
	return base
}

func (icons *Icons) AddIconsToDesktop(index int, desktop string) error {
	if index < 0 || index >= len(*icons) {
		return os.ErrNotExist
	}

	icon := (*icons)[index]
	fullpath := filepath.Join(desktop, icon.Name + ".desktop")

	dst, err := os.OpenFile(fullpath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, icon.Data); err != nil {
		return err
	}

	// habilitar confiança do arquivo
	cmd := exec.Command(
		"gio",
		"set",
		fullpath,
		"metadata::trusted",
		"true",
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
}