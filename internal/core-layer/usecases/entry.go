package usecases

import "gobanking/internal/data-layer/store"

type entryService struct {
	Data *store.DataStore
}
