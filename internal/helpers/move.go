package helpers

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

func recreateFile(from, to string, noOverwrite bool) error {
	_, err := os.Stat(to)
	if err == nil && noOverwrite {
		AppLogger.Trace("File %v already exists and no overwrite is %v so not recreating file", to, noOverwrite)
		return nil
	}

	stats, err := os.Stat(from)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(from)
	if err != nil {
		return err
	}

	AppLogger.Trace("Copying file %v -> %v", from, to)
	return os.WriteFile(to, data, stats.Mode())
}

func recreateDirectory(from, to string, noOverwrite bool) error {
	_, err := os.Stat(to)
	if err == nil && noOverwrite {
		AppLogger.Trace("Directory %v already exists and no overwrite is %v so not recreating directory", to, noOverwrite)
		return nil
	}

	oldStats, err := os.Stat(from)
	if err != nil {
		return err
	}

	AppLogger.Trace("Creating Directory %v -> %v", from, to)
	return os.Mkdir(to, oldStats.Mode())
}

func recreateSymlink(from, to string, noOverwrite bool) error {
	_, err := os.Stat(to)
	if err == nil && noOverwrite {
		AppLogger.Trace("Symlink %v already exists and no overwrite is %v so not recreating symlink", to, noOverwrite)
		return nil
	}

	realFile, err := os.Readlink(from)
	if err != nil {
		return err
	}

	AppLogger.Trace("Creating Symlink %v -> %v for %v", from, to, realFile)
	return os.Symlink(realFile, to)
}

func MoveDir(from string, to string, noOverwrite bool) error {

	fileSystem := os.DirFS(from)
	root := "."

	err := fs.WalkDir(fileSystem, root, func(name string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		oldFile, err := filepath.Abs(path.Join(from, name))
		if err != nil {
			return err
		}

		newFile, err := filepath.Abs(path.Join(to, name))
		if err != nil {
			return err
		}

		switch d.Type() {
		case fs.ModeDir:
			{
				err := recreateDirectory(oldFile, newFile, noOverwrite)
				if err != nil {
					return err
				}
			}
		case fs.ModeSymlink:
			{
				err := recreateSymlink(oldFile, newFile, noOverwrite)
				if err != nil {
					return err
				}
			}
		default:
			if d.Type().IsRegular() {
				err := recreateFile(oldFile, newFile, noOverwrite)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return os.RemoveAll(from)
}
