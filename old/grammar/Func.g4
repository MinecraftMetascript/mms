grammar Func;

import Core_Lang;

FuncParamType: 'int' | 'float';

Variable: '$' [a-zA-Z] [a-z0-9A-Z_]*;
paramDef: Variable ':' FuncParamType;
funcOf: 'func' NL* '(' NL* (paramDef NL* ',' NL*)* NL* (paramDef)? NL* ')' NL* '->';
param: Identifier | number;
funcCall: Identifier '(' (param ',')* (param)? ')';
