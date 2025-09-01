parser grammar Lang_Parser;

options {
    tokenVocab= Main_Lexer;
}

resourceReference: (Identifier Colon)? Identifier;
number: Int | Float;


