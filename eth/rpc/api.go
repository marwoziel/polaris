// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package rpc

import (
	"github.com/ethereum/go-ethereum/eth/filters"

	"pkg.berachain.dev/polaris/eth/rpc/api"
)

// GetAPIs returns a list of all available APIs.
func GetAPIs(apiBackend PolarisBackend) []API {
	nonceLock := new(AddrLocker)
	return []API{
		{
			Namespace: "eth",
			Service:   NewEthereumAPI(apiBackend),
		}, {
			Namespace: "eth",
			Service:   NewBlockChainAPI(apiBackend),
		},
		{
			Namespace: "eth",
			Service:   NewTransactionAPI(apiBackend, nonceLock),
		},
		{
			Namespace: "eth",
			Service:   api.NewEthashAPI(apiBackend),
		},
		{
			Namespace: "eth",
			// TODO: config must be setup properly.
			Service: NewFilterAPI(filters.NewFilterSystem(apiBackend, filters.Config{}), false),
		},
		{
			Namespace: "txpool",
			Service:   NewTxPoolAPI(apiBackend),
		},
		{
			Namespace: "debug",
			Service:   NewDebugAPI(apiBackend),
		},
		{
			Namespace: "net",
			Service:   api.NewNetAPI(apiBackend),
		},
		{
			Namespace: "web3",
			Service:   api.NewWeb3API(apiBackend),
		},
	}
}
