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
    ) external view returns (bool isValid);

    /**
     * @dev Verifies a BLS signature and returns a human-readable result
     * @param message The message that was signed (as a string)
     * @param signature The BLS signature
     * @param publicKey The BLS public key
     * @return result A string indicating the verification result
     */
    function verifyAndFormatBLSSignature(
        string calldata message,
        bytes calldata signature,
        bytes calldata publicKey
    ) external view returns (string memory result);
}