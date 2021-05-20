package components

import (
	"AmphionLand/generated/res"
	owm "github.com/briandowns/openweathermap"
	"github.com/cadmean-ru/amphion/engine"
	"github.com/cadmean-ru/amphion/engine/builtin"
	"strconv"
)

type WeatherController struct {
	engine.ComponentImpl
	temp    float64
	rain    float64
	wind    float64
	imageURL string
}

func (s *WeatherController) OnInit(ctx engine.InitContext) {
	s.ComponentImpl.OnInit(ctx)
	s.temp = 0
	s.rain = 0
	s.wind = 0
	s.imageURL = ""
	apiKey, err := s.Engine.GetResourceManager().ReadFile(res.Strings_definetlynotkey)
	if err==nil {
		w, apiErr := owm.NewCurrent("C", "ru", string(apiKey))
		if apiErr == nil {
			_ = w.CurrentByName("Moscow")
			s.temp  = w.Main.Temp
			s.rain = w.Rain.OneH
			s.wind = w.Wind.Speed
			s.imageURL = "http://openweathermap.org/img/wn/" + w.Weather[0].Icon + ".png"
		} else {
			engine.LogDebug(apiErr.Error())
		}
	} else {
		engine.LogDebug(err.Error())
	}

	name := s.SceneObject.GetChildByName("City Name")
	name.GetComponentByName("TextView").(*builtin.TextView).SetText("Москва")

	wg := s.SceneObject.GetChildByName("Weather Grid")
	wg.GetChildByName("Temperature").GetComponentByName("TextView").(*builtin.TextView).
		SetText(strconv.FormatFloat(s.temp, 'f',3,32))
	wg.GetChildByName("Weather Icon").GetComponentByName("ImageView").(*builtin.ImageView).SetImageUrl(s.imageURL)
	wg.GetChildByName("Rain").GetComponentByName("TextView").(*builtin.TextView).
		SetText(strconv.FormatFloat(s.rain, 'f',3,32))
	wg.GetChildByName("Wind").GetComponentByName("TextView").(*builtin.TextView).
		SetText(strconv.FormatFloat(s.wind, 'f',3,32))
}

func (s *WeatherController) GetName() string {
	return engine.NameOfComponent(s)
}