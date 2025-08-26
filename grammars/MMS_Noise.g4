parser grammar MMS_Noise;
options {
    tokenVocab = MMSLexer;
}

import MMS_Lang_Parsers;

worldgenDeclaration: Keyword_WorldGen CurlyOpen NL* ((noiseDeclaration NL*)*) noiseDeclaration NL* CurlyClose;

noiseDeclaration: Keyword_Noise Identifier NL* noiseDefinition;
noiseDefinition: number NL* (SquareOpen NL* (number NL*)+ SquareClose)?;