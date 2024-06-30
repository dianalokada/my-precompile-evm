// SPDX-License-Identifier: MIT
pragma solidity >=0.8.18;

interface IBLSSignatureVerifier {
    /**
     * @dev Verifies a BLS signature
     * @param message The message that was signed (as a string)
     * @param signature The BLS signature
     * @param publicKey The BLS public key
     * @return isValid True if the signature is valid, false otherwise
     */
    
    function verifyBLSSignature(
        string calldata message,
        bytes calldata signature,
        bytes calldata publicKey

        // bytes calldata data
    ) external view returns (bool isValid);
}

