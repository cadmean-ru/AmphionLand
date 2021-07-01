package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/common/a"
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
	items []CheckItem
	initialized bool
}

func (s *CheckBoxGroup) OnInit(ctx engine.InitContext) {
	s.GridLayout.OnInit(ctx)
	s.Cols = 1

	s.initialized = true
}

func (s *CheckBoxGroup) OnStart() {
	s.updateItems()
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

	s.selectedIndexes[indexPos] = s.selectedIndexes[len(s.selectedIndexes) - 1]
	s.selectedIndexes[len(s.selectedIndexes) - 1] = -1
	s.selectedIndexes = s.selectedIndexes[:len(s.selectedIndexes) - 1]

	if !s.initialized {
		return
	}

	s.updateItems()
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

				check := obj.GetComponentByName("CheckBox").(*CheckBox)
				if s.IsSelected(check.item.index) {
					s.SetNotSelected(check.item.index)
				} else {
					s.SetSelected(check.item.index)
				}

				return true
			}))
			itemObj.SetSizeXy(a.MatchParent, 20)
			s.SceneObject.AddChild(itemObj)
		}
	} else if len(s.items) < s.SceneObject.GetChildrenCount() {
		for i := s.SceneObject.GetChildrenCount()-1; i >= len(s.items); i-- {
			c := s.SceneObject.GetChildren()
			s.SceneObject.RemoveChild(c[i])
		}
	}

	children := s.SceneObject.GetChildren()
	for i, item := range s.items {
		c := children[i]
		b := c.GetComponentByName("CheckBox", true).(*CheckBox)
		b.setItem(item)
	}
}

func (s *CheckBoxGroup) GetName() string {
	return engine.NameOfComponent(s)
}

func NewCheckBoxGroup() *CheckBoxGroup {
	return &CheckBoxGroup{
		selectedIndexes: []int{},
		items:         make([]CheckItem, 0, 3),
	}
}

type CheckBox struct {
	engine.ViewImpl
	item   CheckItem
	group *CheckBoxGroup
	circle int
	text int
	rNode *rendering.Node
}

func (s *CheckBox) OnInit(ctx engine.InitContext) {
	s.ViewImpl.OnInit(ctx)

	s.rNode = ctx.GetRenderingNode()
	s.group = s.SceneObject.GetParent().GetComponentByName("CheckBoxGroup", true).(*CheckBoxGroup)
}

func (s *CheckBox) OnStart() {
	s.circle = s.rNode.AddPrimitive()
	s.text = s.rNode.AddPrimitive()
}

func (s *CheckBox) setItem(item CheckItem) {
	s.item = item
	s.ShouldRedraw = true
	engine.RequestRendering()
}

func (s *CheckBox) OnDraw(ctx engine.DrawingContext) {
	circlePrimitive := rendering.NewGeometryPrimitive(rendering.PrimitiveEllipse)
	rect := s.SceneObject.Transform.GetGlobalRect()
	pos := s.SceneObject.Transform.GetGlobalTopLeftPosition()

	circlePrimitive.Transform = rendering.NewTransform()
	circlePrimitive.Transform.Position = pos.Round()
	circlePrimitive.Transform.Size = a.NewIntVector3(10, 10, 0)
	circlePrimitive.Appearance.StrokeColor = a.BlackColor()
	if s.group.IsSelected(s.item.index) {
		circlePrimitive.Appearance.FillColor = a.PinkColor()
	} else {
		circlePrimitive.Appearance.FillColor = a.WhiteColor()
	}
	s.rNode.SetPrimitive(s.circle, circlePrimitive)

	textPrimitive := rendering.NewTextPrimitive(s.item.text)
	textPrimitive.Transform = rendering.NewTransform()
	textPrimitive.Transform.Position = pos.Add(a.NewVector3(15, 0, 0)).Round()
	textPrimitive.Transform.Size = rect.GetSize().Sub(a.NewVector3(15, 0, 0)).Round()
	if s.group.IsSelected(s.item.index) {
		textPrimitive.Appearance.FillColor = a.BlackColor()
	} else {
		textPrimitive.Appearance.FillColor = a.NewColor("#aaa")
	}
	s.rNode.SetPrimitive(s.text, textPrimitive)

	s.ShouldRedraw = false
}

func (s *CheckBox) GetName() string {
	return engine.NameOfComponent(s)
}

func NewCheckBox(item CheckItem) *CheckBox {
	return &CheckBox{
		item: item,
	}
}
