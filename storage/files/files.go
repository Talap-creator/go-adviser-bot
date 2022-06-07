package files

import "go-adviser-bot/storage"

type Storage struct {
	basePath string
}

func New(basePatch string) Storage {
	return Storage{basePath: basePatch}

}

func (s Storage) Save(page *storage.Page) error {

}
