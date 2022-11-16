# Arcology's Customized EVM

Arcology's customized EVM allows you to utilize Arcology's **Concurrency** library within EVM.

To achieve this purpose, we expose a **KernelAPI** interface to EVM. With this interface, EVM can redirect all the **Concurrency** calls to Arcology's unique parallel execution framework.

An extra object **'api'** that implements the **KernelAPI** interface is provided when initializing the EVM instance. This **'api'** knows whether a call to a contract is a classical Ethereum call or a brand new Arcology **Concurrency** call.

Besides that, this customized EVM is identical to the standard EVM, all the standard Ethereum smart contract can run on it seamlessly.
