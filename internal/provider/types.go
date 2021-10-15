package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type BoolMarshalerType struct{}

func (BoolMarshalerType) TerraformType(context.Context) tftypes.Type {
	return tftypes.Bool
}

func (BoolMarshalerType) ValueFromTerraform(ctx context.Context, val tftypes.Value) (attr.Value, error) {
	inner, err := types.BoolType.ValueFromTerraform(ctx, val)
	if err != nil {
		return nil, err
	}
	return BoolMarshaler{inner.(types.Bool)}, nil
}

func (b BoolMarshalerType) Equal(other attr.Type) bool {
	_, ok := other.(BoolMarshalerType)
	return ok
}

func (b BoolMarshalerType) String() string {
	return "BoolMarshalerType"
}

func (b BoolMarshalerType) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	return nil, fmt.Errorf("cannot apply AttributePathStep %T to %s", step, b.String())
}

type BoolMarshaler struct {
	types.Bool
}

func (b BoolMarshaler) Type(context.Context) attr.Type {
	return BoolMarshalerType{}
}

func (b BoolMarshaler) MarshalJSON() ([]byte, error) {
	if b.Null || b.Unknown {
		return json.Marshal((*bool)(nil))
	}
	return json.Marshal(b.Value)
}

func (b *BoolMarshaler) UnmarshalJSON(data []byte) error {
	var bPtr *bool
	if err := json.Unmarshal(data, &bPtr); err != nil {
		return err
	}
	b.Unknown = false
	if bPtr == nil {
		b.Value = false
		b.Null = true
	} else {
		b.Value = *bPtr
		b.Null = false
	}
	return nil
}
