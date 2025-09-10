grammar MinecraftMetascript;

import Surface,DensityFunctions, Noise, NoiseRouter, NoiseSettings, Core_Lang;

script: NL* (namespace NL*)*;

namespaceDeclaration: 'Namespace' Identifier;
namespace: namespaceDeclaration NL* '{' NL* (contentBlocks NL*)* '}';

contentBlocks: surfaceBlock | noiseBlock | densityFnBlock | noiseRouterBlock | noiseSettingsBlock;