package entries

import "gobanking/internal/common/types"

type EntryService interface {
	Create(*types.CreateEntryRequest) (types.CreateEntryResponse, error)
	GetAll() ([]*types.GetEntryResponse, error)
	GetById(int64) ([]*types.GetEntryResponse, error)
}
