lexer grammar Main_Lexer;
// We are using an import for core functions like "Identifier" so that they don't supercede all other imports
import Lang_Lexer;

Keyword_Namespace: 'namespace';

Keyword_Absolute: 'Absolute';
Keyword_AboveBottom: 'AboveBottom';
Keyword_BelowTop: 'BelowTop';

Keyword_SurfaceRule: 'SurfaceRule';
Keyword_SurfaceCondition: 'SurfaceCondition';
Keyword_Block: 'Block';
Keyword_If: 'If';
Keyword_Sequence: 'Sequence';
Keyword_Bandlands: 'Bandlands';


SurfaceRule_Head_Block: Keyword_SurfaceRule Dot Keyword_Block;
SurfaceRule_Head_Bandlands: Keyword_SurfaceRule Dot Keyword_Bandlands;
SurfaceRule_Head_Condition: Keyword_SurfaceRule Dot Keyword_If;
SurfaceRule_Head_Sequence: Keyword_SurfaceRule Dot Keyword_Sequence;

SurfaceCondition_Head_AboveSurface: Keyword_SurfaceCondition Dot Keyword_AboveSurface;
SurfaceCondition_Head_Biome: Keyword_SurfaceCondition Dot Keyword_Biome;
SurfaceCondition_Head_Hole: Keyword_SurfaceCondition Dot Keyword_Hole;
SurfaceCondition_Head_Noise: Keyword_SurfaceCondition Dot Keyword_Noise;
SurfaceCondition_Head_Steep: Keyword_SurfaceCondition Dot Keyword_Steep;
SurfaceCondition_Head_StoneDepth: Keyword_SurfaceCondition Dot Keyword_StoneDepth;
SurfaceCondition_Head_Freezing: Keyword_SurfaceCondition Dot Keyword_Freezing;
SurfaceCondition_Head_Temperature: Keyword_SurfaceCondition Dot Keyword_Temperature;
SurfaceCondition_Head_VerticalGradient: Keyword_SurfaceCondition Dot Keyword_VerticalGradient;
SurfaceCondition_Head_AboveWater: Keyword_SurfaceCondition Dot Keyword_AboveWater;
SurfaceCondition_Head_YAbove: Keyword_SurfaceCondition Dot Keyword_YAbove;
SurfaceCondition_Head_Floor: Keyword_SurfaceCondition Dot Keyword_Floor;
SurfaceCondition_Head_Ceiling: Keyword_SurfaceCondition Dot Keyword_Ceiling;



Keyword_AboveSurface: 'AboveSurface';
Keyword_Biome: 'Biome';
Keyword_Hole: 'Hole';
Keyword_Noise: 'Noise';
Keyword_Steep: 'Steep';
Keyword_StoneDepth: 'StoneDepth';
Keyword_Freezing: 'Freezing';
Keyword_Temperature: 'Temperature';
Keyword_VerticalGradient: 'VerticalGradient';
Keyword_AboveWater: 'AboveWater';
Keyword_YAbove: 'YAbove';
Keyword_Floor: 'Floor';
Keyword_Ceiling: 'Ceiling';
Keyword_And: 'and';
Keyword_Or: 'or';
Keyword_Add: 'add';
Keyword_Sub: 'sub';
///
