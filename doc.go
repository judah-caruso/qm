// Package qm is a performance aware math library for games.
//
// The goal is to create the highest performance math library for Go.
// Because of this, results for certain operations are less precise than
// their standard library counterpart. However, this isn't an issue as
// games do not require incredibly precise math in most cases.
//
// # Assembly
//
// In cases where it's meaningful to do so[1], qm uses assembly to improve
// performance. These assembly procedures take advantage of platform-specific
// instructions and are only available when compiling for certain platforms.
// Slower, pure Go alternatives are used when compiling for unimplemented
// targets (i.e. WebAssembly).
//
// Generally, qm uses avo[2] to generate assembly code, but checks these files
// into source control so users don't need to generate them manually.
// However, generating them manually can be done by navigating to the root directory
// of this package and running:
//
//	go generate -v ./internal
//
// # Notes
//
// As this library is in active development, performance should be
// expected to improve over time. In cases where a specific math function
// hasn't been implemented yet, qm uses Go's math package as a fallback to
// maintain coverage. Ideally, qm will have no dependency on the math package.
//
// [1] Because the Go compiler does not inline assembly functions,
// assembly is only used when:
//  1. Function call overhead does not meaningfully affect the performance;
//  2. The compiler fails to optimize a piece of code, and performance can be
//     improved even when the pure Go version is inlined;
//  3. The assembly version prevents unnecessary runtime checks inserted by the Go compiler;
//
// [2] avo only supports x86_64 assembly at the time of writing, so assembly for other platforms
// is written manually and follows the same pattern: [module]_[buildtag].[go|s].
package qm
