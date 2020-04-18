// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package asserter

import (
	"fmt"
	"reflect"

	"github.com/coinbase/rosetta-sdk-go/types"
)

// containsCurrency returns a boolean indicating if a
// *types.Currency is contained within a slice of
// *types.Currency. The check for equality takes
// into account everything within the types.Currency
// struct (including currency.Metadata).
func containsCurrency(currencies []*types.Currency, currency *types.Currency) bool {
	for _, curr := range currencies {
		if reflect.DeepEqual(curr, currency) {
			return true
		}
	}

	return false
}

// assertBalanceAmounts returns an error if a slice
// of types.Amount returned as an types.AccountIdentifier's
// balance is invalid. It is considered invalid if the same
// currency is returned multiple times (these shoould be
// consolidated) or if a types.Amount is considered invalid.
func assertBalanceAmounts(amounts []*types.Amount) error {
	currencies := make([]*types.Currency, 0)
	for _, amount := range amounts {
		// Ensure a currency is used at most once in balance.Amounts
		if containsCurrency(currencies, amount.Currency) {
			return fmt.Errorf("currency %+v used in balance multiple times", amount.Currency)
		}
		currencies = append(currencies, amount.Currency)

		if err := Amount(amount); err != nil {
			return err
		}
	}

	return nil
}

// AccountBalanceResponse returns an error if the provided
// types.BlockIdentifier is invalid, if the requestBlock
// is not nil and not equal to the response block, or
// if the same currency is present in multiple amounts.
func AccountBalanceResponse(
	requestBlock *types.PartialBlockIdentifier,
	responseBlock *types.BlockIdentifier,
	balances []*types.Amount,
) error {
	if err := BlockIdentifier(responseBlock); err != nil {
		return err
	}

	if err := assertBalanceAmounts(balances); err != nil {
		return err
	}

	if requestBlock == nil {
		return nil
	}

	if requestBlock.Hash != nil && *requestBlock.Hash != responseBlock.Hash {
		return fmt.Errorf(
			"request block hash %s does not match response block hash %s",
			*requestBlock.Hash,
			responseBlock.Hash,
		)
	}

	if requestBlock.Index != nil && *requestBlock.Index != responseBlock.Index {
		return fmt.Errorf(
			"request block index %d does not match response block index %d",
			*requestBlock.Index,
			responseBlock.Index,
		)
	}

	return nil
}
