package pipeline

import (
	"reflect"
)

//Pipeline representation a process of work needs to be run by order and will be stopped if a step is failed.
type Pipeline struct {
	First    *Stage
	IsPassed bool
	values   map[string][]reflect.Value
}

//Run start the pipeline, success when the pipeline is passed when all stages are passed.
func (p *Pipeline) Run() error {
	var current *Stage = p.First
	p.IsPassed = false
	p.values = make(map[string][]reflect.Value)

	for current != nil {
		result := current.task.Call(nil)

		// err := error(result[1].Interface())
		// err := current.Run()
		err := result[1].Interface()

		if err != nil {
			return err.(error)
		}

		indirect := reflect.Indirect(result[0])
		for loop := 0; loop < indirect.NumField(); loop++ {
			name := result[0].Type().Field(loop).Name
			value := indirect.Field(loop)

			p.values[name] = append(p.values[name], value)
		}

		current = current.stages
	}

	p.IsPassed = true
	return nil
}

func (p Pipeline) convert(embryo interface{}) {
	convert := func(in []reflect.Value) []reflect.Value {
		if !p.IsPassed {
			return nil
		}

		return p.values[in[0].String()]
	}

	fn := reflect.ValueOf(embryo).Elem()
	v := reflect.MakeFunc(fn.Type(), convert)

	fn.Set(v)
}

func (p Pipeline) GetFloat(field string) []float64 {
	// var embryo func(string) []float64
	// p.convert(embryo)

	// result := embryo(field)

	if !p.IsPassed {
		return nil
	}

	var result []float64 = []float64{}
	for _, value := range p.values[field] {
		result = append(result, value.Float())
	}

	return result
}

func (p Pipeline) GetString(field string) []string {
	if !p.IsPassed {
		return nil
	}

	var result []string = []string{}
	for _, value := range p.values[field] {
		result = append(result, value.String())
	}

	return result
}

func (p Pipeline) GetInt(field string) []int64 {
	if !p.IsPassed {
		return nil
	}

	var result []int64 = []int64{}
	for _, value := range p.values[field] {
		result = append(result, value.Int())
	}

	return result
}

func (p Pipeline) GetBool(field string) []bool {
	if !p.IsPassed {
		return nil
	}

	var result []bool = []bool{}
	for _, value := range p.values[field] {
		result = append(result, value.Bool())
	}

	return result
}

//Constructor for creating new Pipeline
func NewPipeline() *Pipeline {
	var pipe *Pipeline = new(Pipeline)
	pipe.values = make(map[string][]reflect.Value)

	return pipe
}
