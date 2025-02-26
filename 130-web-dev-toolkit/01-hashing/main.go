package main

// NOTE: 🔒 Comprehensive Comparison of Cryptographic Methods (Encryption, Hashing, HMAC, JWT, etc.)
/*
	This table covers HMAC, Bcrypt, SHA, MD5, AES, RSA, PBKDF2, Argon2, JWT, and more—detailing
	their purpose, security, speed, and best use cases.
*/

// NOTE: 🔹 Quick Overview
/*
	Algorithm	Type						Key Required?							Fixed Output?			 Reversible?	  Main Use Case
	HMAC		Keyed Hashing				✅ Yes (Secret Key)						✅ Yes					❌ No			API authentication, message integrity
	JWT			Token-based Authentication	✅ Yes (Secret or Public/Private Key)	❌ No					❌ No			Securely passing identity claims
	Bcrypt		Password Hashing			❌ No (uses Salt)						❌ No (variable-length)	❌ No			Password storage
	SHA-1		Hashing						❌ No									✅ Yes (160-bit)			❌ No			Data integrity (deprecated)
	SHA-256		Hashing						❌ No									✅ Yes (256-bit)			❌ No			File integrity, digital signatures
	MD5			Hashing						❌ No									✅ Yes (128-bit)			❌ No			File checksums (deprecated)
	AES			Symmetric Encryption		✅ Yes (Secret Key)						❌ No					✅ Yes			Encrypting files, SSL/TLS
	RSA			Asymmetric Encryption		✅ Yes (Public/Private Key)				❌ No					✅ Yes			Digital signatures, key exchange
	PBKDF2		Password Hashing			✅ Yes (Salt)							❌ No					❌ No			Key derivation, password hashing
	Argon2		Password Hashing			✅ Yes (Salt)							❌ No					❌ No			Modern password hashing
	ECDSA		Digital Signatures			✅ Yes (Public/Private Key)				❌ No					✅ Yes			Digital signatures, blockchain
*/

// NOTE: 🔹 In-Depth Comparison
/*
	Feature			HMAC			JWT				Bcrypt				SHA-1		SHA-256		MD5			AES						RSA						 PBKDF2				Argon2				ECDSA
	Type			Keyed Hashing	Token Auth		Password Hashing	Hashing		Hashing		Hashing		Symmetric Encryption	Asymmetric Encryption	 Key Derivation		Password Hashing	Digital Signatures
	Key Required?	✅ Yes			✅ Yes		   ❌ No				  ❌ No		 ❌ No		 ❌ No		✅ Yes					✅ Yes					✅ Yes			  ✅ Yes		 		  ✅ Yes
	Output Size		256-bit			Variable		Variable			160-bit		256-bit		128-bit		Variable				Variable				 Variable			Variable			Variable
	Reversible?		❌ No			❌ No		   ❌ No				  ❌ No		 ❌ No		 ❌ No		✅ Yes					✅ Yes					❌ No			  ❌ No		 		  ✅ Yes
	Secure?			✅ Yes			✅ Yes		   ✅ Yes			  ❌ No		 ✅ Yes		 ❌ No		✅ Yes					✅ Yes					✅ Yes			  ✅ Yes		 		  ✅ Yes
	Use Case		API auth		Secure claims	Passwords			Legacy		Integrity	Checksums	File encryption			Digital signatures		 Passwords			Passwords			Blockchain
*/

// NOTE: 🔹 Explanation of Each Algorithm
/*
	1️⃣ HMAC (Hash-based Message Authentication Code)
		•	Use Case: Ensures message integrity and authentication (e.g., API requests).
		•	How It Works: Combines a message and a secret key with a hash function (SHA-256, SHA-512, etc.).
		•	Example: Used in OAuth, JWT signing, and secure APIs.

	2️⃣ JWT (JSON Web Token)
		•	Use Case: Securely transmitting claims (e.g., user authentication).
		•	How It Works: Uses HMAC (symmetric) or RSA/ECDSA (asymmetric) for signing.
		•	Example: SSO (Single Sign-On), API authentication.
		•	Format: Header.Payload.Signature (Base64 encoded).

	3️⃣ Bcrypt
		•	Use Case: Secure password hashing.
		•	How It Works: Uses Blowfish-based hashing with automatic salt and multiple iterations.
		•	Why Secure?: Slow hashing makes brute-force attacks harder.
		•	Example: Used in password storage (e.g., databases).

	4️⃣ SHA-1 (Secure Hash Algorithm 1)
		•	Use Case: File integrity checks (but now insecure).
		•	❌ DO NOT use for security purposes—it is vulnerable to collision attacks.
		•	Example: Older digital signatures, deprecated in SSL/TLS.

	5️⃣ SHA-256 (Part of SHA-2 Family)
		•	Use Case: Secure file integrity, digital signatures.
		•	Why Secure?: Large 256-bit output, resistant to collisions.
		•	Example: Used in blockchain (Bitcoin), SSL/TLS.

	6️⃣ MD5 (Message Digest 5)
		•	Use Case: Legacy checksum verification (not secure for cryptographic use).
		•	❌ DO NOT use for passwords—highly vulnerable to collision attacks.
		•	Example: Used in file integrity checksums (but replaced by SHA-256).

	7️⃣ AES (Advanced Encryption Standard)
		•	Use Case: Encrypting sensitive data (e.g., files, SSL/TLS).
		•	How It Works: Symmetric encryption (same key for encryption & decryption).
		•	Example: Wi-Fi encryption (WPA2), VPNs, disk encryption.

	8️⃣ RSA (Rivest-Shamir-Adleman)
		•	Use Case: Secure key exchange, digital signatures.
		•	How It Works: Uses public/private key pairs (asymmetric encryption).
		•	Example: SSL/TLS certificates, secure email (PGP/GPG).

	9️⃣ PBKDF2 (Password-Based Key Derivation Function 2)
		•	Use Case: Password hashing & key derivation.
		•	How It Works: Uses many iterations of SHA-256 or SHA-512 for slow hashing.
		•	Example: Used in Wi-Fi WPA2-PSK, secure password storage.

	🔟 Argon2 (Winner of the Password Hashing Competition)
		•	Use Case: Modern, secure password hashing.
		•	How It Works: Uses memory-hard functions (harder to brute-force with GPUs).
		•	Example: Next-gen password hashing standard.

	1️⃣1️⃣ ECDSA (Elliptic Curve Digital Signature Algorithm)
		•	Use Case: Digital signatures in secure communications & blockchain.
		•	How It Works: Uses elliptic curve cryptography for efficient signing & verification.
		•	Example: Used in Bitcoin, SSL/TLS, SSH keys.
*/

// NOTE: 🔹 When to Use Each
/*
	✅ HMAC → Secure API authentication, message integrity.
	✅ JWT → Token-based authentication (e.g., OAuth, SSO).
	✅ Bcrypt / Argon2 / PBKDF2 → Secure password hashing (Argon2 > Bcrypt).
	✅ SHA-256 → File integrity, blockchain (better than MD5/SHA-1).
	✅ AES → Encrypting sensitive data (fast & secure).
	✅ RSA → Asymmetric encryption (SSL/TLS, digital signatures).
	✅ ECDSA → Blockchain transactions, digital signatures.

	❌ Avoid MD5 & SHA-1 for security—they are broken.
*/

// NOTE: 🔹 Speed Comparison (General)
/*
	Algorithm		Type					Speed (Faster ⏩ to Slower 🐢)		Main Factor Affecting Speed
	HMAC			Keyed Hashing			⏩ Very Fast							Hash function used (SHA-256, SHA-512, etc.)
	JWT				Token-based Auth		⏩ Fast								Depends on signature method (HMAC is faster than RSA)
	Bcrypt			Password Hashing		🐢 Slow								 Intentionally slow (work factor)
	SHA-1			Hashing					⏩ Fast								Fixed-length processing
	SHA-256			Hashing					⏩ Fast								More rounds than SHA-1
	MD5				Hashing					⏩ Very Fast							Simple computation
	AES				Symmetric Encryption	⏩ Very Fast							Block size, key length
	RSA				Asymmetric Encryption	🐢 Very Slow						 Large key size (2048-bit+)
	PBKDF2			Password Hashing		🐢 Slow								 Number of iterations
	Argon2			Password Hashing		🐢 Slowest							 Memory-hard function
	ECDSA			Digital Signatures		⏩ Faster than RSA					Curve size
*/

// NOTE: 🔹 Performance Benchmarks
/*
	Algorithm					Speed (Ops/Sec)						Security Level
	MD5							🔥 Fastest (~1M ops/sec)			❌ Insecure (broken)
	SHA-1						⚡ Fast (~500K ops/sec)			   ❌ Insecure (collision attacks)
	SHA-256						⚡ Fast (~250K ops/sec)			   ✅ Secure
	HMAC-SHA256					⚡ Fast (~200K ops/sec)			   ✅ Secure
	AES-128						⚡ Very Fast (~500K ops/sec)		   ✅ Secure
	AES-256						⚡ Fast (~250K ops/sec)			   ✅ Secure
	RSA-2048 Encryption			🐢 Very Slow (~500 ops/sec)			✅ Secure
	RSA-2048 Signing			🐢 Very Slow (~100 ops/sec)			✅ Secure
	ECDSA Signing				⚡ Faster than RSA (~1K ops/sec)	   ✅ Secure
	Bcrypt (Cost=10)			🐢 Slow (~100 ops/sec)				✅ Secure
	PBKDF2 (100K iterations)	🐢 Slow (~10 ops/sec)				✅ Secure
	Argon2 (High Memory Mode)	🐢 Slowest (~5 ops/sec)				✅ Most Secure
*/

// NOTE: 🔹 Which One to Use?
/*
	✅ Fast & Secure? → Use AES (encryption), SHA-256 (hashing), or HMAC.
	✅ Strong Password Hashing? → Use Argon2 > Bcrypt > PBKDF2.
	✅ Best for Signatures? → Use ECDSA > RSA (ECDSA is faster).
	✅ Avoid MD5 & SHA-1 for security—they are broken.
*/

func main() {

}
