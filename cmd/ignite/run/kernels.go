package run

import (
	api "github.com/weaveworks/ignite/pkg/apis/ignite/v1alpha1"
	"github.com/weaveworks/ignite/pkg/client"
	"github.com/weaveworks/ignite/pkg/filter"
	"github.com/weaveworks/ignite/pkg/util"
)

type kernelsOptions struct {
	allKernels []*api.Kernel
}

func NewKernelsOptions() (ko *kernelsOptions, err error) {
	ko = &kernelsOptions{}
	ko.allKernels, err = client.Kernels().FindAll(filter.NewAllFilter())
	return
}

func Kernels(ko *kernelsOptions) error {
	o := util.NewOutput()
	defer o.Flush()

	o.Write("KERNEL ID", "NAME", "CREATED", "SIZE", "VERSION")
	for _, kernel := range ko.allKernels {
		o.Write(kernel.GetUID(), kernel.GetName(), kernel.GetCreated(), kernel.Status.OCISource.Size.String(), kernel.Status.Version)
	}

	return nil
}
