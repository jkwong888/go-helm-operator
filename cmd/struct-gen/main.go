package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

func generateStruct(parentName string, name string, v map[interface{}]interface{}) string {
	mymap := make(map[string]string)
	s := ""

	// depth first search
	for k, vv := range v {
		switch vv.(type) {
		case bool:
			mymap[k.(string)] = "bool"
		case string:
			mymap[k.(string)] = "string"
		case int:
			mymap[k.(string)] = "int"
		case []interface{}:
			// special case: an empty array may indicate flexible input, just provide a an array of maps here for now
			if len(vv.([]interface{})) == 0 {
				mymap[k.(string)] = "[]map[string]string"
			} else {
				mymap[k.(string)] = fmt.Sprintf("[]%s%s%sSpec", strings.Title(parentName), strings.Title(name), strings.Title(k.(string)))
			}
		case map[interface{}]interface{}:
			// special case: an empty map may indicate flexible input, just provide a map here for now
			if len(vv.(map[interface{}]interface{})) == 0 {
				mymap[k.(string)] = "map[string]string"
			} else {
				mymap[k.(string)] = fmt.Sprintf("%s%s%sSpec", strings.Title(parentName), strings.Title(name), strings.Title(k.(string)))
				s += generateStruct(fmt.Sprintf("%s%s", strings.Title(parentName), strings.Title(name)), k.(string), vv.(map[interface{}]interface{}))
			}
		default:
			//fmt.Printf("nil\n")
		}
	}

	if len(mymap) > 0 {
		s += fmt.Sprintf("type %s%sSpec struct {\n", strings.Title(parentName), strings.Title(name))
		for k, vv := range mymap {
			s += fmt.Sprintf("  %s *%s `json:\"%s,omitempty\" yaml:\"%s,omitempty\"`\n", strings.Title(k), vv, k, k)
		}
		s += fmt.Sprintf("}\n")
	}

	return s
}

func generateDefaultSpec(prefix string, specName string, v map[interface{}]interface{}) string {
	// generate a spec with default values from v
	s := fmt.Sprintf("%sSpec{\n", specName)

	for k, v := range v {
		switch v.(type) {
		case bool:
			s += fmt.Sprintf("  %s%s: %v,\n", prefix, strings.Title(k.(string)), v.(bool))
		case string:
			if len(v.(string)) != 0 {
				s += fmt.Sprintf("  %s%s: \"%s\",\n", prefix, strings.Title(k.(string)), v.(string))
			}
		case int:
			s += fmt.Sprintf("  %s%s: %v,\n", prefix, strings.Title(k.(string)), v.(int))
		case []interface{}:
			// special case: an empty array may indicate flexible input, just provide a an array of maps here for now
			if len(v.([]interface{})) != 0 {
				s += fmt.Sprintf("  %s%s: %v,\n", prefix, strings.Title(k.(string)), v.([]interface{}))
				//} else {
				//mymap[k.(string)] = fmt.Sprintf("[]%s%s%sSpec", strings.Title(parentName), strings.Title(name), strings.Title(k.(string)))
			}
		case map[interface{}]interface{}:
			// special case: an empty map may indicate flexible input, just provide a map here for now
			if len(v.(map[interface{}]interface{})) == 0 {
				//mymap[k.(string)] = "map[string]string"
			} else {
				//mymap[k.(string)] = fmt.Sprintf("%s%s%sSpec", strings.Title(parentName), strings.Title(name), strings.Title(k.(string)))
				s += fmt.Sprintf("  %s%s: %v,\n", prefix, strings.Title(k.(string)),
					generateDefaultSpec(fmt.Sprintf("  %s", prefix), fmt.Sprintf("%s%s", strings.Title(specName), strings.Title(k.(string))), v.(map[interface{}]interface{})))
				//generateDefaultSpec(fmt.Sprintf("%s%s", strings.Title(parentName), strings.Title(name)), k.(string), vv.(map[interface{}]interface{}))
			}
		default:
			//fmt.Printf("nil\n")
		}
	}
	s += fmt.Sprintf("%s}", prefix)

	return s
}

func generateNewObjFunc(specName string, v map[interface{}]interface{}) string {
	// generate a go function that produces a new spec with default values from values.yaml
	// the caller should then overwrite them with what appears in the CRD.
	s := fmt.Sprintf("func New%sSpec() *%sSpec{\n", specName, specName)
	s += fmt.Sprintf("  return &%s\n", generateDefaultSpec("  ", specName, v))
	s += fmt.Sprintf("}\n")

	return s
}

func main() {
	if len(os.Args) < 4 {
		log.Printf("usage:   %s <packageName> <specName> <valuesYaml> [<outputFile>]", os.Args[0])
		log.Printf("example: %s v1alpha1 LibertyApp values.yaml pkg/apis/liberty/v1alpha/libertyappspec.go", os.Args[0])
		os.Exit(0)
	}

	specName := os.Args[2]
	valuesYaml := os.Args[3]

	data, err := ioutil.ReadFile(valuesYaml)
	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}

	s := fmt.Sprintf("package %s\n\n", os.Args[1])
	s += generateStruct("", specName, m)
	//s += generateNewObjFunc(specName, m)

	if len(os.Args) > 4 {
		f, err := os.Create(os.Args[4])
		if err != nil {
			log.Fatalf("error: %v", err)
			os.Exit(1)
		}

		i, err := f.WriteString(s)
		if err != nil {
			log.Fatalf("error: %v", err)
			os.Exit(1)
		}

		fmt.Printf("%d bytes written to %s. ", i, os.Args[4])

		f.Close()

		os.Exit(0)
	}

	fmt.Printf(s)
}
