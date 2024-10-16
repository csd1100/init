package helpers

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

func recreateFile(from, to string) error {
	AppLogger.Trace("Copying file %v -> %v", from, to)
	stats, err := os.Stat(from)
	if err != nil {
		return err
	}

	data, err := os.ReadFile(from)
	if err != nil {
		return err
	}

	return os.WriteFile(to, data, stats.Mode())
}

func recreateDirectory(from, to string) error {
	AppLogger.Trace("Creating Directory %v -> %v", from, to)

	oldStats, err := os.Stat(from)
	if err != nil {
		return err
	}

	return os.Mkdir(to, oldStats.Mode())
}

func recreateSymlink(from, to string) error {
	realFile, err := os.Readlink(from)
	if err != nil {
		return err
	}

	AppLogger.Trace("Creating Symlink %v -> %v for %v", from, to, realFile)

	return os.Symlink(realFile, to)
}

func MoveDir(from string, to string) error {

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
				err := recreateDirectory(oldFile, newFile)
				if err != nil {
					return err
				}
			}
		case fs.ModeSymlink:
			{
				err := recreateSymlink(oldFile, newFile)
				if err != nil {
					return err
				}
			}
		default:
			if d.Type().IsRegular() {
				err := recreateFile(oldFile, newFile)
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
