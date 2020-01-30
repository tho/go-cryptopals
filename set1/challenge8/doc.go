// Detect AES in ECB mode.
//
// In file https://cryptopals.com/static/challenge-data/8.txt are a bunch of
// hex-encoded ciphertexts.
//
// One of them has been encrypted with ECB.
//
// Detect it.
//
// Remember that the problem with ECB is that it is stateless and deterministic;
// the same 16 byte plaintext block will always produce the same 16 byte
// ciphertext.
package challenge8
