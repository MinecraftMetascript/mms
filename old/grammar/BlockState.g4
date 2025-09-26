grammar BlockState;
import Core_Lang;

// Name
blockState: 'Block' NL* '(' NL* resourceReference NL* ')' NL* (blockState_Builder NL*)*;
blockState_Builder: '.' 'Property' '(' String ',' String ')';

