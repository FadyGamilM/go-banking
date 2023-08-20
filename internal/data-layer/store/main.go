package store

import (
	account_ports "gobanking/internal/core-layer/ports/accounts"
	entry_ports "gobanking/internal/core-layer/ports/entries"
	transfer_ports "gobanking/internal/core-layer/ports/transfers"
)

type DataStore struct {
	Account  *account_ports.AccountRepository
	Entry    *entry_ports.EntryRepository
	Transfer *transfer_ports.TransferRepository
}

func NewDataStore(acc *account_ports.AccountRepository, entry *entry_ports.EntryRepository, transfer *transfer_ports.TransferRepository) *DataStore {
	return &DataStore{
		Account:  acc,
		Entry:    entry,
		Transfer: transfer,
	}
}
