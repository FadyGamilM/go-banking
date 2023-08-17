package usecases

import "gobanking/internal/data-layer/store"

type transferStore struct {
	Data *store.DataStore
}
