grammar MinecraftMetascript;

// Imports maintain priority order
// (e.g. Core_Lang has the lowest priority, and only matches when no other grammars contain matches)
import
    Surface,
    DensityFunctions,
    Noise,
    NoiseRouter,
    NoiseSettings,
    Core_Lang;

script: NL* (namespace NL*)*;

namespaceDeclaration: 'Namespace' Identifier;
namespace: namespaceDeclaration NL* '{' NL* (contentBlocks NL*)* '}';

contentBlocks: surfaceBlock | noiseBlock | densityFnBlock | noiseRouterBlock | noiseSettingsBlock;