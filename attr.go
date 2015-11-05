// Copyright 2015 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tcell

// AttrMask represents a mask of text attributes, apart from color.
// Note that support for attributes may vary widely across terminals.
type AttrMask int

// NB: the colors listed here are in the order that ANSI terminals expect.

const (
	AttrBold AttrMask = 1 << (25 + iota)
	AttrBlink
	AttrReverse
	AttrUnderline
	AttrDim

	// AttrNone is just normal text.
	AttrNone AttrMask = 0
)

const attrAll = AttrBold | AttrBlink | AttrReverse | AttrUnderline | AttrDim
