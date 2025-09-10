grammar NoiseSettings;

import Core_Lang, BlockState, NoiseRouter, Surface;

noiseSettingsBlock: 'NoiseSettings' NL* '{' NL* (noiseSettingsDeclaration NL*)* NL* '}';

noiseSettingsDeclaration: declare noiseSettings;
noiseSettings: 'NoiseSettings' NL* '(' NL* ')' NL* (noiseSettings_Builder NL*)*;


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
builder_NoiseSize: '.NoiseSize' '(' Int ',' Int ')';
builder_NoiseRouter: '.NoiseRouter' '(' noiseRouter ')';
builder_SeaLevel: '.SeaLevel' '(' Int ')';
builder_DisableCreatures: '.DisableCreatures' '(' ')';
builder_DisableVeins: '.DisableVeins' '(' ')';
builder_DisableAquifers: '.DisableAquifers' '(' ')';
builder_LegacyRandomSource: '.LegacyRandomSource' '(' ')';
builder_DefaultBlock: '.DefaultBlock' '(' (blockState | resourceReference) ')';
builder_DefaultFluid: '.DefaultFluid' '(' (blockState | resourceReference) ')';
builder_SpawnTarget: '.SpawnTarget' '(' /* TODO: Implement */ ')';
builder_MinY: '.MinY' '(' Int ')';
builder_Height: '.Height' '(' Int ')';

builder_SurfaceRule: '.SurfaceRule' '(' surfaceRule ')';
