// Copyright 2017 The ebakus/node Authors
// This file is part of the ebakus/node library.
//
// The ebakus/node library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The ebakus/node library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the ebakus/node library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xf8E\xc2\x01\x01\xc2\x02\x01\xc2\x03\x01\xc2\x04\x01\xc2\x05\x01\xc2\x06\x01\xc2\a\x01\xc2\b\x01\xc2\t\x01\u0102\x01\x01\x01\u0102\x01\x02\x01\u07d4S\xe5\x18.y\x8c]*\x93\xc5&\xf7\xc9tZi7\xa7\x1f\u03c965\u026d\xc5\u07a0\x00\x00"
const testnetAllocData = "\xf8E\xc2\x01\x01\xc2\x02\x01\xc2\x03\x01\xc2\x04\x01\xc2\x05\x01\xc2\x06\x01\xc2\a\x01\xc2\b\x01\xc2\t\x01\u0102\x01\x01\x01\u0102\x01\x02\x01\u07d4\xd5=\u70f1/\x12\xb7\x85#H\xf2\xe3,!\x99\aF\xad\x02\x8965\u026d\xc5\u07a0\x00\x00"