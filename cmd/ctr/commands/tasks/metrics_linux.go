// +build linux

/*
   Copyright The containerd Authors.

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

package tasks

import (
	"errors"
	"text/tabwriter"

	v1 "github.com/containerd/cgroups/stats/v1"
	v2 "github.com/containerd/cgroups/v2/stats"
)

func printMetricsAsTable(w *tabwriter.Writer, anydata interface{}) error {
	var (
		data  *v1.Metrics
		data2 *v2.Metrics
	)
	switch v := anydata.(type) {
	case *v1.Metrics:
		data = v
	case *v2.Metrics:
		data2 = v
	default:
		return errors.New("cannot convert metric data to cgroups.Metrics")
	}

	if data != nil {
		printCgroupMetricsTable(w, data)
	} else {
		printCgroup2MetricsTable(w, data2)
	}
	return nil
}
