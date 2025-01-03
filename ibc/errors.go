// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/aizel/aizel/blob/main/LICENSE)
package ibc

import "errors"

var (
	ErrNoIBCVoucherDenom  = errors.New("denom is not an IBC voucher")
	ErrDenomTraceNotFound = errors.New("denom trace not found")
	ErrInvalidBaseDenom   = errors.New("invalid base denomination")
)
