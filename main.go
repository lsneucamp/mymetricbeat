package main

import (
	"os"

	"github.com/elastic/beats/libbeat/cmd/instance"
	"github.com/elastic/beats/metricbeat/beater"

	// Comment out the following line to exclude all official metricbeat modules and metricsets
	_ "github.com/elastic/beats/metricbeat/include"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/lsneucamp/mymetricbeat/include"
)

var Name = "mymetricbeat"
var IndexPrefix = "mymetricbeat"

func main() {
	if err := instance.Run(Name, IndexPrefix, "", beater.DefaultCreator()); err != nil {
		os.Exit(1)
	}
}
