package main

import (
	"log"
	"net/http"
)

// NOTE: What is TLS?
/*
	TLS (Transport Layer Security) is a cryptographic protocol that provides encryption, authentication, and data integrity for secure communication over the internet. It is the successor to SSL (Secure Sockets Layer) and is used in HTTPS, email (SMTP, IMAP, POP3), and other secure protocols.

	TLS ensures:
		1.	Encryption: Protects data from being intercepted and read by unauthorized parties.
		2.	Authentication: Ensures that the communicating parties are who they claim to be.
		3.	Integrity: Prevents data from being tampered with during transmission.
*/

// NOTE: TLS Handshake Process
/*
	The TLS handshake is a process that establishes a secure connection between a client (e.g., a web browser) and a server (e.g., a website). It negotiates encryption settings, authenticates the server (and optionally the client), and securely exchanges keys for encrypting data.

	Steps of the TLS Handshake

	The handshake process varies slightly between TLS 1.2 and TLS 1.3, but the general steps are:

	1. Client Hello
		•	The client (browser) initiates the handshake by sending a ClientHello message to the server.
		•	This message includes:
		•	Supported TLS versions (e.g., TLS 1.2, TLS 1.3).
		•	List of supported cipher suites (encryption algorithms like AES, ChaCha20).
		•	A random number (used for secure key generation).

	2. Server Hello
		•	The server responds with a ServerHello message.
		•	This message contains:
		•	The TLS version chosen by the server.
		•	The cipher suite selected from the client’s list.
		•	Another random number (used in key generation).
		•	The server’s SSL/TLS certificate (to prove its identity).

	3. Server Certificate and Authentication
		•	The server sends its digital certificate, issued by a Certificate Authority (CA).
		•	The certificate contains:
		•	The server’s public key.
		•	Information about the CA that issued the certificate.
		•	The server’s domain name.
		•	The client verifies the certificate’s authenticity (using CA trust chains).

	4. Key Exchange (Pre-Master Secret)
		•	The client and server agree on a shared secret key for encrypting data.
		•	The method depends on the TLS version:
		•	TLS 1.2: Uses RSA or Diffie-Hellman (DH/ECDH) for key exchange.
		•	TLS 1.3: Uses Elliptic Curve Diffie-Hellman (ECDHE) for Perfect Forward Secrecy (PFS).

	5. Finished Messages
		•	Both client and server send a “Finished” message to confirm the handshake.
		•	All further communication is encrypted using the agreed-upon key.
*/

// NOTE: TLS 1.2 vs. TLS 1.3 Handshake Differences
/*
	Feature						TLS 1.2								TLS 1.3
	Number of Handshake Steps	4+									2
	Cipher Suites				Supports many, including RSA & DH	Only secure ones (ECDHE, AES-GCM, ChaCha20)
	Key Exchange				RSA, DH, ECDH						ECDHE (Perfect Forward Secrecy)
	Handshake Time				Slower								Faster (reduces round trips)
	Security					Older algorithms (SHA-1, MD5)		Stronger encryption, removes weak ciphers
*/

// NOTE: Why TLS 1.3 is Better?
/*
	•	Faster: Reduces handshake steps, making HTTPS connections quicker.
	•	More Secure: Removes weak encryption methods (RSA, SHA-1).
	•	Forward Secrecy: Uses ECDHE, ensuring past communications remain secure even if the private key is compromised.
*/

// NOTE: Conclusion
/*
	The TLS handshake is the foundation of secure web communication, ensuring that data sent over the internet is encrypted, authenticated, and tamper-proof. TLS 1.3 improves security and performance over TLS 1.2, making it the preferred protocol for modern web security.
*/

// NOTE: TLS vs SSL
/*
	TLS (Transport Layer Security) and SSL (Secure Sockets Layer) are cryptographic protocols designed to secure communication over networks, but TLS is the more modern and secure version.

	Key Differences Between TLS and SSL:
		1.	Evolution:
		•	SSL is the older protocol, with versions SSL 1.0, 2.0, and 3.0 (SSL 1.0 was never publicly released).
		•	TLS is the successor of SSL, with the first version being TLS 1.0 (essentially SSL 3.1). TLS has evolved through versions 1.1, 1.2, and the most recent 1.3, which is now the preferred protocol for secure communications.
		2.	Security Improvements:
		•	TLS fixes many vulnerabilities present in SSL, such as weaknesses in the handshake process, encryption algorithms, and padding.
		•	SSL 3.0 is considered obsolete and insecure due to various attacks (e.g., POODLE).
		•	TLS versions 1.1 and above offer stronger security mechanisms, and TLS 1.2 and 1.3 are widely used today.
		3.	Protocol Changes:
		•	TLS has changes to its encryption algorithms and mechanisms for message integrity, offering better cryptographic strength.
		•	SSL and TLS differ in their message formats, handshake processes, and supported cryptographic algorithms.
		4.	Deprecation:
		•	Most modern applications and services have migrated to TLS and no longer support SSL. SSL 2.0 and 3.0 have been deprecated by the IETF and are no longer recommended for use.

	Underlying Hashing Algorithms in TLS and SSL:

	Both SSL and TLS use hashing algorithms to ensure the integrity of the data being transmitted and to verify the authenticity of the message. However, TLS generally supports more secure hashing algorithms than SSL.

	SSL:
		•	SSL 3.0 uses the MD5 (Message Digest Algorithm 5) hashing algorithm in combination with SHA-1 for some parts of the handshake process. MD5 is now considered weak and vulnerable to collision attacks.

	TLS:
		•	TLS 1.0 to 1.2: In these versions, the main hashing algorithms used are SHA-1 and SHA-256, as part of the HMAC (Hash-based Message Authentication Code). However, SHA-1 has also been found vulnerable to collision attacks, and its use is being deprecated.
		•	TLS 1.3:
		•	SHA-256 is typically used for the key derivation and authentication process in TLS 1.3.
		•	TLS 1.3 does not use MD5 or SHA-1 at all. It favors stronger and more modern cryptographic algorithms.
		•	The new key exchange mechanism in TLS 1.3 also makes use of elliptic curve cryptography (ECC) and supports algorithms like SHA-256 and SHA-384 for hashing.

	Summary of Hashing Algorithms:
		•	SSL: Primarily MD5 (weak, deprecated) and SHA-1.
		•	TLS 1.0 to 1.2: SHA-1, SHA-256.
		•	TLS 1.3: SHA-256, SHA-384 (stronger, more secure).

	In conclusion, TLS is more secure than SSL due to its stronger cryptographic mechanisms, including the use of more secure hashing algorithms. TLS 1.2 and TLS 1.3 are widely recommended today for secure communications.
*/

// * create a cert
/*
	```
	go env GOROOT
	go run /usr/local/go/src/crypto/tls/generate_cert.go -host "localhost"

	or

	go run $(go env GOROOT)/src/crypto/tls/generate_cert.go -host "localhost"
	```
*/

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServeTLS(":10443", "08-https/cert.pem", "08-https/key.pem", nil)) // * 443 for production
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}
