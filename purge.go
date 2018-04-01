package goml

import (
       "fmt"
       "github.com/smallfish/simpleyaml"
       "gopkg.in/yaml.v2"
       "reflect"
       "strings"
)

var data = `
Data:
    - name: "foo"
      bar1: 0
      k1: val1
      k2:
         val2
         val3
      bar2: 1
      k3: val4
      k4: val5
      k5: val5
      k6: val6
`

func printVal(v interface{}, depth int) {
    typ := reflect.TypeOf(v).Kind()
    if typ == reflect.Int || typ == reflect.String {
        fmt.Printf("%s%v\n", strings.Repeat(" ", depth), v)
    } else if typ == reflect.Slice {
        fmt.Printf("\n")
        printSlice(v.([]interface{}), depth+1)
    } else if typ == reflect.Map {
        fmt.Printf("\n")
        printMap(v.(map[interface{}]interface{}), depth+1)
    }

}

func printMap(m map[interface{}]interface{}, depth int) {
    for k, v := range m {
        fmt.Printf("%sKey:%s", strings.Repeat(" ", depth), k.(string))
        printVal(v, depth+1)
    }
}

func printSlice(slc []interface{}, depth int) {
    for _, v := range slc {
        printVal(v, depth+1)
    }
}

func PurgeToFile(from string, fromPath string, to string, toPath string) error {
	fromYaml, err := ReadYamlFromFile(from)
	if err != nil {
		return err
	}

	toYaml, err := ReadYamlFromFile(to)
	if err != nil {
		return err
	}

	err = Purge(fromYaml, fromPath, toYaml, toPath)
	if err != nil {
		return err
	}

	return WriteYaml(toYaml, to)
}

func Purge(fromYaml *simpleyaml.Yaml, fromPath string, toYaml *simpleyaml.Yaml, toPath string) error {
        m := make(map[string]interface{})

        err := yaml.Unmarshal([]byte(data), &m)
        if err != nil {
            panic(err)
        }
        for k, v := range m {
            fmt.Printf("Key:%s ", k)
            printVal(v, 1)
        }

    	value, err := Get(fromYaml, fromPath)
	if err != nil {
		return err
	}

	return Set(toYaml, toPath, value)
}
