package pipeline

import "reflect"

//GetMapString return a map string of a specified field which are received from every stages.
//
//When multiple maps which the same field returned. GetMapString will merge all of them into one.
func (p Pipeline) GetMapString(field string) map[string]string {
	if !p.IsPassed {
		return nil
	}

	var result map[string]string = make(map[string]string)
	for _, value := range p.values[field] {
		if value.Kind() == reflect.Map {
			keys := value.MapKeys()
			for _, key := range keys {
				index := value.MapIndex(key)
				if index.Kind() == reflect.String {
					result[key.String()] = value.MapIndex(key).String()
				}
			}
		}
	}

	return result
}

//GetMapFloat return a map float of a specified field which are received from every stages.
//
//When multiple maps which the same field returned. GetMapFloat will merge all of them into one.
func (p Pipeline) GetMapFloat(field string) map[string]float64 {
	if !p.IsPassed {
		return nil
	}

	var result map[string]float64 = make(map[string]float64)
	for _, value := range p.values[field] {
		if value.Kind() == reflect.Map {
			keys := value.MapKeys()
			for _, key := range keys {
				index := value.MapIndex(key)
				if index.Kind() == reflect.Float64 {
					result[key.String()] = value.MapIndex(key).Float()
				}
			}
		}
	}

	return result
}

//GetMapInt return a map float of a specified field which are received from every stages.
//
//When multiple maps which the same field returned. GetMapInt will merge all of them into one.
func (p Pipeline) GetMapInt(field string) map[string]int64 {
	if !p.IsPassed {
		return nil
	}

	var result map[string]int64 = make(map[string]int64)
	for _, value := range p.values[field] {
		if value.Kind() == reflect.Map {
			keys := value.MapKeys()
			for _, key := range keys {
				index := value.MapIndex(key)
				if index.Kind() == reflect.Int64 {
					result[key.String()] = value.MapIndex(key).Int()
				}
			}
		}
	}

	return result
}
