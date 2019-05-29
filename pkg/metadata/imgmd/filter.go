package imgmd

import (
	"fmt"
	"github.com/luxas/ignite/pkg/filter"
	"strings"
)

// Compile-time assert to verify interface compatibility
var _ filter.Filter = &ImageFilter{}

type ImageFilter struct {
	prefix string
}

func NewImageFilter(p string) *ImageFilter {
	return &ImageFilter{
		prefix: p,
	}
}

func (n *ImageFilter) Filter(f filter.Filterable) (bool, error) {
	md, ok := f.(*ImageMetadata)
	if !ok {
		return false, fmt.Errorf("failed to assert Filterable %v to ImageMetadata", f)
	}

	return strings.HasPrefix(md.ID, n.prefix) || strings.HasPrefix(md.Name, n.prefix), nil
}

func LoadImageMetadata(id string) (*ImageMetadata, error) {
	md := NewImageMetadata(id, "-") // A blank name triggers an unnecessary name generation
	err := md.Load()
	return md, err
}

func LoadImageMetadataFilterable(id string) (filter.Filterable, error) {
	return LoadImageMetadata(id)
}