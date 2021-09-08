package provider

import (
	"context"
)

type emptyDescriptions struct {
}

func (emptyDescriptions) Description(ctx context.Context) string {
	return ""
}

func (emptyDescriptions) MarkdownDescription(ctx context.Context) string {
	return ""
}
