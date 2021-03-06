/*
Copyright 2019 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Copyright 2020 CyVerse
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/iychoi/parcel-csi-driver/pkg/driver"

	"k8s.io/klog"
)

var (
	conf driver.Config
)

func main() {
	var version bool

	// Parse parameters
	flag.StringVar(&conf.Endpoint, "endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	flag.StringVar(&conf.NodeID, "nodeid", "", "node id")
	flag.StringVar(&conf.SecretPath, "secretpath", "/etc/parcel-csi-dirver", "Secret mount path")

	flag.BoolVar(&version, "version", false, "Print driver version information")

	klog.InitFlags(nil)
	flag.Parse()

	// Handle Version
	if version {
		info, err := driver.GetVersionJSON()

		if err != nil {
			klog.Fatalln(err)
		}

		fmt.Println(info)
		os.Exit(0)
	}

	klog.V(1).Infof("Driver version: %s", driver.GetDriverVersion())

	if conf.NodeID == "" {
		klog.Fatalln("Node ID is not given")
	}

	drv := driver.NewDriver(&conf)
	if err := drv.Run(); err != nil {
		klog.Fatalln(err)
	}

	os.Exit(0)
}
