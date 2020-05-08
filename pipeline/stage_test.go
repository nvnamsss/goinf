package pipeline

import (
	"testing"
)

func TestNextStage(t *testing.T) {

	fn := func() (struct {
		Field string
		Meo   int
	}, error) {

		return struct {
			Field string
			Meo   int
		}{"meomeocute", 1}, nil
	}
	stage := NewStage(fn)
	stage2 := NewStage(func() (str struct{ Meo []string }, e error) {
		str.Meo = []string{"Abc", "123", "meomeocute"}

		return
	})

	stage.NextStage(stage2)
	var p *Pipeline = NewPipeline()
	p.First = stage
	err := p.Run()

	if err != nil {
		t.Errorf("The pipeline is run wrong, problem: %v", err.Error())
	}

	s := p.GetString("Field")[0]
	f := p.GetInt("Meo")[0]
	s2 := p.GetString("Meo")[0]
	if s != "meomeocute" {
		t.Errorf("pipeline returned wrong string value: got %v want %v",
			s, "meomeocute")
	} else {
		t.Logf("pipeline returned string value passed")
	}

	if f != 1 {
		t.Errorf("pipeline returned wrong float value: got %v want %v",
			f, 1)
	} else {
		t.Logf("pipeline returned float value passed")
	}

	if s2 != "meomeocute" {
		t.Errorf("pipeline returned wrong string value: got %v want %v", s2, "meomeocute")
	} else {
		t.Logf("pipeline returned string value in array passed")
	}
}
