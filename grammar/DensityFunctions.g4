grammar DensityFunctions;

import Core_Lang, Noise;

sharedBuilder_XzScale: '.XZScale(' number ')';
sharedBuilder_YScale: '.YScale(' number ')';
sharedBuilder_XzFactor: '.XZFactor(' number ')';
sharedBuilder_YFactor: '.YFactor(' number ')';


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

densityFn_Noise: (
    ('Noise' NL* '(' NL* (resourceReference) NL* ')') | noise
) (densityFn_NoiseBuilder NL*)*;

densityFn_NoiseBuilder: sharedBuilder_XzScale | sharedBuilder_YScale;

DensityFn_CacheKind: '2d' | 'Once' | 'All';
densityFn_Cache: 'Cache' NL* '(' NL* DensityFn_CacheKind NL* ',' NL* densityFn NL* ')';

densityFn_DualInput: (
    | 'Min'
    | 'Max'
) NL* '(' NL* densityFn ',' NL* densityFn NL* ')';



densityFn_OldBlendedNoise: 'OldBlendedNoise' '(' ')' (densityFn_OldBlendedNoiseBuilder NL*)*;
densityFn_OldBlendedNoiseBuilder:
      sharedBuilder_XzScale
    | sharedBuilder_YScale
    | sharedBuilder_XzFactor
    | sharedBuilder_YFactor
    | densityFn_OldBlendedNoiseBuilder_Smear;

densityFn_OldBlendedNoiseBuilder_Smear: 'SmearMul' '(' number ')';

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