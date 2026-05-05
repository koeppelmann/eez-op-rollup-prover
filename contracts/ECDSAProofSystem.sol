// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {IProofSystem} from "./IProofSystem.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

/// @title ECDSAProofSystem
/// @notice IProofSystem implementation that accepts a 65-byte ECDSA signature over the raw
///         publicInputsHash digest (no EIP-191 prefix). Same role tmpECDSAVerifier had under
///         the old IZKVerifier interface — adapted to the multi-prover model where each
///         rollup's manager registers an IProofSystem address + vkey in its allowed set.
/// @dev `proof` layout: `abi.encodePacked(r, s, v)` with v ∈ {27, 28}.
contract ECDSAProofSystem is IProofSystem, Ownable {
    address public signer;

    event SignerUpdated(address indexed previousSigner, address indexed newSigner);

    constructor(address initialOwner, address initialSigner) Ownable(initialOwner) {
        signer = initialSigner;
    }

    function setSigner(address newSigner) external onlyOwner {
        emit SignerUpdated(signer, newSigner);
        signer = newSigner;
    }

    function verify(bytes calldata proof, bytes32 publicInputsHash) external view returns (bool) {
        return ECDSA.recover(publicInputsHash, proof) == signer;
    }
}
