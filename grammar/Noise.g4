grammar Noise;
import Core_Lang;

noiseBlock: 'Noise' NL* '{' NL* (noiseDeclaration NL*)* NL* '}';

noiseDeclaration: Identifier '=' noiseDefinition;

noiseDefinition: 'Noise(' NL* Int NL* ')' NL* (noiseDefinition_Builder NL*)*;

noiseDefinition_Builder: noiseDefinition_Builder_Octaves;
noiseDefinition_Builder_Octaves: '.Octaves(' (number ',')* number  ')';