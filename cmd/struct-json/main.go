package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jkwong888/websphere-liberty-operator/pkg/apis/liberty/v1alpha1"
	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

func main() {
	//	o.Spec = *v1alpha1.NewLibertyAppSpec()

	b, err := ioutil.ReadFile(os.Args[1])
	newApp := &v1alpha1.LibertyApp{}
	err = yaml.Unmarshal(b, newApp)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	//err = overwriteSpec(&o.Spec, newApp.Spec)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	oMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(newApp)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	o := &unstructured.Unstructured{
		Object: oMap,
	}

	b, err = yaml.Marshal(o)

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	fmt.Printf(string(b))
}
