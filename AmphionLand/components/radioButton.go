package components

import (
	"fmt"
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/common/atext"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"github.com/cadmean-ru/amphion/rendering"
)

type RadioItem struct {
	index int
	text  string
}

func (s RadioItem) Index() int {
	return s.index
}

func (s RadioItem) Text() string {
	return s.text
}

type RadioButtonGroup struct {
	builtin.GridLayout
	selectedIndex int
	items         []RadioItem
	initialized   bool
}

func (s *RadioButtonGroup) OnInit(ctx engine.InitContext) {
	s.GridLayout.OnInit(ctx)
	s.AddColumn(a.WrapContent)

	s.initialized = true
}

func (s *RadioButtonGroup) OnStart() {
	s.updateItems()
}

func (s *RadioButtonGroup) AddItem(text string) {
	s.items = append(s.items, RadioItem{
		index: len(s.items),
		text:  text,
	})

	if !s.initialized {
		return
	}

	s.updateItems()
}

func (s *RadioButtonGroup) IsSelected(index int) bool {
	return s.selectedIndex == index
}

func (s *RadioButtonGroup) SetSelected(index int) {
	s.selectedIndex = index

	if !s.initialized {
		return
	}

	s.updateItems()
}

func (s *RadioButtonGroup) updateItems() {
	if len(s.items) > s.SceneObject.GetChildrenCount() {
		for i := s.SceneObject.GetChildrenCount(); i < len(s.items); i++ {
			itemObj := engine.NewSceneObject(fmt.Sprintf("Item %d", i))
			itemObj.AddComponent(NewRadioButton(s.items[i]))
			itemObj.AddComponent(builtin.NewRectBoundary())
			itemObj.AddComponent(builtin.NewEventListener(engine.EventMouseDown, func(event engine.AmphionEvent) bool {
				obj, ok := event.Sender.(*engine.SceneObject)
				if !ok {
					return true
				}

				radio := obj.GetComponentByName("RadioButton").(*RadioButton)
				s.SetSelected(radio.item.index)

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
		b := c.GetComponentByName("RadioButton", true).(*RadioButton)
		b.setItem(item)
	}
}

func (s *RadioButtonGroup) GetName() string {
	return engine.NameOfComponent(s)
}

func NewRadioButtonGroup() *RadioButtonGroup {
	return &RadioButtonGroup{
		selectedIndex: -1,
		items:         make([]RadioItem, 0, 3),
	}
}

type RadioButton struct {
	engine.ViewImpl
	item   RadioItem
	group  *RadioButtonGroup
	circle int
	text   int
	rNode  *rendering.Node
	aFace  *atext.Face
	aText  *atext.Text
}

func (s *RadioButton) OnInit(ctx engine.InitContext) {
	s.ViewImpl.OnInit(ctx)

	s.rNode = ctx.GetRenderingNode()
	s.group = s.SceneObject.GetParent().GetComponentByName("RadioButtonGroup", true).(*RadioButtonGroup)

	font, _ := atext.ParseFont(atext.DefaultFontData)
	s.aFace = font.NewFace(14)
	s.layoutText()
}

func (s *RadioButton) OnStart() {
	s.circle = s.rNode.AddPrimitive()
	s.text = s.rNode.AddPrimitive()
}

func (s *RadioButton) OnStop() {
	s.rNode.RemovePrimitive(s.circle)
	s.rNode.RemovePrimitive(s.text)
}

func (s *RadioButton) setItem(item RadioItem) {
	s.item = item
	s.ShouldRedraw = true
	s.layoutText()
	engine.RequestRendering()
}

func (s *RadioButton) layoutText() {
	s.aText = atext.LayoutRunes(s.aFace, []rune(s.item.text), s.SceneObject.Transform.GetGlobalRect(), atext.LayoutOptions{})
}

func (s *RadioButton) OnDraw(_ engine.DrawingContext) {
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

	textPrimitive := rendering.NewTextPrimitive(s.item.text, s.aText)
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

func (s *RadioButton) GetName() string {
	return engine.NameOfComponent(s)
}

func NewRadioButton(item RadioItem) *RadioButton {
	return &RadioButton{
		item: item,
	}
}

func (s *RadioButtonGroup) SelectedItemText() string {
	if s.selectedIndex != -1 {
		return s.items[s.selectedIndex].text
	} else {
		return ""
	}
}
