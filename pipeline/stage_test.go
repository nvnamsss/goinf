package pipeline

import (
	"testing"
)

func TestNextStage(t *testing.T) {

	fn := func() (struct {
		Field string
		Meo   int64
	}, error) {

		return struct {
			Field string
			Meo   int64
		}{"meomeocute", 1}, nil
	}
	stage := NewStage(fn)
	stage2 := NewStage(func() (str struct{ Meo []string }, e error) {
		str.Meo = []string{"Abc", "123", "meomeocute"}

		return
	})

	stage3 := NewStage(func() (str struct {
		Data map[string]string
	}, e error) {
		str.Data = make(map[string]string)
		str.Data["meo"] = "abc"
		str.Data["ma"] = "ma"
		str.Data["may"] = "meomeocute"

		return
	})

	stage.NextStage(stage2)
	stage2.NextStage(stage3)
	var p *Pipeline = NewPipeline()
	p.First = stage
	err := p.Run()

	if err != nil {
		t.Errorf("The pipeline is run wrong, problem: %v", err.Error())
	}

	s := p.GetString("Field")[0]
	f := p.GetInt("Meo")[0]
	s2 := p.GetString("Meo")[2]
	m := p.GetMapString("Data")
	d := p.GetIntFirstOrDefault("anhon")

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

	if m["may"] != "meomeocute" {
		t.Errorf("pipeline returned wrong string value: got %v want %v", m["may"], "meomeocute")
	} else {
		t.Logf("pipeline returned string value in map passed")
	}

	if d != 0 {
		t.Errorf("pipeline returned wrong int default value: got %v want %v", d, 0)
	} else {
		t.Logf("pipeline returned default in value passed")
	}
}
