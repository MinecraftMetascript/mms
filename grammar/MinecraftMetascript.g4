grammar MinecraftMetascript;

import Surface, Core_Lang;

script: (namespace NL*)*;

namespaceDeclaration: 'namespace' Identifier;
namespace: namespaceDeclaration NL* '{' NL* (contentBlocks NL*)* NL* '}';

contentBlocks: surface;