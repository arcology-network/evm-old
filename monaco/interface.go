package monaco

import (
	"github.com/HPISTechnologies/evm/common"
)

// KernelAPI provides system level function calls supported by Monaco platform.
type KernelAPI interface {
	AddLog(key, value string)
	IsKernelAPI(addr common.Address) bool
	Call(caller, callee common.Address, input []byte, origin common.Address, nonce uint64, blockhash common.Hash) ([]byte, bool)
}
