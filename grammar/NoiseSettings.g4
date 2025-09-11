grammar NoiseSettings;

import Core_Lang, BlockState, NoiseRouter, Surface;

noiseSettingsBlock: 'NoiseSettings' NL* '{' NL* (noiseSettingsDeclaration NL*)* NL* '}';

noiseSettingsDeclaration: declare noiseSettings;
noiseSettings: 'NoiseSettings' NL* '(' NL* NL* ')' NL* (noiseSettings_Builder NL*)*;


noiseSettings_Builder:
      builder_SeaLevel
    | builder_DisableCreatures
    | builder_DisableVeins
    | builder_DisableAquifers
    | builder_LegacyRandomSource
    | builder_DefaultBlock
    | builder_DefaultFluid
    | builder_SpawnTarget
    | builder_MinY
    | builder_Height
    | builder_NoiseSize
    | builder_NoiseRouter
    | builder_SurfaceRule
    ;
builder_NoiseSize: '.NoiseSize' '(' NL* Int NL* ',' NL* Int NL* ')';
builder_NoiseRouter: '.NoiseRouter' '(' NL* noiseRouter NL* ')';
builder_SeaLevel: '.SeaLevel' '(' NL* Int NL* ')';
builder_DisableCreatures: '.DisableCreatures' '(' NL* NL* ')';
builder_DisableVeins: '.DisableVeins' '(' NL* NL* ')';
builder_DisableAquifers: '.DisableAquifers' '(' NL* NL* ')';
builder_LegacyRandomSource: '.LegacyRandomSource' '(' NL* NL* ')';
builder_DefaultBlock: '.DefaultBlock' '(' NL* (blockState | resourceReference) NL* ')';
builder_DefaultFluid: '.DefaultFluid' '(' NL* (blockState | resourceReference) NL* ')';
builder_SpawnTarget: '.SpawnTarget' '(' NL* /* TODO: Implement */ NL* ')';
builder_MinY: '.MinY' '(' NL* Int NL* ')';
builder_Height: '.Height' '(' NL* Int NL* ')';

builder_SurfaceRule: '.SurfaceRule' '(' NL* surfaceRule NL* ')';
