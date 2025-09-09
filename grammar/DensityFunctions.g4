grammar DensityFunctions;

import Core_Lang, Noise;

densityFnBlock: 'DensityFn' NL* '{' NL* (densityFnDeclaration NL*)* NL* '}';

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
    | densityFn_ShiftedNoise
    | densityFn_RangeChoice
    | densityFn_Clamp
    | densityFn_YClampedGradient
    | densityFn_SplineFn
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

densityFn_NoiseBuilder: builder_XZScale | builder_YScale;

DensityFn_CacheKind: '2d' | 'Once' | 'All';
densityFn_Cache: 'Cache' NL* '(' NL* DensityFn_CacheKind NL* ',' NL* densityFn NL* ')';

densityFn_DualInput: (
    | 'Min'
    | 'Max'
) NL* '(' NL* densityFn ',' NL* densityFn NL* ')';

densityFn_OldBlendedNoise: 'OldBlendedNoise' NL* '(' NL* ')' NL* (densityFn_OldBlendedNoiseBuilder NL*)*;
densityFn_OldBlendedNoiseBuilder:
      builder_XZScale
    | builder_YScale
    | builder_XZFactor
    | builder_YFactor
    | builder_Smear;


densityFn_WierdScaledSampler: 'WeirdScaledSampler' NL* '(' NL* densityFn NL* ')' NL* (densityFn_WierdScaledSamplerBuilder NL*)*;

densityFn_WierdScaledSamplerBuilder: builder_Type1 | builder_Type2 | builder_Noise;

densityFn_ShiftedNoise: 'ShiftedNoise' NL* '(' NL* ')' NL* (densityFn_ShiftedNoiseBuilder NL*)*;
densityFn_ShiftedNoiseBuilder: builder_Noise | builder_XZScale | builder_YScale | builder_ShiftX | builder_ShiftY | builder_ShiftZ;

densityFn_RangeChoice: 'RangeChoice' NL* '(' densityFn ')' NL* (densityFn_RangeChoiceBuilder NL*)*;
densityFn_RangeChoiceBuilder: builder_Min | builder_Max | builder_InRange | builder_OutRange;

densityFn_Clamp: 'Clamp' NL* '(' densityFn ')' NL* (densityFn_ClampBuilder NL*)*;
densityFn_ClampBuilder: builder_Min | builder_Max;

densityFn_YClampedGradient: 'YClampedGradient' NL* '(' NL* ')' NL* (densityFn_YClampedGradientBuilder NL*)*;
densityFn_YClampedGradientBuilder: builder_TopLiteral | builder_BottomLiteral | builder_Min | builder_Max;

densityFn_SplineFn: densityFn_Spline;
densityFn_Spline: 'Spline' NL* (('(' NL* densityFn NL* ')' NL* (densityFn_SplinePoint NL*)* ) | '(' number ')');
densityFn_SplinePoint: '.Point' '(' number ',' (number | densityFn_Spline | resourceReference) ',' number ')';

densityFn_Constant: number;
densityFn_Reference: resourceReference;
densityFn_Math: ('+' | '*') densityFn;


/*
5.15	spline
*/