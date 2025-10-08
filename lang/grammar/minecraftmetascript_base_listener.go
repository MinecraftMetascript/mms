// Code generated from ./grammar/MinecraftMetascript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // MinecraftMetascript
import "github.com/antlr4-go/antlr/v4"

// BaseMinecraftMetascriptListener is a complete listener for a parse tree produced by MinecraftMetascriptParser.
type BaseMinecraftMetascriptListener struct{}

var _ MinecraftMetascriptListener = &BaseMinecraftMetascriptListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMinecraftMetascriptListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMinecraftMetascriptListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMinecraftMetascriptListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMinecraftMetascriptListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseMinecraftMetascriptListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseMinecraftMetascriptListener) ExitFile(ctx *FileContext) {}

// EnterNamedBlock is called when production namedBlock is entered.
func (s *BaseMinecraftMetascriptListener) EnterNamedBlock(ctx *NamedBlockContext) {}

// ExitNamedBlock is called when production namedBlock is exited.
func (s *BaseMinecraftMetascriptListener) ExitNamedBlock(ctx *NamedBlockContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseMinecraftMetascriptListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseMinecraftMetascriptListener) ExitBlock(ctx *BlockContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseMinecraftMetascriptListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseMinecraftMetascriptListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterResourceReference is called when production resourceReference is entered.
func (s *BaseMinecraftMetascriptListener) EnterResourceReference(ctx *ResourceReferenceContext) {}

// ExitResourceReference is called when production resourceReference is exited.
func (s *BaseMinecraftMetascriptListener) ExitResourceReference(ctx *ResourceReferenceContext) {}

// EnterFn is called when production fn is entered.
func (s *BaseMinecraftMetascriptListener) EnterFn(ctx *FnContext) {}

// ExitFn is called when production fn is exited.
func (s *BaseMinecraftMetascriptListener) ExitFn(ctx *FnContext) {}

// EnterFnArgBody is called when production fnArgBody is entered.
func (s *BaseMinecraftMetascriptListener) EnterFnArgBody(ctx *FnArgBodyContext) {}

// ExitFnArgBody is called when production fnArgBody is exited.
func (s *BaseMinecraftMetascriptListener) ExitFnArgBody(ctx *FnArgBodyContext) {}

// EnterValue is called when production value is entered.
func (s *BaseMinecraftMetascriptListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseMinecraftMetascriptListener) ExitValue(ctx *ValueContext) {}

// EnterCondGrouped is called when production condGrouped is entered.
func (s *BaseMinecraftMetascriptListener) EnterCondGrouped(ctx *CondGroupedContext) {}

// ExitCondGrouped is called when production condGrouped is exited.
func (s *BaseMinecraftMetascriptListener) ExitCondGrouped(ctx *CondGroupedContext) {}

// EnterCondNegate is called when production condNegate is entered.
func (s *BaseMinecraftMetascriptListener) EnterCondNegate(ctx *CondNegateContext) {}

// ExitCondNegate is called when production condNegate is exited.
func (s *BaseMinecraftMetascriptListener) ExitCondNegate(ctx *CondNegateContext) {}

// EnterCondAnd is called when production condAnd is entered.
func (s *BaseMinecraftMetascriptListener) EnterCondAnd(ctx *CondAndContext) {}

// ExitCondAnd is called when production condAnd is exited.
func (s *BaseMinecraftMetascriptListener) ExitCondAnd(ctx *CondAndContext) {}

// EnterCondOr is called when production condOr is entered.
func (s *BaseMinecraftMetascriptListener) EnterCondOr(ctx *CondOrContext) {}

// ExitCondOr is called when production condOr is exited.
func (s *BaseMinecraftMetascriptListener) ExitCondOr(ctx *CondOrContext) {}

// EnterCondPrimary is called when production condPrimary is entered.
func (s *BaseMinecraftMetascriptListener) EnterCondPrimary(ctx *CondPrimaryContext) {}

// ExitCondPrimary is called when production condPrimary is exited.
func (s *BaseMinecraftMetascriptListener) ExitCondPrimary(ctx *CondPrimaryContext) {}

// EnterRootCondition is called when production rootCondition is entered.
func (s *BaseMinecraftMetascriptListener) EnterRootCondition(ctx *RootConditionContext) {}

// ExitRootCondition is called when production rootCondition is exited.
func (s *BaseMinecraftMetascriptListener) ExitRootCondition(ctx *RootConditionContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseMinecraftMetascriptListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseMinecraftMetascriptListener) ExitConditional(ctx *ConditionalContext) {}

// EnterConditionalBody is called when production conditionalBody is entered.
func (s *BaseMinecraftMetascriptListener) EnterConditionalBody(ctx *ConditionalBodyContext) {}

// ExitConditionalBody is called when production conditionalBody is exited.
func (s *BaseMinecraftMetascriptListener) ExitConditionalBody(ctx *ConditionalBodyContext) {}

// EnterList is called when production list is entered.
func (s *BaseMinecraftMetascriptListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseMinecraftMetascriptListener) ExitList(ctx *ListContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseMinecraftMetascriptListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseMinecraftMetascriptListener) ExitNumber(ctx *NumberContext) {}
