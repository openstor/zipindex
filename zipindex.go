// SPDX-FileCopyrightText: 2025 openstor contributors
// SPDX-FileCopyrightText: 2015-2025 MinIO, Inc.
// SPDX-License-Identifier: Apache-2.0

// Package zipindex provides a size optimized representation of a zip file to allow
// decompressing the file without reading the zip file index.
//
// It will only provide the minimal needed data for successful decompression and CRC checks.
//
// Custom metadata can be stored per file and filtering can be performed on the incoming files.
package zipindex

import "errors"

// MaxFiles is the maximum number of files inside a zip file.
const MaxFiles = 1_000_000_000

// ErrTooManyFiles is returned when a zip file contains too many files.
var ErrTooManyFiles = errors.New("too many files")

// MaxCustomEntries is the maximum number of custom entries per file.
const MaxCustomEntries = 1000

// ErrTooManyCustomEntries is returned when a zip file custom
// entry has too many entries.
var ErrTooManyCustomEntries = errors.New("custom entry count exceeded")

// MaxIndexSize is the maximum index size, uncompressed.
const MaxIndexSize = 128 << 20

// ErrMaxSizeExceeded is returned if the maximum size of data is exceeded.
var ErrMaxSizeExceeded = errors.New("index maximum size exceeded")
