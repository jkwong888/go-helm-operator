package options

import (
	"github.com/spf13/pflag"
)

// ReconcilePeriod contains the necessary values to create a new controller
var ReconcilePeriod string

func init() {
	pflag.StringVar(&ReconcilePeriod, "reconcilePeriod", "15s", "Reconcile Period")
}
