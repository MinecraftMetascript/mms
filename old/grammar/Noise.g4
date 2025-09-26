grammar Noise;
import Core_Lang;

noiseBlock: 'Noise' NL* '{' NL* (noiseDeclaration NL*)* NL* '}';

noiseDeclaration: declare noise;

noise: 'Noise' noiseDefinition;

noiseDefinition: '(' NL* Int NL* ')' NL* (noise_Builder NL*)*;

noise_Builder: builder_Amplitudes;
