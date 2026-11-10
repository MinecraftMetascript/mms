package sema

import (
	"strings"
	"testing"

	"github.com/minecraftmetascript/mms/engine/jsonv"
	"github.com/minecraftmetascript/mms/engine/spec"
)

// TestValidateUnionReportsInside checks that a value in a union is
// reported against the member it was written as. A block state is a
// string id or { id, properties }; an object with the wrong fields must
// say so rather than "expected a string", and an inline placed feature in
// a biome must point inside the feature rather than at the id member.
func TestValidateUnionReportsInside(t *testing.T) {
	p := spec.MustDefault()
	kinds := NewKinds(p)
	for _, tc := range []struct{ res, json, want string }{
		{
			"worldgen/feature",
			`{"type":"minecraft:tree","minimum_size":{"type":"minecraft:two_layers_feature_size"},"below_trunk_provider":"minecraft:soil_beneath_tree","decorators":[],
			  "foliage_placer":{"type":"minecraft:blob_foliage_placer","height":3,"offset":0,"radius":2},"foliage_provider":{"id":"minecraft:oak_leaves"},
			  "trunk_placer":{"type":"minecraft:straight_trunk_placer","base_height":4,"height_rand_a":2,"height_rand_b":0},
			  "trunk_provider":{"type":"minecraft:simple","state":{"Name":"minecraft:oak_log"}}}`,
			"$.trunk_provider.state: missing field 'id'",
		},
		{
			"worldgen/biome",
			`{"temperature":0,"downfall":0,"has_precipitation":false,"effects":{"water_color":1},"carvers":[],
			  "features":[[{"feature":{"type":"minecraft:no_op","config":{}},"placement":[{"type":"minecraft:count","count":"many"}]}]]}`,
			"$.features[0][0].placement[0].count",
		},
	} {
		v, err := jsonv.Parse([]byte(tc.json))
		if err != nil {
			t.Fatal(err)
		}
		errs := Validate(p, kinds, tc.res, v)
		if len(errs) == 0 || !strings.HasPrefix(errs[0], tc.want) {
			t.Errorf("%s: got %q, want an error starting %q", tc.res, errs, tc.want)
		}
	}
}
