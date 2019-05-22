package yamlutil

import (
	"io"
	"strings"

	"gopkg.in/yaml.v2"
)

type yamldoc map[interface{}]interface{}

// SplitYaml split string with multiple yamls in it into a slice of string yamls
func SplitYaml(yamlStr string) ([]string, error) {
	var toRet []string

	strReader := strings.NewReader(yamlStr)
	decoder := yaml.NewDecoder(strReader)

	for true {
		intf := &yamldoc{}

		err := decoder.Decode(intf)
		if err == io.EOF {
			break
		} else if err != nil {
			// all non-EOF errors
			return nil, err
		}

		out, err := yaml.Marshal(intf)
		if err != nil {
			return nil, err
		}

		toRet = append(toRet, string(out))
	}

	return toRet, nil
}

// ToYaml convert a list of runtime objects into a manifest yaml containing all objects
func ToYaml(objects []interface{}) (string, error) {
	strBuilder := &strings.Builder{}
	encoder := yaml.NewEncoder(strBuilder)

	for _, obj := range objects {
		err := encoder.Encode(obj)
		if err != nil {
			return "", err
		}
	}

	return strBuilder.String(), nil
}
