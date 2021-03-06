// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build js

package monogame

import (
	"github.com/hajimehoshi/ebiten/internal/driver"
	"github.com/hajimehoshi/ebiten/internal/graphics"
)

type RenderTarget2D interface {
	SetAsDestination(viewportWidth, viewportHeight int)
	SetAsSource()
	ReplacePixels(args []*driver.ReplacePixelsArgs)
	Dispose()
	IsScreen() bool
}

type Image struct {
	v      RenderTarget2D
	g      *Graphics
	width  int
	height int
}

func (i *Image) Dispose() {
	i.v.Dispose()
}

func (*Image) IsInvalidated() bool {
	return false
}

func (*Image) Pixels() ([]byte, error) {
	panic("monogame: Pixels is not implemented yet")
	return nil, nil
}

func (i *Image) SetAsDestination() {
	w, h := i.width, i.height
	if !i.v.IsScreen() {
		w, h = graphics.InternalImageSize(w), graphics.InternalImageSize(h)
	}
	i.v.SetAsDestination(w, h)
}

func (i *Image) SetAsSource() {
	i.v.SetAsSource()
}

func (i *Image) ReplacePixels(args []*driver.ReplacePixelsArgs) {
	i.v.ReplacePixels(args)
}
