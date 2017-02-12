package govaluate

import (
	"errors"
	"time"
	"fmt"
)

var stageSymbolMap = map[OperatorSymbol]evaluationOperator{
	EQ:             equalStage,
	NEQ:            notEqualStage,
	GT:             gtStage,
	LT:             ltStage,
	GTE:            gteStage,
	LTE:            lteStage,
	REQ:            regexStage,
	NREQ:           notRegexStage,
	AND:            andStage,
	OR:             orStage,
	IN:             inStage,
	BITWISE_OR:     bitwiseOrStage,
	BITWISE_AND:    bitwiseAndStage,
	BITWISE_XOR:    bitwiseXORStage,
	BITWISE_LSHIFT: leftShiftStage,
	BITWISE_RSHIFT: rightShiftStage,
	PLUS:           addStage,
	MINUS:          subtractStage,
	MULTIPLY:       multiplyStage,
	DIVIDE:         divideStage,
	MODULUS:        modulusStage,
	EXPONENT:       exponentStage,
	NEGATE:         negateStage,
	INVERT:         invertStage,
	BITWISE_NOT:    bitwiseNotStage,
	TERNARY_TRUE:   ternaryIfStage,
	TERNARY_FALSE:  ternaryElseStage,
	COALESCE:       ternaryElseStage,
	SEPARATE:       separatorStage,
}

/*
	A "precedent" is a function which will recursively parse new evaluateionStages from a given stream of tokens.
	It's called a `precedent` because it is expected to handle exactly what precedence of operator,
	and defer to other `precedent`s for other operators.
*/
type precedent func(stream *tokenStream) (*evaluationStage, error)

/*
	A convenience function for specifying the behavior of a `precedent`.
	Most `precedent` functions can be described by the same function, just with different type checks, symbols, and error formats.
	This struct is passed to `makePrecedentFromPlanner` to create a `precedent` function.
*/
type precedencePlanner struct {
	validSymbols map[string]OperatorSymbol
	validKinds   []TokenKind

	typeErrorFormat string

	next      precedent
	nextRight precedent
}

var planPrefix precedent
var planExponential precedent
var planMultiplicative precedent
var planAdditive precedent
var planBitwise precedent
var planShift precedent
var planComparator precedent
var planLogical precedent
var planTernary precedent
var planSeparator precedent

func init() {

	// all these stages can use the same code (in `planPrecedenceLevel`) to execute,
	// they simply need different type checks, symbols, and recursive precedents.
	// While not all precedent phases are listed here, most can be represented this way.
	planPrefix = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols:    PREFIX_SYMBOLS,
		validKinds:      []TokenKind{PREFIX},
		typeErrorFormat: TYPEERROR_PREFIX,
		nextRight:       planFunction,
	})
	planExponential = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols:    EXPONENTIAL_SYMBOLS,
		validKinds:      []TokenKind{MODIFIER},
		typeErrorFormat: TYPEERROR_MODIFIER,
		next:            planFunction,
	})
	planMultiplicative = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols:    MULTIPLICATIVE_SYMBOLS,
		validKinds:      []TokenKind{MODIFIER},
		typeErrorFormat: TYPEERROR_MODIFIER,
		next:            planExponential,
	})
	planAdditive = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols:    ADDITIVE_SYMBOLS,
		validKinds:      []TokenKind{MODIFIER},
		typeErrorFormat: TYPEERROR_MODIFIER,
		next:            planMultiplicative,
	})
	planShift = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols:    BITWISE_SHIFT_SYMBOLS,
		validKinds:      []TokenKind{MODIFIER},
		typeErrorFormat: TYPEERROR_MODIFIER,
		next:            planAdditive,
	})
	planBitwise = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols:    BITWISE_SYMBOLS,
		validKinds:      []TokenKind{MODIFIER},
		typeErrorFormat: TYPEERROR_MODIFIER,
		next:            planShift,
	})
	planComparator = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols:    COMPARATOR_SYMBOLS,
		validKinds:      []TokenKind{COMPARATOR},
		typeErrorFormat: TYPEERROR_COMPARATOR,
		next:            planBitwise,
	})
	planLogical = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols:    LOGICAL_SYMBOLS,
		validKinds:      []TokenKind{LOGICALOP},
		typeErrorFormat: TYPEERROR_LOGICAL,
		next:            planComparator,
	})
	planTernary = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols:    TERNARY_SYMBOLS,
		validKinds:      []TokenKind{TERNARY},
		typeErrorFormat: TYPEERROR_TERNARY,
		next:            planLogical,
	})
	planSeparator = makePrecedentFromPlanner(&precedencePlanner{
		validSymbols: SEPARATOR_SYMBOLS,
		validKinds:   []TokenKind{SEPARATOR},
		next:         planTernary,
	})
}

/*
	Given a planner, creates a function which will evaluate a specific precedence level of operators,
	and link it to other `precedent`s which recurse to parse other precedence levels.
*/
func makePrecedentFromPlanner(planner *precedencePlanner) precedent {

	var generated precedent
	var nextRight precedent

	generated = func(stream *tokenStream) (*evaluationStage, error) {
		return planPrecedenceLevel(
			stream,
			planner.typeErrorFormat,
			planner.validSymbols,
			planner.validKinds,
			nextRight,
			planner.next,
		)
	}

	if planner.nextRight != nil {
		nextRight = planner.nextRight
	} else {
		nextRight = generated
	}

	return generated
}

/*
	Creates a `evaluationStageList` object which represents an execution plan (or tree)
	which is used to completely evaluate a set of tokens at evaluation-time.
	The three stages of evaluation can be thought of as parsing strings to tokens, then tokens to a stage list, then evaluation with parameters.
*/
func planStages(tokens []ExpressionToken) (*evaluationStage, error) {

	stream := newTokenStream(tokens)

	stage, err := planTokens(stream)
	if err != nil {
		return nil, err
	}

	// while we're now fully-planned, we now need to re-order same-precedence operators.
	// this could probably be avoided with a different planning method
	reorderStages(stage)
	return stage, nil
}

func planTokens(stream *tokenStream) (*evaluationStage, error) {

	if !stream.hasNext() {
		return nil, nil
	}

	return planSeparator(stream)
}

/*
	The most usual method of parsing an evaluation stage for a given precedence.
	Most stages use the same logic
*/
func planPrecedenceLevel(
	stream *tokenStream,
	typeErrorFormat string,
	validSymbols map[string]OperatorSymbol,
	validKinds []TokenKind,
	rightPrecedent precedent,
	leftPrecedent precedent) (*evaluationStage, error) {

	var token ExpressionToken
	var symbol OperatorSymbol
	var leftStage, rightStage *evaluationStage
	var checks typeChecks
	var err error
	var keyFound bool

	if leftPrecedent != nil {

		leftStage, err = leftPrecedent(stream)
		if err != nil {
			return nil, err
		}
	}

	for stream.hasNext() {

		token = stream.next()

		if len(validKinds) > 0 {

			keyFound = false
			for _, kind := range validKinds {
				if kind == token.Kind {
					keyFound = true
					break
				}
			}

			if !keyFound {
				break
			}
		}

		if validSymbols != nil {

			if !isString(token.Value) {
				break
			}

			symbol, keyFound = validSymbols[token.Value.(string)]
			if !keyFound {
				break
			}
		}

		if rightPrecedent != nil {
			rightStage, err = rightPrecedent(stream)
			if err != nil {
				return nil, err
			}
		}

		checks = findTypeChecks(symbol)

		return &evaluationStage{

			symbol:     symbol,
			leftStage:  leftStage,
			rightStage: rightStage,
			operator:   stageSymbolMap[symbol],

			leftTypeCheck:   checks.left,
			rightTypeCheck:  checks.right,
			typeCheck:       checks.combined,
			typeErrorFormat: typeErrorFormat,
		}, nil
	}

	stream.rewind()
	return leftStage, nil
}

/*
	A special case where functions need to be of higher precedence than values, and need a special wrapped execution stage operator.
*/
func planFunction(stream *tokenStream) (*evaluationStage, error) {

	var token ExpressionToken
	var rightStage *evaluationStage
	var err error

	token = stream.next()

	if token.Kind != FUNCTION {
		stream.rewind()
		return planValue(stream)
	}

	rightStage, err = planValue(stream)
	if err != nil {
		return nil, err
	}

	return &evaluationStage{

		symbol:          FUNCTIONAL,
		rightStage:      rightStage,
		operator:        makeFunctionStage(token.Value.(ExpressionFunction)),
		typeErrorFormat: "Unable to run function '%v': %v",
	}, nil
}

/*
	A truly special precedence function, this handles all the "lowest-case" errata of the process, including literals, parmeters,
	clauses, and prefixes.
*/
func planValue(stream *tokenStream) (*evaluationStage, error) {

	var token ExpressionToken
	var ret *evaluationStage
	var operator evaluationOperator
	var err error

	token = stream.next()

	switch token.Kind {

	case CLAUSE:

		ret, err = planTokens(stream)
		if err != nil {
			return nil, err
		}

		// advance past the CLAUSE_CLOSE token. We know that it's a CLAUSE_CLOSE, because at parse-time we check for unbalanced parens.
		stream.next()
		return ret, nil

	case CLAUSE_CLOSE:

		// when functions have empty params, this will be hit. In this case, we don't have any evaluation stage to do,
		// so we just return nil so that the stage planner continues on its way.
		stream.rewind()
		return nil, nil

	case VARIABLE:
		operator = makeParameterStage(token.Value.(string))

	case NUMERIC:
		fallthrough
	case STRING:
		fallthrough
	case PATTERN:
		fallthrough
	case BOOLEAN:
		operator = makeLiteralStage(token.Value)
	case TIME:
		operator = makeLiteralStage(float64(token.Value.(time.Time).Unix()))

	case PREFIX:
		stream.rewind()
		return planPrefix(stream)
	}

	if operator == nil {
		errorMsg := fmt.Sprintf("Unable to plan token kind: '%s', value: '%v'", GetTokenKindString(token.Kind), token.Value)
		return nil, errors.New(errorMsg)
	}

	return &evaluationStage{
		operator: operator,
	}, nil
}

/*
	Convenience function to pass a triplet of typechecks between `findTypeChecks` and `planPrecedenceLevel`.
	Each of these members may be nil, which indicates that type does not matter for that value.
*/
type typeChecks struct {
	left     stageTypeCheck
	right    stageTypeCheck
	combined stageCombinedTypeCheck
}

/*
	Maps a given [symbol] to a set of typechecks to be used during runtime.
*/
func findTypeChecks(symbol OperatorSymbol) typeChecks {

	switch symbol {
	case GT:
		fallthrough
	case LT:
		fallthrough
	case GTE:
		fallthrough
	case LTE:
		return typeChecks{
			combined: comparatorTypeCheck,
		}
	case REQ:
		fallthrough
	case NREQ:
		return typeChecks{
			left:  isString,
			right: isRegexOrString,
		}
	case AND:
		fallthrough
	case OR:
		return typeChecks{
			left:  isBool,
			right: isBool,
		}
	case IN:
		return typeChecks{
			right: isArray,
		}
	case BITWISE_LSHIFT:
		fallthrough
	case BITWISE_RSHIFT:
		fallthrough
	case BITWISE_OR:
		fallthrough
	case BITWISE_AND:
		fallthrough
	case BITWISE_XOR:
		return typeChecks{
			left:  isFloat64,
			right: isFloat64,
		}
	case PLUS:
		return typeChecks{
			combined: additionTypeCheck,
		}
	case MINUS:
		fallthrough
	case MULTIPLY:
		fallthrough
	case DIVIDE:
		fallthrough
	case MODULUS:
		fallthrough
	case EXPONENT:
		return typeChecks{
			left:  isFloat64,
			right: isFloat64,
		}
	case NEGATE:
		return typeChecks{
			right: isFloat64,
		}
	case INVERT:
		return typeChecks{
			right: isBool,
		}
	case BITWISE_NOT:
		return typeChecks{
			right: isFloat64,
		}
	case TERNARY_TRUE:
		return typeChecks{
			left: isBool,
		}

	// unchecked cases
	case EQ:
		fallthrough
	case NEQ:
		return typeChecks{}
	case TERNARY_FALSE:
		fallthrough
	case COALESCE:
		fallthrough
	default:
		return typeChecks{}
	}
}

/*
	During stage planning, stages of equal precedence are parsed such that they'll be evaluated in reverse order.
	For commutative operators like "+" or "-", it's no big deal. But for order-specific operators, it ruins the expected result.
*/
func reorderStages(rootStage *evaluationStage) {

	// traverse every rightStage until we find multiples in a row of the same precedence.
	var identicalPrecedences []*evaluationStage
	var currentStage, nextStage, lastStage *evaluationStage
	var precedence, currentPrecedence OperatorPrecedence

	nextStage = rootStage
	precedence = findOperatorPrecedenceForSymbol(rootStage.symbol)

	for nextStage != nil {

		lastStage = currentStage
		currentStage = nextStage
		nextStage = currentStage.rightStage

		currentPrecedence = findOperatorPrecedenceForSymbol(currentStage.symbol)

		if currentPrecedence == COMPARATOR_PRECEDENCE || currentPrecedence == LOGICAL_PRECEDENCE {
			continue
		}

		if currentPrecedence == precedence {
			identicalPrecedences = append(identicalPrecedences, currentStage)
			continue
		}

		// precedence break.
		// See how many in a row we had, and reorder if there's more than one.
		if len(identicalPrecedences) > 1 {
			mirrorStageSubtree(identicalPrecedences)
		} else {
			if lastStage.leftStage != nil {
				reorderStages(lastStage.leftStage)
			}
		}

		identicalPrecedences = []*evaluationStage{currentStage}
		precedence = currentPrecedence
	}

	if len(identicalPrecedences) > 1 {
		mirrorStageSubtree(identicalPrecedences)
	}
}

/*
	Performs a "mirror" on a subtree of stages.
	This mirror functionally inverts the order of execution for all members of the [stages] list.
	That list is assumed to be a root-to-leaf (ordered) list of evaluation stages, where each is a right-hand stage of the last.
*/
func mirrorStageSubtree(stages []*evaluationStage) {

	var rootStage, inverseStage, carryStage, frontStage *evaluationStage

	stagesLength := len(stages)

	// reverse all right/left
	for _, frontStage = range stages {

		carryStage = frontStage.rightStage
		frontStage.rightStage = frontStage.leftStage
		frontStage.leftStage = carryStage
	}

	// end left swaps with root right
	rootStage = stages[0]
	frontStage = stages[stagesLength-1]

	carryStage = frontStage.leftStage
	frontStage.leftStage = rootStage.rightStage
	rootStage.rightStage = carryStage

	// for all non-root non-end stages, right is swapped with inverse stage right in list
	for i := 0; i < (stagesLength-2)/2+1; i++ {

		frontStage = stages[i+1]
		inverseStage = stages[stagesLength-i-1]

		carryStage = frontStage.rightStage
		frontStage.rightStage = inverseStage.rightStage
		inverseStage.rightStage = carryStage
	}

	// swap all other information with inverse stages
	for i := 0; i < stagesLength/2; i++ {

		frontStage = stages[i]
		inverseStage = stages[stagesLength-i-1]
		frontStage.swapWith(inverseStage)
	}
}