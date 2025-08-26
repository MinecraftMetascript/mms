package lib

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

type VerticalAnchorType string

const (
	VerticalAnchorType_Absolute    VerticalAnchorType = "absolute"
	VerticalAnchorType_AboveBottom VerticalAnchorType = "above_bottom"
	VerticalAnchorType_BelowTop    VerticalAnchorType = "below_top"
)

type VerticalAnchor struct {
	json.Marshaler
	Type  VerticalAnchorType
	Value int
}

func (v VerticalAnchor) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"%s\": %d}", v.Type, v.Value)), nil
}

func ParseVerticalAnchor(anchor grammars.IVerticalAnchorContext) (*VerticalAnchor, error) {
	if rule := anchor.VerticalAnchor_Absolute(); rule != nil {
		value, err := strconv.Atoi(rule.Int().GetText())
		if err != nil {
			return nil, err
		}
		return &VerticalAnchor{
			Type:  VerticalAnchorType_Absolute,
			Value: value,
		}, nil
	}
	if rule := anchor.VerticalAnchor_AboveBottom(); rule != nil {
		value, err := strconv.Atoi(rule.Int().GetText())
		if err != nil {
			return nil, err
		}
		return &VerticalAnchor{
			Type:  VerticalAnchorType_AboveBottom,
			Value: value,
		}, nil
	}
	if rule := anchor.VerticalAnchor_BelowTop(); rule != nil {
		value, err := strconv.Atoi(rule.Int().GetText())
		if err != nil {
			return nil, err
		}
		return &VerticalAnchor{
			Type:  VerticalAnchorType_BelowTop,
			Value: value,
		}, nil
	}

	return nil, fmt.Errorf("unknown anchor type: %s", anchor.GetText())
}
