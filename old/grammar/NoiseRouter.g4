grammar NoiseRouter;

import Core_Lang, DensityFunctions;


noiseRouterBlock: 'NoiseRouter' NL* '{' NL* (noiseRouterDeclaration NL*)* NL* '}';

noiseRouterDeclaration: declare noiseRouter;

noiseRouter: 'Router' NL* '(' NL* ')' NL* (noiseRouter_Builder NL*)*;


noiseRouter_Builder: '.' (
      'FinalDensity'
    | 'Barrier'
    | 'FluidLevelFloodedness'
    | 'FluidLevelSpread'
    | 'Lava'
    | 'VeinToggle'
    | 'VeinRidged'
    | 'VeinGap'
    | 'Temperature'
    | 'Vegetation'
    | 'Continents'
    | 'Erosion'
    | 'Depth'
    | 'Ridges') NL* '(' NL* densityFn NL* ')';




