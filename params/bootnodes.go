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
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:30301",
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:0",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:30301",
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:0",
	// @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:30301",
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:0",
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:30301",
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:0",
	// Ethereum Foundation bootnode
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:30301",
	"enode://ee4bb1f622984d88239d5450be20dc27aea1a88e500b4cdf0e2065a7809cbfddab6fff2ed018dbdc4446d79d3884c2d6d61acfd02b57ec8084bd6b28be2f744a@127.0.0.1:0",
}
