package noise

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/minecraftmetascript/mms/lang/grammars"
)

type Noise struct {
	json.Marshaler
	FirstOctave float64
	Amplitudes  []float64
}

func (n Noise) MarshalJSON() ([]byte, error) {
	if n.Amplitudes == nil {
		n.Amplitudes = make([]float64, 0)
	}
	return json.MarshalIndent(struct {
		FirstOctave float64   `json:"firstOctave"`
		Amplitudes  []float64 `json:"amplitudes"`
	}{
		FirstOctave: n.FirstOctave,
		Amplitudes:  n.Amplitudes,
	}, "", "  ")
}

func NewNoise(ctx *grammars.NoiseDefinitionContext) (*Noise, error) {
	firstOctave, err := strconv.ParseFloat(ctx.AllNumber()[0].GetText(), 64)
	if err != nil {
		return nil, err
	}
	additionalOctaves := make([]float64, 0)
	for _, octave := range ctx.AllNumber()[1:] {
		val, err := strconv.ParseFloat(octave.GetText(), 32)
		val = math.Round(val*1e6) / 1e6

		if err != nil {
			return nil, err
		}
		additionalOctaves = append(additionalOctaves, float64(val))
	}

	return &Noise{
		FirstOctave: firstOctave,
		Amplitudes:  additionalOctaves,
	}, nil

}
