package components

import (
	"github.com/cadmean-ru/amphion/common/a"
	"github.com/cadmean-ru/amphion/engine"
)

type EmptyBox struct {
	engine.ComponentImpl

}

func (s *EmptyBox) LayoutChildren() {
	children := s.SceneObject.GetChildren()
	if len(children) == 0 {
		return
	}

	first := children[0]
	first.Transform.Position = a.NewVector3(0, 0, 1)
	first.Transform.Pivot = a.ZeroVector()
	first.Transform.Size = a.NewVector3(a.MatchParent, a.MatchParent, a.MatchParent)

	for i := 1; i < len(children); i++ {
		c := children[i]
		c.Transform.Position = a.ZeroVector()
		c.Transform.Size = a.ZeroVector()
	}
}

func (s *EmptyBox) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
}

func (s *EmptyBox) GetName() string {
	return engine.NameOfComponent(s)
}
