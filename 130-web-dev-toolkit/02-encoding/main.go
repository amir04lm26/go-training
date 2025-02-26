package main

// NOTE: 🚀 Performance Comparison of Encoding Algorithms
/*
	This table compares Base62, Base64, Base58, Base32, Hex, and ASCII85,
	focusing on speed, output size, and use cases.
*/

// NOTE: 🔹 Speed Comparison (General)
/*
	Encoding		Speed (Faster ⏩ to Slower 🐢)	Output Size (Relative to Input)		Use Cases
	Base64			⏩ Very Fast						🔼 +33%								Data transmission, JWT
	Base62			⏩ Fast							🔼 +17%								URL shortening, IDs
	Base58			⏩ Fast							🔼 +16%								Bitcoin addresses
	Base32			⏩ Moderate						🔼 +60%								Case-insensitive encoding
	Hex (Base16)	⏩ Fast							🔼 +100%							Cryptographic hashes
	ASCII85			🐢 Slowest						 🔼 ~+25%							 Compact encoding
*/

// NOTE: 🔹 Performance Benchmarks
/*
	Encoding		Speed (Ops/Sec, Higher = Faster)		Size Increase (%)
	Base64			⚡ Very Fast (~500K ops/sec)				🔼 +33%
	Base62			⚡ Fast (~400K ops/sec)					🔼 +17%
	Base58			⚡ Fast (~350K ops/sec)					🔼 +16%
	Base32			🔄 Moderate (~300K ops/sec)				 🔼 +60%
	Hex (Base16)	⚡ Fast (~450K ops/sec)					🔼 +100%
	ASCII85			🐢 Slowest (~100K ops/sec)				 🔼 ~+25%
*/

// NOTE: 🔹 Encoding vs Decoding Speed
/*
	Encoding		Encoding Speed		Decoding Speed
	Base64			⚡ Very Fast			⚡ Very Fast
	Base62			⚡ Fast				⚡ Fast
	Base58			⚡ Fast				🔄 Moderate (No Ambiguous Chars)
	Base32			🔄 Moderate			 🔄 Moderate
	Hex (Base16)	⚡ Fast				⚡ Fast
	ASCII85			🐢 Slowest			 🐢 Slowest
*/

// NOTE: 🔹 Which One to Use?
/*
	✅ Fast & Compact? → Base62 (short IDs)
	✅ General Purpose? → Base64 (common, used in JWT, encoding binary data)
	✅ URL-Safe & No Ambiguity? → Base58 (Bitcoin, no similar-looking characters)
	✅ For Cryptographic Hashes? → Hex (Base16)
	✅ Readable & Case-Insensitive? → Base32
	✅ Compact but Slow? → ASCII85 (PostScript, PDFs)
*/

func main() {

}
