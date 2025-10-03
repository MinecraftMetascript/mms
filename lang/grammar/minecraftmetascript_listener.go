// Code generated from ./grammar/MinecraftMetascript.g4 by ANTLR 4.13.2. DO NOT EDIT.

package grammar // MinecraftMetascript
import "github.com/antlr4-go/antlr/v4"

// MinecraftMetascriptListener is a complete listener for a parse tree produced by MinecraftMetascriptParser.
type MinecraftMetascriptListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterNamedBlock is called when entering the namedBlock production.
	EnterNamedBlock(c *NamedBlockContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterResourceReference is called when entering the resourceReference production.
	EnterResourceReference(c *ResourceReferenceContext)

	// EnterFn is called when entering the fn production.
	EnterFn(c *FnContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterCondGrouped is called when entering the condGrouped production.
	EnterCondGrouped(c *CondGroupedContext)

	// EnterCondNegate is called when entering the condNegate production.
	EnterCondNegate(c *CondNegateContext)

	// EnterCondAnd is called when entering the condAnd production.
	EnterCondAnd(c *CondAndContext)

	// EnterCondOr is called when entering the condOr production.
	EnterCondOr(c *CondOrContext)

	// EnterCondPrimary is called when entering the condPrimary production.
	EnterCondPrimary(c *CondPrimaryContext)

	// EnterRootCondition is called when entering the rootCondition production.
	EnterRootCondition(c *RootConditionContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitNamedBlock is called when exiting the namedBlock production.
	ExitNamedBlock(c *NamedBlockContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitResourceReference is called when exiting the resourceReference production.
	ExitResourceReference(c *ResourceReferenceContext)

	// ExitFn is called when exiting the fn production.
	ExitFn(c *FnContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitCondGrouped is called when exiting the condGrouped production.
	ExitCondGrouped(c *CondGroupedContext)

	// ExitCondNegate is called when exiting the condNegate production.
	ExitCondNegate(c *CondNegateContext)

	// ExitCondAnd is called when exiting the condAnd production.
	ExitCondAnd(c *CondAndContext)

	// ExitCondOr is called when exiting the condOr production.
	ExitCondOr(c *CondOrContext)

	// ExitCondPrimary is called when exiting the condPrimary production.
	ExitCondPrimary(c *CondPrimaryContext)

	// ExitRootCondition is called when exiting the rootCondition production.
	ExitRootCondition(c *RootConditionContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
