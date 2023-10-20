// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package nativeminter

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

/* NOTE: Events can only be emitted in state-changing functions. So you cannot use events in read-only (view) functions.
Events are generally emitted at the end of a state-changing function with AddLog method of the StateDB. The AddLog method takes 4 arguments:
	1. Address of the contract that emitted the event.
	2. Topic hashes of the event.
	3. Encoded non-indexed data of the event.
	4. Block number at which the event was emitted.
The first argument is the address of the contract that emitted the event.
Topics can be at most 4 elements, the first topic is the hash of the event signature and the rest are the indexed event arguments. There can be at most 3 indexed arguments.
Topics cannot be fully unpacked into their original values since they're 32-bytes hashes.
The non-indexed arguments are encoded using the ABI encoding scheme. The non-indexed arguments can be unpacked into their original values.
You can use the following code to emit an event in your state-changing precompile functions (generated packer might be different)):
topics, data, err := PackMyEvent(
	topic1,
	topic2,
	data1,
	data2,
)
if err != nil {
	return nil, remainingGas, err
}
accessibleState.GetStateDB().AddLog(
	ContractAddress,
	topics,
	data,
	accessibleState.GetBlockContext().Number().Uint64(),
)
*/

// PackMintNativeCoinEvent packs the event into the appropriate arguments for MintNativeCoin.
// It returns topic hashes and the encoded non-indexed data.
func PackMintNativeCoinEvent(sender common.Address, addr common.Address, amount *big.Int) ([]common.Hash, []byte, error) {
	return NativeMinterABI.PackEvent("MintNativeCoin", sender, addr, amount)
}

// UnpackMintNativeCoinEvent won't be generated because the event does not have any non-indexed data.
