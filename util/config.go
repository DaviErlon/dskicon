package util

import (
	"encoding/json"
	"os"
	"path/filepath"
)

var configFilePath string

func init() {
	path, err := initFilePath()
	if err != nil {
		panic(err)
	}
	configFilePath = path
}

func initFilePath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(exe), "dskicon.json"), nil
}

type AppConfig struct {
	DesktopDir string           `json:"desktop"`
	SearchDirs []SearchDirEntry `json:"search"`
}

type SearchDirEntry struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func LoadConfig() (*AppConfig, error) {
	var cfg AppConfig

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			data, err = json.MarshalIndent(cfg, "", "    ")
			if err != nil {
				return nil, err
			}

			if err = os.WriteFile(configFilePath, data, 0644); err != nil {
				return nil, err
			}

			return &cfg, nil
		}
		return nil, err
	}

	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *AppConfig) IsConfigured() bool {
	return cfg.DesktopDir != ""
}

func (cfg *AppConfig) save() error {
	data, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFilePath, data, 0644)
}

func (cfg *AppConfig) SetDesktopDir(dir string) error {
	cfg.DesktopDir = dir
	return cfg.save()
}

func (cfg *AppConfig) AddSearchDir(name, dirPath string) error {
	cfg.SearchDirs = append(cfg.SearchDirs, SearchDirEntry{
		Name: name,
		Path: dirPath,
	})
	return cfg.save()
}

func (cfg *AppConfig) UpdateSearchDir(index int, dirPath string) error {
	if index < 0 || index >= len(cfg.SearchDirs) {
		return os.ErrNotExist
	}
	cfg.SearchDirs[index].Path = dirPath
	return cfg.save()
}

func (cfg *AppConfig) RemoveSearchDir(index int) error {
	if index < 0 || index >= len(cfg.SearchDirs) {
		return os.ErrNotExist
	}
	cfg.SearchDirs = append(
		cfg.SearchDirs[:index],
		cfg.SearchDirs[index+1:]...,
	)
	return cfg.save()
}
