// Code generated by scripts/gen_go_std_tables.go; DO NOT EDIT.

// Generated from Go versions [go1.22.6].

package main

var runtimeAndDeps = map[string]bool{
	"internal/abi":             true, // go1.22.6
	"internal/bytealg":         true, // go1.22.6
	"internal/chacha8rand":     true, // go1.22.6
	"internal/coverage/rtcov":  true, // go1.22.6
	"internal/cpu":             true, // go1.22.6
	"internal/goarch":          true, // go1.22.6
	"internal/godebugs":        true, // go1.22.6
	"internal/goexperiment":    true, // go1.22.6
	"internal/goos":            true, // go1.22.6
	"runtime":                  true, // go1.22.6
	"runtime/internal/atomic":  true, // go1.22.6
	"runtime/internal/math":    true, // go1.22.6
	"runtime/internal/sys":     true, // go1.22.6
	"runtime/internal/syscall": true, // go1.22.6
	"unsafe":                   true, // go1.22.6
}

var runtimeLinknamed = []string{
	"arena",                          // go1.22.6
	"crypto/internal/boring",         // go1.22.6
	"crypto/internal/boring/bcache",  // go1.22.6
	"crypto/internal/boring/fipstls", // go1.22.6
	"crypto/x509/internal/macos",     // go1.22.6
	"internal/godebug",               // go1.22.6
	"internal/poll",                  // go1.22.6
	"internal/reflectlite",           // go1.22.6
	"internal/syscall/unix",          // go1.22.6
	"internal/syscall/windows",       // go1.22.6
	"maps",                           // go1.22.6
	"os",                             // go1.22.6
	"os/signal",                      // go1.22.6
	"plugin",                         // go1.22.6
	"reflect",                        // go1.22.6
	"runtime/coverage",               // go1.22.6
	"runtime/debug",                  // go1.22.6
	"runtime/metrics",                // go1.22.6
	"runtime/pprof",                  // go1.22.6
	"runtime/trace",                  // go1.22.6
	"sync",                           // go1.22.6
	"sync/atomic",                    // go1.22.6
	"syscall",                        // go1.22.6
	"syscall/js",                     // go1.22.6
	"time",                           // go1.22.6
	// The net package linknames to the runtime, not the other way around.
	// TODO: support this automatically via our script.
	"net",
}

var compilerIntrinsics = map[string]map[string]bool{
	"math": {
		"Abs":         true, // go1.22.6
		"Ceil":        true, // go1.22.6
		"Copysign":    true, // go1.22.6
		"FMA":         true, // go1.22.6
		"Floor":       true, // go1.22.6
		"Round":       true, // go1.22.6
		"RoundToEven": true, // go1.22.6
		"Trunc":       true, // go1.22.6
		"sqrt":        true, // go1.22.6
	},
	"math/big": {
		"mulWW": true, // go1.22.6
	},
	"math/bits": {
		"Add":             true, // go1.22.6
		"Add64":           true, // go1.22.6
		"Div":             true, // go1.22.6
		"Div64":           true, // go1.22.6
		"Len":             true, // go1.22.6
		"Len16":           true, // go1.22.6
		"Len32":           true, // go1.22.6
		"Len64":           true, // go1.22.6
		"Len8":            true, // go1.22.6
		"Mul":             true, // go1.22.6
		"Mul64":           true, // go1.22.6
		"OnesCount":       true, // go1.22.6
		"OnesCount16":     true, // go1.22.6
		"OnesCount32":     true, // go1.22.6
		"OnesCount64":     true, // go1.22.6
		"OnesCount8":      true, // go1.22.6
		"Reverse":         true, // go1.22.6
		"Reverse16":       true, // go1.22.6
		"Reverse32":       true, // go1.22.6
		"Reverse64":       true, // go1.22.6
		"Reverse8":        true, // go1.22.6
		"ReverseBytes16":  true, // go1.22.6
		"ReverseBytes32":  true, // go1.22.6
		"ReverseBytes64":  true, // go1.22.6
		"RotateLeft":      true, // go1.22.6
		"RotateLeft16":    true, // go1.22.6
		"RotateLeft32":    true, // go1.22.6
		"RotateLeft64":    true, // go1.22.6
		"RotateLeft8":     true, // go1.22.6
		"Sub":             true, // go1.22.6
		"Sub64":           true, // go1.22.6
		"TrailingZeros16": true, // go1.22.6
		"TrailingZeros32": true, // go1.22.6
		"TrailingZeros64": true, // go1.22.6
		"TrailingZeros8":  true, // go1.22.6
	},
	"runtime": {
		"publicationBarrier": true, // go1.22.6
	},
	"runtime/internal/atomic": {
		"And":             true, // go1.22.6
		"And8":            true, // go1.22.6
		"Cas":             true, // go1.22.6
		"Cas64":           true, // go1.22.6
		"CasRel":          true, // go1.22.6
		"Casint32":        true, // go1.22.6
		"Casint64":        true, // go1.22.6
		"Casp1":           true, // go1.22.6
		"Casuintptr":      true, // go1.22.6
		"Load":            true, // go1.22.6
		"Load64":          true, // go1.22.6
		"Load8":           true, // go1.22.6
		"LoadAcq":         true, // go1.22.6
		"LoadAcq64":       true, // go1.22.6
		"LoadAcquintptr":  true, // go1.22.6
		"Loadint32":       true, // go1.22.6
		"Loadint64":       true, // go1.22.6
		"Loadp":           true, // go1.22.6
		"Loaduint":        true, // go1.22.6
		"Loaduintptr":     true, // go1.22.6
		"Or":              true, // go1.22.6
		"Or8":             true, // go1.22.6
		"Store":           true, // go1.22.6
		"Store64":         true, // go1.22.6
		"Store8":          true, // go1.22.6
		"StoreRel":        true, // go1.22.6
		"StoreRel64":      true, // go1.22.6
		"StoreReluintptr": true, // go1.22.6
		"Storeint32":      true, // go1.22.6
		"Storeint64":      true, // go1.22.6
		"StorepNoWB":      true, // go1.22.6
		"Storeuintptr":    true, // go1.22.6
		"Xadd":            true, // go1.22.6
		"Xadd64":          true, // go1.22.6
		"Xaddint32":       true, // go1.22.6
		"Xaddint64":       true, // go1.22.6
		"Xadduintptr":     true, // go1.22.6
		"Xchg":            true, // go1.22.6
		"Xchg64":          true, // go1.22.6
		"Xchgint32":       true, // go1.22.6
		"Xchgint64":       true, // go1.22.6
		"Xchguintptr":     true, // go1.22.6
	},
	"runtime/internal/math": {
		"MulUintptr": true, // go1.22.6
	},
	"runtime/internal/sys": {
		"Bswap32":          true, // go1.22.6
		"Bswap64":          true, // go1.22.6
		"Len64":            true, // go1.22.6
		"Len8":             true, // go1.22.6
		"OnesCount64":      true, // go1.22.6
		"Prefetch":         true, // go1.22.6
		"PrefetchStreamed": true, // go1.22.6
		"TrailingZeros32":  true, // go1.22.6
		"TrailingZeros64":  true, // go1.22.6
		"TrailingZeros8":   true, // go1.22.6
	},
	"sync": {
		"runtime_LoadAcquintptr":  true, // go1.22.6
		"runtime_StoreReluintptr": true, // go1.22.6
	},
	"sync/atomic": {
		"AddInt32":              true, // go1.22.6
		"AddInt64":              true, // go1.22.6
		"AddUint32":             true, // go1.22.6
		"AddUint64":             true, // go1.22.6
		"AddUintptr":            true, // go1.22.6
		"CompareAndSwapInt32":   true, // go1.22.6
		"CompareAndSwapInt64":   true, // go1.22.6
		"CompareAndSwapUint32":  true, // go1.22.6
		"CompareAndSwapUint64":  true, // go1.22.6
		"CompareAndSwapUintptr": true, // go1.22.6
		"LoadInt32":             true, // go1.22.6
		"LoadInt64":             true, // go1.22.6
		"LoadPointer":           true, // go1.22.6
		"LoadUint32":            true, // go1.22.6
		"LoadUint64":            true, // go1.22.6
		"LoadUintptr":           true, // go1.22.6
		"StoreInt32":            true, // go1.22.6
		"StoreInt64":            true, // go1.22.6
		"StoreUint32":           true, // go1.22.6
		"StoreUint64":           true, // go1.22.6
		"StoreUintptr":          true, // go1.22.6
		"SwapInt32":             true, // go1.22.6
		"SwapInt64":             true, // go1.22.6
		"SwapUint32":            true, // go1.22.6
		"SwapUint64":            true, // go1.22.6
		"SwapUintptr":           true, // go1.22.6
	},
}

var reflectSkipPkg = map[string]bool{
	"fmt": true,
}
