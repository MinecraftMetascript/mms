parser grammar VerticalAnchor_Parser;

options {
    tokenVocab = Main_Lexer;
}
import Lang_Parser;

verticalAnchorDefinition: Keyword_Anchor CurlyOpen verticalAnchor CurlyClose;
verticalAnchor: verticalAnchor_Absolute | verticalAnchor_AboveBottom | verticalAnchor_BelowTop | verticalAnchor_Reference;
verticalAnchor_Absolute: Int;
verticalAnchor_AboveBottom: Tilde Int;
verticalAnchor_BelowTop: Tilde Dash Int;
verticalAnchor_Reference: resourceReference;