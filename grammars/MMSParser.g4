parser grammar MMSParser;
options {
	tokenVocab = MMSLexer;
}

import MMS_SurfaceRules, MMS_Lang_Parsers, MMS_Noise;

namespaceDeclaration: Keyword_Namespace resourceReference SemiColon;

statement: surfaceDeclaration | worldgenDeclaration;

mmsFile: NL* namespaceDeclaration NL* (statement NL+)* statement? NL* EOF;