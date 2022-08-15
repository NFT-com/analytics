package storage

import (
	"fmt"
)

// CurrencySymbol returns the symbol for the specified currency.
func (s *Storage) CurrencySymbol(currencyID string) (string, error) {

	query := s.db.
		Table("currencies").
		Select("symbol").
		Where("id = ?", currencyID).
		Limit(1)

	var symbol string
	err := query.Take(&symbol).Error
	if err != nil {
		return "", fmt.Errorf("could not retrieve symbol for the currency: %w", err)
	}

	return symbol, nil
}
