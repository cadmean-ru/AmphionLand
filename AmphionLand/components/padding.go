package components

import "github.com/cadmean-ru/amphion/engine"

type Padding struct {
	engine.ComponentImpl
	LeftX float32
	RightX float32
	UpY float32
	DownY float32
}

func (s *Padding) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
}


func (s *Padding) GetName() string {
	return engine.NameOfComponent(s)
}

func NewPadding() *Padding {
	return &Padding{
		LeftX: 0,
		RightX: 0,
		UpY: 0,
		DownY: 0,
	}
}

func (s *Padding) UpdatePadding() {
	child := s.SceneObject.GetChildren()[0]

	s.SceneObject.Transform.Size = child.Transform.Size

	s.SceneObject.Transform.Size.X += s.LeftX
	child.Transform.Position.X += s.LeftX

	s.SceneObject.Transform.Size.X += s.RightX

	s.SceneObject.Transform.Size.Y += s.UpY
	child.Transform.Position.Y += s.UpY

	s.SceneObject.Transform.Size.Y += s.DownY

}
