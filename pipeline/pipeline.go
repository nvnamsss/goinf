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
	if !p.IsPassed {
		return nil
	}

	var result []float64 = []float64{}
	for _, value := range p.values[field] {
		switch value.Kind() {
		case reflect.Float32, reflect.Float64:
			result = append(result, value.Float())
		}
	}

	return result
}

//GetFloatFirstOrDefault return the first item from GetFloat but automatically add a default value (0) if the field is not existed in the pipeline
//
//GetFloatFirstOrDefault should use in case you want to use the default value if the field is not existed without checking for empty result.
func (p Pipeline) GetFloatFirstOrDefault(field string) (result float64) {
	rs := p.GetFloat(field)

	if len(rs) == 0 {
		return 0
	}

	return rs[0]
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

//GetStringFirstOrDefault return the first item from GetString but automatically return a default value ("") if the field is not existed in the pipeline
//
//GetStringFirstOrDefault should use in case you want to use the default value if the field is not existed without checking for empty result.
func (p Pipeline) GetStringFirstOrDefault(field string) (result string) {
	rs := p.GetString(field)

	if len(rs) == 0 {
		return ""
	}

	return rs[0]
}

//GetInt return int values of a specified field which are received from every stages.
func (p Pipeline) GetInt(field string) []int64 {
	if !p.IsPassed {
		return nil
	}

	var result []int64 = []int64{}
	for _, value := range p.values[field] {
		kind := value.Kind()
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result = append(result, value.Int())
		}
	}

	return result
}

//GetIntFirstOrDefault return the first item from GetInt but automatically return a default value (zero) if the field is not existed in the pipeline
//
//GetIntFirstOrDefault should use in case you want to use the default value if the field is not existed without checking for empty result.
func (p Pipeline) GetIntFirstOrDefault(field string) (result int64) {
	rs := p.GetInt(field)

	if len(rs) == 0 {
		return 0
	}

	return rs[0]
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

//GetBoolFirstOrDefault return the first item from GetBool but automatically return a default value (false) if the field is not existed in the pipeline
//
//GetBoolFirstOrDefault should use in case you want to use the default value if the field is not existed without checking for empty result.
func (p Pipeline) GetBoolFirstOrDefault(field string) (result bool) {
	rs := p.GetBool(field)

	if len(rs) == 0 {
		return false
	}

	return rs[0]
}

//Constructor for creating new Pipeline
func NewPipeline() *Pipeline {
	var pipe *Pipeline = new(Pipeline)
	pipe.values = make(map[string][]reflect.Value)

	return pipe
}
