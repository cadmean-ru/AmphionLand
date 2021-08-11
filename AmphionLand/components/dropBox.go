package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/common"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/common/atext"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"github.com/cadmean-ru/amphion/rendering"
)

type DropItem struct {
	index int
	text  string
}

func (s DropItem) Index() int {
	return s.index
}

func (s DropItem) Text() string {
	return s.text
}

type DropBoxGroup struct {
	builtin.GridLayout
	selectedIndex int
	items         []DropItem
	initialized   bool
	fullMode      bool
	listener      func(item DropItem)
}

func (s *DropBoxGroup) OnInit(ctx engine.InitContext) {
	s.GridLayout.OnInit(ctx)
	s.AddColumn(a.FillParent)
	s.AutoExpansion = true
	s.AutoShrinking = true

	s.fullMode = false

	s.initialized = true
}

func (s *DropBoxGroup) OnStart() {
	s.updateItems()
}

func (s *DropBoxGroup) SetOnItemSelectedListener(listener func(item DropItem)) {
	s.listener = listener
}

func (s *DropBoxGroup) AddItem(text string) {
	s.items = append(s.items, DropItem{
		index: len(s.items),
		text:  text,
	})

	if !s.initialized {
		return
	}

	s.updateItems()
}

func (s *DropBoxGroup) IsSelected(index int) bool {
	if s.selectedIndex == index {
		return true
	} else {
		return false
	}
}

func (s *DropBoxGroup) SetSelected(index int) {
	s.selectedIndex = index

	if !s.initialized {
		return
	}

	s.updateItems()
}

func (s *DropBoxGroup) GetItemIndex(text string) int {
	for i := 0; i < len(s.items); i++ {
		if s.items[i].text == text {
			return s.items[i].index
		}
	}
	return -1
}

func (s *DropBoxGroup) AddDropItem (i int){
	itemObj := engine.NewSceneObject(fmt.Sprintf("Item %d", i))
	itemObj.AddComponent(NewDropBox(s.items[i]))
	itemObj.AddComponent(builtin.NewRectBoundary())
	itemObj.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
		obj, ok := event.Sender.(*engine.SceneObject)
		if !ok {
			return true
		}

		drop := GetDropBox(obj)
		s.fullMode = false
		s.SetSelected(drop.item.index)

		if s.listener != nil {
			s.listener(drop.item)
		}

		return true
	}))
	itemObj.SetSizeXy(a.MatchParent, 20)
	s.SceneObject.AddChild(itemObj)
}

func (s *DropBoxGroup) updateItems() {
	if s.fullMode {
		if len(s.items) + 1 > s.SceneObject.GetChildrenCount() {
			//s.AddDropItem(s.selectedIndex)
			for i := s.SceneObject.GetChildrenCount(); i < len(s.items); i++ {
				s.AddDropItem(i)
			}
		} else if len(s.items) + 1 < s.SceneObject.GetChildrenCount() {
			for i := s.SceneObject.GetChildrenCount() - 1; i >= len(s.items); i-- {
				c := s.SceneObject.GetChildren()
				s.SceneObject.RemoveChild(c[i])
			}
		}

		children := s.SceneObject.GetChildren()
		for i, item := range s.items {
			c := children[i]
			b := GetDropBox(c, true)
			b.setItem(item)
		}
	} else {
		//if s.SceneObject.GetChildrenCount() > 1 {
		//	for i := s.SceneObject.GetChildrenCount() - 1; i > 1; i-- {
		//		c := s.SceneObject.GetChildren()
		//		s.SceneObject.RemoveChild(c[i])
		//	}
		//}
		itemObj := engine.NewSceneObject(fmt.Sprintf("Item %d", s.selectedIndex))
		itemObj.AddComponent(NewDropBox(s.items[s.selectedIndex]))
		itemObj.AddComponent(builtin.NewRectBoundary())
		itemObj.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
			_, ok := event.Sender.(*engine.SceneObject)
			if !ok {
				return true
			}

			s.fullMode = true
			engine.LogDebug("Short click")
			s.updateItems()

			return true
		}))
		itemObj.SetSizeXy(a.MatchParent, 20)
		s.SceneObject.AddChild(itemObj)

		b := GetDropBox(s.SceneObject.GetChildren()[0], true)
		b.setItem(s.items[s.selectedIndex])
	}
}

func NewDropBoxGroup() *DropBoxGroup {
	return &DropBoxGroup{
		selectedIndex: -1,
		items:         make([]DropItem, 0, 3),
	}
}

type DropBox struct {
	engine.ViewImpl
	item          DropItem
	group         *DropBoxGroup
	triangleID    int
	squareId      int
	smallSquareId int
	textId        int
	rNode         *rendering.Node
	aFace         *atext.Face
	aText         *atext.Text
	initialized   bool
	textOffset    a.Vector3
	prevTransform engine.Transform
}

func (s *DropBox) OnInit(ctx engine.InitContext) {
	s.ViewImpl.OnInit(ctx)

	s.textOffset = a.NewVector3(15, 0, 0)

	s.rNode = ctx.GetRenderingNode()
	s.group = GetDropBoxGroup(s.SceneObject.GetParent())

	font, _ := atext.ParseFont(atext.DefaultFontData)
	s.aFace = font.NewFace(14)
	s.layoutText()

	s.initialized = true
}

func (s *DropBox) OnStart() {
	s.triangleID = s.rNode.AddPrimitive()
	s.textId = s.rNode.AddPrimitive()
}

func (s *DropBox) OnStop() {
	s.rNode.RemovePrimitive(s.triangleID)
	s.rNode.RemovePrimitive(s.textId)
}

func (s *DropBox) setItem(item DropItem) {
	s.item = item
	s.ShouldRedraw = true

	if !s.initialized {
		return
	}

	s.layoutText()
	engine.RequestRendering()
}

func (s *DropBox) layoutText() {
	pos := s.SceneObject.Transform.GetGlobalTopLeftPosition().Add(s.textOffset)
	size := s.SceneObject.Transform.GetSize().Sub(s.textOffset)
	rect := common.NewRectBoundaryFromPositionAndSize(pos, size)
	s.aText = atext.LayoutRunes(s.aFace, []rune(s.item.text), rect, atext.LayoutOptions{
		HTextAlign: a.TextAlignLeft,
		VTextAlign: a.TextAlignCenter,
		SingleLine: true,
	})
}

func (s *DropBox) GetAText() *atext.Text {
	return s.aText
}

func (s *DropBox) OnUpdate(_ engine.UpdateContext) {
	if !s.ShouldDraw() && s.prevTransform.Equals(s.SceneObject.Transform) {
		return
	}

	s.prevTransform = s.SceneObject.Transform

	s.layoutText()
}

func (s *DropBox) OnDraw(_ engine.DrawingContext) {
	rect := s.SceneObject.Transform.GetGlobalRect()
	pos := s.SceneObject.Transform.GetGlobalTopLeftPosition()

	textPrimitive := rendering.NewTextPrimitive(s.item.text, s)
	textPrimitive.Transform = rendering.NewTransform()
	textPrimitive.Transform.Position = pos.Add(a.NewVector3(15, 0, 0)).Round()
	textPrimitive.Transform.Size = rect.GetSize().Sub(a.NewVector3(15, 0, 0)).Round()
	if s.group.IsSelected(s.item.index) {
		textPrimitive.Appearance.FillColor = a.BlueColor()
	} else {
		textPrimitive.Appearance.FillColor = a.BlackColor()
	}
	s.rNode.SetPrimitive(s.textId, textPrimitive)
	check := 0
	if !s.group.fullMode{
		check = s.group.selectedIndex
	}
	if s.item.index == check {
		trianglePrimitive := rendering.NewGeometryPrimitive(rendering.PrimitiveEllipse)
		trianglePrimitive.Transform = rendering.NewTransform()
		//trianglePrimitive.Transform.Position = pos.Add(a.NewVector3(0, s.SceneObject.Transform.GetSize().Y/2-5, 0)).Round()
		//trianglePrimitive.Transform.Position = pos.Add(a.NewVector3(float32(textPrimitive.Transform.Position.X), float32(textPrimitive.Transform.Position.Y), 0)).Round()
		trianglePrimitive.Transform.Position = pos.Add(a.NewVector3(float32(textPrimitive.Transform.Size.X), float32(textPrimitive.Transform.Size.Y / 4), 0)).Round()
		trianglePrimitive.Transform.Size = a.NewIntVector3(10, 10, 0)
		trianglePrimitive.Appearance.FillColor = a.BlackColor()
		trianglePrimitive.Appearance.StrokeColor = a.BlackColor()

		s.rNode.SetPrimitive(s.triangleID, trianglePrimitive)
	}

	s.ShouldRedraw = false
}

func NewDropBox(item DropItem) *DropBox {
	return &DropBox{
		item: item,
	}
}

func (s *DropBoxGroup) SelectedItemText() string {
	if s.selectedIndex != -1 {
		return s.items[s.selectedIndex].text
	} else {
		return ""
	}
}
