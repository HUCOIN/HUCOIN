// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://21e167d604bfb5c145c4c2aad1b594742a8f3f801b3e2b0f110791517680ffa0f81c21fa25b5b393d9e84d8a3b59f001cd06580a1aa2f51c2d62c5dac6cc1f47@46.101.114.157:58968",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://21e167d604bfb5c145c4c2aad1b594742a8f3f801b3e2b0f110791517680ffa0f81c21fa25b5b393d9e84d8a3b59f001cd06580a1aa2f51c2d62c5dac6cc1f47@46.101.114.157:58968",
	// @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://21e167d604bfb5c145c4c2aad1b594742a8f3f801b3e2b0f110791517680ffa0f81c21fa25b5b393d9e84d8a3b59f001cd06580a1aa2f51c2d62c5dac6cc1f47@46.101.114.157:58968",
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://21e167d604bfb5c145c4c2aad1b594742a8f3f801b3e2b0f110791517680ffa0f81c21fa25b5b393d9e84d8a3b59f001cd06580a1aa2f51c2d62c5dac6cc1f47@46.101.114.157:58968",
	// Ethereum Foundation bootnode
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://21e167d604bfb5c145c4c2aad1b594742a8f3f801b3e2b0f110791517680ffa0f81c21fa25b5b393d9e84d8a3b59f001cd06580a1aa2f51c2d62c5dac6cc1f47@46.101.114.157:58968",
}
