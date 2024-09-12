// Package qm is a performance aware math library for games.
//
// The goal is to create the highest performance math library for Go.
// Because of this, results for certain operations are less precise than
// their standard library counterpart. However, this isn't an issue as
// games do not require incredibly precise math in most cases.
//
// # Assembly
//
// In cases where it's meaningful to do so, qm uses assembly to improve
// performance. These assembly procedures take advantage of platform-specific
// instructions and are only available when compiling for certain platforms.
// Slower, pure Go alternatives are used when compiling for second-class targets (i.e. WebAssembly).
//
// Generally, qm uses avo to generate assembly code, but checks these files
// into source control so users don't need to generate them manually.
// However, generating them manually can be done by navigating to the root directory
// of this package and running:
//
//	go generate -v ./generate
//
// # Notes
//
// As this library is in active development, performance should be
// expected to improve over time. In cases where a specific math function
// hasn't been implemented yet, qm uses Go's math package as a fallback to
// maintain coverage. Ideally, qm will have no dependency on the math package.
package qm
