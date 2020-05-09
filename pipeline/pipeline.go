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

			if value.Kind() == reflect.Slice || value.Kind() == reflect.Array {
				len := value.Len()
				for loop := 0; loop < len; loop++ {
					p.values[name] = append(p.values[name], value.Index(loop))
				}
			} else {
				p.values[name] = append(p.values[name], value)
			}

		}

		current = current.stages
	}

	p.IsPassed = true
	return nil
}

//GetFloat return float values of a specified field which are received from every stages.
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

//GetString return string values of a specified field which are received from every stages.
func (p Pipeline) GetString(field string) []string {
	if !p.IsPassed {
		return nil
	}

	var result []string = []string{}
	for _, value := range p.values[field] {
		if value.Kind() == reflect.String {
			result = append(result, value.String())
		}
	}

	return result
}

//GetInt return int values of a specified field which are received from every stages.
func (p Pipeline) GetInt(field string) []int64 {
	if !p.IsPassed {
		return nil
	}

	var result []int64 = []int64{}
	for _, value := range p.values[field] {

		if value.Kind() == reflect.Int {
			result = append(result, value.Int())
		}
	}

	return result
}

//GetBool return bool values of a specified field which are received from every stages.
func (p Pipeline) GetBool(field string) []bool {
	if !p.IsPassed {
		return nil
	}

	var result []bool = []bool{}
	for _, value := range p.values[field] {
		if value.Kind() == reflect.Bool {
			result = append(result, value.Bool())
		}
	}

	return result
}

//GetValue return values of a specified field which are received from every stages.
func (p Pipeline) GetValue(field string) []reflect.Value {
	if !p.IsPassed {
		return nil
	}

	return p.values[field]
}

func (p Pipeline) GetValueGeneric(field string, T reflect.Type) []interface{} {
	if !p.IsPassed {
		return nil
	}

	var result []interface{} = []interface{}{}

	for _, value := range p.values[field] {
		t := reflect.TypeOf(value)

		if t == T {
			result = append(result, value.Interface())
		}
	}

	return result
}

//Constructor for creating new Pipeline
func NewPipeline() *Pipeline {
	var pipe *Pipeline = new(Pipeline)
	pipe.values = make(map[string][]reflect.Value)

	return pipe
}
