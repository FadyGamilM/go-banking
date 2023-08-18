package store

import (
	ports "gobanking/internal/core-layer/ports/accounts"
	ports "gobanking/internal/core-layer/ports/entries"
	ports "gobanking/internal/core-layer/ports/transfers"
)

type DataStore struct {
	Account  *ports.AccountRepository
	Entry    *ports.EntryRepository
	Transfer *ports.TransferRepository
}
