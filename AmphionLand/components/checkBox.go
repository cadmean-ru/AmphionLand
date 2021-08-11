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

type CheckItem struct {
	index int
	text  string
}

func (s CheckItem) Index() int {
	return s.index
}

func (s CheckItem) Text() string {
	return s.text
}

type CheckBoxGroup struct {
	builtin.GridLayout
	selectedIndexes []int
	items           []CheckItem
	initialized     bool
	listener        func(item CheckItem)
}

func (s *CheckBoxGroup) OnInit(ctx engine.InitContext) {
	s.GridLayout.OnInit(ctx)
	s.AddColumn(a.FillParent)
	s.AutoExpansion = true
	s.AutoShrinking = true

	s.initialized = true
}

func (s *CheckBoxGroup) OnStart() {
	s.updateItems()
}

func (s *CheckBoxGroup) SetOnItemSelectedListener(listener func(item CheckItem)) {
	s.listener = listener
}

func (s *CheckBoxGroup) AddItem(text string) {
	s.items = append(s.items, CheckItem{
		index: len(s.items),
		text:  text,
	})

	if !s.initialized {
		return
	}

	s.updateItems()
}

func (s *CheckBoxGroup) IsSelected(index int) bool {
	for _, i := range s.selectedIndexes {
		if i == index {
			return true
		}
	}
	return false
}

func (s *CheckBoxGroup) SetSelected(index int) {
	s.selectedIndexes = append(s.selectedIndexes, index)

	if !s.initialized {
		return
	}

	s.updateItems()
}

func (s *CheckBoxGroup) SetNotSelected(index int) {
	indexPos := -1
	for i := 0; i < len(s.selectedIndexes); i++ {
		if s.selectedIndexes[i] == index {
			indexPos = i
		}
	}

	if indexPos == -1 {
		return
	}

	s.selectedIndexes[indexPos] = s.selectedIndexes[len(s.selectedIndexes)-1]
	s.selectedIndexes[len(s.selectedIndexes)-1] = -1
	s.selectedIndexes = s.selectedIndexes[:len(s.selectedIndexes)-1]

	if !s.initialized {
		return
	}

	s.updateItems()
}

func (s *CheckBoxGroup) GetItemIndex(text string) int {
	for i:=0; i< len(s.items); i++ {
		if s.items[i].text == text {
			return s.items[i].index
		}
	}
	return -1
}

func (s *CheckBoxGroup) updateItems() {
	if len(s.items) > s.SceneObject.GetChildrenCount() {
		for i := s.SceneObject.GetChildrenCount(); i < len(s.items); i++ {
			itemObj := engine.NewSceneObject(fmt.Sprintf("Item %d", i))
			itemObj.AddComponent(NewCheckBox(s.items[i]))
			itemObj.AddComponent(builtin.NewRectBoundary())
			itemObj.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
				obj, ok := event.Sender.(*engine.SceneObject)
				if !ok {
					return true
				}

				check := GetCheckBox(obj)
				if s.IsSelected(check.item.index) {
					s.SetNotSelected(check.item.index)
				} else {
					s.SetSelected(check.item.index)
				}

				if s.listener != nil {
					s.listener(check.item)
				}

				return true
			}))
			itemObj.SetSizeXy(a.MatchParent, 20)
			s.SceneObject.AddChild(itemObj)
		}
	} else if len(s.items) < s.SceneObject.GetChildrenCount() {
		for i := s.SceneObject.GetChildrenCount() - 1; i >= len(s.items); i-- {
			c := s.SceneObject.GetChildren()
			s.SceneObject.RemoveChild(c[i])
		}
	}

	children := s.SceneObject.GetChildren()
	for i, item := range s.items {
		c := children[i]
		b := GetCheckBox(c, true)
		b.setItem(item)
	}
}

func NewCheckBoxGroup() *CheckBoxGroup {
	return &CheckBoxGroup{
		selectedIndexes: []int{},
		items:           make([]CheckItem, 0, 3),
	}
}

type CheckBox struct {
	engine.ViewImpl
	item     CheckItem
	group    *CheckBoxGroup
	circleId int
	textId   int
	rNode    *rendering.Node
	aFace    *atext.Face
	aText    *atext.Text
	initialized   bool
	textOffset    a.Vector3
	prevTransform engine.Transform
}

func (s *CheckBox) OnInit(ctx engine.InitContext) {
	s.ViewImpl.OnInit(ctx)

	s.textOffset = a.NewVector3(15, 0, 0)

	s.rNode = ctx.GetRenderingNode()
	s.group = GetCheckBoxGroup(s.SceneObject.GetParent())

	font, _ := atext.ParseFont(atext.DefaultFontData)
	s.aFace = font.NewFace(14)
	s.layoutText()

	s.initialized = true
}

func (s *CheckBox) OnStart() {
	s.circleId = s.rNode.AddPrimitive()
	s.textId = s.rNode.AddPrimitive()
}

func (s *CheckBox) OnStop() {
	s.rNode.RemovePrimitive(s.circleId)
	s.rNode.RemovePrimitive(s.textId)
}

func (s *CheckBox) setItem(item CheckItem) {
	s.item = item
	s.ShouldRedraw = true

	if !s.initialized {
		return
	}

	s.layoutText()
	engine.RequestRendering()
}

func (s *CheckBox) layoutText() {
	pos := s.SceneObject.Transform.GetGlobalTopLeftPosition().Add(s.textOffset)
	size := s.SceneObject.Transform.GetSize().Sub(s.textOffset)
	rect := common.NewRectBoundaryFromPositionAndSize(pos, size)
	s.aText = atext.LayoutRunes(s.aFace, []rune(s.item.text), rect, atext.LayoutOptions{
		HTextAlign: a.TextAlignLeft,
		VTextAlign: a.TextAlignCenter,
		SingleLine: true,
	})
}

func (s *CheckBox) GetAText() *atext.Text {
	return s.aText
}

func (s *CheckBox) OnUpdate(_ engine.UpdateContext) {
	if !s.ShouldDraw() && s.prevTransform.Equals(s.SceneObject.Transform) {
		return
	}

	s.prevTransform = s.SceneObject.Transform

	s.layoutText()
}

func (s *CheckBox) OnDraw(_ engine.DrawingContext) {
	rect := s.SceneObject.Transform.GetGlobalRect()
	pos := s.SceneObject.Transform.GetGlobalTopLeftPosition()

	circlePrimitive := rendering.NewGeometryPrimitive(rendering.PrimitiveEllipse)
	circlePrimitive.Transform = rendering.NewTransform()
	circlePrimitive.Transform.Position = pos.Add(a.NewVector3(0, s.SceneObject.Transform.GetSize().Y/2-5, 0)).Round()
	circlePrimitive.Transform.Size = a.NewIntVector3(10, 10, 0)
	circlePrimitive.Appearance.StrokeColor = a.BlackColor()
	if s.group.IsSelected(s.item.index) {
		circlePrimitive.Appearance.FillColor = a.PinkColor()
	} else {
		circlePrimitive.Appearance.FillColor = a.WhiteColor()
	}
	s.rNode.SetPrimitive(s.circleId, circlePrimitive)

	textPrimitive := rendering.NewTextPrimitive(s.item.text, s)
	textPrimitive.Transform = rendering.NewTransform()
	textPrimitive.Transform.Position = pos.Add(a.NewVector3(15, 0, 0)).Round()
	textPrimitive.Transform.Size = rect.GetSize().Sub(a.NewVector3(15, 0, 0)).Round()
	if s.group.IsSelected(s.item.index) {
		textPrimitive.Appearance.FillColor = a.BlackColor()
	} else {
		textPrimitive.Appearance.FillColor = a.NewColor("#aaa")
	}
	s.rNode.SetPrimitive(s.textId, textPrimitive)

	s.ShouldRedraw = false
}

func NewCheckBox(item CheckItem) *CheckBox {
	return &CheckBox{
		item: item,
	}
}
