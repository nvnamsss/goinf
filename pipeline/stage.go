package pipeline

import (
	"reflect"
)

//Stage representation a step in the Pipeline, a Stage will lead to another Stage.
//If there is no Stage, this Stage will be the end of Pipeline
type Stage struct {
	Name   string
	stages *Stage
	task   reflect.Value
}

//NextStage set the left Stage as the next stage of right Stage
//In the Pipeline, left Stage will run when the right is done
func (r *Stage) NextStage(l *Stage) {
	r.stages = l
}

//Next create new stage which is run the task. New stage will be appended to the current stage
func (r *Stage) Next(task interface{}) {
	stage := NewStage(task)

	r.stages = stage
}

//Set using to set the task of stage, when the pipeline run the task of stage will be called
func (r *Stage) Set(task interface{}) {
	t := reflect.ValueOf(task).Type()
	olen := t.NumOut()
	if olen == 1 {
		var fn func() (struct{}, error)

		r.task = reflect.MakeFunc(reflect.ValueOf(fn).Type(), func(args []reflect.Value) (results []reflect.Value) {
			result := reflect.ValueOf(task).Call(nil)
			result = append([]reflect.Value{reflect.ValueOf(struct{}{})}, result...)

			return result
		})

	}

	if olen == 2 {
		if t.Out(0).Kind() != reflect.Struct {
			panic("The first return type must be struct")
		}

		if t.Out(1).Kind() != reflect.Interface {
			panic("The second return type must be interface")
		}
		r.task = reflect.ValueOf(task)
	}
}

func NewStage(task interface{}) *Stage {
	var stage *Stage = new(Stage)
	stage.Name = ""
	stage.Set(task)

	return stage
}
