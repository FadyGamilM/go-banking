package store

import "gobanking/internal/core/ports"

type DataStore struct {
	Account  *ports.AccountRepository
	Entry    *ports.EntryRepository
	Transfer *ports.TransferRepository
}
