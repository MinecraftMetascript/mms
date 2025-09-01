parser grammar Main_Parser;

import SurfaceRules_Parser, VerticalAnchor_Parser;

options {
    tokenVocab = Main_Lexer;
}

script: NL* namespaceDefinition NL* (declaration NL*)+;

namespaceDefinition: Keyword_Namespace Identifier SemiColon;
declaration: Identifier Operator_Declare NL* definition;
definition: surfaceRuleDefinition | surfaceConditionDefinition | verticalAnchorDefinition;

