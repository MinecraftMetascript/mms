parser grammar MMS_Keyword_Rule;
options {
	tokenVocab = MMSLexer;
}
keyword:
	Keyword_Surface
	| Keyword_Rule
	| Keyword_Condition
	| Keyword_Sequence
	| Keyword_Block
	| Keyword_Bandlands
	| Keyword_AbovePreliminarySurface
	| Keyword_Biome
	| Keyword_Hole
	| Keyword_Noise
	| Keyword_Steep
	| Keyword_StoneDepth
	| Keyword_Freezing
	| Keyword_Temperature
	| Keyword_VerticalGradient
	| Keyword_AboveWater
	| Keyword_YAbove
	| Keyword_Floor
	| Keyword_Ceiling
	| Keyword_And
	| Keyword_Add
	| Keyword_Sub
	| Keyword_Absolute
	| Keyword_AboveBottom
	| Keyword_BelowTop
	| Keyword_Namespace
	| Keyword_If
	| Keyword_Else
	| Keyword_In;
