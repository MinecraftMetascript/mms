grammar MinecraftMetascript;

import Surface, Noise, Core_Lang;

script: NL* (namespace NL*)*;

namespaceDeclaration: 'namespace' Identifier;
namespace: namespaceDeclaration NL* '{' NL* (contentBlocks NL*)* NL* '}';

contentBlocks: surfaceBlock | noiseBlock;