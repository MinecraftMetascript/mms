grammar DensityFunctions;

import Core_Lang, Noise;

densityFnBlock: 'DensityFn' NL* '{' NL* (densityFnDeclaration NL*)* NL* '}';

sharedBuilder_XzScale: '.XZScale(' number ')';
sharedBuilder_YScale: '.YScale(' number ')';
sharedBuilder_XzFactor: '.XZFactor(' number ')';
sharedBuilder_YFactor: '.YFactor(' number ')';



densityFnDeclaration: Identifier NL* '=' NL* densityFn;
densityFn: (
      densityFn_SingleInput
    | densityFn_Cache
    | densityFn_DualInput
    | densityFn_Constant
    | densityFn_Noise
    | densityFn_NoInput
    | densityFn_OldBlendedNoise
    | densityFn_WierdScaledSampler
    | densityFn_Reference // should be last
) densityFn_Math?;

densityFn_NoInput: (
      'BlendAlpha'
    | 'BlendOffset'
    | 'EndIslands'
);

densityFn_SingleInput: (
      'Interpolated'
    | 'BlendDensity'
    | 'FlatCache'
    | 'Abs'
    | 'Square'
    | 'Cube'
    | 'HalfNeg'
    | 'QuarterNeg'
    | 'Squeeze'
    | 'Shift'
    | 'ShiftA'
    | 'ShiftB'
) NL* '(' NL* densityFn NL* ')';

densityFn_InlineNoise: noise;

densityFn_Noise: (
    ('Noise' NL* '(' NL* (resourceReference) NL* ')') | densityFn_InlineNoise
) NL* (densityFn_NoiseBuilder NL*)*;

densityFn_NoiseBuilder: sharedBuilder_XzScale | sharedBuilder_YScale;

DensityFn_CacheKind: '2d' | 'Once' | 'All';
densityFn_Cache: 'Cache' NL* '(' NL* DensityFn_CacheKind NL* ',' NL* densityFn NL* ')';

densityFn_DualInput: (
    | 'Min'
    | 'Max'
) NL* '(' NL* densityFn ',' NL* densityFn NL* ')';

densityFn_OldBlendedNoise: 'OldBlendedNoise' '(' ')' NL* (densityFn_OldBlendedNoiseBuilder NL*)*;
densityFn_OldBlendedNoiseBuilder:
      sharedBuilder_XzScale
    | sharedBuilder_YScale
    | sharedBuilder_XzFactor
    | sharedBuilder_YFactor
    | densityFn_OldBlendedNoiseBuilder_Smear;

densityFn_OldBlendedNoiseBuilder_Smear: 'SmearMul' '(' number ')';

densityFn_WierdScaledSampler: 'WeirdScaledSampler' '(' densityFn ')' NL* (densityFn_WierdScaledSamplerBuilder NL*)*;

densityFn_WierdScaledSamplerBuilder:   densityFn_WierdScaledSamplerBuilder_Type1
                                     | densityFn_WierdScaledSamplerBuilder_Type2
                                     | densityFn_WierdScaledSamplerBuilder_Noise
                                    ;
densityFn_WierdScaledSamplerBuilder_Type1: '.Type1' '(' ')';
densityFn_WierdScaledSamplerBuilder_Type2: '.Type2' '(' ')';
densityFn_WierdScaledSamplerBuilder_Noise: '.Noise' (noiseDefinition | ('(' resourceReference ')'));



densityFn_Constant: number;
densityFn_Reference: resourceReference;
densityFn_Math: ('+' | '*') densityFn;


/*
5.8	weird_scaled_sampler
5.9	shifted_noise
5.10	range_choice
5.14	clamp
5.15	spline
5.17	y_clamped_gradient
*/