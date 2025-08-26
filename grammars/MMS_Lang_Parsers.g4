parser grammar MMS_Lang_Parsers;

options {
	tokenVocab = MMSLexer;
}

import MMS_Keyword_Rule;

reference: (keyword | Identifier) Colon (keyword | Identifier);

resourceReference: reference | (keyword | Identifier);

number: Int | Float;