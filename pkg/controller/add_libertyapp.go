package controller

import (
	"github.com/jkwong888/websphere-liberty-operator/pkg/controller/libertyapp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, libertyapp.Add)
}
