parser grammar VerticalAnchor_Parser;

options {
    tokenVocab = Main_Lexer;
}
import Lang_Parser;

verticalAnchorDefinition: Keyword_Anchor CurlyOpen verticalAnchor CurlyClose;
verticalAnchor: verticalAnchor_Absolute | verticalAnchor_AboveBottom | verticalAnchor_BelowTop | verticalAnchor_Reference;
verticalAnchor_Absolute: Keyword_Absolute RoundOpen Int RoundClose;
verticalAnchor_AboveBottom: Keyword_AboveBottom RoundOpen Int RoundClose;
verticalAnchor_BelowTop: Keyword_BelowTop RoundOpen Int RoundClose;
verticalAnchor_Reference: resourceReference;