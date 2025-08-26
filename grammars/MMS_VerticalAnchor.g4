parser grammar MMS_VerticalAnchor;

options {
	tokenVocab = MMSLexer;
}

import MMS_Lang_Parsers;

verticalAnchor: verticalAnchor_Absolute | verticalAnchor_AboveBottom | verticalAnchor_BelowTop;
verticalAnchor_Absolute: Keyword_Absolute Int;
verticalAnchor_AboveBottom: Keyword_AboveBottom Int;
verticalAnchor_BelowTop: Keyword_BelowTop Int;
