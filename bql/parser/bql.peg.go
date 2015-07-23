package parser

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const end_symbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleStatements
	ruleStatement
	ruleSourceStmt
	ruleSinkStmt
	ruleStateStmt
	ruleStreamStmt
	ruleSelectStmt
	ruleCreateStreamAsSelectStmt
	ruleCreateSourceStmt
	ruleCreateSinkStmt
	ruleCreateStateStmt
	ruleUpdateStateStmt
	ruleUpdateSourceStmt
	ruleUpdateSinkStmt
	ruleInsertIntoSelectStmt
	ruleInsertIntoFromStmt
	rulePauseSourceStmt
	ruleResumeSourceStmt
	ruleRewindSourceStmt
	ruleDropSourceStmt
	ruleDropStreamStmt
	ruleDropSinkStmt
	ruleDropStateStmt
	ruleEmitter
	ruleProjections
	ruleProjection
	ruleAliasExpression
	ruleWindowedFrom
	ruleInterval
	ruleTimeInterval
	ruleTuplesInterval
	ruleRelations
	ruleFilter
	ruleGrouping
	ruleGroupList
	ruleHaving
	ruleRelationLike
	ruleAliasedStreamWindow
	ruleStreamWindow
	ruleStreamLike
	ruleUDSFFuncApp
	ruleSourceSinkSpecs
	ruleUpdateSourceSinkSpecs
	ruleSourceSinkParam
	ruleSourceSinkParamVal
	rulePausedOpt
	ruleExpression
	ruleorExpr
	ruleandExpr
	rulenotExpr
	rulecomparisonExpr
	ruleotherOpExpr
	ruleisExpr
	ruletermExpr
	ruleproductExpr
	ruleminusExpr
	rulecastExpr
	rulebaseExpr
	ruleFuncTypeCast
	ruleFuncApp
	ruleFuncParams
	ruleLiteral
	ruleComparisonOp
	ruleOtherOp
	ruleIsOp
	rulePlusMinusOp
	ruleMultDivOp
	ruleStream
	ruleRowMeta
	ruleRowTimestamp
	ruleRowValue
	ruleNumericLiteral
	ruleFloatLiteral
	ruleFunction
	ruleNullLiteral
	ruleBooleanLiteral
	ruleTRUE
	ruleFALSE
	ruleWildcard
	ruleStringLiteral
	ruleISTREAM
	ruleDSTREAM
	ruleRSTREAM
	ruleTUPLES
	ruleSECONDS
	ruleStreamIdentifier
	ruleSourceSinkType
	ruleSourceSinkParamKey
	rulePaused
	ruleUnpaused
	ruleType
	ruleBool
	ruleInt
	ruleFloat
	ruleString
	ruleBlob
	ruleTimestamp
	ruleArray
	ruleMap
	ruleOr
	ruleAnd
	ruleNot
	ruleEqual
	ruleLess
	ruleLessOrEqual
	ruleGreater
	ruleGreaterOrEqual
	ruleNotEqual
	ruleConcat
	ruleIs
	ruleIsNot
	rulePlus
	ruleMinus
	ruleMultiply
	ruleDivide
	ruleModulo
	ruleUnaryMinus
	ruleIdentifier
	ruleTargetIdentifier
	ruleident
	rulejsonPath
	rulesp
	rulecomment
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6
	ruleAction7
	ruleAction8
	ruleAction9
	ruleAction10
	ruleAction11
	ruleAction12
	ruleAction13
	ruleAction14
	ruleAction15
	ruleAction16
	ruleAction17
	rulePegText
	ruleAction18
	ruleAction19
	ruleAction20
	ruleAction21
	ruleAction22
	ruleAction23
	ruleAction24
	ruleAction25
	ruleAction26
	ruleAction27
	ruleAction28
	ruleAction29
	ruleAction30
	ruleAction31
	ruleAction32
	ruleAction33
	ruleAction34
	ruleAction35
	ruleAction36
	ruleAction37
	ruleAction38
	ruleAction39
	ruleAction40
	ruleAction41
	ruleAction42
	ruleAction43
	ruleAction44
	ruleAction45
	ruleAction46
	ruleAction47
	ruleAction48
	ruleAction49
	ruleAction50
	ruleAction51
	ruleAction52
	ruleAction53
	ruleAction54
	ruleAction55
	ruleAction56
	ruleAction57
	ruleAction58
	ruleAction59
	ruleAction60
	ruleAction61
	ruleAction62
	ruleAction63
	ruleAction64
	ruleAction65
	ruleAction66
	ruleAction67
	ruleAction68
	ruleAction69
	ruleAction70
	ruleAction71
	ruleAction72
	ruleAction73
	ruleAction74
	ruleAction75
	ruleAction76
	ruleAction77
	ruleAction78
	ruleAction79
	ruleAction80
	ruleAction81
	ruleAction82
	ruleAction83
	ruleAction84
	ruleAction85
	ruleAction86
	ruleAction87
	ruleAction88
	ruleAction89
	ruleAction90
	ruleAction91
	ruleAction92
	ruleAction93
	ruleAction94
	ruleAction95

	rulePre_
	rule_In_
	rule_Suf
)

var rul3s = [...]string{
	"Unknown",
	"Statements",
	"Statement",
	"SourceStmt",
	"SinkStmt",
	"StateStmt",
	"StreamStmt",
	"SelectStmt",
	"CreateStreamAsSelectStmt",
	"CreateSourceStmt",
	"CreateSinkStmt",
	"CreateStateStmt",
	"UpdateStateStmt",
	"UpdateSourceStmt",
	"UpdateSinkStmt",
	"InsertIntoSelectStmt",
	"InsertIntoFromStmt",
	"PauseSourceStmt",
	"ResumeSourceStmt",
	"RewindSourceStmt",
	"DropSourceStmt",
	"DropStreamStmt",
	"DropSinkStmt",
	"DropStateStmt",
	"Emitter",
	"Projections",
	"Projection",
	"AliasExpression",
	"WindowedFrom",
	"Interval",
	"TimeInterval",
	"TuplesInterval",
	"Relations",
	"Filter",
	"Grouping",
	"GroupList",
	"Having",
	"RelationLike",
	"AliasedStreamWindow",
	"StreamWindow",
	"StreamLike",
	"UDSFFuncApp",
	"SourceSinkSpecs",
	"UpdateSourceSinkSpecs",
	"SourceSinkParam",
	"SourceSinkParamVal",
	"PausedOpt",
	"Expression",
	"orExpr",
	"andExpr",
	"notExpr",
	"comparisonExpr",
	"otherOpExpr",
	"isExpr",
	"termExpr",
	"productExpr",
	"minusExpr",
	"castExpr",
	"baseExpr",
	"FuncTypeCast",
	"FuncApp",
	"FuncParams",
	"Literal",
	"ComparisonOp",
	"OtherOp",
	"IsOp",
	"PlusMinusOp",
	"MultDivOp",
	"Stream",
	"RowMeta",
	"RowTimestamp",
	"RowValue",
	"NumericLiteral",
	"FloatLiteral",
	"Function",
	"NullLiteral",
	"BooleanLiteral",
	"TRUE",
	"FALSE",
	"Wildcard",
	"StringLiteral",
	"ISTREAM",
	"DSTREAM",
	"RSTREAM",
	"TUPLES",
	"SECONDS",
	"StreamIdentifier",
	"SourceSinkType",
	"SourceSinkParamKey",
	"Paused",
	"Unpaused",
	"Type",
	"Bool",
	"Int",
	"Float",
	"String",
	"Blob",
	"Timestamp",
	"Array",
	"Map",
	"Or",
	"And",
	"Not",
	"Equal",
	"Less",
	"LessOrEqual",
	"Greater",
	"GreaterOrEqual",
	"NotEqual",
	"Concat",
	"Is",
	"IsNot",
	"Plus",
	"Minus",
	"Multiply",
	"Divide",
	"Modulo",
	"UnaryMinus",
	"Identifier",
	"TargetIdentifier",
	"ident",
	"jsonPath",
	"sp",
	"comment",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
	"Action6",
	"Action7",
	"Action8",
	"Action9",
	"Action10",
	"Action11",
	"Action12",
	"Action13",
	"Action14",
	"Action15",
	"Action16",
	"Action17",
	"PegText",
	"Action18",
	"Action19",
	"Action20",
	"Action21",
	"Action22",
	"Action23",
	"Action24",
	"Action25",
	"Action26",
	"Action27",
	"Action28",
	"Action29",
	"Action30",
	"Action31",
	"Action32",
	"Action33",
	"Action34",
	"Action35",
	"Action36",
	"Action37",
	"Action38",
	"Action39",
	"Action40",
	"Action41",
	"Action42",
	"Action43",
	"Action44",
	"Action45",
	"Action46",
	"Action47",
	"Action48",
	"Action49",
	"Action50",
	"Action51",
	"Action52",
	"Action53",
	"Action54",
	"Action55",
	"Action56",
	"Action57",
	"Action58",
	"Action59",
	"Action60",
	"Action61",
	"Action62",
	"Action63",
	"Action64",
	"Action65",
	"Action66",
	"Action67",
	"Action68",
	"Action69",
	"Action70",
	"Action71",
	"Action72",
	"Action73",
	"Action74",
	"Action75",
	"Action76",
	"Action77",
	"Action78",
	"Action79",
	"Action80",
	"Action81",
	"Action82",
	"Action83",
	"Action84",
	"Action85",
	"Action86",
	"Action87",
	"Action88",
	"Action89",
	"Action90",
	"Action91",
	"Action92",
	"Action93",
	"Action94",
	"Action95",

	"Pre_",
	"_In_",
	"_Suf",
}

type tokenTree interface {
	Print()
	PrintSyntax()
	PrintSyntaxTree(buffer string)
	Add(rule pegRule, begin, end, next uint32, depth int)
	Expand(index int) tokenTree
	Tokens() <-chan token32
	AST() *node32
	Error() []token32
	trim(length int)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(depth int, buffer string) {
	for node != nil {
		for c := 0; c < depth; c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[node.pegRule], strconv.Quote(string(([]rune(buffer)[node.begin:node.end]))))
		if node.up != nil {
			node.up.print(depth+1, buffer)
		}
		node = node.next
	}
}

func (ast *node32) Print(buffer string) {
	ast.print(0, buffer)
}

type element struct {
	node *node32
	down *element
}

/* ${@} bit structure for abstract syntax tree */
type token32 struct {
	pegRule
	begin, end, next uint32
}

func (t *token32) isZero() bool {
	return t.pegRule == ruleUnknown && t.begin == 0 && t.end == 0 && t.next == 0
}

func (t *token32) isParentOf(u token32) bool {
	return t.begin <= u.begin && t.end >= u.end && t.next > u.next
}

func (t *token32) getToken32() token32 {
	return token32{pegRule: t.pegRule, begin: uint32(t.begin), end: uint32(t.end), next: uint32(t.next)}
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v %v", rul3s[t.pegRule], t.begin, t.end, t.next)
}

type tokens32 struct {
	tree    []token32
	ordered [][]token32
}

func (t *tokens32) trim(length int) {
	t.tree = t.tree[0:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) Order() [][]token32 {
	if t.ordered != nil {
		return t.ordered
	}

	depths := make([]int32, 1, math.MaxInt16)
	for i, token := range t.tree {
		if token.pegRule == ruleUnknown {
			t.tree = t.tree[:i]
			break
		}
		depth := int(token.next)
		if length := len(depths); depth >= length {
			depths = depths[:depth+1]
		}
		depths[depth]++
	}
	depths = append(depths, 0)

	ordered, pool := make([][]token32, len(depths)), make([]token32, len(t.tree)+len(depths))
	for i, depth := range depths {
		depth++
		ordered[i], pool, depths[i] = pool[:depth], pool[depth:], 0
	}

	for i, token := range t.tree {
		depth := token.next
		token.next = uint32(i)
		ordered[depth][depths[depth]] = token
		depths[depth]++
	}
	t.ordered = ordered
	return ordered
}

type state32 struct {
	token32
	depths []int32
	leaf   bool
}

func (t *tokens32) AST() *node32 {
	tokens := t.Tokens()
	stack := &element{node: &node32{token32: <-tokens}}
	for token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	return stack.node
}

func (t *tokens32) PreOrder() (<-chan state32, [][]token32) {
	s, ordered := make(chan state32, 6), t.Order()
	go func() {
		var states [8]state32
		for i, _ := range states {
			states[i].depths = make([]int32, len(ordered))
		}
		depths, state, depth := make([]int32, len(ordered)), 0, 1
		write := func(t token32, leaf bool) {
			S := states[state]
			state, S.pegRule, S.begin, S.end, S.next, S.leaf = (state+1)%8, t.pegRule, t.begin, t.end, uint32(depth), leaf
			copy(S.depths, depths)
			s <- S
		}

		states[state].token32 = ordered[0][0]
		depths[0]++
		state++
		a, b := ordered[depth-1][depths[depth-1]-1], ordered[depth][depths[depth]]
	depthFirstSearch:
		for {
			for {
				if i := depths[depth]; i > 0 {
					if c, j := ordered[depth][i-1], depths[depth-1]; a.isParentOf(c) &&
						(j < 2 || !ordered[depth-1][j-2].isParentOf(c)) {
						if c.end != b.begin {
							write(token32{pegRule: rule_In_, begin: c.end, end: b.begin}, true)
						}
						break
					}
				}

				if a.begin < b.begin {
					write(token32{pegRule: rulePre_, begin: a.begin, end: b.begin}, true)
				}
				break
			}

			next := depth + 1
			if c := ordered[next][depths[next]]; c.pegRule != ruleUnknown && b.isParentOf(c) {
				write(b, false)
				depths[depth]++
				depth, a, b = next, b, c
				continue
			}

			write(b, true)
			depths[depth]++
			c, parent := ordered[depth][depths[depth]], true
			for {
				if c.pegRule != ruleUnknown && a.isParentOf(c) {
					b = c
					continue depthFirstSearch
				} else if parent && b.end != a.end {
					write(token32{pegRule: rule_Suf, begin: b.end, end: a.end}, true)
				}

				depth--
				if depth > 0 {
					a, b, c = ordered[depth-1][depths[depth-1]-1], a, ordered[depth][depths[depth]]
					parent = a.isParentOf(b)
					continue
				}

				break depthFirstSearch
			}
		}

		close(s)
	}()
	return s, ordered
}

func (t *tokens32) PrintSyntax() {
	tokens, ordered := t.PreOrder()
	max := -1
	for token := range tokens {
		if !token.leaf {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[36m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[36m%v\x1B[m\n", rul3s[token.pegRule])
		} else if token.begin == token.end {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[31m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[31m%v\x1B[m\n", rul3s[token.pegRule])
		} else {
			for c, end := token.begin, token.end; c < end; c++ {
				if i := int(c); max+1 < i {
					for j := max; j < i; j++ {
						fmt.Printf("skip %v %v\n", j, token.String())
					}
					max = i
				} else if i := int(c); i <= max {
					for j := i; j <= max; j++ {
						fmt.Printf("dupe %v %v\n", j, token.String())
					}
				} else {
					max = int(c)
				}
				fmt.Printf("%v", c)
				for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
					fmt.Printf(" \x1B[34m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
				}
				fmt.Printf(" \x1B[34m%v\x1B[m\n", rul3s[token.pegRule])
			}
			fmt.Printf("\n")
		}
	}
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	tokens, _ := t.PreOrder()
	for token := range tokens {
		for c := 0; c < int(token.next); c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[token.pegRule], strconv.Quote(string(([]rune(buffer)[token.begin:token.end]))))
	}
}

func (t *tokens32) Add(rule pegRule, begin, end, depth uint32, index int) {
	t.tree[index] = token32{pegRule: rule, begin: uint32(begin), end: uint32(end), next: uint32(depth)}
}

func (t *tokens32) Tokens() <-chan token32 {
	s := make(chan token32, 16)
	go func() {
		for _, v := range t.tree {
			s <- v.getToken32()
		}
		close(s)
	}()
	return s
}

func (t *tokens32) Error() []token32 {
	ordered := t.Order()
	length := len(ordered)
	tokens, length := make([]token32, length), length-1
	for i, _ := range tokens {
		o := ordered[length-i]
		if len(o) > 1 {
			tokens[i] = o[len(o)-2].getToken32()
		}
	}
	return tokens
}

/*func (t *tokens16) Expand(index int) tokenTree {
	tree := t.tree
	if index >= len(tree) {
		expanded := make([]token32, 2 * len(tree))
		for i, v := range tree {
			expanded[i] = v.getToken32()
		}
		return &tokens32{tree: expanded}
	}
	return nil
}*/

func (t *tokens32) Expand(index int) tokenTree {
	tree := t.tree
	if index >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	return nil
}

type bqlPeg struct {
	parseStack

	Buffer string
	buffer []rune
	rules  [221]func() bool
	Parse  func(rule ...int) error
	Reset  func()
	tokenTree
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer string, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer[0:] {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p *bqlPeg
}

func (e *parseError) Error() string {
	tokens, error := e.p.tokenTree.Error(), "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.Buffer, positions)
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf("parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n",
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			/*strconv.Quote(*/ e.p.Buffer[begin:end] /*)*/)
	}

	return error
}

func (p *bqlPeg) PrintSyntaxTree() {
	p.tokenTree.PrintSyntaxTree(p.Buffer)
}

func (p *bqlPeg) Highlighter() {
	p.tokenTree.PrintSyntax()
}

func (p *bqlPeg) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for token := range p.tokenTree.Tokens() {
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:

			p.AssembleSelect()

		case ruleAction1:

			p.AssembleCreateStreamAsSelect()

		case ruleAction2:

			p.AssembleCreateSource()

		case ruleAction3:

			p.AssembleCreateSink()

		case ruleAction4:

			p.AssembleCreateState()

		case ruleAction5:

			p.AssembleUpdateState()

		case ruleAction6:

			p.AssembleUpdateSource()

		case ruleAction7:

			p.AssembleUpdateSink()

		case ruleAction8:

			p.AssembleInsertIntoSelect()

		case ruleAction9:

			p.AssembleInsertIntoFrom()

		case ruleAction10:

			p.AssemblePauseSource()

		case ruleAction11:

			p.AssembleResumeSource()

		case ruleAction12:

			p.AssembleRewindSource()

		case ruleAction13:

			p.AssembleDropSource()

		case ruleAction14:

			p.AssembleDropStream()

		case ruleAction15:

			p.AssembleDropSink()

		case ruleAction16:

			p.AssembleDropState()

		case ruleAction17:

			p.AssembleEmitter()

		case ruleAction18:

			p.AssembleProjections(begin, end)

		case ruleAction19:

			p.AssembleAlias()

		case ruleAction20:

			// This is *always* executed, even if there is no
			// FROM clause present in the statement.
			p.AssembleWindowedFrom(begin, end)

		case ruleAction21:

			p.AssembleInterval()

		case ruleAction22:

			p.AssembleInterval()

		case ruleAction23:

			// This is *always* executed, even if there is no
			// WHERE clause present in the statement.
			p.AssembleFilter(begin, end)

		case ruleAction24:

			// This is *always* executed, even if there is no
			// GROUP BY clause present in the statement.
			p.AssembleGrouping(begin, end)

		case ruleAction25:

			// This is *always* executed, even if there is no
			// HAVING clause present in the statement.
			p.AssembleHaving(begin, end)

		case ruleAction26:

			p.EnsureAliasedStreamWindow()

		case ruleAction27:

			p.AssembleAliasedStreamWindow()

		case ruleAction28:

			p.AssembleStreamWindow()

		case ruleAction29:

			p.AssembleUDSFFuncApp()

		case ruleAction30:

			p.AssembleSourceSinkSpecs(begin, end)

		case ruleAction31:

			p.AssembleSourceSinkSpecs(begin, end)

		case ruleAction32:

			p.AssembleSourceSinkParam()

		case ruleAction33:

			p.EnsureKeywordPresent(begin, end)

		case ruleAction34:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction35:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction36:

			p.AssembleUnaryPrefixOperation(begin, end)

		case ruleAction37:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction38:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction39:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction40:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction41:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction42:

			p.AssembleUnaryPrefixOperation(begin, end)

		case ruleAction43:

			p.AssembleTypeCast(begin, end)

		case ruleAction44:

			p.AssembleTypeCast(begin, end)

		case ruleAction45:

			p.AssembleFuncApp()

		case ruleAction46:

			p.AssembleExpressions(begin, end)

		case ruleAction47:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewStream(substr))

		case ruleAction48:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewRowMeta(substr, TimestampMeta))

		case ruleAction49:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewRowValue(substr))

		case ruleAction50:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewNumericLiteral(substr))

		case ruleAction51:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewFloatLiteral(substr))

		case ruleAction52:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, FuncName(substr))

		case ruleAction53:

			p.PushComponent(begin, end, NewNullLiteral())

		case ruleAction54:

			p.PushComponent(begin, end, NewBoolLiteral(true))

		case ruleAction55:

			p.PushComponent(begin, end, NewBoolLiteral(false))

		case ruleAction56:

			p.PushComponent(begin, end, NewWildcard())

		case ruleAction57:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewStringLiteral(substr))

		case ruleAction58:

			p.PushComponent(begin, end, Istream)

		case ruleAction59:

			p.PushComponent(begin, end, Dstream)

		case ruleAction60:

			p.PushComponent(begin, end, Rstream)

		case ruleAction61:

			p.PushComponent(begin, end, Tuples)

		case ruleAction62:

			p.PushComponent(begin, end, Seconds)

		case ruleAction63:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, StreamIdentifier(substr))

		case ruleAction64:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, SourceSinkType(substr))

		case ruleAction65:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, SourceSinkParamKey(substr))

		case ruleAction66:

			p.PushComponent(begin, end, Yes)

		case ruleAction67:

			p.PushComponent(begin, end, No)

		case ruleAction68:

			p.PushComponent(begin, end, Bool)

		case ruleAction69:

			p.PushComponent(begin, end, Int)

		case ruleAction70:

			p.PushComponent(begin, end, Float)

		case ruleAction71:

			p.PushComponent(begin, end, String)

		case ruleAction72:

			p.PushComponent(begin, end, Blob)

		case ruleAction73:

			p.PushComponent(begin, end, Timestamp)

		case ruleAction74:

			p.PushComponent(begin, end, Array)

		case ruleAction75:

			p.PushComponent(begin, end, Map)

		case ruleAction76:

			p.PushComponent(begin, end, Or)

		case ruleAction77:

			p.PushComponent(begin, end, And)

		case ruleAction78:

			p.PushComponent(begin, end, Not)

		case ruleAction79:

			p.PushComponent(begin, end, Equal)

		case ruleAction80:

			p.PushComponent(begin, end, Less)

		case ruleAction81:

			p.PushComponent(begin, end, LessOrEqual)

		case ruleAction82:

			p.PushComponent(begin, end, Greater)

		case ruleAction83:

			p.PushComponent(begin, end, GreaterOrEqual)

		case ruleAction84:

			p.PushComponent(begin, end, NotEqual)

		case ruleAction85:

			p.PushComponent(begin, end, Concat)

		case ruleAction86:

			p.PushComponent(begin, end, Is)

		case ruleAction87:

			p.PushComponent(begin, end, IsNot)

		case ruleAction88:

			p.PushComponent(begin, end, Plus)

		case ruleAction89:

			p.PushComponent(begin, end, Minus)

		case ruleAction90:

			p.PushComponent(begin, end, Multiply)

		case ruleAction91:

			p.PushComponent(begin, end, Divide)

		case ruleAction92:

			p.PushComponent(begin, end, Modulo)

		case ruleAction93:

			p.PushComponent(begin, end, UnaryMinus)

		case ruleAction94:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, Identifier(substr))

		case ruleAction95:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, Identifier(substr))

		}
	}
	_, _, _, _ = buffer, text, begin, end
}

func (p *bqlPeg) Init() {
	p.buffer = []rune(p.Buffer)
	if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != end_symbol {
		p.buffer = append(p.buffer, end_symbol)
	}

	var tree tokenTree = &tokens32{tree: make([]token32, math.MaxInt16)}
	position, depth, tokenIndex, buffer, _rules := uint32(0), uint32(0), 0, p.buffer, p.rules

	p.Parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokenTree = tree
		if matches {
			p.tokenTree.trim(tokenIndex)
			return nil
		}
		return &parseError{p}
	}

	p.Reset = func() {
		position, tokenIndex, depth = 0, 0, 0
	}

	add := func(rule pegRule, begin uint32) {
		if t := tree.Expand(tokenIndex); t != nil {
			tree = t
		}
		tree.Add(rule, begin, position, depth, tokenIndex)
		tokenIndex++
	}

	matchDot := func() bool {
		if buffer[position] != end_symbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 Statements <- <(sp ((Statement sp ';' .*) / Statement) !.)> */
		func() bool {
			position0, tokenIndex0, depth0 := position, tokenIndex, depth
			{
				position1 := position
				depth++
				if !_rules[rulesp]() {
					goto l0
				}
				{
					position2, tokenIndex2, depth2 := position, tokenIndex, depth
					if !_rules[ruleStatement]() {
						goto l3
					}
					if !_rules[rulesp]() {
						goto l3
					}
					if buffer[position] != rune(';') {
						goto l3
					}
					position++
				l4:
					{
						position5, tokenIndex5, depth5 := position, tokenIndex, depth
						if !matchDot() {
							goto l5
						}
						goto l4
					l5:
						position, tokenIndex, depth = position5, tokenIndex5, depth5
					}
					goto l2
				l3:
					position, tokenIndex, depth = position2, tokenIndex2, depth2
					if !_rules[ruleStatement]() {
						goto l0
					}
				}
			l2:
				{
					position6, tokenIndex6, depth6 := position, tokenIndex, depth
					if !matchDot() {
						goto l6
					}
					goto l0
				l6:
					position, tokenIndex, depth = position6, tokenIndex6, depth6
				}
				depth--
				add(ruleStatements, position1)
			}
			return true
		l0:
			position, tokenIndex, depth = position0, tokenIndex0, depth0
			return false
		},
		/* 1 Statement <- <(SelectStmt / SourceStmt / SinkStmt / StateStmt / StreamStmt)> */
		func() bool {
			position7, tokenIndex7, depth7 := position, tokenIndex, depth
			{
				position8 := position
				depth++
				{
					position9, tokenIndex9, depth9 := position, tokenIndex, depth
					if !_rules[ruleSelectStmt]() {
						goto l10
					}
					goto l9
				l10:
					position, tokenIndex, depth = position9, tokenIndex9, depth9
					if !_rules[ruleSourceStmt]() {
						goto l11
					}
					goto l9
				l11:
					position, tokenIndex, depth = position9, tokenIndex9, depth9
					if !_rules[ruleSinkStmt]() {
						goto l12
					}
					goto l9
				l12:
					position, tokenIndex, depth = position9, tokenIndex9, depth9
					if !_rules[ruleStateStmt]() {
						goto l13
					}
					goto l9
				l13:
					position, tokenIndex, depth = position9, tokenIndex9, depth9
					if !_rules[ruleStreamStmt]() {
						goto l7
					}
				}
			l9:
				depth--
				add(ruleStatement, position8)
			}
			return true
		l7:
			position, tokenIndex, depth = position7, tokenIndex7, depth7
			return false
		},
		/* 2 SourceStmt <- <(CreateSourceStmt / UpdateSourceStmt / DropSourceStmt / PauseSourceStmt / ResumeSourceStmt / RewindSourceStmt)> */
		func() bool {
			position14, tokenIndex14, depth14 := position, tokenIndex, depth
			{
				position15 := position
				depth++
				{
					position16, tokenIndex16, depth16 := position, tokenIndex, depth
					if !_rules[ruleCreateSourceStmt]() {
						goto l17
					}
					goto l16
				l17:
					position, tokenIndex, depth = position16, tokenIndex16, depth16
					if !_rules[ruleUpdateSourceStmt]() {
						goto l18
					}
					goto l16
				l18:
					position, tokenIndex, depth = position16, tokenIndex16, depth16
					if !_rules[ruleDropSourceStmt]() {
						goto l19
					}
					goto l16
				l19:
					position, tokenIndex, depth = position16, tokenIndex16, depth16
					if !_rules[rulePauseSourceStmt]() {
						goto l20
					}
					goto l16
				l20:
					position, tokenIndex, depth = position16, tokenIndex16, depth16
					if !_rules[ruleResumeSourceStmt]() {
						goto l21
					}
					goto l16
				l21:
					position, tokenIndex, depth = position16, tokenIndex16, depth16
					if !_rules[ruleRewindSourceStmt]() {
						goto l14
					}
				}
			l16:
				depth--
				add(ruleSourceStmt, position15)
			}
			return true
		l14:
			position, tokenIndex, depth = position14, tokenIndex14, depth14
			return false
		},
		/* 3 SinkStmt <- <(CreateSinkStmt / UpdateSinkStmt / DropSinkStmt)> */
		func() bool {
			position22, tokenIndex22, depth22 := position, tokenIndex, depth
			{
				position23 := position
				depth++
				{
					position24, tokenIndex24, depth24 := position, tokenIndex, depth
					if !_rules[ruleCreateSinkStmt]() {
						goto l25
					}
					goto l24
				l25:
					position, tokenIndex, depth = position24, tokenIndex24, depth24
					if !_rules[ruleUpdateSinkStmt]() {
						goto l26
					}
					goto l24
				l26:
					position, tokenIndex, depth = position24, tokenIndex24, depth24
					if !_rules[ruleDropSinkStmt]() {
						goto l22
					}
				}
			l24:
				depth--
				add(ruleSinkStmt, position23)
			}
			return true
		l22:
			position, tokenIndex, depth = position22, tokenIndex22, depth22
			return false
		},
		/* 4 StateStmt <- <(CreateStateStmt / UpdateStateStmt / DropStateStmt)> */
		func() bool {
			position27, tokenIndex27, depth27 := position, tokenIndex, depth
			{
				position28 := position
				depth++
				{
					position29, tokenIndex29, depth29 := position, tokenIndex, depth
					if !_rules[ruleCreateStateStmt]() {
						goto l30
					}
					goto l29
				l30:
					position, tokenIndex, depth = position29, tokenIndex29, depth29
					if !_rules[ruleUpdateStateStmt]() {
						goto l31
					}
					goto l29
				l31:
					position, tokenIndex, depth = position29, tokenIndex29, depth29
					if !_rules[ruleDropStateStmt]() {
						goto l27
					}
				}
			l29:
				depth--
				add(ruleStateStmt, position28)
			}
			return true
		l27:
			position, tokenIndex, depth = position27, tokenIndex27, depth27
			return false
		},
		/* 5 StreamStmt <- <(CreateStreamAsSelectStmt / DropStreamStmt / InsertIntoSelectStmt / InsertIntoFromStmt)> */
		func() bool {
			position32, tokenIndex32, depth32 := position, tokenIndex, depth
			{
				position33 := position
				depth++
				{
					position34, tokenIndex34, depth34 := position, tokenIndex, depth
					if !_rules[ruleCreateStreamAsSelectStmt]() {
						goto l35
					}
					goto l34
				l35:
					position, tokenIndex, depth = position34, tokenIndex34, depth34
					if !_rules[ruleDropStreamStmt]() {
						goto l36
					}
					goto l34
				l36:
					position, tokenIndex, depth = position34, tokenIndex34, depth34
					if !_rules[ruleInsertIntoSelectStmt]() {
						goto l37
					}
					goto l34
				l37:
					position, tokenIndex, depth = position34, tokenIndex34, depth34
					if !_rules[ruleInsertIntoFromStmt]() {
						goto l32
					}
				}
			l34:
				depth--
				add(ruleStreamStmt, position33)
			}
			return true
		l32:
			position, tokenIndex, depth = position32, tokenIndex32, depth32
			return false
		},
		/* 6 SelectStmt <- <(('s' / 'S') ('e' / 'E') ('l' / 'L') ('e' / 'E') ('c' / 'C') ('t' / 'T') sp Emitter sp Projections sp WindowedFrom sp Filter sp Grouping sp Having sp Action0)> */
		func() bool {
			position38, tokenIndex38, depth38 := position, tokenIndex, depth
			{
				position39 := position
				depth++
				{
					position40, tokenIndex40, depth40 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l41
					}
					position++
					goto l40
				l41:
					position, tokenIndex, depth = position40, tokenIndex40, depth40
					if buffer[position] != rune('S') {
						goto l38
					}
					position++
				}
			l40:
				{
					position42, tokenIndex42, depth42 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l43
					}
					position++
					goto l42
				l43:
					position, tokenIndex, depth = position42, tokenIndex42, depth42
					if buffer[position] != rune('E') {
						goto l38
					}
					position++
				}
			l42:
				{
					position44, tokenIndex44, depth44 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l45
					}
					position++
					goto l44
				l45:
					position, tokenIndex, depth = position44, tokenIndex44, depth44
					if buffer[position] != rune('L') {
						goto l38
					}
					position++
				}
			l44:
				{
					position46, tokenIndex46, depth46 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l47
					}
					position++
					goto l46
				l47:
					position, tokenIndex, depth = position46, tokenIndex46, depth46
					if buffer[position] != rune('E') {
						goto l38
					}
					position++
				}
			l46:
				{
					position48, tokenIndex48, depth48 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l49
					}
					position++
					goto l48
				l49:
					position, tokenIndex, depth = position48, tokenIndex48, depth48
					if buffer[position] != rune('C') {
						goto l38
					}
					position++
				}
			l48:
				{
					position50, tokenIndex50, depth50 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l51
					}
					position++
					goto l50
				l51:
					position, tokenIndex, depth = position50, tokenIndex50, depth50
					if buffer[position] != rune('T') {
						goto l38
					}
					position++
				}
			l50:
				if !_rules[rulesp]() {
					goto l38
				}
				if !_rules[ruleEmitter]() {
					goto l38
				}
				if !_rules[rulesp]() {
					goto l38
				}
				if !_rules[ruleProjections]() {
					goto l38
				}
				if !_rules[rulesp]() {
					goto l38
				}
				if !_rules[ruleWindowedFrom]() {
					goto l38
				}
				if !_rules[rulesp]() {
					goto l38
				}
				if !_rules[ruleFilter]() {
					goto l38
				}
				if !_rules[rulesp]() {
					goto l38
				}
				if !_rules[ruleGrouping]() {
					goto l38
				}
				if !_rules[rulesp]() {
					goto l38
				}
				if !_rules[ruleHaving]() {
					goto l38
				}
				if !_rules[rulesp]() {
					goto l38
				}
				if !_rules[ruleAction0]() {
					goto l38
				}
				depth--
				add(ruleSelectStmt, position39)
			}
			return true
		l38:
			position, tokenIndex, depth = position38, tokenIndex38, depth38
			return false
		},
		/* 7 CreateStreamAsSelectStmt <- <(('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M')) sp StreamIdentifier sp (('a' / 'A') ('s' / 'S')) sp (('s' / 'S') ('e' / 'E') ('l' / 'L') ('e' / 'E') ('c' / 'C') ('t' / 'T')) sp Emitter sp Projections sp WindowedFrom sp Filter sp Grouping sp Having sp Action1)> */
		func() bool {
			position52, tokenIndex52, depth52 := position, tokenIndex, depth
			{
				position53 := position
				depth++
				{
					position54, tokenIndex54, depth54 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l55
					}
					position++
					goto l54
				l55:
					position, tokenIndex, depth = position54, tokenIndex54, depth54
					if buffer[position] != rune('C') {
						goto l52
					}
					position++
				}
			l54:
				{
					position56, tokenIndex56, depth56 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l57
					}
					position++
					goto l56
				l57:
					position, tokenIndex, depth = position56, tokenIndex56, depth56
					if buffer[position] != rune('R') {
						goto l52
					}
					position++
				}
			l56:
				{
					position58, tokenIndex58, depth58 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l59
					}
					position++
					goto l58
				l59:
					position, tokenIndex, depth = position58, tokenIndex58, depth58
					if buffer[position] != rune('E') {
						goto l52
					}
					position++
				}
			l58:
				{
					position60, tokenIndex60, depth60 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l61
					}
					position++
					goto l60
				l61:
					position, tokenIndex, depth = position60, tokenIndex60, depth60
					if buffer[position] != rune('A') {
						goto l52
					}
					position++
				}
			l60:
				{
					position62, tokenIndex62, depth62 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l63
					}
					position++
					goto l62
				l63:
					position, tokenIndex, depth = position62, tokenIndex62, depth62
					if buffer[position] != rune('T') {
						goto l52
					}
					position++
				}
			l62:
				{
					position64, tokenIndex64, depth64 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l65
					}
					position++
					goto l64
				l65:
					position, tokenIndex, depth = position64, tokenIndex64, depth64
					if buffer[position] != rune('E') {
						goto l52
					}
					position++
				}
			l64:
				if !_rules[rulesp]() {
					goto l52
				}
				{
					position66, tokenIndex66, depth66 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l67
					}
					position++
					goto l66
				l67:
					position, tokenIndex, depth = position66, tokenIndex66, depth66
					if buffer[position] != rune('S') {
						goto l52
					}
					position++
				}
			l66:
				{
					position68, tokenIndex68, depth68 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l69
					}
					position++
					goto l68
				l69:
					position, tokenIndex, depth = position68, tokenIndex68, depth68
					if buffer[position] != rune('T') {
						goto l52
					}
					position++
				}
			l68:
				{
					position70, tokenIndex70, depth70 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l71
					}
					position++
					goto l70
				l71:
					position, tokenIndex, depth = position70, tokenIndex70, depth70
					if buffer[position] != rune('R') {
						goto l52
					}
					position++
				}
			l70:
				{
					position72, tokenIndex72, depth72 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l73
					}
					position++
					goto l72
				l73:
					position, tokenIndex, depth = position72, tokenIndex72, depth72
					if buffer[position] != rune('E') {
						goto l52
					}
					position++
				}
			l72:
				{
					position74, tokenIndex74, depth74 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l75
					}
					position++
					goto l74
				l75:
					position, tokenIndex, depth = position74, tokenIndex74, depth74
					if buffer[position] != rune('A') {
						goto l52
					}
					position++
				}
			l74:
				{
					position76, tokenIndex76, depth76 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l77
					}
					position++
					goto l76
				l77:
					position, tokenIndex, depth = position76, tokenIndex76, depth76
					if buffer[position] != rune('M') {
						goto l52
					}
					position++
				}
			l76:
				if !_rules[rulesp]() {
					goto l52
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l52
				}
				if !_rules[rulesp]() {
					goto l52
				}
				{
					position78, tokenIndex78, depth78 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l79
					}
					position++
					goto l78
				l79:
					position, tokenIndex, depth = position78, tokenIndex78, depth78
					if buffer[position] != rune('A') {
						goto l52
					}
					position++
				}
			l78:
				{
					position80, tokenIndex80, depth80 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l81
					}
					position++
					goto l80
				l81:
					position, tokenIndex, depth = position80, tokenIndex80, depth80
					if buffer[position] != rune('S') {
						goto l52
					}
					position++
				}
			l80:
				if !_rules[rulesp]() {
					goto l52
				}
				{
					position82, tokenIndex82, depth82 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l83
					}
					position++
					goto l82
				l83:
					position, tokenIndex, depth = position82, tokenIndex82, depth82
					if buffer[position] != rune('S') {
						goto l52
					}
					position++
				}
			l82:
				{
					position84, tokenIndex84, depth84 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l85
					}
					position++
					goto l84
				l85:
					position, tokenIndex, depth = position84, tokenIndex84, depth84
					if buffer[position] != rune('E') {
						goto l52
					}
					position++
				}
			l84:
				{
					position86, tokenIndex86, depth86 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l87
					}
					position++
					goto l86
				l87:
					position, tokenIndex, depth = position86, tokenIndex86, depth86
					if buffer[position] != rune('L') {
						goto l52
					}
					position++
				}
			l86:
				{
					position88, tokenIndex88, depth88 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l89
					}
					position++
					goto l88
				l89:
					position, tokenIndex, depth = position88, tokenIndex88, depth88
					if buffer[position] != rune('E') {
						goto l52
					}
					position++
				}
			l88:
				{
					position90, tokenIndex90, depth90 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l91
					}
					position++
					goto l90
				l91:
					position, tokenIndex, depth = position90, tokenIndex90, depth90
					if buffer[position] != rune('C') {
						goto l52
					}
					position++
				}
			l90:
				{
					position92, tokenIndex92, depth92 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l93
					}
					position++
					goto l92
				l93:
					position, tokenIndex, depth = position92, tokenIndex92, depth92
					if buffer[position] != rune('T') {
						goto l52
					}
					position++
				}
			l92:
				if !_rules[rulesp]() {
					goto l52
				}
				if !_rules[ruleEmitter]() {
					goto l52
				}
				if !_rules[rulesp]() {
					goto l52
				}
				if !_rules[ruleProjections]() {
					goto l52
				}
				if !_rules[rulesp]() {
					goto l52
				}
				if !_rules[ruleWindowedFrom]() {
					goto l52
				}
				if !_rules[rulesp]() {
					goto l52
				}
				if !_rules[ruleFilter]() {
					goto l52
				}
				if !_rules[rulesp]() {
					goto l52
				}
				if !_rules[ruleGrouping]() {
					goto l52
				}
				if !_rules[rulesp]() {
					goto l52
				}
				if !_rules[ruleHaving]() {
					goto l52
				}
				if !_rules[rulesp]() {
					goto l52
				}
				if !_rules[ruleAction1]() {
					goto l52
				}
				depth--
				add(ruleCreateStreamAsSelectStmt, position53)
			}
			return true
		l52:
			position, tokenIndex, depth = position52, tokenIndex52, depth52
			return false
		},
		/* 8 CreateSourceStmt <- <(('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp PausedOpt sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier sp (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) sp SourceSinkType sp SourceSinkSpecs Action2)> */
		func() bool {
			position94, tokenIndex94, depth94 := position, tokenIndex, depth
			{
				position95 := position
				depth++
				{
					position96, tokenIndex96, depth96 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l97
					}
					position++
					goto l96
				l97:
					position, tokenIndex, depth = position96, tokenIndex96, depth96
					if buffer[position] != rune('C') {
						goto l94
					}
					position++
				}
			l96:
				{
					position98, tokenIndex98, depth98 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l99
					}
					position++
					goto l98
				l99:
					position, tokenIndex, depth = position98, tokenIndex98, depth98
					if buffer[position] != rune('R') {
						goto l94
					}
					position++
				}
			l98:
				{
					position100, tokenIndex100, depth100 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l101
					}
					position++
					goto l100
				l101:
					position, tokenIndex, depth = position100, tokenIndex100, depth100
					if buffer[position] != rune('E') {
						goto l94
					}
					position++
				}
			l100:
				{
					position102, tokenIndex102, depth102 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l103
					}
					position++
					goto l102
				l103:
					position, tokenIndex, depth = position102, tokenIndex102, depth102
					if buffer[position] != rune('A') {
						goto l94
					}
					position++
				}
			l102:
				{
					position104, tokenIndex104, depth104 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l105
					}
					position++
					goto l104
				l105:
					position, tokenIndex, depth = position104, tokenIndex104, depth104
					if buffer[position] != rune('T') {
						goto l94
					}
					position++
				}
			l104:
				{
					position106, tokenIndex106, depth106 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l107
					}
					position++
					goto l106
				l107:
					position, tokenIndex, depth = position106, tokenIndex106, depth106
					if buffer[position] != rune('E') {
						goto l94
					}
					position++
				}
			l106:
				if !_rules[rulesp]() {
					goto l94
				}
				if !_rules[rulePausedOpt]() {
					goto l94
				}
				if !_rules[rulesp]() {
					goto l94
				}
				{
					position108, tokenIndex108, depth108 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l109
					}
					position++
					goto l108
				l109:
					position, tokenIndex, depth = position108, tokenIndex108, depth108
					if buffer[position] != rune('S') {
						goto l94
					}
					position++
				}
			l108:
				{
					position110, tokenIndex110, depth110 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l111
					}
					position++
					goto l110
				l111:
					position, tokenIndex, depth = position110, tokenIndex110, depth110
					if buffer[position] != rune('O') {
						goto l94
					}
					position++
				}
			l110:
				{
					position112, tokenIndex112, depth112 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l113
					}
					position++
					goto l112
				l113:
					position, tokenIndex, depth = position112, tokenIndex112, depth112
					if buffer[position] != rune('U') {
						goto l94
					}
					position++
				}
			l112:
				{
					position114, tokenIndex114, depth114 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l115
					}
					position++
					goto l114
				l115:
					position, tokenIndex, depth = position114, tokenIndex114, depth114
					if buffer[position] != rune('R') {
						goto l94
					}
					position++
				}
			l114:
				{
					position116, tokenIndex116, depth116 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l117
					}
					position++
					goto l116
				l117:
					position, tokenIndex, depth = position116, tokenIndex116, depth116
					if buffer[position] != rune('C') {
						goto l94
					}
					position++
				}
			l116:
				{
					position118, tokenIndex118, depth118 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l119
					}
					position++
					goto l118
				l119:
					position, tokenIndex, depth = position118, tokenIndex118, depth118
					if buffer[position] != rune('E') {
						goto l94
					}
					position++
				}
			l118:
				if !_rules[rulesp]() {
					goto l94
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l94
				}
				if !_rules[rulesp]() {
					goto l94
				}
				{
					position120, tokenIndex120, depth120 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l121
					}
					position++
					goto l120
				l121:
					position, tokenIndex, depth = position120, tokenIndex120, depth120
					if buffer[position] != rune('T') {
						goto l94
					}
					position++
				}
			l120:
				{
					position122, tokenIndex122, depth122 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l123
					}
					position++
					goto l122
				l123:
					position, tokenIndex, depth = position122, tokenIndex122, depth122
					if buffer[position] != rune('Y') {
						goto l94
					}
					position++
				}
			l122:
				{
					position124, tokenIndex124, depth124 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l125
					}
					position++
					goto l124
				l125:
					position, tokenIndex, depth = position124, tokenIndex124, depth124
					if buffer[position] != rune('P') {
						goto l94
					}
					position++
				}
			l124:
				{
					position126, tokenIndex126, depth126 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l127
					}
					position++
					goto l126
				l127:
					position, tokenIndex, depth = position126, tokenIndex126, depth126
					if buffer[position] != rune('E') {
						goto l94
					}
					position++
				}
			l126:
				if !_rules[rulesp]() {
					goto l94
				}
				if !_rules[ruleSourceSinkType]() {
					goto l94
				}
				if !_rules[rulesp]() {
					goto l94
				}
				if !_rules[ruleSourceSinkSpecs]() {
					goto l94
				}
				if !_rules[ruleAction2]() {
					goto l94
				}
				depth--
				add(ruleCreateSourceStmt, position95)
			}
			return true
		l94:
			position, tokenIndex, depth = position94, tokenIndex94, depth94
			return false
		},
		/* 9 CreateSinkStmt <- <(('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('i' / 'I') ('n' / 'N') ('k' / 'K')) sp StreamIdentifier sp (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) sp SourceSinkType sp SourceSinkSpecs Action3)> */
		func() bool {
			position128, tokenIndex128, depth128 := position, tokenIndex, depth
			{
				position129 := position
				depth++
				{
					position130, tokenIndex130, depth130 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l131
					}
					position++
					goto l130
				l131:
					position, tokenIndex, depth = position130, tokenIndex130, depth130
					if buffer[position] != rune('C') {
						goto l128
					}
					position++
				}
			l130:
				{
					position132, tokenIndex132, depth132 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l133
					}
					position++
					goto l132
				l133:
					position, tokenIndex, depth = position132, tokenIndex132, depth132
					if buffer[position] != rune('R') {
						goto l128
					}
					position++
				}
			l132:
				{
					position134, tokenIndex134, depth134 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l135
					}
					position++
					goto l134
				l135:
					position, tokenIndex, depth = position134, tokenIndex134, depth134
					if buffer[position] != rune('E') {
						goto l128
					}
					position++
				}
			l134:
				{
					position136, tokenIndex136, depth136 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l137
					}
					position++
					goto l136
				l137:
					position, tokenIndex, depth = position136, tokenIndex136, depth136
					if buffer[position] != rune('A') {
						goto l128
					}
					position++
				}
			l136:
				{
					position138, tokenIndex138, depth138 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l139
					}
					position++
					goto l138
				l139:
					position, tokenIndex, depth = position138, tokenIndex138, depth138
					if buffer[position] != rune('T') {
						goto l128
					}
					position++
				}
			l138:
				{
					position140, tokenIndex140, depth140 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l141
					}
					position++
					goto l140
				l141:
					position, tokenIndex, depth = position140, tokenIndex140, depth140
					if buffer[position] != rune('E') {
						goto l128
					}
					position++
				}
			l140:
				if !_rules[rulesp]() {
					goto l128
				}
				{
					position142, tokenIndex142, depth142 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l143
					}
					position++
					goto l142
				l143:
					position, tokenIndex, depth = position142, tokenIndex142, depth142
					if buffer[position] != rune('S') {
						goto l128
					}
					position++
				}
			l142:
				{
					position144, tokenIndex144, depth144 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l145
					}
					position++
					goto l144
				l145:
					position, tokenIndex, depth = position144, tokenIndex144, depth144
					if buffer[position] != rune('I') {
						goto l128
					}
					position++
				}
			l144:
				{
					position146, tokenIndex146, depth146 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l147
					}
					position++
					goto l146
				l147:
					position, tokenIndex, depth = position146, tokenIndex146, depth146
					if buffer[position] != rune('N') {
						goto l128
					}
					position++
				}
			l146:
				{
					position148, tokenIndex148, depth148 := position, tokenIndex, depth
					if buffer[position] != rune('k') {
						goto l149
					}
					position++
					goto l148
				l149:
					position, tokenIndex, depth = position148, tokenIndex148, depth148
					if buffer[position] != rune('K') {
						goto l128
					}
					position++
				}
			l148:
				if !_rules[rulesp]() {
					goto l128
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l128
				}
				if !_rules[rulesp]() {
					goto l128
				}
				{
					position150, tokenIndex150, depth150 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l151
					}
					position++
					goto l150
				l151:
					position, tokenIndex, depth = position150, tokenIndex150, depth150
					if buffer[position] != rune('T') {
						goto l128
					}
					position++
				}
			l150:
				{
					position152, tokenIndex152, depth152 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l153
					}
					position++
					goto l152
				l153:
					position, tokenIndex, depth = position152, tokenIndex152, depth152
					if buffer[position] != rune('Y') {
						goto l128
					}
					position++
				}
			l152:
				{
					position154, tokenIndex154, depth154 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l155
					}
					position++
					goto l154
				l155:
					position, tokenIndex, depth = position154, tokenIndex154, depth154
					if buffer[position] != rune('P') {
						goto l128
					}
					position++
				}
			l154:
				{
					position156, tokenIndex156, depth156 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l157
					}
					position++
					goto l156
				l157:
					position, tokenIndex, depth = position156, tokenIndex156, depth156
					if buffer[position] != rune('E') {
						goto l128
					}
					position++
				}
			l156:
				if !_rules[rulesp]() {
					goto l128
				}
				if !_rules[ruleSourceSinkType]() {
					goto l128
				}
				if !_rules[rulesp]() {
					goto l128
				}
				if !_rules[ruleSourceSinkSpecs]() {
					goto l128
				}
				if !_rules[ruleAction3]() {
					goto l128
				}
				depth--
				add(ruleCreateSinkStmt, position129)
			}
			return true
		l128:
			position, tokenIndex, depth = position128, tokenIndex128, depth128
			return false
		},
		/* 10 CreateStateStmt <- <(('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('e' / 'E')) sp StreamIdentifier sp (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) sp SourceSinkType sp SourceSinkSpecs Action4)> */
		func() bool {
			position158, tokenIndex158, depth158 := position, tokenIndex, depth
			{
				position159 := position
				depth++
				{
					position160, tokenIndex160, depth160 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l161
					}
					position++
					goto l160
				l161:
					position, tokenIndex, depth = position160, tokenIndex160, depth160
					if buffer[position] != rune('C') {
						goto l158
					}
					position++
				}
			l160:
				{
					position162, tokenIndex162, depth162 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l163
					}
					position++
					goto l162
				l163:
					position, tokenIndex, depth = position162, tokenIndex162, depth162
					if buffer[position] != rune('R') {
						goto l158
					}
					position++
				}
			l162:
				{
					position164, tokenIndex164, depth164 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l165
					}
					position++
					goto l164
				l165:
					position, tokenIndex, depth = position164, tokenIndex164, depth164
					if buffer[position] != rune('E') {
						goto l158
					}
					position++
				}
			l164:
				{
					position166, tokenIndex166, depth166 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l167
					}
					position++
					goto l166
				l167:
					position, tokenIndex, depth = position166, tokenIndex166, depth166
					if buffer[position] != rune('A') {
						goto l158
					}
					position++
				}
			l166:
				{
					position168, tokenIndex168, depth168 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l169
					}
					position++
					goto l168
				l169:
					position, tokenIndex, depth = position168, tokenIndex168, depth168
					if buffer[position] != rune('T') {
						goto l158
					}
					position++
				}
			l168:
				{
					position170, tokenIndex170, depth170 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l171
					}
					position++
					goto l170
				l171:
					position, tokenIndex, depth = position170, tokenIndex170, depth170
					if buffer[position] != rune('E') {
						goto l158
					}
					position++
				}
			l170:
				if !_rules[rulesp]() {
					goto l158
				}
				{
					position172, tokenIndex172, depth172 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l173
					}
					position++
					goto l172
				l173:
					position, tokenIndex, depth = position172, tokenIndex172, depth172
					if buffer[position] != rune('S') {
						goto l158
					}
					position++
				}
			l172:
				{
					position174, tokenIndex174, depth174 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l175
					}
					position++
					goto l174
				l175:
					position, tokenIndex, depth = position174, tokenIndex174, depth174
					if buffer[position] != rune('T') {
						goto l158
					}
					position++
				}
			l174:
				{
					position176, tokenIndex176, depth176 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l177
					}
					position++
					goto l176
				l177:
					position, tokenIndex, depth = position176, tokenIndex176, depth176
					if buffer[position] != rune('A') {
						goto l158
					}
					position++
				}
			l176:
				{
					position178, tokenIndex178, depth178 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l179
					}
					position++
					goto l178
				l179:
					position, tokenIndex, depth = position178, tokenIndex178, depth178
					if buffer[position] != rune('T') {
						goto l158
					}
					position++
				}
			l178:
				{
					position180, tokenIndex180, depth180 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l181
					}
					position++
					goto l180
				l181:
					position, tokenIndex, depth = position180, tokenIndex180, depth180
					if buffer[position] != rune('E') {
						goto l158
					}
					position++
				}
			l180:
				if !_rules[rulesp]() {
					goto l158
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l158
				}
				if !_rules[rulesp]() {
					goto l158
				}
				{
					position182, tokenIndex182, depth182 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l183
					}
					position++
					goto l182
				l183:
					position, tokenIndex, depth = position182, tokenIndex182, depth182
					if buffer[position] != rune('T') {
						goto l158
					}
					position++
				}
			l182:
				{
					position184, tokenIndex184, depth184 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l185
					}
					position++
					goto l184
				l185:
					position, tokenIndex, depth = position184, tokenIndex184, depth184
					if buffer[position] != rune('Y') {
						goto l158
					}
					position++
				}
			l184:
				{
					position186, tokenIndex186, depth186 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l187
					}
					position++
					goto l186
				l187:
					position, tokenIndex, depth = position186, tokenIndex186, depth186
					if buffer[position] != rune('P') {
						goto l158
					}
					position++
				}
			l186:
				{
					position188, tokenIndex188, depth188 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l189
					}
					position++
					goto l188
				l189:
					position, tokenIndex, depth = position188, tokenIndex188, depth188
					if buffer[position] != rune('E') {
						goto l158
					}
					position++
				}
			l188:
				if !_rules[rulesp]() {
					goto l158
				}
				if !_rules[ruleSourceSinkType]() {
					goto l158
				}
				if !_rules[rulesp]() {
					goto l158
				}
				if !_rules[ruleSourceSinkSpecs]() {
					goto l158
				}
				if !_rules[ruleAction4]() {
					goto l158
				}
				depth--
				add(ruleCreateStateStmt, position159)
			}
			return true
		l158:
			position, tokenIndex, depth = position158, tokenIndex158, depth158
			return false
		},
		/* 11 UpdateStateStmt <- <(('u' / 'U') ('p' / 'P') ('d' / 'D') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('e' / 'E')) sp StreamIdentifier sp UpdateSourceSinkSpecs Action5)> */
		func() bool {
			position190, tokenIndex190, depth190 := position, tokenIndex, depth
			{
				position191 := position
				depth++
				{
					position192, tokenIndex192, depth192 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l193
					}
					position++
					goto l192
				l193:
					position, tokenIndex, depth = position192, tokenIndex192, depth192
					if buffer[position] != rune('U') {
						goto l190
					}
					position++
				}
			l192:
				{
					position194, tokenIndex194, depth194 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l195
					}
					position++
					goto l194
				l195:
					position, tokenIndex, depth = position194, tokenIndex194, depth194
					if buffer[position] != rune('P') {
						goto l190
					}
					position++
				}
			l194:
				{
					position196, tokenIndex196, depth196 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l197
					}
					position++
					goto l196
				l197:
					position, tokenIndex, depth = position196, tokenIndex196, depth196
					if buffer[position] != rune('D') {
						goto l190
					}
					position++
				}
			l196:
				{
					position198, tokenIndex198, depth198 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l199
					}
					position++
					goto l198
				l199:
					position, tokenIndex, depth = position198, tokenIndex198, depth198
					if buffer[position] != rune('A') {
						goto l190
					}
					position++
				}
			l198:
				{
					position200, tokenIndex200, depth200 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l201
					}
					position++
					goto l200
				l201:
					position, tokenIndex, depth = position200, tokenIndex200, depth200
					if buffer[position] != rune('T') {
						goto l190
					}
					position++
				}
			l200:
				{
					position202, tokenIndex202, depth202 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l203
					}
					position++
					goto l202
				l203:
					position, tokenIndex, depth = position202, tokenIndex202, depth202
					if buffer[position] != rune('E') {
						goto l190
					}
					position++
				}
			l202:
				if !_rules[rulesp]() {
					goto l190
				}
				{
					position204, tokenIndex204, depth204 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l205
					}
					position++
					goto l204
				l205:
					position, tokenIndex, depth = position204, tokenIndex204, depth204
					if buffer[position] != rune('S') {
						goto l190
					}
					position++
				}
			l204:
				{
					position206, tokenIndex206, depth206 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l207
					}
					position++
					goto l206
				l207:
					position, tokenIndex, depth = position206, tokenIndex206, depth206
					if buffer[position] != rune('T') {
						goto l190
					}
					position++
				}
			l206:
				{
					position208, tokenIndex208, depth208 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l209
					}
					position++
					goto l208
				l209:
					position, tokenIndex, depth = position208, tokenIndex208, depth208
					if buffer[position] != rune('A') {
						goto l190
					}
					position++
				}
			l208:
				{
					position210, tokenIndex210, depth210 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l211
					}
					position++
					goto l210
				l211:
					position, tokenIndex, depth = position210, tokenIndex210, depth210
					if buffer[position] != rune('T') {
						goto l190
					}
					position++
				}
			l210:
				{
					position212, tokenIndex212, depth212 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l213
					}
					position++
					goto l212
				l213:
					position, tokenIndex, depth = position212, tokenIndex212, depth212
					if buffer[position] != rune('E') {
						goto l190
					}
					position++
				}
			l212:
				if !_rules[rulesp]() {
					goto l190
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l190
				}
				if !_rules[rulesp]() {
					goto l190
				}
				if !_rules[ruleUpdateSourceSinkSpecs]() {
					goto l190
				}
				if !_rules[ruleAction5]() {
					goto l190
				}
				depth--
				add(ruleUpdateStateStmt, position191)
			}
			return true
		l190:
			position, tokenIndex, depth = position190, tokenIndex190, depth190
			return false
		},
		/* 12 UpdateSourceStmt <- <(('u' / 'U') ('p' / 'P') ('d' / 'D') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier sp UpdateSourceSinkSpecs Action6)> */
		func() bool {
			position214, tokenIndex214, depth214 := position, tokenIndex, depth
			{
				position215 := position
				depth++
				{
					position216, tokenIndex216, depth216 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l217
					}
					position++
					goto l216
				l217:
					position, tokenIndex, depth = position216, tokenIndex216, depth216
					if buffer[position] != rune('U') {
						goto l214
					}
					position++
				}
			l216:
				{
					position218, tokenIndex218, depth218 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l219
					}
					position++
					goto l218
				l219:
					position, tokenIndex, depth = position218, tokenIndex218, depth218
					if buffer[position] != rune('P') {
						goto l214
					}
					position++
				}
			l218:
				{
					position220, tokenIndex220, depth220 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l221
					}
					position++
					goto l220
				l221:
					position, tokenIndex, depth = position220, tokenIndex220, depth220
					if buffer[position] != rune('D') {
						goto l214
					}
					position++
				}
			l220:
				{
					position222, tokenIndex222, depth222 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l223
					}
					position++
					goto l222
				l223:
					position, tokenIndex, depth = position222, tokenIndex222, depth222
					if buffer[position] != rune('A') {
						goto l214
					}
					position++
				}
			l222:
				{
					position224, tokenIndex224, depth224 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l225
					}
					position++
					goto l224
				l225:
					position, tokenIndex, depth = position224, tokenIndex224, depth224
					if buffer[position] != rune('T') {
						goto l214
					}
					position++
				}
			l224:
				{
					position226, tokenIndex226, depth226 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l227
					}
					position++
					goto l226
				l227:
					position, tokenIndex, depth = position226, tokenIndex226, depth226
					if buffer[position] != rune('E') {
						goto l214
					}
					position++
				}
			l226:
				if !_rules[rulesp]() {
					goto l214
				}
				{
					position228, tokenIndex228, depth228 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l229
					}
					position++
					goto l228
				l229:
					position, tokenIndex, depth = position228, tokenIndex228, depth228
					if buffer[position] != rune('S') {
						goto l214
					}
					position++
				}
			l228:
				{
					position230, tokenIndex230, depth230 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l231
					}
					position++
					goto l230
				l231:
					position, tokenIndex, depth = position230, tokenIndex230, depth230
					if buffer[position] != rune('O') {
						goto l214
					}
					position++
				}
			l230:
				{
					position232, tokenIndex232, depth232 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l233
					}
					position++
					goto l232
				l233:
					position, tokenIndex, depth = position232, tokenIndex232, depth232
					if buffer[position] != rune('U') {
						goto l214
					}
					position++
				}
			l232:
				{
					position234, tokenIndex234, depth234 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l235
					}
					position++
					goto l234
				l235:
					position, tokenIndex, depth = position234, tokenIndex234, depth234
					if buffer[position] != rune('R') {
						goto l214
					}
					position++
				}
			l234:
				{
					position236, tokenIndex236, depth236 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l237
					}
					position++
					goto l236
				l237:
					position, tokenIndex, depth = position236, tokenIndex236, depth236
					if buffer[position] != rune('C') {
						goto l214
					}
					position++
				}
			l236:
				{
					position238, tokenIndex238, depth238 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l239
					}
					position++
					goto l238
				l239:
					position, tokenIndex, depth = position238, tokenIndex238, depth238
					if buffer[position] != rune('E') {
						goto l214
					}
					position++
				}
			l238:
				if !_rules[rulesp]() {
					goto l214
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l214
				}
				if !_rules[rulesp]() {
					goto l214
				}
				if !_rules[ruleUpdateSourceSinkSpecs]() {
					goto l214
				}
				if !_rules[ruleAction6]() {
					goto l214
				}
				depth--
				add(ruleUpdateSourceStmt, position215)
			}
			return true
		l214:
			position, tokenIndex, depth = position214, tokenIndex214, depth214
			return false
		},
		/* 13 UpdateSinkStmt <- <(('u' / 'U') ('p' / 'P') ('d' / 'D') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('i' / 'I') ('n' / 'N') ('k' / 'K')) sp StreamIdentifier sp UpdateSourceSinkSpecs Action7)> */
		func() bool {
			position240, tokenIndex240, depth240 := position, tokenIndex, depth
			{
				position241 := position
				depth++
				{
					position242, tokenIndex242, depth242 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l243
					}
					position++
					goto l242
				l243:
					position, tokenIndex, depth = position242, tokenIndex242, depth242
					if buffer[position] != rune('U') {
						goto l240
					}
					position++
				}
			l242:
				{
					position244, tokenIndex244, depth244 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l245
					}
					position++
					goto l244
				l245:
					position, tokenIndex, depth = position244, tokenIndex244, depth244
					if buffer[position] != rune('P') {
						goto l240
					}
					position++
				}
			l244:
				{
					position246, tokenIndex246, depth246 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l247
					}
					position++
					goto l246
				l247:
					position, tokenIndex, depth = position246, tokenIndex246, depth246
					if buffer[position] != rune('D') {
						goto l240
					}
					position++
				}
			l246:
				{
					position248, tokenIndex248, depth248 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l249
					}
					position++
					goto l248
				l249:
					position, tokenIndex, depth = position248, tokenIndex248, depth248
					if buffer[position] != rune('A') {
						goto l240
					}
					position++
				}
			l248:
				{
					position250, tokenIndex250, depth250 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l251
					}
					position++
					goto l250
				l251:
					position, tokenIndex, depth = position250, tokenIndex250, depth250
					if buffer[position] != rune('T') {
						goto l240
					}
					position++
				}
			l250:
				{
					position252, tokenIndex252, depth252 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l253
					}
					position++
					goto l252
				l253:
					position, tokenIndex, depth = position252, tokenIndex252, depth252
					if buffer[position] != rune('E') {
						goto l240
					}
					position++
				}
			l252:
				if !_rules[rulesp]() {
					goto l240
				}
				{
					position254, tokenIndex254, depth254 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l255
					}
					position++
					goto l254
				l255:
					position, tokenIndex, depth = position254, tokenIndex254, depth254
					if buffer[position] != rune('S') {
						goto l240
					}
					position++
				}
			l254:
				{
					position256, tokenIndex256, depth256 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l257
					}
					position++
					goto l256
				l257:
					position, tokenIndex, depth = position256, tokenIndex256, depth256
					if buffer[position] != rune('I') {
						goto l240
					}
					position++
				}
			l256:
				{
					position258, tokenIndex258, depth258 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l259
					}
					position++
					goto l258
				l259:
					position, tokenIndex, depth = position258, tokenIndex258, depth258
					if buffer[position] != rune('N') {
						goto l240
					}
					position++
				}
			l258:
				{
					position260, tokenIndex260, depth260 := position, tokenIndex, depth
					if buffer[position] != rune('k') {
						goto l261
					}
					position++
					goto l260
				l261:
					position, tokenIndex, depth = position260, tokenIndex260, depth260
					if buffer[position] != rune('K') {
						goto l240
					}
					position++
				}
			l260:
				if !_rules[rulesp]() {
					goto l240
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l240
				}
				if !_rules[rulesp]() {
					goto l240
				}
				if !_rules[ruleUpdateSourceSinkSpecs]() {
					goto l240
				}
				if !_rules[ruleAction7]() {
					goto l240
				}
				depth--
				add(ruleUpdateSinkStmt, position241)
			}
			return true
		l240:
			position, tokenIndex, depth = position240, tokenIndex240, depth240
			return false
		},
		/* 14 InsertIntoSelectStmt <- <(('i' / 'I') ('n' / 'N') ('s' / 'S') ('e' / 'E') ('r' / 'R') ('t' / 'T') sp (('i' / 'I') ('n' / 'N') ('t' / 'T') ('o' / 'O')) sp StreamIdentifier sp SelectStmt Action8)> */
		func() bool {
			position262, tokenIndex262, depth262 := position, tokenIndex, depth
			{
				position263 := position
				depth++
				{
					position264, tokenIndex264, depth264 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l265
					}
					position++
					goto l264
				l265:
					position, tokenIndex, depth = position264, tokenIndex264, depth264
					if buffer[position] != rune('I') {
						goto l262
					}
					position++
				}
			l264:
				{
					position266, tokenIndex266, depth266 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l267
					}
					position++
					goto l266
				l267:
					position, tokenIndex, depth = position266, tokenIndex266, depth266
					if buffer[position] != rune('N') {
						goto l262
					}
					position++
				}
			l266:
				{
					position268, tokenIndex268, depth268 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l269
					}
					position++
					goto l268
				l269:
					position, tokenIndex, depth = position268, tokenIndex268, depth268
					if buffer[position] != rune('S') {
						goto l262
					}
					position++
				}
			l268:
				{
					position270, tokenIndex270, depth270 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l271
					}
					position++
					goto l270
				l271:
					position, tokenIndex, depth = position270, tokenIndex270, depth270
					if buffer[position] != rune('E') {
						goto l262
					}
					position++
				}
			l270:
				{
					position272, tokenIndex272, depth272 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l273
					}
					position++
					goto l272
				l273:
					position, tokenIndex, depth = position272, tokenIndex272, depth272
					if buffer[position] != rune('R') {
						goto l262
					}
					position++
				}
			l272:
				{
					position274, tokenIndex274, depth274 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l275
					}
					position++
					goto l274
				l275:
					position, tokenIndex, depth = position274, tokenIndex274, depth274
					if buffer[position] != rune('T') {
						goto l262
					}
					position++
				}
			l274:
				if !_rules[rulesp]() {
					goto l262
				}
				{
					position276, tokenIndex276, depth276 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l277
					}
					position++
					goto l276
				l277:
					position, tokenIndex, depth = position276, tokenIndex276, depth276
					if buffer[position] != rune('I') {
						goto l262
					}
					position++
				}
			l276:
				{
					position278, tokenIndex278, depth278 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l279
					}
					position++
					goto l278
				l279:
					position, tokenIndex, depth = position278, tokenIndex278, depth278
					if buffer[position] != rune('N') {
						goto l262
					}
					position++
				}
			l278:
				{
					position280, tokenIndex280, depth280 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l281
					}
					position++
					goto l280
				l281:
					position, tokenIndex, depth = position280, tokenIndex280, depth280
					if buffer[position] != rune('T') {
						goto l262
					}
					position++
				}
			l280:
				{
					position282, tokenIndex282, depth282 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l283
					}
					position++
					goto l282
				l283:
					position, tokenIndex, depth = position282, tokenIndex282, depth282
					if buffer[position] != rune('O') {
						goto l262
					}
					position++
				}
			l282:
				if !_rules[rulesp]() {
					goto l262
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l262
				}
				if !_rules[rulesp]() {
					goto l262
				}
				if !_rules[ruleSelectStmt]() {
					goto l262
				}
				if !_rules[ruleAction8]() {
					goto l262
				}
				depth--
				add(ruleInsertIntoSelectStmt, position263)
			}
			return true
		l262:
			position, tokenIndex, depth = position262, tokenIndex262, depth262
			return false
		},
		/* 15 InsertIntoFromStmt <- <(('i' / 'I') ('n' / 'N') ('s' / 'S') ('e' / 'E') ('r' / 'R') ('t' / 'T') sp (('i' / 'I') ('n' / 'N') ('t' / 'T') ('o' / 'O')) sp StreamIdentifier sp (('f' / 'F') ('r' / 'R') ('o' / 'O') ('m' / 'M')) sp StreamIdentifier Action9)> */
		func() bool {
			position284, tokenIndex284, depth284 := position, tokenIndex, depth
			{
				position285 := position
				depth++
				{
					position286, tokenIndex286, depth286 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l287
					}
					position++
					goto l286
				l287:
					position, tokenIndex, depth = position286, tokenIndex286, depth286
					if buffer[position] != rune('I') {
						goto l284
					}
					position++
				}
			l286:
				{
					position288, tokenIndex288, depth288 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l289
					}
					position++
					goto l288
				l289:
					position, tokenIndex, depth = position288, tokenIndex288, depth288
					if buffer[position] != rune('N') {
						goto l284
					}
					position++
				}
			l288:
				{
					position290, tokenIndex290, depth290 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l291
					}
					position++
					goto l290
				l291:
					position, tokenIndex, depth = position290, tokenIndex290, depth290
					if buffer[position] != rune('S') {
						goto l284
					}
					position++
				}
			l290:
				{
					position292, tokenIndex292, depth292 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l293
					}
					position++
					goto l292
				l293:
					position, tokenIndex, depth = position292, tokenIndex292, depth292
					if buffer[position] != rune('E') {
						goto l284
					}
					position++
				}
			l292:
				{
					position294, tokenIndex294, depth294 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l295
					}
					position++
					goto l294
				l295:
					position, tokenIndex, depth = position294, tokenIndex294, depth294
					if buffer[position] != rune('R') {
						goto l284
					}
					position++
				}
			l294:
				{
					position296, tokenIndex296, depth296 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l297
					}
					position++
					goto l296
				l297:
					position, tokenIndex, depth = position296, tokenIndex296, depth296
					if buffer[position] != rune('T') {
						goto l284
					}
					position++
				}
			l296:
				if !_rules[rulesp]() {
					goto l284
				}
				{
					position298, tokenIndex298, depth298 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l299
					}
					position++
					goto l298
				l299:
					position, tokenIndex, depth = position298, tokenIndex298, depth298
					if buffer[position] != rune('I') {
						goto l284
					}
					position++
				}
			l298:
				{
					position300, tokenIndex300, depth300 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l301
					}
					position++
					goto l300
				l301:
					position, tokenIndex, depth = position300, tokenIndex300, depth300
					if buffer[position] != rune('N') {
						goto l284
					}
					position++
				}
			l300:
				{
					position302, tokenIndex302, depth302 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l303
					}
					position++
					goto l302
				l303:
					position, tokenIndex, depth = position302, tokenIndex302, depth302
					if buffer[position] != rune('T') {
						goto l284
					}
					position++
				}
			l302:
				{
					position304, tokenIndex304, depth304 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l305
					}
					position++
					goto l304
				l305:
					position, tokenIndex, depth = position304, tokenIndex304, depth304
					if buffer[position] != rune('O') {
						goto l284
					}
					position++
				}
			l304:
				if !_rules[rulesp]() {
					goto l284
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l284
				}
				if !_rules[rulesp]() {
					goto l284
				}
				{
					position306, tokenIndex306, depth306 := position, tokenIndex, depth
					if buffer[position] != rune('f') {
						goto l307
					}
					position++
					goto l306
				l307:
					position, tokenIndex, depth = position306, tokenIndex306, depth306
					if buffer[position] != rune('F') {
						goto l284
					}
					position++
				}
			l306:
				{
					position308, tokenIndex308, depth308 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l309
					}
					position++
					goto l308
				l309:
					position, tokenIndex, depth = position308, tokenIndex308, depth308
					if buffer[position] != rune('R') {
						goto l284
					}
					position++
				}
			l308:
				{
					position310, tokenIndex310, depth310 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l311
					}
					position++
					goto l310
				l311:
					position, tokenIndex, depth = position310, tokenIndex310, depth310
					if buffer[position] != rune('O') {
						goto l284
					}
					position++
				}
			l310:
				{
					position312, tokenIndex312, depth312 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l313
					}
					position++
					goto l312
				l313:
					position, tokenIndex, depth = position312, tokenIndex312, depth312
					if buffer[position] != rune('M') {
						goto l284
					}
					position++
				}
			l312:
				if !_rules[rulesp]() {
					goto l284
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l284
				}
				if !_rules[ruleAction9]() {
					goto l284
				}
				depth--
				add(ruleInsertIntoFromStmt, position285)
			}
			return true
		l284:
			position, tokenIndex, depth = position284, tokenIndex284, depth284
			return false
		},
		/* 16 PauseSourceStmt <- <(('p' / 'P') ('a' / 'A') ('u' / 'U') ('s' / 'S') ('e' / 'E') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier Action10)> */
		func() bool {
			position314, tokenIndex314, depth314 := position, tokenIndex, depth
			{
				position315 := position
				depth++
				{
					position316, tokenIndex316, depth316 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l317
					}
					position++
					goto l316
				l317:
					position, tokenIndex, depth = position316, tokenIndex316, depth316
					if buffer[position] != rune('P') {
						goto l314
					}
					position++
				}
			l316:
				{
					position318, tokenIndex318, depth318 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l319
					}
					position++
					goto l318
				l319:
					position, tokenIndex, depth = position318, tokenIndex318, depth318
					if buffer[position] != rune('A') {
						goto l314
					}
					position++
				}
			l318:
				{
					position320, tokenIndex320, depth320 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l321
					}
					position++
					goto l320
				l321:
					position, tokenIndex, depth = position320, tokenIndex320, depth320
					if buffer[position] != rune('U') {
						goto l314
					}
					position++
				}
			l320:
				{
					position322, tokenIndex322, depth322 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l323
					}
					position++
					goto l322
				l323:
					position, tokenIndex, depth = position322, tokenIndex322, depth322
					if buffer[position] != rune('S') {
						goto l314
					}
					position++
				}
			l322:
				{
					position324, tokenIndex324, depth324 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l325
					}
					position++
					goto l324
				l325:
					position, tokenIndex, depth = position324, tokenIndex324, depth324
					if buffer[position] != rune('E') {
						goto l314
					}
					position++
				}
			l324:
				if !_rules[rulesp]() {
					goto l314
				}
				{
					position326, tokenIndex326, depth326 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l327
					}
					position++
					goto l326
				l327:
					position, tokenIndex, depth = position326, tokenIndex326, depth326
					if buffer[position] != rune('S') {
						goto l314
					}
					position++
				}
			l326:
				{
					position328, tokenIndex328, depth328 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l329
					}
					position++
					goto l328
				l329:
					position, tokenIndex, depth = position328, tokenIndex328, depth328
					if buffer[position] != rune('O') {
						goto l314
					}
					position++
				}
			l328:
				{
					position330, tokenIndex330, depth330 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l331
					}
					position++
					goto l330
				l331:
					position, tokenIndex, depth = position330, tokenIndex330, depth330
					if buffer[position] != rune('U') {
						goto l314
					}
					position++
				}
			l330:
				{
					position332, tokenIndex332, depth332 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l333
					}
					position++
					goto l332
				l333:
					position, tokenIndex, depth = position332, tokenIndex332, depth332
					if buffer[position] != rune('R') {
						goto l314
					}
					position++
				}
			l332:
				{
					position334, tokenIndex334, depth334 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l335
					}
					position++
					goto l334
				l335:
					position, tokenIndex, depth = position334, tokenIndex334, depth334
					if buffer[position] != rune('C') {
						goto l314
					}
					position++
				}
			l334:
				{
					position336, tokenIndex336, depth336 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l337
					}
					position++
					goto l336
				l337:
					position, tokenIndex, depth = position336, tokenIndex336, depth336
					if buffer[position] != rune('E') {
						goto l314
					}
					position++
				}
			l336:
				if !_rules[rulesp]() {
					goto l314
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l314
				}
				if !_rules[ruleAction10]() {
					goto l314
				}
				depth--
				add(rulePauseSourceStmt, position315)
			}
			return true
		l314:
			position, tokenIndex, depth = position314, tokenIndex314, depth314
			return false
		},
		/* 17 ResumeSourceStmt <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ('u' / 'U') ('m' / 'M') ('e' / 'E') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier Action11)> */
		func() bool {
			position338, tokenIndex338, depth338 := position, tokenIndex, depth
			{
				position339 := position
				depth++
				{
					position340, tokenIndex340, depth340 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l341
					}
					position++
					goto l340
				l341:
					position, tokenIndex, depth = position340, tokenIndex340, depth340
					if buffer[position] != rune('R') {
						goto l338
					}
					position++
				}
			l340:
				{
					position342, tokenIndex342, depth342 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l343
					}
					position++
					goto l342
				l343:
					position, tokenIndex, depth = position342, tokenIndex342, depth342
					if buffer[position] != rune('E') {
						goto l338
					}
					position++
				}
			l342:
				{
					position344, tokenIndex344, depth344 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l345
					}
					position++
					goto l344
				l345:
					position, tokenIndex, depth = position344, tokenIndex344, depth344
					if buffer[position] != rune('S') {
						goto l338
					}
					position++
				}
			l344:
				{
					position346, tokenIndex346, depth346 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l347
					}
					position++
					goto l346
				l347:
					position, tokenIndex, depth = position346, tokenIndex346, depth346
					if buffer[position] != rune('U') {
						goto l338
					}
					position++
				}
			l346:
				{
					position348, tokenIndex348, depth348 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l349
					}
					position++
					goto l348
				l349:
					position, tokenIndex, depth = position348, tokenIndex348, depth348
					if buffer[position] != rune('M') {
						goto l338
					}
					position++
				}
			l348:
				{
					position350, tokenIndex350, depth350 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l351
					}
					position++
					goto l350
				l351:
					position, tokenIndex, depth = position350, tokenIndex350, depth350
					if buffer[position] != rune('E') {
						goto l338
					}
					position++
				}
			l350:
				if !_rules[rulesp]() {
					goto l338
				}
				{
					position352, tokenIndex352, depth352 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l353
					}
					position++
					goto l352
				l353:
					position, tokenIndex, depth = position352, tokenIndex352, depth352
					if buffer[position] != rune('S') {
						goto l338
					}
					position++
				}
			l352:
				{
					position354, tokenIndex354, depth354 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l355
					}
					position++
					goto l354
				l355:
					position, tokenIndex, depth = position354, tokenIndex354, depth354
					if buffer[position] != rune('O') {
						goto l338
					}
					position++
				}
			l354:
				{
					position356, tokenIndex356, depth356 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l357
					}
					position++
					goto l356
				l357:
					position, tokenIndex, depth = position356, tokenIndex356, depth356
					if buffer[position] != rune('U') {
						goto l338
					}
					position++
				}
			l356:
				{
					position358, tokenIndex358, depth358 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l359
					}
					position++
					goto l358
				l359:
					position, tokenIndex, depth = position358, tokenIndex358, depth358
					if buffer[position] != rune('R') {
						goto l338
					}
					position++
				}
			l358:
				{
					position360, tokenIndex360, depth360 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l361
					}
					position++
					goto l360
				l361:
					position, tokenIndex, depth = position360, tokenIndex360, depth360
					if buffer[position] != rune('C') {
						goto l338
					}
					position++
				}
			l360:
				{
					position362, tokenIndex362, depth362 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l363
					}
					position++
					goto l362
				l363:
					position, tokenIndex, depth = position362, tokenIndex362, depth362
					if buffer[position] != rune('E') {
						goto l338
					}
					position++
				}
			l362:
				if !_rules[rulesp]() {
					goto l338
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l338
				}
				if !_rules[ruleAction11]() {
					goto l338
				}
				depth--
				add(ruleResumeSourceStmt, position339)
			}
			return true
		l338:
			position, tokenIndex, depth = position338, tokenIndex338, depth338
			return false
		},
		/* 18 RewindSourceStmt <- <(('r' / 'R') ('e' / 'E') ('w' / 'W') ('i' / 'I') ('n' / 'N') ('d' / 'D') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier Action12)> */
		func() bool {
			position364, tokenIndex364, depth364 := position, tokenIndex, depth
			{
				position365 := position
				depth++
				{
					position366, tokenIndex366, depth366 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l367
					}
					position++
					goto l366
				l367:
					position, tokenIndex, depth = position366, tokenIndex366, depth366
					if buffer[position] != rune('R') {
						goto l364
					}
					position++
				}
			l366:
				{
					position368, tokenIndex368, depth368 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l369
					}
					position++
					goto l368
				l369:
					position, tokenIndex, depth = position368, tokenIndex368, depth368
					if buffer[position] != rune('E') {
						goto l364
					}
					position++
				}
			l368:
				{
					position370, tokenIndex370, depth370 := position, tokenIndex, depth
					if buffer[position] != rune('w') {
						goto l371
					}
					position++
					goto l370
				l371:
					position, tokenIndex, depth = position370, tokenIndex370, depth370
					if buffer[position] != rune('W') {
						goto l364
					}
					position++
				}
			l370:
				{
					position372, tokenIndex372, depth372 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l373
					}
					position++
					goto l372
				l373:
					position, tokenIndex, depth = position372, tokenIndex372, depth372
					if buffer[position] != rune('I') {
						goto l364
					}
					position++
				}
			l372:
				{
					position374, tokenIndex374, depth374 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l375
					}
					position++
					goto l374
				l375:
					position, tokenIndex, depth = position374, tokenIndex374, depth374
					if buffer[position] != rune('N') {
						goto l364
					}
					position++
				}
			l374:
				{
					position376, tokenIndex376, depth376 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l377
					}
					position++
					goto l376
				l377:
					position, tokenIndex, depth = position376, tokenIndex376, depth376
					if buffer[position] != rune('D') {
						goto l364
					}
					position++
				}
			l376:
				if !_rules[rulesp]() {
					goto l364
				}
				{
					position378, tokenIndex378, depth378 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l379
					}
					position++
					goto l378
				l379:
					position, tokenIndex, depth = position378, tokenIndex378, depth378
					if buffer[position] != rune('S') {
						goto l364
					}
					position++
				}
			l378:
				{
					position380, tokenIndex380, depth380 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l381
					}
					position++
					goto l380
				l381:
					position, tokenIndex, depth = position380, tokenIndex380, depth380
					if buffer[position] != rune('O') {
						goto l364
					}
					position++
				}
			l380:
				{
					position382, tokenIndex382, depth382 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l383
					}
					position++
					goto l382
				l383:
					position, tokenIndex, depth = position382, tokenIndex382, depth382
					if buffer[position] != rune('U') {
						goto l364
					}
					position++
				}
			l382:
				{
					position384, tokenIndex384, depth384 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l385
					}
					position++
					goto l384
				l385:
					position, tokenIndex, depth = position384, tokenIndex384, depth384
					if buffer[position] != rune('R') {
						goto l364
					}
					position++
				}
			l384:
				{
					position386, tokenIndex386, depth386 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l387
					}
					position++
					goto l386
				l387:
					position, tokenIndex, depth = position386, tokenIndex386, depth386
					if buffer[position] != rune('C') {
						goto l364
					}
					position++
				}
			l386:
				{
					position388, tokenIndex388, depth388 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l389
					}
					position++
					goto l388
				l389:
					position, tokenIndex, depth = position388, tokenIndex388, depth388
					if buffer[position] != rune('E') {
						goto l364
					}
					position++
				}
			l388:
				if !_rules[rulesp]() {
					goto l364
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l364
				}
				if !_rules[ruleAction12]() {
					goto l364
				}
				depth--
				add(ruleRewindSourceStmt, position365)
			}
			return true
		l364:
			position, tokenIndex, depth = position364, tokenIndex364, depth364
			return false
		},
		/* 19 DropSourceStmt <- <(('d' / 'D') ('r' / 'R') ('o' / 'O') ('p' / 'P') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier Action13)> */
		func() bool {
			position390, tokenIndex390, depth390 := position, tokenIndex, depth
			{
				position391 := position
				depth++
				{
					position392, tokenIndex392, depth392 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l393
					}
					position++
					goto l392
				l393:
					position, tokenIndex, depth = position392, tokenIndex392, depth392
					if buffer[position] != rune('D') {
						goto l390
					}
					position++
				}
			l392:
				{
					position394, tokenIndex394, depth394 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l395
					}
					position++
					goto l394
				l395:
					position, tokenIndex, depth = position394, tokenIndex394, depth394
					if buffer[position] != rune('R') {
						goto l390
					}
					position++
				}
			l394:
				{
					position396, tokenIndex396, depth396 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l397
					}
					position++
					goto l396
				l397:
					position, tokenIndex, depth = position396, tokenIndex396, depth396
					if buffer[position] != rune('O') {
						goto l390
					}
					position++
				}
			l396:
				{
					position398, tokenIndex398, depth398 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l399
					}
					position++
					goto l398
				l399:
					position, tokenIndex, depth = position398, tokenIndex398, depth398
					if buffer[position] != rune('P') {
						goto l390
					}
					position++
				}
			l398:
				if !_rules[rulesp]() {
					goto l390
				}
				{
					position400, tokenIndex400, depth400 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l401
					}
					position++
					goto l400
				l401:
					position, tokenIndex, depth = position400, tokenIndex400, depth400
					if buffer[position] != rune('S') {
						goto l390
					}
					position++
				}
			l400:
				{
					position402, tokenIndex402, depth402 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l403
					}
					position++
					goto l402
				l403:
					position, tokenIndex, depth = position402, tokenIndex402, depth402
					if buffer[position] != rune('O') {
						goto l390
					}
					position++
				}
			l402:
				{
					position404, tokenIndex404, depth404 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l405
					}
					position++
					goto l404
				l405:
					position, tokenIndex, depth = position404, tokenIndex404, depth404
					if buffer[position] != rune('U') {
						goto l390
					}
					position++
				}
			l404:
				{
					position406, tokenIndex406, depth406 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l407
					}
					position++
					goto l406
				l407:
					position, tokenIndex, depth = position406, tokenIndex406, depth406
					if buffer[position] != rune('R') {
						goto l390
					}
					position++
				}
			l406:
				{
					position408, tokenIndex408, depth408 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l409
					}
					position++
					goto l408
				l409:
					position, tokenIndex, depth = position408, tokenIndex408, depth408
					if buffer[position] != rune('C') {
						goto l390
					}
					position++
				}
			l408:
				{
					position410, tokenIndex410, depth410 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l411
					}
					position++
					goto l410
				l411:
					position, tokenIndex, depth = position410, tokenIndex410, depth410
					if buffer[position] != rune('E') {
						goto l390
					}
					position++
				}
			l410:
				if !_rules[rulesp]() {
					goto l390
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l390
				}
				if !_rules[ruleAction13]() {
					goto l390
				}
				depth--
				add(ruleDropSourceStmt, position391)
			}
			return true
		l390:
			position, tokenIndex, depth = position390, tokenIndex390, depth390
			return false
		},
		/* 20 DropStreamStmt <- <(('d' / 'D') ('r' / 'R') ('o' / 'O') ('p' / 'P') sp (('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M')) sp StreamIdentifier Action14)> */
		func() bool {
			position412, tokenIndex412, depth412 := position, tokenIndex, depth
			{
				position413 := position
				depth++
				{
					position414, tokenIndex414, depth414 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l415
					}
					position++
					goto l414
				l415:
					position, tokenIndex, depth = position414, tokenIndex414, depth414
					if buffer[position] != rune('D') {
						goto l412
					}
					position++
				}
			l414:
				{
					position416, tokenIndex416, depth416 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l417
					}
					position++
					goto l416
				l417:
					position, tokenIndex, depth = position416, tokenIndex416, depth416
					if buffer[position] != rune('R') {
						goto l412
					}
					position++
				}
			l416:
				{
					position418, tokenIndex418, depth418 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l419
					}
					position++
					goto l418
				l419:
					position, tokenIndex, depth = position418, tokenIndex418, depth418
					if buffer[position] != rune('O') {
						goto l412
					}
					position++
				}
			l418:
				{
					position420, tokenIndex420, depth420 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l421
					}
					position++
					goto l420
				l421:
					position, tokenIndex, depth = position420, tokenIndex420, depth420
					if buffer[position] != rune('P') {
						goto l412
					}
					position++
				}
			l420:
				if !_rules[rulesp]() {
					goto l412
				}
				{
					position422, tokenIndex422, depth422 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l423
					}
					position++
					goto l422
				l423:
					position, tokenIndex, depth = position422, tokenIndex422, depth422
					if buffer[position] != rune('S') {
						goto l412
					}
					position++
				}
			l422:
				{
					position424, tokenIndex424, depth424 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l425
					}
					position++
					goto l424
				l425:
					position, tokenIndex, depth = position424, tokenIndex424, depth424
					if buffer[position] != rune('T') {
						goto l412
					}
					position++
				}
			l424:
				{
					position426, tokenIndex426, depth426 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l427
					}
					position++
					goto l426
				l427:
					position, tokenIndex, depth = position426, tokenIndex426, depth426
					if buffer[position] != rune('R') {
						goto l412
					}
					position++
				}
			l426:
				{
					position428, tokenIndex428, depth428 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l429
					}
					position++
					goto l428
				l429:
					position, tokenIndex, depth = position428, tokenIndex428, depth428
					if buffer[position] != rune('E') {
						goto l412
					}
					position++
				}
			l428:
				{
					position430, tokenIndex430, depth430 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l431
					}
					position++
					goto l430
				l431:
					position, tokenIndex, depth = position430, tokenIndex430, depth430
					if buffer[position] != rune('A') {
						goto l412
					}
					position++
				}
			l430:
				{
					position432, tokenIndex432, depth432 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l433
					}
					position++
					goto l432
				l433:
					position, tokenIndex, depth = position432, tokenIndex432, depth432
					if buffer[position] != rune('M') {
						goto l412
					}
					position++
				}
			l432:
				if !_rules[rulesp]() {
					goto l412
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l412
				}
				if !_rules[ruleAction14]() {
					goto l412
				}
				depth--
				add(ruleDropStreamStmt, position413)
			}
			return true
		l412:
			position, tokenIndex, depth = position412, tokenIndex412, depth412
			return false
		},
		/* 21 DropSinkStmt <- <(('d' / 'D') ('r' / 'R') ('o' / 'O') ('p' / 'P') sp (('s' / 'S') ('i' / 'I') ('n' / 'N') ('k' / 'K')) sp StreamIdentifier Action15)> */
		func() bool {
			position434, tokenIndex434, depth434 := position, tokenIndex, depth
			{
				position435 := position
				depth++
				{
					position436, tokenIndex436, depth436 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l437
					}
					position++
					goto l436
				l437:
					position, tokenIndex, depth = position436, tokenIndex436, depth436
					if buffer[position] != rune('D') {
						goto l434
					}
					position++
				}
			l436:
				{
					position438, tokenIndex438, depth438 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l439
					}
					position++
					goto l438
				l439:
					position, tokenIndex, depth = position438, tokenIndex438, depth438
					if buffer[position] != rune('R') {
						goto l434
					}
					position++
				}
			l438:
				{
					position440, tokenIndex440, depth440 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l441
					}
					position++
					goto l440
				l441:
					position, tokenIndex, depth = position440, tokenIndex440, depth440
					if buffer[position] != rune('O') {
						goto l434
					}
					position++
				}
			l440:
				{
					position442, tokenIndex442, depth442 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l443
					}
					position++
					goto l442
				l443:
					position, tokenIndex, depth = position442, tokenIndex442, depth442
					if buffer[position] != rune('P') {
						goto l434
					}
					position++
				}
			l442:
				if !_rules[rulesp]() {
					goto l434
				}
				{
					position444, tokenIndex444, depth444 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l445
					}
					position++
					goto l444
				l445:
					position, tokenIndex, depth = position444, tokenIndex444, depth444
					if buffer[position] != rune('S') {
						goto l434
					}
					position++
				}
			l444:
				{
					position446, tokenIndex446, depth446 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l447
					}
					position++
					goto l446
				l447:
					position, tokenIndex, depth = position446, tokenIndex446, depth446
					if buffer[position] != rune('I') {
						goto l434
					}
					position++
				}
			l446:
				{
					position448, tokenIndex448, depth448 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l449
					}
					position++
					goto l448
				l449:
					position, tokenIndex, depth = position448, tokenIndex448, depth448
					if buffer[position] != rune('N') {
						goto l434
					}
					position++
				}
			l448:
				{
					position450, tokenIndex450, depth450 := position, tokenIndex, depth
					if buffer[position] != rune('k') {
						goto l451
					}
					position++
					goto l450
				l451:
					position, tokenIndex, depth = position450, tokenIndex450, depth450
					if buffer[position] != rune('K') {
						goto l434
					}
					position++
				}
			l450:
				if !_rules[rulesp]() {
					goto l434
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l434
				}
				if !_rules[ruleAction15]() {
					goto l434
				}
				depth--
				add(ruleDropSinkStmt, position435)
			}
			return true
		l434:
			position, tokenIndex, depth = position434, tokenIndex434, depth434
			return false
		},
		/* 22 DropStateStmt <- <(('d' / 'D') ('r' / 'R') ('o' / 'O') ('p' / 'P') sp (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('e' / 'E')) sp StreamIdentifier Action16)> */
		func() bool {
			position452, tokenIndex452, depth452 := position, tokenIndex, depth
			{
				position453 := position
				depth++
				{
					position454, tokenIndex454, depth454 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l455
					}
					position++
					goto l454
				l455:
					position, tokenIndex, depth = position454, tokenIndex454, depth454
					if buffer[position] != rune('D') {
						goto l452
					}
					position++
				}
			l454:
				{
					position456, tokenIndex456, depth456 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l457
					}
					position++
					goto l456
				l457:
					position, tokenIndex, depth = position456, tokenIndex456, depth456
					if buffer[position] != rune('R') {
						goto l452
					}
					position++
				}
			l456:
				{
					position458, tokenIndex458, depth458 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l459
					}
					position++
					goto l458
				l459:
					position, tokenIndex, depth = position458, tokenIndex458, depth458
					if buffer[position] != rune('O') {
						goto l452
					}
					position++
				}
			l458:
				{
					position460, tokenIndex460, depth460 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l461
					}
					position++
					goto l460
				l461:
					position, tokenIndex, depth = position460, tokenIndex460, depth460
					if buffer[position] != rune('P') {
						goto l452
					}
					position++
				}
			l460:
				if !_rules[rulesp]() {
					goto l452
				}
				{
					position462, tokenIndex462, depth462 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l463
					}
					position++
					goto l462
				l463:
					position, tokenIndex, depth = position462, tokenIndex462, depth462
					if buffer[position] != rune('S') {
						goto l452
					}
					position++
				}
			l462:
				{
					position464, tokenIndex464, depth464 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l465
					}
					position++
					goto l464
				l465:
					position, tokenIndex, depth = position464, tokenIndex464, depth464
					if buffer[position] != rune('T') {
						goto l452
					}
					position++
				}
			l464:
				{
					position466, tokenIndex466, depth466 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l467
					}
					position++
					goto l466
				l467:
					position, tokenIndex, depth = position466, tokenIndex466, depth466
					if buffer[position] != rune('A') {
						goto l452
					}
					position++
				}
			l466:
				{
					position468, tokenIndex468, depth468 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l469
					}
					position++
					goto l468
				l469:
					position, tokenIndex, depth = position468, tokenIndex468, depth468
					if buffer[position] != rune('T') {
						goto l452
					}
					position++
				}
			l468:
				{
					position470, tokenIndex470, depth470 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l471
					}
					position++
					goto l470
				l471:
					position, tokenIndex, depth = position470, tokenIndex470, depth470
					if buffer[position] != rune('E') {
						goto l452
					}
					position++
				}
			l470:
				if !_rules[rulesp]() {
					goto l452
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l452
				}
				if !_rules[ruleAction16]() {
					goto l452
				}
				depth--
				add(ruleDropStateStmt, position453)
			}
			return true
		l452:
			position, tokenIndex, depth = position452, tokenIndex452, depth452
			return false
		},
		/* 23 Emitter <- <((ISTREAM / DSTREAM / RSTREAM) Action17)> */
		func() bool {
			position472, tokenIndex472, depth472 := position, tokenIndex, depth
			{
				position473 := position
				depth++
				{
					position474, tokenIndex474, depth474 := position, tokenIndex, depth
					if !_rules[ruleISTREAM]() {
						goto l475
					}
					goto l474
				l475:
					position, tokenIndex, depth = position474, tokenIndex474, depth474
					if !_rules[ruleDSTREAM]() {
						goto l476
					}
					goto l474
				l476:
					position, tokenIndex, depth = position474, tokenIndex474, depth474
					if !_rules[ruleRSTREAM]() {
						goto l472
					}
				}
			l474:
				if !_rules[ruleAction17]() {
					goto l472
				}
				depth--
				add(ruleEmitter, position473)
			}
			return true
		l472:
			position, tokenIndex, depth = position472, tokenIndex472, depth472
			return false
		},
		/* 24 Projections <- <(<(Projection sp (',' sp Projection)*)> Action18)> */
		func() bool {
			position477, tokenIndex477, depth477 := position, tokenIndex, depth
			{
				position478 := position
				depth++
				{
					position479 := position
					depth++
					if !_rules[ruleProjection]() {
						goto l477
					}
					if !_rules[rulesp]() {
						goto l477
					}
				l480:
					{
						position481, tokenIndex481, depth481 := position, tokenIndex, depth
						if buffer[position] != rune(',') {
							goto l481
						}
						position++
						if !_rules[rulesp]() {
							goto l481
						}
						if !_rules[ruleProjection]() {
							goto l481
						}
						goto l480
					l481:
						position, tokenIndex, depth = position481, tokenIndex481, depth481
					}
					depth--
					add(rulePegText, position479)
				}
				if !_rules[ruleAction18]() {
					goto l477
				}
				depth--
				add(ruleProjections, position478)
			}
			return true
		l477:
			position, tokenIndex, depth = position477, tokenIndex477, depth477
			return false
		},
		/* 25 Projection <- <(AliasExpression / Expression / Wildcard)> */
		func() bool {
			position482, tokenIndex482, depth482 := position, tokenIndex, depth
			{
				position483 := position
				depth++
				{
					position484, tokenIndex484, depth484 := position, tokenIndex, depth
					if !_rules[ruleAliasExpression]() {
						goto l485
					}
					goto l484
				l485:
					position, tokenIndex, depth = position484, tokenIndex484, depth484
					if !_rules[ruleExpression]() {
						goto l486
					}
					goto l484
				l486:
					position, tokenIndex, depth = position484, tokenIndex484, depth484
					if !_rules[ruleWildcard]() {
						goto l482
					}
				}
			l484:
				depth--
				add(ruleProjection, position483)
			}
			return true
		l482:
			position, tokenIndex, depth = position482, tokenIndex482, depth482
			return false
		},
		/* 26 AliasExpression <- <((Expression / Wildcard) sp (('a' / 'A') ('s' / 'S')) sp TargetIdentifier Action19)> */
		func() bool {
			position487, tokenIndex487, depth487 := position, tokenIndex, depth
			{
				position488 := position
				depth++
				{
					position489, tokenIndex489, depth489 := position, tokenIndex, depth
					if !_rules[ruleExpression]() {
						goto l490
					}
					goto l489
				l490:
					position, tokenIndex, depth = position489, tokenIndex489, depth489
					if !_rules[ruleWildcard]() {
						goto l487
					}
				}
			l489:
				if !_rules[rulesp]() {
					goto l487
				}
				{
					position491, tokenIndex491, depth491 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l492
					}
					position++
					goto l491
				l492:
					position, tokenIndex, depth = position491, tokenIndex491, depth491
					if buffer[position] != rune('A') {
						goto l487
					}
					position++
				}
			l491:
				{
					position493, tokenIndex493, depth493 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l494
					}
					position++
					goto l493
				l494:
					position, tokenIndex, depth = position493, tokenIndex493, depth493
					if buffer[position] != rune('S') {
						goto l487
					}
					position++
				}
			l493:
				if !_rules[rulesp]() {
					goto l487
				}
				if !_rules[ruleTargetIdentifier]() {
					goto l487
				}
				if !_rules[ruleAction19]() {
					goto l487
				}
				depth--
				add(ruleAliasExpression, position488)
			}
			return true
		l487:
			position, tokenIndex, depth = position487, tokenIndex487, depth487
			return false
		},
		/* 27 WindowedFrom <- <(<(('f' / 'F') ('r' / 'R') ('o' / 'O') ('m' / 'M') sp Relations sp)?> Action20)> */
		func() bool {
			position495, tokenIndex495, depth495 := position, tokenIndex, depth
			{
				position496 := position
				depth++
				{
					position497 := position
					depth++
					{
						position498, tokenIndex498, depth498 := position, tokenIndex, depth
						{
							position500, tokenIndex500, depth500 := position, tokenIndex, depth
							if buffer[position] != rune('f') {
								goto l501
							}
							position++
							goto l500
						l501:
							position, tokenIndex, depth = position500, tokenIndex500, depth500
							if buffer[position] != rune('F') {
								goto l498
							}
							position++
						}
					l500:
						{
							position502, tokenIndex502, depth502 := position, tokenIndex, depth
							if buffer[position] != rune('r') {
								goto l503
							}
							position++
							goto l502
						l503:
							position, tokenIndex, depth = position502, tokenIndex502, depth502
							if buffer[position] != rune('R') {
								goto l498
							}
							position++
						}
					l502:
						{
							position504, tokenIndex504, depth504 := position, tokenIndex, depth
							if buffer[position] != rune('o') {
								goto l505
							}
							position++
							goto l504
						l505:
							position, tokenIndex, depth = position504, tokenIndex504, depth504
							if buffer[position] != rune('O') {
								goto l498
							}
							position++
						}
					l504:
						{
							position506, tokenIndex506, depth506 := position, tokenIndex, depth
							if buffer[position] != rune('m') {
								goto l507
							}
							position++
							goto l506
						l507:
							position, tokenIndex, depth = position506, tokenIndex506, depth506
							if buffer[position] != rune('M') {
								goto l498
							}
							position++
						}
					l506:
						if !_rules[rulesp]() {
							goto l498
						}
						if !_rules[ruleRelations]() {
							goto l498
						}
						if !_rules[rulesp]() {
							goto l498
						}
						goto l499
					l498:
						position, tokenIndex, depth = position498, tokenIndex498, depth498
					}
				l499:
					depth--
					add(rulePegText, position497)
				}
				if !_rules[ruleAction20]() {
					goto l495
				}
				depth--
				add(ruleWindowedFrom, position496)
			}
			return true
		l495:
			position, tokenIndex, depth = position495, tokenIndex495, depth495
			return false
		},
		/* 28 Interval <- <(TimeInterval / TuplesInterval)> */
		func() bool {
			position508, tokenIndex508, depth508 := position, tokenIndex, depth
			{
				position509 := position
				depth++
				{
					position510, tokenIndex510, depth510 := position, tokenIndex, depth
					if !_rules[ruleTimeInterval]() {
						goto l511
					}
					goto l510
				l511:
					position, tokenIndex, depth = position510, tokenIndex510, depth510
					if !_rules[ruleTuplesInterval]() {
						goto l508
					}
				}
			l510:
				depth--
				add(ruleInterval, position509)
			}
			return true
		l508:
			position, tokenIndex, depth = position508, tokenIndex508, depth508
			return false
		},
		/* 29 TimeInterval <- <(NumericLiteral sp SECONDS Action21)> */
		func() bool {
			position512, tokenIndex512, depth512 := position, tokenIndex, depth
			{
				position513 := position
				depth++
				if !_rules[ruleNumericLiteral]() {
					goto l512
				}
				if !_rules[rulesp]() {
					goto l512
				}
				if !_rules[ruleSECONDS]() {
					goto l512
				}
				if !_rules[ruleAction21]() {
					goto l512
				}
				depth--
				add(ruleTimeInterval, position513)
			}
			return true
		l512:
			position, tokenIndex, depth = position512, tokenIndex512, depth512
			return false
		},
		/* 30 TuplesInterval <- <(NumericLiteral sp TUPLES Action22)> */
		func() bool {
			position514, tokenIndex514, depth514 := position, tokenIndex, depth
			{
				position515 := position
				depth++
				if !_rules[ruleNumericLiteral]() {
					goto l514
				}
				if !_rules[rulesp]() {
					goto l514
				}
				if !_rules[ruleTUPLES]() {
					goto l514
				}
				if !_rules[ruleAction22]() {
					goto l514
				}
				depth--
				add(ruleTuplesInterval, position515)
			}
			return true
		l514:
			position, tokenIndex, depth = position514, tokenIndex514, depth514
			return false
		},
		/* 31 Relations <- <(RelationLike sp (',' sp RelationLike)*)> */
		func() bool {
			position516, tokenIndex516, depth516 := position, tokenIndex, depth
			{
				position517 := position
				depth++
				if !_rules[ruleRelationLike]() {
					goto l516
				}
				if !_rules[rulesp]() {
					goto l516
				}
			l518:
				{
					position519, tokenIndex519, depth519 := position, tokenIndex, depth
					if buffer[position] != rune(',') {
						goto l519
					}
					position++
					if !_rules[rulesp]() {
						goto l519
					}
					if !_rules[ruleRelationLike]() {
						goto l519
					}
					goto l518
				l519:
					position, tokenIndex, depth = position519, tokenIndex519, depth519
				}
				depth--
				add(ruleRelations, position517)
			}
			return true
		l516:
			position, tokenIndex, depth = position516, tokenIndex516, depth516
			return false
		},
		/* 32 Filter <- <(<(('w' / 'W') ('h' / 'H') ('e' / 'E') ('r' / 'R') ('e' / 'E') sp Expression)?> Action23)> */
		func() bool {
			position520, tokenIndex520, depth520 := position, tokenIndex, depth
			{
				position521 := position
				depth++
				{
					position522 := position
					depth++
					{
						position523, tokenIndex523, depth523 := position, tokenIndex, depth
						{
							position525, tokenIndex525, depth525 := position, tokenIndex, depth
							if buffer[position] != rune('w') {
								goto l526
							}
							position++
							goto l525
						l526:
							position, tokenIndex, depth = position525, tokenIndex525, depth525
							if buffer[position] != rune('W') {
								goto l523
							}
							position++
						}
					l525:
						{
							position527, tokenIndex527, depth527 := position, tokenIndex, depth
							if buffer[position] != rune('h') {
								goto l528
							}
							position++
							goto l527
						l528:
							position, tokenIndex, depth = position527, tokenIndex527, depth527
							if buffer[position] != rune('H') {
								goto l523
							}
							position++
						}
					l527:
						{
							position529, tokenIndex529, depth529 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l530
							}
							position++
							goto l529
						l530:
							position, tokenIndex, depth = position529, tokenIndex529, depth529
							if buffer[position] != rune('E') {
								goto l523
							}
							position++
						}
					l529:
						{
							position531, tokenIndex531, depth531 := position, tokenIndex, depth
							if buffer[position] != rune('r') {
								goto l532
							}
							position++
							goto l531
						l532:
							position, tokenIndex, depth = position531, tokenIndex531, depth531
							if buffer[position] != rune('R') {
								goto l523
							}
							position++
						}
					l531:
						{
							position533, tokenIndex533, depth533 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l534
							}
							position++
							goto l533
						l534:
							position, tokenIndex, depth = position533, tokenIndex533, depth533
							if buffer[position] != rune('E') {
								goto l523
							}
							position++
						}
					l533:
						if !_rules[rulesp]() {
							goto l523
						}
						if !_rules[ruleExpression]() {
							goto l523
						}
						goto l524
					l523:
						position, tokenIndex, depth = position523, tokenIndex523, depth523
					}
				l524:
					depth--
					add(rulePegText, position522)
				}
				if !_rules[ruleAction23]() {
					goto l520
				}
				depth--
				add(ruleFilter, position521)
			}
			return true
		l520:
			position, tokenIndex, depth = position520, tokenIndex520, depth520
			return false
		},
		/* 33 Grouping <- <(<(('g' / 'G') ('r' / 'R') ('o' / 'O') ('u' / 'U') ('p' / 'P') sp (('b' / 'B') ('y' / 'Y')) sp GroupList)?> Action24)> */
		func() bool {
			position535, tokenIndex535, depth535 := position, tokenIndex, depth
			{
				position536 := position
				depth++
				{
					position537 := position
					depth++
					{
						position538, tokenIndex538, depth538 := position, tokenIndex, depth
						{
							position540, tokenIndex540, depth540 := position, tokenIndex, depth
							if buffer[position] != rune('g') {
								goto l541
							}
							position++
							goto l540
						l541:
							position, tokenIndex, depth = position540, tokenIndex540, depth540
							if buffer[position] != rune('G') {
								goto l538
							}
							position++
						}
					l540:
						{
							position542, tokenIndex542, depth542 := position, tokenIndex, depth
							if buffer[position] != rune('r') {
								goto l543
							}
							position++
							goto l542
						l543:
							position, tokenIndex, depth = position542, tokenIndex542, depth542
							if buffer[position] != rune('R') {
								goto l538
							}
							position++
						}
					l542:
						{
							position544, tokenIndex544, depth544 := position, tokenIndex, depth
							if buffer[position] != rune('o') {
								goto l545
							}
							position++
							goto l544
						l545:
							position, tokenIndex, depth = position544, tokenIndex544, depth544
							if buffer[position] != rune('O') {
								goto l538
							}
							position++
						}
					l544:
						{
							position546, tokenIndex546, depth546 := position, tokenIndex, depth
							if buffer[position] != rune('u') {
								goto l547
							}
							position++
							goto l546
						l547:
							position, tokenIndex, depth = position546, tokenIndex546, depth546
							if buffer[position] != rune('U') {
								goto l538
							}
							position++
						}
					l546:
						{
							position548, tokenIndex548, depth548 := position, tokenIndex, depth
							if buffer[position] != rune('p') {
								goto l549
							}
							position++
							goto l548
						l549:
							position, tokenIndex, depth = position548, tokenIndex548, depth548
							if buffer[position] != rune('P') {
								goto l538
							}
							position++
						}
					l548:
						if !_rules[rulesp]() {
							goto l538
						}
						{
							position550, tokenIndex550, depth550 := position, tokenIndex, depth
							if buffer[position] != rune('b') {
								goto l551
							}
							position++
							goto l550
						l551:
							position, tokenIndex, depth = position550, tokenIndex550, depth550
							if buffer[position] != rune('B') {
								goto l538
							}
							position++
						}
					l550:
						{
							position552, tokenIndex552, depth552 := position, tokenIndex, depth
							if buffer[position] != rune('y') {
								goto l553
							}
							position++
							goto l552
						l553:
							position, tokenIndex, depth = position552, tokenIndex552, depth552
							if buffer[position] != rune('Y') {
								goto l538
							}
							position++
						}
					l552:
						if !_rules[rulesp]() {
							goto l538
						}
						if !_rules[ruleGroupList]() {
							goto l538
						}
						goto l539
					l538:
						position, tokenIndex, depth = position538, tokenIndex538, depth538
					}
				l539:
					depth--
					add(rulePegText, position537)
				}
				if !_rules[ruleAction24]() {
					goto l535
				}
				depth--
				add(ruleGrouping, position536)
			}
			return true
		l535:
			position, tokenIndex, depth = position535, tokenIndex535, depth535
			return false
		},
		/* 34 GroupList <- <(Expression sp (',' sp Expression)*)> */
		func() bool {
			position554, tokenIndex554, depth554 := position, tokenIndex, depth
			{
				position555 := position
				depth++
				if !_rules[ruleExpression]() {
					goto l554
				}
				if !_rules[rulesp]() {
					goto l554
				}
			l556:
				{
					position557, tokenIndex557, depth557 := position, tokenIndex, depth
					if buffer[position] != rune(',') {
						goto l557
					}
					position++
					if !_rules[rulesp]() {
						goto l557
					}
					if !_rules[ruleExpression]() {
						goto l557
					}
					goto l556
				l557:
					position, tokenIndex, depth = position557, tokenIndex557, depth557
				}
				depth--
				add(ruleGroupList, position555)
			}
			return true
		l554:
			position, tokenIndex, depth = position554, tokenIndex554, depth554
			return false
		},
		/* 35 Having <- <(<(('h' / 'H') ('a' / 'A') ('v' / 'V') ('i' / 'I') ('n' / 'N') ('g' / 'G') sp Expression)?> Action25)> */
		func() bool {
			position558, tokenIndex558, depth558 := position, tokenIndex, depth
			{
				position559 := position
				depth++
				{
					position560 := position
					depth++
					{
						position561, tokenIndex561, depth561 := position, tokenIndex, depth
						{
							position563, tokenIndex563, depth563 := position, tokenIndex, depth
							if buffer[position] != rune('h') {
								goto l564
							}
							position++
							goto l563
						l564:
							position, tokenIndex, depth = position563, tokenIndex563, depth563
							if buffer[position] != rune('H') {
								goto l561
							}
							position++
						}
					l563:
						{
							position565, tokenIndex565, depth565 := position, tokenIndex, depth
							if buffer[position] != rune('a') {
								goto l566
							}
							position++
							goto l565
						l566:
							position, tokenIndex, depth = position565, tokenIndex565, depth565
							if buffer[position] != rune('A') {
								goto l561
							}
							position++
						}
					l565:
						{
							position567, tokenIndex567, depth567 := position, tokenIndex, depth
							if buffer[position] != rune('v') {
								goto l568
							}
							position++
							goto l567
						l568:
							position, tokenIndex, depth = position567, tokenIndex567, depth567
							if buffer[position] != rune('V') {
								goto l561
							}
							position++
						}
					l567:
						{
							position569, tokenIndex569, depth569 := position, tokenIndex, depth
							if buffer[position] != rune('i') {
								goto l570
							}
							position++
							goto l569
						l570:
							position, tokenIndex, depth = position569, tokenIndex569, depth569
							if buffer[position] != rune('I') {
								goto l561
							}
							position++
						}
					l569:
						{
							position571, tokenIndex571, depth571 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l572
							}
							position++
							goto l571
						l572:
							position, tokenIndex, depth = position571, tokenIndex571, depth571
							if buffer[position] != rune('N') {
								goto l561
							}
							position++
						}
					l571:
						{
							position573, tokenIndex573, depth573 := position, tokenIndex, depth
							if buffer[position] != rune('g') {
								goto l574
							}
							position++
							goto l573
						l574:
							position, tokenIndex, depth = position573, tokenIndex573, depth573
							if buffer[position] != rune('G') {
								goto l561
							}
							position++
						}
					l573:
						if !_rules[rulesp]() {
							goto l561
						}
						if !_rules[ruleExpression]() {
							goto l561
						}
						goto l562
					l561:
						position, tokenIndex, depth = position561, tokenIndex561, depth561
					}
				l562:
					depth--
					add(rulePegText, position560)
				}
				if !_rules[ruleAction25]() {
					goto l558
				}
				depth--
				add(ruleHaving, position559)
			}
			return true
		l558:
			position, tokenIndex, depth = position558, tokenIndex558, depth558
			return false
		},
		/* 36 RelationLike <- <(AliasedStreamWindow / (StreamWindow Action26))> */
		func() bool {
			position575, tokenIndex575, depth575 := position, tokenIndex, depth
			{
				position576 := position
				depth++
				{
					position577, tokenIndex577, depth577 := position, tokenIndex, depth
					if !_rules[ruleAliasedStreamWindow]() {
						goto l578
					}
					goto l577
				l578:
					position, tokenIndex, depth = position577, tokenIndex577, depth577
					if !_rules[ruleStreamWindow]() {
						goto l575
					}
					if !_rules[ruleAction26]() {
						goto l575
					}
				}
			l577:
				depth--
				add(ruleRelationLike, position576)
			}
			return true
		l575:
			position, tokenIndex, depth = position575, tokenIndex575, depth575
			return false
		},
		/* 37 AliasedStreamWindow <- <(StreamWindow sp (('a' / 'A') ('s' / 'S')) sp Identifier Action27)> */
		func() bool {
			position579, tokenIndex579, depth579 := position, tokenIndex, depth
			{
				position580 := position
				depth++
				if !_rules[ruleStreamWindow]() {
					goto l579
				}
				if !_rules[rulesp]() {
					goto l579
				}
				{
					position581, tokenIndex581, depth581 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l582
					}
					position++
					goto l581
				l582:
					position, tokenIndex, depth = position581, tokenIndex581, depth581
					if buffer[position] != rune('A') {
						goto l579
					}
					position++
				}
			l581:
				{
					position583, tokenIndex583, depth583 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l584
					}
					position++
					goto l583
				l584:
					position, tokenIndex, depth = position583, tokenIndex583, depth583
					if buffer[position] != rune('S') {
						goto l579
					}
					position++
				}
			l583:
				if !_rules[rulesp]() {
					goto l579
				}
				if !_rules[ruleIdentifier]() {
					goto l579
				}
				if !_rules[ruleAction27]() {
					goto l579
				}
				depth--
				add(ruleAliasedStreamWindow, position580)
			}
			return true
		l579:
			position, tokenIndex, depth = position579, tokenIndex579, depth579
			return false
		},
		/* 38 StreamWindow <- <(StreamLike sp '[' sp (('r' / 'R') ('a' / 'A') ('n' / 'N') ('g' / 'G') ('e' / 'E')) sp Interval sp ']' Action28)> */
		func() bool {
			position585, tokenIndex585, depth585 := position, tokenIndex, depth
			{
				position586 := position
				depth++
				if !_rules[ruleStreamLike]() {
					goto l585
				}
				if !_rules[rulesp]() {
					goto l585
				}
				if buffer[position] != rune('[') {
					goto l585
				}
				position++
				if !_rules[rulesp]() {
					goto l585
				}
				{
					position587, tokenIndex587, depth587 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l588
					}
					position++
					goto l587
				l588:
					position, tokenIndex, depth = position587, tokenIndex587, depth587
					if buffer[position] != rune('R') {
						goto l585
					}
					position++
				}
			l587:
				{
					position589, tokenIndex589, depth589 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l590
					}
					position++
					goto l589
				l590:
					position, tokenIndex, depth = position589, tokenIndex589, depth589
					if buffer[position] != rune('A') {
						goto l585
					}
					position++
				}
			l589:
				{
					position591, tokenIndex591, depth591 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l592
					}
					position++
					goto l591
				l592:
					position, tokenIndex, depth = position591, tokenIndex591, depth591
					if buffer[position] != rune('N') {
						goto l585
					}
					position++
				}
			l591:
				{
					position593, tokenIndex593, depth593 := position, tokenIndex, depth
					if buffer[position] != rune('g') {
						goto l594
					}
					position++
					goto l593
				l594:
					position, tokenIndex, depth = position593, tokenIndex593, depth593
					if buffer[position] != rune('G') {
						goto l585
					}
					position++
				}
			l593:
				{
					position595, tokenIndex595, depth595 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l596
					}
					position++
					goto l595
				l596:
					position, tokenIndex, depth = position595, tokenIndex595, depth595
					if buffer[position] != rune('E') {
						goto l585
					}
					position++
				}
			l595:
				if !_rules[rulesp]() {
					goto l585
				}
				if !_rules[ruleInterval]() {
					goto l585
				}
				if !_rules[rulesp]() {
					goto l585
				}
				if buffer[position] != rune(']') {
					goto l585
				}
				position++
				if !_rules[ruleAction28]() {
					goto l585
				}
				depth--
				add(ruleStreamWindow, position586)
			}
			return true
		l585:
			position, tokenIndex, depth = position585, tokenIndex585, depth585
			return false
		},
		/* 39 StreamLike <- <(UDSFFuncApp / Stream)> */
		func() bool {
			position597, tokenIndex597, depth597 := position, tokenIndex, depth
			{
				position598 := position
				depth++
				{
					position599, tokenIndex599, depth599 := position, tokenIndex, depth
					if !_rules[ruleUDSFFuncApp]() {
						goto l600
					}
					goto l599
				l600:
					position, tokenIndex, depth = position599, tokenIndex599, depth599
					if !_rules[ruleStream]() {
						goto l597
					}
				}
			l599:
				depth--
				add(ruleStreamLike, position598)
			}
			return true
		l597:
			position, tokenIndex, depth = position597, tokenIndex597, depth597
			return false
		},
		/* 40 UDSFFuncApp <- <(FuncApp Action29)> */
		func() bool {
			position601, tokenIndex601, depth601 := position, tokenIndex, depth
			{
				position602 := position
				depth++
				if !_rules[ruleFuncApp]() {
					goto l601
				}
				if !_rules[ruleAction29]() {
					goto l601
				}
				depth--
				add(ruleUDSFFuncApp, position602)
			}
			return true
		l601:
			position, tokenIndex, depth = position601, tokenIndex601, depth601
			return false
		},
		/* 41 SourceSinkSpecs <- <(<(('w' / 'W') ('i' / 'I') ('t' / 'T') ('h' / 'H') sp SourceSinkParam sp (',' sp SourceSinkParam)*)?> Action30)> */
		func() bool {
			position603, tokenIndex603, depth603 := position, tokenIndex, depth
			{
				position604 := position
				depth++
				{
					position605 := position
					depth++
					{
						position606, tokenIndex606, depth606 := position, tokenIndex, depth
						{
							position608, tokenIndex608, depth608 := position, tokenIndex, depth
							if buffer[position] != rune('w') {
								goto l609
							}
							position++
							goto l608
						l609:
							position, tokenIndex, depth = position608, tokenIndex608, depth608
							if buffer[position] != rune('W') {
								goto l606
							}
							position++
						}
					l608:
						{
							position610, tokenIndex610, depth610 := position, tokenIndex, depth
							if buffer[position] != rune('i') {
								goto l611
							}
							position++
							goto l610
						l611:
							position, tokenIndex, depth = position610, tokenIndex610, depth610
							if buffer[position] != rune('I') {
								goto l606
							}
							position++
						}
					l610:
						{
							position612, tokenIndex612, depth612 := position, tokenIndex, depth
							if buffer[position] != rune('t') {
								goto l613
							}
							position++
							goto l612
						l613:
							position, tokenIndex, depth = position612, tokenIndex612, depth612
							if buffer[position] != rune('T') {
								goto l606
							}
							position++
						}
					l612:
						{
							position614, tokenIndex614, depth614 := position, tokenIndex, depth
							if buffer[position] != rune('h') {
								goto l615
							}
							position++
							goto l614
						l615:
							position, tokenIndex, depth = position614, tokenIndex614, depth614
							if buffer[position] != rune('H') {
								goto l606
							}
							position++
						}
					l614:
						if !_rules[rulesp]() {
							goto l606
						}
						if !_rules[ruleSourceSinkParam]() {
							goto l606
						}
						if !_rules[rulesp]() {
							goto l606
						}
					l616:
						{
							position617, tokenIndex617, depth617 := position, tokenIndex, depth
							if buffer[position] != rune(',') {
								goto l617
							}
							position++
							if !_rules[rulesp]() {
								goto l617
							}
							if !_rules[ruleSourceSinkParam]() {
								goto l617
							}
							goto l616
						l617:
							position, tokenIndex, depth = position617, tokenIndex617, depth617
						}
						goto l607
					l606:
						position, tokenIndex, depth = position606, tokenIndex606, depth606
					}
				l607:
					depth--
					add(rulePegText, position605)
				}
				if !_rules[ruleAction30]() {
					goto l603
				}
				depth--
				add(ruleSourceSinkSpecs, position604)
			}
			return true
		l603:
			position, tokenIndex, depth = position603, tokenIndex603, depth603
			return false
		},
		/* 42 UpdateSourceSinkSpecs <- <(<(('s' / 'S') ('e' / 'E') ('t' / 'T') sp SourceSinkParam sp (',' sp SourceSinkParam)*)> Action31)> */
		func() bool {
			position618, tokenIndex618, depth618 := position, tokenIndex, depth
			{
				position619 := position
				depth++
				{
					position620 := position
					depth++
					{
						position621, tokenIndex621, depth621 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l622
						}
						position++
						goto l621
					l622:
						position, tokenIndex, depth = position621, tokenIndex621, depth621
						if buffer[position] != rune('S') {
							goto l618
						}
						position++
					}
				l621:
					{
						position623, tokenIndex623, depth623 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l624
						}
						position++
						goto l623
					l624:
						position, tokenIndex, depth = position623, tokenIndex623, depth623
						if buffer[position] != rune('E') {
							goto l618
						}
						position++
					}
				l623:
					{
						position625, tokenIndex625, depth625 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l626
						}
						position++
						goto l625
					l626:
						position, tokenIndex, depth = position625, tokenIndex625, depth625
						if buffer[position] != rune('T') {
							goto l618
						}
						position++
					}
				l625:
					if !_rules[rulesp]() {
						goto l618
					}
					if !_rules[ruleSourceSinkParam]() {
						goto l618
					}
					if !_rules[rulesp]() {
						goto l618
					}
				l627:
					{
						position628, tokenIndex628, depth628 := position, tokenIndex, depth
						if buffer[position] != rune(',') {
							goto l628
						}
						position++
						if !_rules[rulesp]() {
							goto l628
						}
						if !_rules[ruleSourceSinkParam]() {
							goto l628
						}
						goto l627
					l628:
						position, tokenIndex, depth = position628, tokenIndex628, depth628
					}
					depth--
					add(rulePegText, position620)
				}
				if !_rules[ruleAction31]() {
					goto l618
				}
				depth--
				add(ruleUpdateSourceSinkSpecs, position619)
			}
			return true
		l618:
			position, tokenIndex, depth = position618, tokenIndex618, depth618
			return false
		},
		/* 43 SourceSinkParam <- <(SourceSinkParamKey '=' SourceSinkParamVal Action32)> */
		func() bool {
			position629, tokenIndex629, depth629 := position, tokenIndex, depth
			{
				position630 := position
				depth++
				if !_rules[ruleSourceSinkParamKey]() {
					goto l629
				}
				if buffer[position] != rune('=') {
					goto l629
				}
				position++
				if !_rules[ruleSourceSinkParamVal]() {
					goto l629
				}
				if !_rules[ruleAction32]() {
					goto l629
				}
				depth--
				add(ruleSourceSinkParam, position630)
			}
			return true
		l629:
			position, tokenIndex, depth = position629, tokenIndex629, depth629
			return false
		},
		/* 44 SourceSinkParamVal <- <(BooleanLiteral / Literal)> */
		func() bool {
			position631, tokenIndex631, depth631 := position, tokenIndex, depth
			{
				position632 := position
				depth++
				{
					position633, tokenIndex633, depth633 := position, tokenIndex, depth
					if !_rules[ruleBooleanLiteral]() {
						goto l634
					}
					goto l633
				l634:
					position, tokenIndex, depth = position633, tokenIndex633, depth633
					if !_rules[ruleLiteral]() {
						goto l631
					}
				}
			l633:
				depth--
				add(ruleSourceSinkParamVal, position632)
			}
			return true
		l631:
			position, tokenIndex, depth = position631, tokenIndex631, depth631
			return false
		},
		/* 45 PausedOpt <- <(<(Paused / Unpaused)?> Action33)> */
		func() bool {
			position635, tokenIndex635, depth635 := position, tokenIndex, depth
			{
				position636 := position
				depth++
				{
					position637 := position
					depth++
					{
						position638, tokenIndex638, depth638 := position, tokenIndex, depth
						{
							position640, tokenIndex640, depth640 := position, tokenIndex, depth
							if !_rules[rulePaused]() {
								goto l641
							}
							goto l640
						l641:
							position, tokenIndex, depth = position640, tokenIndex640, depth640
							if !_rules[ruleUnpaused]() {
								goto l638
							}
						}
					l640:
						goto l639
					l638:
						position, tokenIndex, depth = position638, tokenIndex638, depth638
					}
				l639:
					depth--
					add(rulePegText, position637)
				}
				if !_rules[ruleAction33]() {
					goto l635
				}
				depth--
				add(rulePausedOpt, position636)
			}
			return true
		l635:
			position, tokenIndex, depth = position635, tokenIndex635, depth635
			return false
		},
		/* 46 Expression <- <orExpr> */
		func() bool {
			position642, tokenIndex642, depth642 := position, tokenIndex, depth
			{
				position643 := position
				depth++
				if !_rules[ruleorExpr]() {
					goto l642
				}
				depth--
				add(ruleExpression, position643)
			}
			return true
		l642:
			position, tokenIndex, depth = position642, tokenIndex642, depth642
			return false
		},
		/* 47 orExpr <- <(<(andExpr sp (Or sp andExpr)?)> Action34)> */
		func() bool {
			position644, tokenIndex644, depth644 := position, tokenIndex, depth
			{
				position645 := position
				depth++
				{
					position646 := position
					depth++
					if !_rules[ruleandExpr]() {
						goto l644
					}
					if !_rules[rulesp]() {
						goto l644
					}
					{
						position647, tokenIndex647, depth647 := position, tokenIndex, depth
						if !_rules[ruleOr]() {
							goto l647
						}
						if !_rules[rulesp]() {
							goto l647
						}
						if !_rules[ruleandExpr]() {
							goto l647
						}
						goto l648
					l647:
						position, tokenIndex, depth = position647, tokenIndex647, depth647
					}
				l648:
					depth--
					add(rulePegText, position646)
				}
				if !_rules[ruleAction34]() {
					goto l644
				}
				depth--
				add(ruleorExpr, position645)
			}
			return true
		l644:
			position, tokenIndex, depth = position644, tokenIndex644, depth644
			return false
		},
		/* 48 andExpr <- <(<(notExpr sp (And sp notExpr)?)> Action35)> */
		func() bool {
			position649, tokenIndex649, depth649 := position, tokenIndex, depth
			{
				position650 := position
				depth++
				{
					position651 := position
					depth++
					if !_rules[rulenotExpr]() {
						goto l649
					}
					if !_rules[rulesp]() {
						goto l649
					}
					{
						position652, tokenIndex652, depth652 := position, tokenIndex, depth
						if !_rules[ruleAnd]() {
							goto l652
						}
						if !_rules[rulesp]() {
							goto l652
						}
						if !_rules[rulenotExpr]() {
							goto l652
						}
						goto l653
					l652:
						position, tokenIndex, depth = position652, tokenIndex652, depth652
					}
				l653:
					depth--
					add(rulePegText, position651)
				}
				if !_rules[ruleAction35]() {
					goto l649
				}
				depth--
				add(ruleandExpr, position650)
			}
			return true
		l649:
			position, tokenIndex, depth = position649, tokenIndex649, depth649
			return false
		},
		/* 49 notExpr <- <(<((Not sp)? comparisonExpr)> Action36)> */
		func() bool {
			position654, tokenIndex654, depth654 := position, tokenIndex, depth
			{
				position655 := position
				depth++
				{
					position656 := position
					depth++
					{
						position657, tokenIndex657, depth657 := position, tokenIndex, depth
						if !_rules[ruleNot]() {
							goto l657
						}
						if !_rules[rulesp]() {
							goto l657
						}
						goto l658
					l657:
						position, tokenIndex, depth = position657, tokenIndex657, depth657
					}
				l658:
					if !_rules[rulecomparisonExpr]() {
						goto l654
					}
					depth--
					add(rulePegText, position656)
				}
				if !_rules[ruleAction36]() {
					goto l654
				}
				depth--
				add(rulenotExpr, position655)
			}
			return true
		l654:
			position, tokenIndex, depth = position654, tokenIndex654, depth654
			return false
		},
		/* 50 comparisonExpr <- <(<(otherOpExpr sp (ComparisonOp sp otherOpExpr)?)> Action37)> */
		func() bool {
			position659, tokenIndex659, depth659 := position, tokenIndex, depth
			{
				position660 := position
				depth++
				{
					position661 := position
					depth++
					if !_rules[ruleotherOpExpr]() {
						goto l659
					}
					if !_rules[rulesp]() {
						goto l659
					}
					{
						position662, tokenIndex662, depth662 := position, tokenIndex, depth
						if !_rules[ruleComparisonOp]() {
							goto l662
						}
						if !_rules[rulesp]() {
							goto l662
						}
						if !_rules[ruleotherOpExpr]() {
							goto l662
						}
						goto l663
					l662:
						position, tokenIndex, depth = position662, tokenIndex662, depth662
					}
				l663:
					depth--
					add(rulePegText, position661)
				}
				if !_rules[ruleAction37]() {
					goto l659
				}
				depth--
				add(rulecomparisonExpr, position660)
			}
			return true
		l659:
			position, tokenIndex, depth = position659, tokenIndex659, depth659
			return false
		},
		/* 51 otherOpExpr <- <(<(isExpr sp (OtherOp sp isExpr sp)*)> Action38)> */
		func() bool {
			position664, tokenIndex664, depth664 := position, tokenIndex, depth
			{
				position665 := position
				depth++
				{
					position666 := position
					depth++
					if !_rules[ruleisExpr]() {
						goto l664
					}
					if !_rules[rulesp]() {
						goto l664
					}
				l667:
					{
						position668, tokenIndex668, depth668 := position, tokenIndex, depth
						if !_rules[ruleOtherOp]() {
							goto l668
						}
						if !_rules[rulesp]() {
							goto l668
						}
						if !_rules[ruleisExpr]() {
							goto l668
						}
						if !_rules[rulesp]() {
							goto l668
						}
						goto l667
					l668:
						position, tokenIndex, depth = position668, tokenIndex668, depth668
					}
					depth--
					add(rulePegText, position666)
				}
				if !_rules[ruleAction38]() {
					goto l664
				}
				depth--
				add(ruleotherOpExpr, position665)
			}
			return true
		l664:
			position, tokenIndex, depth = position664, tokenIndex664, depth664
			return false
		},
		/* 52 isExpr <- <(<(termExpr sp (IsOp sp NullLiteral)?)> Action39)> */
		func() bool {
			position669, tokenIndex669, depth669 := position, tokenIndex, depth
			{
				position670 := position
				depth++
				{
					position671 := position
					depth++
					if !_rules[ruletermExpr]() {
						goto l669
					}
					if !_rules[rulesp]() {
						goto l669
					}
					{
						position672, tokenIndex672, depth672 := position, tokenIndex, depth
						if !_rules[ruleIsOp]() {
							goto l672
						}
						if !_rules[rulesp]() {
							goto l672
						}
						if !_rules[ruleNullLiteral]() {
							goto l672
						}
						goto l673
					l672:
						position, tokenIndex, depth = position672, tokenIndex672, depth672
					}
				l673:
					depth--
					add(rulePegText, position671)
				}
				if !_rules[ruleAction39]() {
					goto l669
				}
				depth--
				add(ruleisExpr, position670)
			}
			return true
		l669:
			position, tokenIndex, depth = position669, tokenIndex669, depth669
			return false
		},
		/* 53 termExpr <- <(<(productExpr sp (PlusMinusOp sp productExpr sp)*)> Action40)> */
		func() bool {
			position674, tokenIndex674, depth674 := position, tokenIndex, depth
			{
				position675 := position
				depth++
				{
					position676 := position
					depth++
					if !_rules[ruleproductExpr]() {
						goto l674
					}
					if !_rules[rulesp]() {
						goto l674
					}
				l677:
					{
						position678, tokenIndex678, depth678 := position, tokenIndex, depth
						if !_rules[rulePlusMinusOp]() {
							goto l678
						}
						if !_rules[rulesp]() {
							goto l678
						}
						if !_rules[ruleproductExpr]() {
							goto l678
						}
						if !_rules[rulesp]() {
							goto l678
						}
						goto l677
					l678:
						position, tokenIndex, depth = position678, tokenIndex678, depth678
					}
					depth--
					add(rulePegText, position676)
				}
				if !_rules[ruleAction40]() {
					goto l674
				}
				depth--
				add(ruletermExpr, position675)
			}
			return true
		l674:
			position, tokenIndex, depth = position674, tokenIndex674, depth674
			return false
		},
		/* 54 productExpr <- <(<(minusExpr sp (MultDivOp sp minusExpr sp)*)> Action41)> */
		func() bool {
			position679, tokenIndex679, depth679 := position, tokenIndex, depth
			{
				position680 := position
				depth++
				{
					position681 := position
					depth++
					if !_rules[ruleminusExpr]() {
						goto l679
					}
					if !_rules[rulesp]() {
						goto l679
					}
				l682:
					{
						position683, tokenIndex683, depth683 := position, tokenIndex, depth
						if !_rules[ruleMultDivOp]() {
							goto l683
						}
						if !_rules[rulesp]() {
							goto l683
						}
						if !_rules[ruleminusExpr]() {
							goto l683
						}
						if !_rules[rulesp]() {
							goto l683
						}
						goto l682
					l683:
						position, tokenIndex, depth = position683, tokenIndex683, depth683
					}
					depth--
					add(rulePegText, position681)
				}
				if !_rules[ruleAction41]() {
					goto l679
				}
				depth--
				add(ruleproductExpr, position680)
			}
			return true
		l679:
			position, tokenIndex, depth = position679, tokenIndex679, depth679
			return false
		},
		/* 55 minusExpr <- <(<((UnaryMinus sp)? castExpr)> Action42)> */
		func() bool {
			position684, tokenIndex684, depth684 := position, tokenIndex, depth
			{
				position685 := position
				depth++
				{
					position686 := position
					depth++
					{
						position687, tokenIndex687, depth687 := position, tokenIndex, depth
						if !_rules[ruleUnaryMinus]() {
							goto l687
						}
						if !_rules[rulesp]() {
							goto l687
						}
						goto l688
					l687:
						position, tokenIndex, depth = position687, tokenIndex687, depth687
					}
				l688:
					if !_rules[rulecastExpr]() {
						goto l684
					}
					depth--
					add(rulePegText, position686)
				}
				if !_rules[ruleAction42]() {
					goto l684
				}
				depth--
				add(ruleminusExpr, position685)
			}
			return true
		l684:
			position, tokenIndex, depth = position684, tokenIndex684, depth684
			return false
		},
		/* 56 castExpr <- <(<(baseExpr (sp (':' ':') sp Type sp)?)> Action43)> */
		func() bool {
			position689, tokenIndex689, depth689 := position, tokenIndex, depth
			{
				position690 := position
				depth++
				{
					position691 := position
					depth++
					if !_rules[rulebaseExpr]() {
						goto l689
					}
					{
						position692, tokenIndex692, depth692 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l692
						}
						if buffer[position] != rune(':') {
							goto l692
						}
						position++
						if buffer[position] != rune(':') {
							goto l692
						}
						position++
						if !_rules[rulesp]() {
							goto l692
						}
						if !_rules[ruleType]() {
							goto l692
						}
						if !_rules[rulesp]() {
							goto l692
						}
						goto l693
					l692:
						position, tokenIndex, depth = position692, tokenIndex692, depth692
					}
				l693:
					depth--
					add(rulePegText, position691)
				}
				if !_rules[ruleAction43]() {
					goto l689
				}
				depth--
				add(rulecastExpr, position690)
			}
			return true
		l689:
			position, tokenIndex, depth = position689, tokenIndex689, depth689
			return false
		},
		/* 57 baseExpr <- <(('(' sp Expression sp ')') / BooleanLiteral / NullLiteral / RowMeta / FuncTypeCast / FuncApp / RowValue / Literal)> */
		func() bool {
			position694, tokenIndex694, depth694 := position, tokenIndex, depth
			{
				position695 := position
				depth++
				{
					position696, tokenIndex696, depth696 := position, tokenIndex, depth
					if buffer[position] != rune('(') {
						goto l697
					}
					position++
					if !_rules[rulesp]() {
						goto l697
					}
					if !_rules[ruleExpression]() {
						goto l697
					}
					if !_rules[rulesp]() {
						goto l697
					}
					if buffer[position] != rune(')') {
						goto l697
					}
					position++
					goto l696
				l697:
					position, tokenIndex, depth = position696, tokenIndex696, depth696
					if !_rules[ruleBooleanLiteral]() {
						goto l698
					}
					goto l696
				l698:
					position, tokenIndex, depth = position696, tokenIndex696, depth696
					if !_rules[ruleNullLiteral]() {
						goto l699
					}
					goto l696
				l699:
					position, tokenIndex, depth = position696, tokenIndex696, depth696
					if !_rules[ruleRowMeta]() {
						goto l700
					}
					goto l696
				l700:
					position, tokenIndex, depth = position696, tokenIndex696, depth696
					if !_rules[ruleFuncTypeCast]() {
						goto l701
					}
					goto l696
				l701:
					position, tokenIndex, depth = position696, tokenIndex696, depth696
					if !_rules[ruleFuncApp]() {
						goto l702
					}
					goto l696
				l702:
					position, tokenIndex, depth = position696, tokenIndex696, depth696
					if !_rules[ruleRowValue]() {
						goto l703
					}
					goto l696
				l703:
					position, tokenIndex, depth = position696, tokenIndex696, depth696
					if !_rules[ruleLiteral]() {
						goto l694
					}
				}
			l696:
				depth--
				add(rulebaseExpr, position695)
			}
			return true
		l694:
			position, tokenIndex, depth = position694, tokenIndex694, depth694
			return false
		},
		/* 58 FuncTypeCast <- <(<(('c' / 'C') ('a' / 'A') ('s' / 'S') ('t' / 'T') sp '(' sp Expression sp (('a' / 'A') ('s' / 'S')) sp Type sp ')')> Action44)> */
		func() bool {
			position704, tokenIndex704, depth704 := position, tokenIndex, depth
			{
				position705 := position
				depth++
				{
					position706 := position
					depth++
					{
						position707, tokenIndex707, depth707 := position, tokenIndex, depth
						if buffer[position] != rune('c') {
							goto l708
						}
						position++
						goto l707
					l708:
						position, tokenIndex, depth = position707, tokenIndex707, depth707
						if buffer[position] != rune('C') {
							goto l704
						}
						position++
					}
				l707:
					{
						position709, tokenIndex709, depth709 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l710
						}
						position++
						goto l709
					l710:
						position, tokenIndex, depth = position709, tokenIndex709, depth709
						if buffer[position] != rune('A') {
							goto l704
						}
						position++
					}
				l709:
					{
						position711, tokenIndex711, depth711 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l712
						}
						position++
						goto l711
					l712:
						position, tokenIndex, depth = position711, tokenIndex711, depth711
						if buffer[position] != rune('S') {
							goto l704
						}
						position++
					}
				l711:
					{
						position713, tokenIndex713, depth713 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l714
						}
						position++
						goto l713
					l714:
						position, tokenIndex, depth = position713, tokenIndex713, depth713
						if buffer[position] != rune('T') {
							goto l704
						}
						position++
					}
				l713:
					if !_rules[rulesp]() {
						goto l704
					}
					if buffer[position] != rune('(') {
						goto l704
					}
					position++
					if !_rules[rulesp]() {
						goto l704
					}
					if !_rules[ruleExpression]() {
						goto l704
					}
					if !_rules[rulesp]() {
						goto l704
					}
					{
						position715, tokenIndex715, depth715 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l716
						}
						position++
						goto l715
					l716:
						position, tokenIndex, depth = position715, tokenIndex715, depth715
						if buffer[position] != rune('A') {
							goto l704
						}
						position++
					}
				l715:
					{
						position717, tokenIndex717, depth717 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l718
						}
						position++
						goto l717
					l718:
						position, tokenIndex, depth = position717, tokenIndex717, depth717
						if buffer[position] != rune('S') {
							goto l704
						}
						position++
					}
				l717:
					if !_rules[rulesp]() {
						goto l704
					}
					if !_rules[ruleType]() {
						goto l704
					}
					if !_rules[rulesp]() {
						goto l704
					}
					if buffer[position] != rune(')') {
						goto l704
					}
					position++
					depth--
					add(rulePegText, position706)
				}
				if !_rules[ruleAction44]() {
					goto l704
				}
				depth--
				add(ruleFuncTypeCast, position705)
			}
			return true
		l704:
			position, tokenIndex, depth = position704, tokenIndex704, depth704
			return false
		},
		/* 59 FuncApp <- <(Function sp '(' sp FuncParams sp ')' Action45)> */
		func() bool {
			position719, tokenIndex719, depth719 := position, tokenIndex, depth
			{
				position720 := position
				depth++
				if !_rules[ruleFunction]() {
					goto l719
				}
				if !_rules[rulesp]() {
					goto l719
				}
				if buffer[position] != rune('(') {
					goto l719
				}
				position++
				if !_rules[rulesp]() {
					goto l719
				}
				if !_rules[ruleFuncParams]() {
					goto l719
				}
				if !_rules[rulesp]() {
					goto l719
				}
				if buffer[position] != rune(')') {
					goto l719
				}
				position++
				if !_rules[ruleAction45]() {
					goto l719
				}
				depth--
				add(ruleFuncApp, position720)
			}
			return true
		l719:
			position, tokenIndex, depth = position719, tokenIndex719, depth719
			return false
		},
		/* 60 FuncParams <- <(<(Wildcard / (Expression sp (',' sp Expression)*)?)> Action46)> */
		func() bool {
			position721, tokenIndex721, depth721 := position, tokenIndex, depth
			{
				position722 := position
				depth++
				{
					position723 := position
					depth++
					{
						position724, tokenIndex724, depth724 := position, tokenIndex, depth
						if !_rules[ruleWildcard]() {
							goto l725
						}
						goto l724
					l725:
						position, tokenIndex, depth = position724, tokenIndex724, depth724
						{
							position726, tokenIndex726, depth726 := position, tokenIndex, depth
							if !_rules[ruleExpression]() {
								goto l726
							}
							if !_rules[rulesp]() {
								goto l726
							}
						l728:
							{
								position729, tokenIndex729, depth729 := position, tokenIndex, depth
								if buffer[position] != rune(',') {
									goto l729
								}
								position++
								if !_rules[rulesp]() {
									goto l729
								}
								if !_rules[ruleExpression]() {
									goto l729
								}
								goto l728
							l729:
								position, tokenIndex, depth = position729, tokenIndex729, depth729
							}
							goto l727
						l726:
							position, tokenIndex, depth = position726, tokenIndex726, depth726
						}
					l727:
					}
				l724:
					depth--
					add(rulePegText, position723)
				}
				if !_rules[ruleAction46]() {
					goto l721
				}
				depth--
				add(ruleFuncParams, position722)
			}
			return true
		l721:
			position, tokenIndex, depth = position721, tokenIndex721, depth721
			return false
		},
		/* 61 Literal <- <(FloatLiteral / NumericLiteral / StringLiteral)> */
		func() bool {
			position730, tokenIndex730, depth730 := position, tokenIndex, depth
			{
				position731 := position
				depth++
				{
					position732, tokenIndex732, depth732 := position, tokenIndex, depth
					if !_rules[ruleFloatLiteral]() {
						goto l733
					}
					goto l732
				l733:
					position, tokenIndex, depth = position732, tokenIndex732, depth732
					if !_rules[ruleNumericLiteral]() {
						goto l734
					}
					goto l732
				l734:
					position, tokenIndex, depth = position732, tokenIndex732, depth732
					if !_rules[ruleStringLiteral]() {
						goto l730
					}
				}
			l732:
				depth--
				add(ruleLiteral, position731)
			}
			return true
		l730:
			position, tokenIndex, depth = position730, tokenIndex730, depth730
			return false
		},
		/* 62 ComparisonOp <- <(Equal / NotEqual / LessOrEqual / Less / GreaterOrEqual / Greater / NotEqual)> */
		func() bool {
			position735, tokenIndex735, depth735 := position, tokenIndex, depth
			{
				position736 := position
				depth++
				{
					position737, tokenIndex737, depth737 := position, tokenIndex, depth
					if !_rules[ruleEqual]() {
						goto l738
					}
					goto l737
				l738:
					position, tokenIndex, depth = position737, tokenIndex737, depth737
					if !_rules[ruleNotEqual]() {
						goto l739
					}
					goto l737
				l739:
					position, tokenIndex, depth = position737, tokenIndex737, depth737
					if !_rules[ruleLessOrEqual]() {
						goto l740
					}
					goto l737
				l740:
					position, tokenIndex, depth = position737, tokenIndex737, depth737
					if !_rules[ruleLess]() {
						goto l741
					}
					goto l737
				l741:
					position, tokenIndex, depth = position737, tokenIndex737, depth737
					if !_rules[ruleGreaterOrEqual]() {
						goto l742
					}
					goto l737
				l742:
					position, tokenIndex, depth = position737, tokenIndex737, depth737
					if !_rules[ruleGreater]() {
						goto l743
					}
					goto l737
				l743:
					position, tokenIndex, depth = position737, tokenIndex737, depth737
					if !_rules[ruleNotEqual]() {
						goto l735
					}
				}
			l737:
				depth--
				add(ruleComparisonOp, position736)
			}
			return true
		l735:
			position, tokenIndex, depth = position735, tokenIndex735, depth735
			return false
		},
		/* 63 OtherOp <- <Concat> */
		func() bool {
			position744, tokenIndex744, depth744 := position, tokenIndex, depth
			{
				position745 := position
				depth++
				if !_rules[ruleConcat]() {
					goto l744
				}
				depth--
				add(ruleOtherOp, position745)
			}
			return true
		l744:
			position, tokenIndex, depth = position744, tokenIndex744, depth744
			return false
		},
		/* 64 IsOp <- <(IsNot / Is)> */
		func() bool {
			position746, tokenIndex746, depth746 := position, tokenIndex, depth
			{
				position747 := position
				depth++
				{
					position748, tokenIndex748, depth748 := position, tokenIndex, depth
					if !_rules[ruleIsNot]() {
						goto l749
					}
					goto l748
				l749:
					position, tokenIndex, depth = position748, tokenIndex748, depth748
					if !_rules[ruleIs]() {
						goto l746
					}
				}
			l748:
				depth--
				add(ruleIsOp, position747)
			}
			return true
		l746:
			position, tokenIndex, depth = position746, tokenIndex746, depth746
			return false
		},
		/* 65 PlusMinusOp <- <(Plus / Minus)> */
		func() bool {
			position750, tokenIndex750, depth750 := position, tokenIndex, depth
			{
				position751 := position
				depth++
				{
					position752, tokenIndex752, depth752 := position, tokenIndex, depth
					if !_rules[rulePlus]() {
						goto l753
					}
					goto l752
				l753:
					position, tokenIndex, depth = position752, tokenIndex752, depth752
					if !_rules[ruleMinus]() {
						goto l750
					}
				}
			l752:
				depth--
				add(rulePlusMinusOp, position751)
			}
			return true
		l750:
			position, tokenIndex, depth = position750, tokenIndex750, depth750
			return false
		},
		/* 66 MultDivOp <- <(Multiply / Divide / Modulo)> */
		func() bool {
			position754, tokenIndex754, depth754 := position, tokenIndex, depth
			{
				position755 := position
				depth++
				{
					position756, tokenIndex756, depth756 := position, tokenIndex, depth
					if !_rules[ruleMultiply]() {
						goto l757
					}
					goto l756
				l757:
					position, tokenIndex, depth = position756, tokenIndex756, depth756
					if !_rules[ruleDivide]() {
						goto l758
					}
					goto l756
				l758:
					position, tokenIndex, depth = position756, tokenIndex756, depth756
					if !_rules[ruleModulo]() {
						goto l754
					}
				}
			l756:
				depth--
				add(ruleMultDivOp, position755)
			}
			return true
		l754:
			position, tokenIndex, depth = position754, tokenIndex754, depth754
			return false
		},
		/* 67 Stream <- <(<ident> Action47)> */
		func() bool {
			position759, tokenIndex759, depth759 := position, tokenIndex, depth
			{
				position760 := position
				depth++
				{
					position761 := position
					depth++
					if !_rules[ruleident]() {
						goto l759
					}
					depth--
					add(rulePegText, position761)
				}
				if !_rules[ruleAction47]() {
					goto l759
				}
				depth--
				add(ruleStream, position760)
			}
			return true
		l759:
			position, tokenIndex, depth = position759, tokenIndex759, depth759
			return false
		},
		/* 68 RowMeta <- <RowTimestamp> */
		func() bool {
			position762, tokenIndex762, depth762 := position, tokenIndex, depth
			{
				position763 := position
				depth++
				if !_rules[ruleRowTimestamp]() {
					goto l762
				}
				depth--
				add(ruleRowMeta, position763)
			}
			return true
		l762:
			position, tokenIndex, depth = position762, tokenIndex762, depth762
			return false
		},
		/* 69 RowTimestamp <- <(<((ident ':')? ('t' 's' '(' ')'))> Action48)> */
		func() bool {
			position764, tokenIndex764, depth764 := position, tokenIndex, depth
			{
				position765 := position
				depth++
				{
					position766 := position
					depth++
					{
						position767, tokenIndex767, depth767 := position, tokenIndex, depth
						if !_rules[ruleident]() {
							goto l767
						}
						if buffer[position] != rune(':') {
							goto l767
						}
						position++
						goto l768
					l767:
						position, tokenIndex, depth = position767, tokenIndex767, depth767
					}
				l768:
					if buffer[position] != rune('t') {
						goto l764
					}
					position++
					if buffer[position] != rune('s') {
						goto l764
					}
					position++
					if buffer[position] != rune('(') {
						goto l764
					}
					position++
					if buffer[position] != rune(')') {
						goto l764
					}
					position++
					depth--
					add(rulePegText, position766)
				}
				if !_rules[ruleAction48]() {
					goto l764
				}
				depth--
				add(ruleRowTimestamp, position765)
			}
			return true
		l764:
			position, tokenIndex, depth = position764, tokenIndex764, depth764
			return false
		},
		/* 70 RowValue <- <(<((ident ':')? jsonPath)> Action49)> */
		func() bool {
			position769, tokenIndex769, depth769 := position, tokenIndex, depth
			{
				position770 := position
				depth++
				{
					position771 := position
					depth++
					{
						position772, tokenIndex772, depth772 := position, tokenIndex, depth
						if !_rules[ruleident]() {
							goto l772
						}
						if buffer[position] != rune(':') {
							goto l772
						}
						position++
						goto l773
					l772:
						position, tokenIndex, depth = position772, tokenIndex772, depth772
					}
				l773:
					if !_rules[rulejsonPath]() {
						goto l769
					}
					depth--
					add(rulePegText, position771)
				}
				if !_rules[ruleAction49]() {
					goto l769
				}
				depth--
				add(ruleRowValue, position770)
			}
			return true
		l769:
			position, tokenIndex, depth = position769, tokenIndex769, depth769
			return false
		},
		/* 71 NumericLiteral <- <(<('-'? [0-9]+)> Action50)> */
		func() bool {
			position774, tokenIndex774, depth774 := position, tokenIndex, depth
			{
				position775 := position
				depth++
				{
					position776 := position
					depth++
					{
						position777, tokenIndex777, depth777 := position, tokenIndex, depth
						if buffer[position] != rune('-') {
							goto l777
						}
						position++
						goto l778
					l777:
						position, tokenIndex, depth = position777, tokenIndex777, depth777
					}
				l778:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l774
					}
					position++
				l779:
					{
						position780, tokenIndex780, depth780 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l780
						}
						position++
						goto l779
					l780:
						position, tokenIndex, depth = position780, tokenIndex780, depth780
					}
					depth--
					add(rulePegText, position776)
				}
				if !_rules[ruleAction50]() {
					goto l774
				}
				depth--
				add(ruleNumericLiteral, position775)
			}
			return true
		l774:
			position, tokenIndex, depth = position774, tokenIndex774, depth774
			return false
		},
		/* 72 FloatLiteral <- <(<('-'? [0-9]+ '.' [0-9]+)> Action51)> */
		func() bool {
			position781, tokenIndex781, depth781 := position, tokenIndex, depth
			{
				position782 := position
				depth++
				{
					position783 := position
					depth++
					{
						position784, tokenIndex784, depth784 := position, tokenIndex, depth
						if buffer[position] != rune('-') {
							goto l784
						}
						position++
						goto l785
					l784:
						position, tokenIndex, depth = position784, tokenIndex784, depth784
					}
				l785:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l781
					}
					position++
				l786:
					{
						position787, tokenIndex787, depth787 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l787
						}
						position++
						goto l786
					l787:
						position, tokenIndex, depth = position787, tokenIndex787, depth787
					}
					if buffer[position] != rune('.') {
						goto l781
					}
					position++
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l781
					}
					position++
				l788:
					{
						position789, tokenIndex789, depth789 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l789
						}
						position++
						goto l788
					l789:
						position, tokenIndex, depth = position789, tokenIndex789, depth789
					}
					depth--
					add(rulePegText, position783)
				}
				if !_rules[ruleAction51]() {
					goto l781
				}
				depth--
				add(ruleFloatLiteral, position782)
			}
			return true
		l781:
			position, tokenIndex, depth = position781, tokenIndex781, depth781
			return false
		},
		/* 73 Function <- <(<ident> Action52)> */
		func() bool {
			position790, tokenIndex790, depth790 := position, tokenIndex, depth
			{
				position791 := position
				depth++
				{
					position792 := position
					depth++
					if !_rules[ruleident]() {
						goto l790
					}
					depth--
					add(rulePegText, position792)
				}
				if !_rules[ruleAction52]() {
					goto l790
				}
				depth--
				add(ruleFunction, position791)
			}
			return true
		l790:
			position, tokenIndex, depth = position790, tokenIndex790, depth790
			return false
		},
		/* 74 NullLiteral <- <(<(('n' / 'N') ('u' / 'U') ('l' / 'L') ('l' / 'L'))> Action53)> */
		func() bool {
			position793, tokenIndex793, depth793 := position, tokenIndex, depth
			{
				position794 := position
				depth++
				{
					position795 := position
					depth++
					{
						position796, tokenIndex796, depth796 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l797
						}
						position++
						goto l796
					l797:
						position, tokenIndex, depth = position796, tokenIndex796, depth796
						if buffer[position] != rune('N') {
							goto l793
						}
						position++
					}
				l796:
					{
						position798, tokenIndex798, depth798 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l799
						}
						position++
						goto l798
					l799:
						position, tokenIndex, depth = position798, tokenIndex798, depth798
						if buffer[position] != rune('U') {
							goto l793
						}
						position++
					}
				l798:
					{
						position800, tokenIndex800, depth800 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l801
						}
						position++
						goto l800
					l801:
						position, tokenIndex, depth = position800, tokenIndex800, depth800
						if buffer[position] != rune('L') {
							goto l793
						}
						position++
					}
				l800:
					{
						position802, tokenIndex802, depth802 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l803
						}
						position++
						goto l802
					l803:
						position, tokenIndex, depth = position802, tokenIndex802, depth802
						if buffer[position] != rune('L') {
							goto l793
						}
						position++
					}
				l802:
					depth--
					add(rulePegText, position795)
				}
				if !_rules[ruleAction53]() {
					goto l793
				}
				depth--
				add(ruleNullLiteral, position794)
			}
			return true
		l793:
			position, tokenIndex, depth = position793, tokenIndex793, depth793
			return false
		},
		/* 75 BooleanLiteral <- <(TRUE / FALSE)> */
		func() bool {
			position804, tokenIndex804, depth804 := position, tokenIndex, depth
			{
				position805 := position
				depth++
				{
					position806, tokenIndex806, depth806 := position, tokenIndex, depth
					if !_rules[ruleTRUE]() {
						goto l807
					}
					goto l806
				l807:
					position, tokenIndex, depth = position806, tokenIndex806, depth806
					if !_rules[ruleFALSE]() {
						goto l804
					}
				}
			l806:
				depth--
				add(ruleBooleanLiteral, position805)
			}
			return true
		l804:
			position, tokenIndex, depth = position804, tokenIndex804, depth804
			return false
		},
		/* 76 TRUE <- <(<(('t' / 'T') ('r' / 'R') ('u' / 'U') ('e' / 'E'))> Action54)> */
		func() bool {
			position808, tokenIndex808, depth808 := position, tokenIndex, depth
			{
				position809 := position
				depth++
				{
					position810 := position
					depth++
					{
						position811, tokenIndex811, depth811 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l812
						}
						position++
						goto l811
					l812:
						position, tokenIndex, depth = position811, tokenIndex811, depth811
						if buffer[position] != rune('T') {
							goto l808
						}
						position++
					}
				l811:
					{
						position813, tokenIndex813, depth813 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l814
						}
						position++
						goto l813
					l814:
						position, tokenIndex, depth = position813, tokenIndex813, depth813
						if buffer[position] != rune('R') {
							goto l808
						}
						position++
					}
				l813:
					{
						position815, tokenIndex815, depth815 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l816
						}
						position++
						goto l815
					l816:
						position, tokenIndex, depth = position815, tokenIndex815, depth815
						if buffer[position] != rune('U') {
							goto l808
						}
						position++
					}
				l815:
					{
						position817, tokenIndex817, depth817 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l818
						}
						position++
						goto l817
					l818:
						position, tokenIndex, depth = position817, tokenIndex817, depth817
						if buffer[position] != rune('E') {
							goto l808
						}
						position++
					}
				l817:
					depth--
					add(rulePegText, position810)
				}
				if !_rules[ruleAction54]() {
					goto l808
				}
				depth--
				add(ruleTRUE, position809)
			}
			return true
		l808:
			position, tokenIndex, depth = position808, tokenIndex808, depth808
			return false
		},
		/* 77 FALSE <- <(<(('f' / 'F') ('a' / 'A') ('l' / 'L') ('s' / 'S') ('e' / 'E'))> Action55)> */
		func() bool {
			position819, tokenIndex819, depth819 := position, tokenIndex, depth
			{
				position820 := position
				depth++
				{
					position821 := position
					depth++
					{
						position822, tokenIndex822, depth822 := position, tokenIndex, depth
						if buffer[position] != rune('f') {
							goto l823
						}
						position++
						goto l822
					l823:
						position, tokenIndex, depth = position822, tokenIndex822, depth822
						if buffer[position] != rune('F') {
							goto l819
						}
						position++
					}
				l822:
					{
						position824, tokenIndex824, depth824 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l825
						}
						position++
						goto l824
					l825:
						position, tokenIndex, depth = position824, tokenIndex824, depth824
						if buffer[position] != rune('A') {
							goto l819
						}
						position++
					}
				l824:
					{
						position826, tokenIndex826, depth826 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l827
						}
						position++
						goto l826
					l827:
						position, tokenIndex, depth = position826, tokenIndex826, depth826
						if buffer[position] != rune('L') {
							goto l819
						}
						position++
					}
				l826:
					{
						position828, tokenIndex828, depth828 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l829
						}
						position++
						goto l828
					l829:
						position, tokenIndex, depth = position828, tokenIndex828, depth828
						if buffer[position] != rune('S') {
							goto l819
						}
						position++
					}
				l828:
					{
						position830, tokenIndex830, depth830 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l831
						}
						position++
						goto l830
					l831:
						position, tokenIndex, depth = position830, tokenIndex830, depth830
						if buffer[position] != rune('E') {
							goto l819
						}
						position++
					}
				l830:
					depth--
					add(rulePegText, position821)
				}
				if !_rules[ruleAction55]() {
					goto l819
				}
				depth--
				add(ruleFALSE, position820)
			}
			return true
		l819:
			position, tokenIndex, depth = position819, tokenIndex819, depth819
			return false
		},
		/* 78 Wildcard <- <(<'*'> Action56)> */
		func() bool {
			position832, tokenIndex832, depth832 := position, tokenIndex, depth
			{
				position833 := position
				depth++
				{
					position834 := position
					depth++
					if buffer[position] != rune('*') {
						goto l832
					}
					position++
					depth--
					add(rulePegText, position834)
				}
				if !_rules[ruleAction56]() {
					goto l832
				}
				depth--
				add(ruleWildcard, position833)
			}
			return true
		l832:
			position, tokenIndex, depth = position832, tokenIndex832, depth832
			return false
		},
		/* 79 StringLiteral <- <(<('\'' (('\'' '\'') / (!'\'' .))* '\'')> Action57)> */
		func() bool {
			position835, tokenIndex835, depth835 := position, tokenIndex, depth
			{
				position836 := position
				depth++
				{
					position837 := position
					depth++
					if buffer[position] != rune('\'') {
						goto l835
					}
					position++
				l838:
					{
						position839, tokenIndex839, depth839 := position, tokenIndex, depth
						{
							position840, tokenIndex840, depth840 := position, tokenIndex, depth
							if buffer[position] != rune('\'') {
								goto l841
							}
							position++
							if buffer[position] != rune('\'') {
								goto l841
							}
							position++
							goto l840
						l841:
							position, tokenIndex, depth = position840, tokenIndex840, depth840
							{
								position842, tokenIndex842, depth842 := position, tokenIndex, depth
								if buffer[position] != rune('\'') {
									goto l842
								}
								position++
								goto l839
							l842:
								position, tokenIndex, depth = position842, tokenIndex842, depth842
							}
							if !matchDot() {
								goto l839
							}
						}
					l840:
						goto l838
					l839:
						position, tokenIndex, depth = position839, tokenIndex839, depth839
					}
					if buffer[position] != rune('\'') {
						goto l835
					}
					position++
					depth--
					add(rulePegText, position837)
				}
				if !_rules[ruleAction57]() {
					goto l835
				}
				depth--
				add(ruleStringLiteral, position836)
			}
			return true
		l835:
			position, tokenIndex, depth = position835, tokenIndex835, depth835
			return false
		},
		/* 80 ISTREAM <- <(<(('i' / 'I') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M'))> Action58)> */
		func() bool {
			position843, tokenIndex843, depth843 := position, tokenIndex, depth
			{
				position844 := position
				depth++
				{
					position845 := position
					depth++
					{
						position846, tokenIndex846, depth846 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l847
						}
						position++
						goto l846
					l847:
						position, tokenIndex, depth = position846, tokenIndex846, depth846
						if buffer[position] != rune('I') {
							goto l843
						}
						position++
					}
				l846:
					{
						position848, tokenIndex848, depth848 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l849
						}
						position++
						goto l848
					l849:
						position, tokenIndex, depth = position848, tokenIndex848, depth848
						if buffer[position] != rune('S') {
							goto l843
						}
						position++
					}
				l848:
					{
						position850, tokenIndex850, depth850 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l851
						}
						position++
						goto l850
					l851:
						position, tokenIndex, depth = position850, tokenIndex850, depth850
						if buffer[position] != rune('T') {
							goto l843
						}
						position++
					}
				l850:
					{
						position852, tokenIndex852, depth852 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l853
						}
						position++
						goto l852
					l853:
						position, tokenIndex, depth = position852, tokenIndex852, depth852
						if buffer[position] != rune('R') {
							goto l843
						}
						position++
					}
				l852:
					{
						position854, tokenIndex854, depth854 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l855
						}
						position++
						goto l854
					l855:
						position, tokenIndex, depth = position854, tokenIndex854, depth854
						if buffer[position] != rune('E') {
							goto l843
						}
						position++
					}
				l854:
					{
						position856, tokenIndex856, depth856 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l857
						}
						position++
						goto l856
					l857:
						position, tokenIndex, depth = position856, tokenIndex856, depth856
						if buffer[position] != rune('A') {
							goto l843
						}
						position++
					}
				l856:
					{
						position858, tokenIndex858, depth858 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l859
						}
						position++
						goto l858
					l859:
						position, tokenIndex, depth = position858, tokenIndex858, depth858
						if buffer[position] != rune('M') {
							goto l843
						}
						position++
					}
				l858:
					depth--
					add(rulePegText, position845)
				}
				if !_rules[ruleAction58]() {
					goto l843
				}
				depth--
				add(ruleISTREAM, position844)
			}
			return true
		l843:
			position, tokenIndex, depth = position843, tokenIndex843, depth843
			return false
		},
		/* 81 DSTREAM <- <(<(('d' / 'D') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M'))> Action59)> */
		func() bool {
			position860, tokenIndex860, depth860 := position, tokenIndex, depth
			{
				position861 := position
				depth++
				{
					position862 := position
					depth++
					{
						position863, tokenIndex863, depth863 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l864
						}
						position++
						goto l863
					l864:
						position, tokenIndex, depth = position863, tokenIndex863, depth863
						if buffer[position] != rune('D') {
							goto l860
						}
						position++
					}
				l863:
					{
						position865, tokenIndex865, depth865 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l866
						}
						position++
						goto l865
					l866:
						position, tokenIndex, depth = position865, tokenIndex865, depth865
						if buffer[position] != rune('S') {
							goto l860
						}
						position++
					}
				l865:
					{
						position867, tokenIndex867, depth867 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l868
						}
						position++
						goto l867
					l868:
						position, tokenIndex, depth = position867, tokenIndex867, depth867
						if buffer[position] != rune('T') {
							goto l860
						}
						position++
					}
				l867:
					{
						position869, tokenIndex869, depth869 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l870
						}
						position++
						goto l869
					l870:
						position, tokenIndex, depth = position869, tokenIndex869, depth869
						if buffer[position] != rune('R') {
							goto l860
						}
						position++
					}
				l869:
					{
						position871, tokenIndex871, depth871 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l872
						}
						position++
						goto l871
					l872:
						position, tokenIndex, depth = position871, tokenIndex871, depth871
						if buffer[position] != rune('E') {
							goto l860
						}
						position++
					}
				l871:
					{
						position873, tokenIndex873, depth873 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l874
						}
						position++
						goto l873
					l874:
						position, tokenIndex, depth = position873, tokenIndex873, depth873
						if buffer[position] != rune('A') {
							goto l860
						}
						position++
					}
				l873:
					{
						position875, tokenIndex875, depth875 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l876
						}
						position++
						goto l875
					l876:
						position, tokenIndex, depth = position875, tokenIndex875, depth875
						if buffer[position] != rune('M') {
							goto l860
						}
						position++
					}
				l875:
					depth--
					add(rulePegText, position862)
				}
				if !_rules[ruleAction59]() {
					goto l860
				}
				depth--
				add(ruleDSTREAM, position861)
			}
			return true
		l860:
			position, tokenIndex, depth = position860, tokenIndex860, depth860
			return false
		},
		/* 82 RSTREAM <- <(<(('r' / 'R') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M'))> Action60)> */
		func() bool {
			position877, tokenIndex877, depth877 := position, tokenIndex, depth
			{
				position878 := position
				depth++
				{
					position879 := position
					depth++
					{
						position880, tokenIndex880, depth880 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l881
						}
						position++
						goto l880
					l881:
						position, tokenIndex, depth = position880, tokenIndex880, depth880
						if buffer[position] != rune('R') {
							goto l877
						}
						position++
					}
				l880:
					{
						position882, tokenIndex882, depth882 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l883
						}
						position++
						goto l882
					l883:
						position, tokenIndex, depth = position882, tokenIndex882, depth882
						if buffer[position] != rune('S') {
							goto l877
						}
						position++
					}
				l882:
					{
						position884, tokenIndex884, depth884 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l885
						}
						position++
						goto l884
					l885:
						position, tokenIndex, depth = position884, tokenIndex884, depth884
						if buffer[position] != rune('T') {
							goto l877
						}
						position++
					}
				l884:
					{
						position886, tokenIndex886, depth886 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l887
						}
						position++
						goto l886
					l887:
						position, tokenIndex, depth = position886, tokenIndex886, depth886
						if buffer[position] != rune('R') {
							goto l877
						}
						position++
					}
				l886:
					{
						position888, tokenIndex888, depth888 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l889
						}
						position++
						goto l888
					l889:
						position, tokenIndex, depth = position888, tokenIndex888, depth888
						if buffer[position] != rune('E') {
							goto l877
						}
						position++
					}
				l888:
					{
						position890, tokenIndex890, depth890 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l891
						}
						position++
						goto l890
					l891:
						position, tokenIndex, depth = position890, tokenIndex890, depth890
						if buffer[position] != rune('A') {
							goto l877
						}
						position++
					}
				l890:
					{
						position892, tokenIndex892, depth892 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l893
						}
						position++
						goto l892
					l893:
						position, tokenIndex, depth = position892, tokenIndex892, depth892
						if buffer[position] != rune('M') {
							goto l877
						}
						position++
					}
				l892:
					depth--
					add(rulePegText, position879)
				}
				if !_rules[ruleAction60]() {
					goto l877
				}
				depth--
				add(ruleRSTREAM, position878)
			}
			return true
		l877:
			position, tokenIndex, depth = position877, tokenIndex877, depth877
			return false
		},
		/* 83 TUPLES <- <(<(('t' / 'T') ('u' / 'U') ('p' / 'P') ('l' / 'L') ('e' / 'E') ('s' / 'S'))> Action61)> */
		func() bool {
			position894, tokenIndex894, depth894 := position, tokenIndex, depth
			{
				position895 := position
				depth++
				{
					position896 := position
					depth++
					{
						position897, tokenIndex897, depth897 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l898
						}
						position++
						goto l897
					l898:
						position, tokenIndex, depth = position897, tokenIndex897, depth897
						if buffer[position] != rune('T') {
							goto l894
						}
						position++
					}
				l897:
					{
						position899, tokenIndex899, depth899 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l900
						}
						position++
						goto l899
					l900:
						position, tokenIndex, depth = position899, tokenIndex899, depth899
						if buffer[position] != rune('U') {
							goto l894
						}
						position++
					}
				l899:
					{
						position901, tokenIndex901, depth901 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l902
						}
						position++
						goto l901
					l902:
						position, tokenIndex, depth = position901, tokenIndex901, depth901
						if buffer[position] != rune('P') {
							goto l894
						}
						position++
					}
				l901:
					{
						position903, tokenIndex903, depth903 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l904
						}
						position++
						goto l903
					l904:
						position, tokenIndex, depth = position903, tokenIndex903, depth903
						if buffer[position] != rune('L') {
							goto l894
						}
						position++
					}
				l903:
					{
						position905, tokenIndex905, depth905 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l906
						}
						position++
						goto l905
					l906:
						position, tokenIndex, depth = position905, tokenIndex905, depth905
						if buffer[position] != rune('E') {
							goto l894
						}
						position++
					}
				l905:
					{
						position907, tokenIndex907, depth907 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l908
						}
						position++
						goto l907
					l908:
						position, tokenIndex, depth = position907, tokenIndex907, depth907
						if buffer[position] != rune('S') {
							goto l894
						}
						position++
					}
				l907:
					depth--
					add(rulePegText, position896)
				}
				if !_rules[ruleAction61]() {
					goto l894
				}
				depth--
				add(ruleTUPLES, position895)
			}
			return true
		l894:
			position, tokenIndex, depth = position894, tokenIndex894, depth894
			return false
		},
		/* 84 SECONDS <- <(<(('s' / 'S') ('e' / 'E') ('c' / 'C') ('o' / 'O') ('n' / 'N') ('d' / 'D') ('s' / 'S'))> Action62)> */
		func() bool {
			position909, tokenIndex909, depth909 := position, tokenIndex, depth
			{
				position910 := position
				depth++
				{
					position911 := position
					depth++
					{
						position912, tokenIndex912, depth912 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l913
						}
						position++
						goto l912
					l913:
						position, tokenIndex, depth = position912, tokenIndex912, depth912
						if buffer[position] != rune('S') {
							goto l909
						}
						position++
					}
				l912:
					{
						position914, tokenIndex914, depth914 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l915
						}
						position++
						goto l914
					l915:
						position, tokenIndex, depth = position914, tokenIndex914, depth914
						if buffer[position] != rune('E') {
							goto l909
						}
						position++
					}
				l914:
					{
						position916, tokenIndex916, depth916 := position, tokenIndex, depth
						if buffer[position] != rune('c') {
							goto l917
						}
						position++
						goto l916
					l917:
						position, tokenIndex, depth = position916, tokenIndex916, depth916
						if buffer[position] != rune('C') {
							goto l909
						}
						position++
					}
				l916:
					{
						position918, tokenIndex918, depth918 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l919
						}
						position++
						goto l918
					l919:
						position, tokenIndex, depth = position918, tokenIndex918, depth918
						if buffer[position] != rune('O') {
							goto l909
						}
						position++
					}
				l918:
					{
						position920, tokenIndex920, depth920 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l921
						}
						position++
						goto l920
					l921:
						position, tokenIndex, depth = position920, tokenIndex920, depth920
						if buffer[position] != rune('N') {
							goto l909
						}
						position++
					}
				l920:
					{
						position922, tokenIndex922, depth922 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l923
						}
						position++
						goto l922
					l923:
						position, tokenIndex, depth = position922, tokenIndex922, depth922
						if buffer[position] != rune('D') {
							goto l909
						}
						position++
					}
				l922:
					{
						position924, tokenIndex924, depth924 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l925
						}
						position++
						goto l924
					l925:
						position, tokenIndex, depth = position924, tokenIndex924, depth924
						if buffer[position] != rune('S') {
							goto l909
						}
						position++
					}
				l924:
					depth--
					add(rulePegText, position911)
				}
				if !_rules[ruleAction62]() {
					goto l909
				}
				depth--
				add(ruleSECONDS, position910)
			}
			return true
		l909:
			position, tokenIndex, depth = position909, tokenIndex909, depth909
			return false
		},
		/* 85 StreamIdentifier <- <(<ident> Action63)> */
		func() bool {
			position926, tokenIndex926, depth926 := position, tokenIndex, depth
			{
				position927 := position
				depth++
				{
					position928 := position
					depth++
					if !_rules[ruleident]() {
						goto l926
					}
					depth--
					add(rulePegText, position928)
				}
				if !_rules[ruleAction63]() {
					goto l926
				}
				depth--
				add(ruleStreamIdentifier, position927)
			}
			return true
		l926:
			position, tokenIndex, depth = position926, tokenIndex926, depth926
			return false
		},
		/* 86 SourceSinkType <- <(<ident> Action64)> */
		func() bool {
			position929, tokenIndex929, depth929 := position, tokenIndex, depth
			{
				position930 := position
				depth++
				{
					position931 := position
					depth++
					if !_rules[ruleident]() {
						goto l929
					}
					depth--
					add(rulePegText, position931)
				}
				if !_rules[ruleAction64]() {
					goto l929
				}
				depth--
				add(ruleSourceSinkType, position930)
			}
			return true
		l929:
			position, tokenIndex, depth = position929, tokenIndex929, depth929
			return false
		},
		/* 87 SourceSinkParamKey <- <(<ident> Action65)> */
		func() bool {
			position932, tokenIndex932, depth932 := position, tokenIndex, depth
			{
				position933 := position
				depth++
				{
					position934 := position
					depth++
					if !_rules[ruleident]() {
						goto l932
					}
					depth--
					add(rulePegText, position934)
				}
				if !_rules[ruleAction65]() {
					goto l932
				}
				depth--
				add(ruleSourceSinkParamKey, position933)
			}
			return true
		l932:
			position, tokenIndex, depth = position932, tokenIndex932, depth932
			return false
		},
		/* 88 Paused <- <(<(('p' / 'P') ('a' / 'A') ('u' / 'U') ('s' / 'S') ('e' / 'E') ('d' / 'D'))> Action66)> */
		func() bool {
			position935, tokenIndex935, depth935 := position, tokenIndex, depth
			{
				position936 := position
				depth++
				{
					position937 := position
					depth++
					{
						position938, tokenIndex938, depth938 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l939
						}
						position++
						goto l938
					l939:
						position, tokenIndex, depth = position938, tokenIndex938, depth938
						if buffer[position] != rune('P') {
							goto l935
						}
						position++
					}
				l938:
					{
						position940, tokenIndex940, depth940 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l941
						}
						position++
						goto l940
					l941:
						position, tokenIndex, depth = position940, tokenIndex940, depth940
						if buffer[position] != rune('A') {
							goto l935
						}
						position++
					}
				l940:
					{
						position942, tokenIndex942, depth942 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l943
						}
						position++
						goto l942
					l943:
						position, tokenIndex, depth = position942, tokenIndex942, depth942
						if buffer[position] != rune('U') {
							goto l935
						}
						position++
					}
				l942:
					{
						position944, tokenIndex944, depth944 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l945
						}
						position++
						goto l944
					l945:
						position, tokenIndex, depth = position944, tokenIndex944, depth944
						if buffer[position] != rune('S') {
							goto l935
						}
						position++
					}
				l944:
					{
						position946, tokenIndex946, depth946 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l947
						}
						position++
						goto l946
					l947:
						position, tokenIndex, depth = position946, tokenIndex946, depth946
						if buffer[position] != rune('E') {
							goto l935
						}
						position++
					}
				l946:
					{
						position948, tokenIndex948, depth948 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l949
						}
						position++
						goto l948
					l949:
						position, tokenIndex, depth = position948, tokenIndex948, depth948
						if buffer[position] != rune('D') {
							goto l935
						}
						position++
					}
				l948:
					depth--
					add(rulePegText, position937)
				}
				if !_rules[ruleAction66]() {
					goto l935
				}
				depth--
				add(rulePaused, position936)
			}
			return true
		l935:
			position, tokenIndex, depth = position935, tokenIndex935, depth935
			return false
		},
		/* 89 Unpaused <- <(<(('u' / 'U') ('n' / 'N') ('p' / 'P') ('a' / 'A') ('u' / 'U') ('s' / 'S') ('e' / 'E') ('d' / 'D'))> Action67)> */
		func() bool {
			position950, tokenIndex950, depth950 := position, tokenIndex, depth
			{
				position951 := position
				depth++
				{
					position952 := position
					depth++
					{
						position953, tokenIndex953, depth953 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l954
						}
						position++
						goto l953
					l954:
						position, tokenIndex, depth = position953, tokenIndex953, depth953
						if buffer[position] != rune('U') {
							goto l950
						}
						position++
					}
				l953:
					{
						position955, tokenIndex955, depth955 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l956
						}
						position++
						goto l955
					l956:
						position, tokenIndex, depth = position955, tokenIndex955, depth955
						if buffer[position] != rune('N') {
							goto l950
						}
						position++
					}
				l955:
					{
						position957, tokenIndex957, depth957 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l958
						}
						position++
						goto l957
					l958:
						position, tokenIndex, depth = position957, tokenIndex957, depth957
						if buffer[position] != rune('P') {
							goto l950
						}
						position++
					}
				l957:
					{
						position959, tokenIndex959, depth959 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l960
						}
						position++
						goto l959
					l960:
						position, tokenIndex, depth = position959, tokenIndex959, depth959
						if buffer[position] != rune('A') {
							goto l950
						}
						position++
					}
				l959:
					{
						position961, tokenIndex961, depth961 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l962
						}
						position++
						goto l961
					l962:
						position, tokenIndex, depth = position961, tokenIndex961, depth961
						if buffer[position] != rune('U') {
							goto l950
						}
						position++
					}
				l961:
					{
						position963, tokenIndex963, depth963 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l964
						}
						position++
						goto l963
					l964:
						position, tokenIndex, depth = position963, tokenIndex963, depth963
						if buffer[position] != rune('S') {
							goto l950
						}
						position++
					}
				l963:
					{
						position965, tokenIndex965, depth965 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l966
						}
						position++
						goto l965
					l966:
						position, tokenIndex, depth = position965, tokenIndex965, depth965
						if buffer[position] != rune('E') {
							goto l950
						}
						position++
					}
				l965:
					{
						position967, tokenIndex967, depth967 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l968
						}
						position++
						goto l967
					l968:
						position, tokenIndex, depth = position967, tokenIndex967, depth967
						if buffer[position] != rune('D') {
							goto l950
						}
						position++
					}
				l967:
					depth--
					add(rulePegText, position952)
				}
				if !_rules[ruleAction67]() {
					goto l950
				}
				depth--
				add(ruleUnpaused, position951)
			}
			return true
		l950:
			position, tokenIndex, depth = position950, tokenIndex950, depth950
			return false
		},
		/* 90 Type <- <(Bool / Int / Float / String / Blob / Timestamp / Array / Map)> */
		func() bool {
			position969, tokenIndex969, depth969 := position, tokenIndex, depth
			{
				position970 := position
				depth++
				{
					position971, tokenIndex971, depth971 := position, tokenIndex, depth
					if !_rules[ruleBool]() {
						goto l972
					}
					goto l971
				l972:
					position, tokenIndex, depth = position971, tokenIndex971, depth971
					if !_rules[ruleInt]() {
						goto l973
					}
					goto l971
				l973:
					position, tokenIndex, depth = position971, tokenIndex971, depth971
					if !_rules[ruleFloat]() {
						goto l974
					}
					goto l971
				l974:
					position, tokenIndex, depth = position971, tokenIndex971, depth971
					if !_rules[ruleString]() {
						goto l975
					}
					goto l971
				l975:
					position, tokenIndex, depth = position971, tokenIndex971, depth971
					if !_rules[ruleBlob]() {
						goto l976
					}
					goto l971
				l976:
					position, tokenIndex, depth = position971, tokenIndex971, depth971
					if !_rules[ruleTimestamp]() {
						goto l977
					}
					goto l971
				l977:
					position, tokenIndex, depth = position971, tokenIndex971, depth971
					if !_rules[ruleArray]() {
						goto l978
					}
					goto l971
				l978:
					position, tokenIndex, depth = position971, tokenIndex971, depth971
					if !_rules[ruleMap]() {
						goto l969
					}
				}
			l971:
				depth--
				add(ruleType, position970)
			}
			return true
		l969:
			position, tokenIndex, depth = position969, tokenIndex969, depth969
			return false
		},
		/* 91 Bool <- <(<(('b' / 'B') ('o' / 'O') ('o' / 'O') ('l' / 'L'))> Action68)> */
		func() bool {
			position979, tokenIndex979, depth979 := position, tokenIndex, depth
			{
				position980 := position
				depth++
				{
					position981 := position
					depth++
					{
						position982, tokenIndex982, depth982 := position, tokenIndex, depth
						if buffer[position] != rune('b') {
							goto l983
						}
						position++
						goto l982
					l983:
						position, tokenIndex, depth = position982, tokenIndex982, depth982
						if buffer[position] != rune('B') {
							goto l979
						}
						position++
					}
				l982:
					{
						position984, tokenIndex984, depth984 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l985
						}
						position++
						goto l984
					l985:
						position, tokenIndex, depth = position984, tokenIndex984, depth984
						if buffer[position] != rune('O') {
							goto l979
						}
						position++
					}
				l984:
					{
						position986, tokenIndex986, depth986 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l987
						}
						position++
						goto l986
					l987:
						position, tokenIndex, depth = position986, tokenIndex986, depth986
						if buffer[position] != rune('O') {
							goto l979
						}
						position++
					}
				l986:
					{
						position988, tokenIndex988, depth988 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l989
						}
						position++
						goto l988
					l989:
						position, tokenIndex, depth = position988, tokenIndex988, depth988
						if buffer[position] != rune('L') {
							goto l979
						}
						position++
					}
				l988:
					depth--
					add(rulePegText, position981)
				}
				if !_rules[ruleAction68]() {
					goto l979
				}
				depth--
				add(ruleBool, position980)
			}
			return true
		l979:
			position, tokenIndex, depth = position979, tokenIndex979, depth979
			return false
		},
		/* 92 Int <- <(<(('i' / 'I') ('n' / 'N') ('t' / 'T'))> Action69)> */
		func() bool {
			position990, tokenIndex990, depth990 := position, tokenIndex, depth
			{
				position991 := position
				depth++
				{
					position992 := position
					depth++
					{
						position993, tokenIndex993, depth993 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l994
						}
						position++
						goto l993
					l994:
						position, tokenIndex, depth = position993, tokenIndex993, depth993
						if buffer[position] != rune('I') {
							goto l990
						}
						position++
					}
				l993:
					{
						position995, tokenIndex995, depth995 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l996
						}
						position++
						goto l995
					l996:
						position, tokenIndex, depth = position995, tokenIndex995, depth995
						if buffer[position] != rune('N') {
							goto l990
						}
						position++
					}
				l995:
					{
						position997, tokenIndex997, depth997 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l998
						}
						position++
						goto l997
					l998:
						position, tokenIndex, depth = position997, tokenIndex997, depth997
						if buffer[position] != rune('T') {
							goto l990
						}
						position++
					}
				l997:
					depth--
					add(rulePegText, position992)
				}
				if !_rules[ruleAction69]() {
					goto l990
				}
				depth--
				add(ruleInt, position991)
			}
			return true
		l990:
			position, tokenIndex, depth = position990, tokenIndex990, depth990
			return false
		},
		/* 93 Float <- <(<(('f' / 'F') ('l' / 'L') ('o' / 'O') ('a' / 'A') ('t' / 'T'))> Action70)> */
		func() bool {
			position999, tokenIndex999, depth999 := position, tokenIndex, depth
			{
				position1000 := position
				depth++
				{
					position1001 := position
					depth++
					{
						position1002, tokenIndex1002, depth1002 := position, tokenIndex, depth
						if buffer[position] != rune('f') {
							goto l1003
						}
						position++
						goto l1002
					l1003:
						position, tokenIndex, depth = position1002, tokenIndex1002, depth1002
						if buffer[position] != rune('F') {
							goto l999
						}
						position++
					}
				l1002:
					{
						position1004, tokenIndex1004, depth1004 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1005
						}
						position++
						goto l1004
					l1005:
						position, tokenIndex, depth = position1004, tokenIndex1004, depth1004
						if buffer[position] != rune('L') {
							goto l999
						}
						position++
					}
				l1004:
					{
						position1006, tokenIndex1006, depth1006 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1007
						}
						position++
						goto l1006
					l1007:
						position, tokenIndex, depth = position1006, tokenIndex1006, depth1006
						if buffer[position] != rune('O') {
							goto l999
						}
						position++
					}
				l1006:
					{
						position1008, tokenIndex1008, depth1008 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1009
						}
						position++
						goto l1008
					l1009:
						position, tokenIndex, depth = position1008, tokenIndex1008, depth1008
						if buffer[position] != rune('A') {
							goto l999
						}
						position++
					}
				l1008:
					{
						position1010, tokenIndex1010, depth1010 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1011
						}
						position++
						goto l1010
					l1011:
						position, tokenIndex, depth = position1010, tokenIndex1010, depth1010
						if buffer[position] != rune('T') {
							goto l999
						}
						position++
					}
				l1010:
					depth--
					add(rulePegText, position1001)
				}
				if !_rules[ruleAction70]() {
					goto l999
				}
				depth--
				add(ruleFloat, position1000)
			}
			return true
		l999:
			position, tokenIndex, depth = position999, tokenIndex999, depth999
			return false
		},
		/* 94 String <- <(<(('s' / 'S') ('t' / 'T') ('r' / 'R') ('i' / 'I') ('n' / 'N') ('g' / 'G'))> Action71)> */
		func() bool {
			position1012, tokenIndex1012, depth1012 := position, tokenIndex, depth
			{
				position1013 := position
				depth++
				{
					position1014 := position
					depth++
					{
						position1015, tokenIndex1015, depth1015 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1016
						}
						position++
						goto l1015
					l1016:
						position, tokenIndex, depth = position1015, tokenIndex1015, depth1015
						if buffer[position] != rune('S') {
							goto l1012
						}
						position++
					}
				l1015:
					{
						position1017, tokenIndex1017, depth1017 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1018
						}
						position++
						goto l1017
					l1018:
						position, tokenIndex, depth = position1017, tokenIndex1017, depth1017
						if buffer[position] != rune('T') {
							goto l1012
						}
						position++
					}
				l1017:
					{
						position1019, tokenIndex1019, depth1019 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1020
						}
						position++
						goto l1019
					l1020:
						position, tokenIndex, depth = position1019, tokenIndex1019, depth1019
						if buffer[position] != rune('R') {
							goto l1012
						}
						position++
					}
				l1019:
					{
						position1021, tokenIndex1021, depth1021 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1022
						}
						position++
						goto l1021
					l1022:
						position, tokenIndex, depth = position1021, tokenIndex1021, depth1021
						if buffer[position] != rune('I') {
							goto l1012
						}
						position++
					}
				l1021:
					{
						position1023, tokenIndex1023, depth1023 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1024
						}
						position++
						goto l1023
					l1024:
						position, tokenIndex, depth = position1023, tokenIndex1023, depth1023
						if buffer[position] != rune('N') {
							goto l1012
						}
						position++
					}
				l1023:
					{
						position1025, tokenIndex1025, depth1025 := position, tokenIndex, depth
						if buffer[position] != rune('g') {
							goto l1026
						}
						position++
						goto l1025
					l1026:
						position, tokenIndex, depth = position1025, tokenIndex1025, depth1025
						if buffer[position] != rune('G') {
							goto l1012
						}
						position++
					}
				l1025:
					depth--
					add(rulePegText, position1014)
				}
				if !_rules[ruleAction71]() {
					goto l1012
				}
				depth--
				add(ruleString, position1013)
			}
			return true
		l1012:
			position, tokenIndex, depth = position1012, tokenIndex1012, depth1012
			return false
		},
		/* 95 Blob <- <(<(('b' / 'B') ('l' / 'L') ('o' / 'O') ('b' / 'B'))> Action72)> */
		func() bool {
			position1027, tokenIndex1027, depth1027 := position, tokenIndex, depth
			{
				position1028 := position
				depth++
				{
					position1029 := position
					depth++
					{
						position1030, tokenIndex1030, depth1030 := position, tokenIndex, depth
						if buffer[position] != rune('b') {
							goto l1031
						}
						position++
						goto l1030
					l1031:
						position, tokenIndex, depth = position1030, tokenIndex1030, depth1030
						if buffer[position] != rune('B') {
							goto l1027
						}
						position++
					}
				l1030:
					{
						position1032, tokenIndex1032, depth1032 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1033
						}
						position++
						goto l1032
					l1033:
						position, tokenIndex, depth = position1032, tokenIndex1032, depth1032
						if buffer[position] != rune('L') {
							goto l1027
						}
						position++
					}
				l1032:
					{
						position1034, tokenIndex1034, depth1034 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1035
						}
						position++
						goto l1034
					l1035:
						position, tokenIndex, depth = position1034, tokenIndex1034, depth1034
						if buffer[position] != rune('O') {
							goto l1027
						}
						position++
					}
				l1034:
					{
						position1036, tokenIndex1036, depth1036 := position, tokenIndex, depth
						if buffer[position] != rune('b') {
							goto l1037
						}
						position++
						goto l1036
					l1037:
						position, tokenIndex, depth = position1036, tokenIndex1036, depth1036
						if buffer[position] != rune('B') {
							goto l1027
						}
						position++
					}
				l1036:
					depth--
					add(rulePegText, position1029)
				}
				if !_rules[ruleAction72]() {
					goto l1027
				}
				depth--
				add(ruleBlob, position1028)
			}
			return true
		l1027:
			position, tokenIndex, depth = position1027, tokenIndex1027, depth1027
			return false
		},
		/* 96 Timestamp <- <(<(('t' / 'T') ('i' / 'I') ('m' / 'M') ('e' / 'E') ('s' / 'S') ('t' / 'T') ('a' / 'A') ('m' / 'M') ('p' / 'P'))> Action73)> */
		func() bool {
			position1038, tokenIndex1038, depth1038 := position, tokenIndex, depth
			{
				position1039 := position
				depth++
				{
					position1040 := position
					depth++
					{
						position1041, tokenIndex1041, depth1041 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1042
						}
						position++
						goto l1041
					l1042:
						position, tokenIndex, depth = position1041, tokenIndex1041, depth1041
						if buffer[position] != rune('T') {
							goto l1038
						}
						position++
					}
				l1041:
					{
						position1043, tokenIndex1043, depth1043 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1044
						}
						position++
						goto l1043
					l1044:
						position, tokenIndex, depth = position1043, tokenIndex1043, depth1043
						if buffer[position] != rune('I') {
							goto l1038
						}
						position++
					}
				l1043:
					{
						position1045, tokenIndex1045, depth1045 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1046
						}
						position++
						goto l1045
					l1046:
						position, tokenIndex, depth = position1045, tokenIndex1045, depth1045
						if buffer[position] != rune('M') {
							goto l1038
						}
						position++
					}
				l1045:
					{
						position1047, tokenIndex1047, depth1047 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1048
						}
						position++
						goto l1047
					l1048:
						position, tokenIndex, depth = position1047, tokenIndex1047, depth1047
						if buffer[position] != rune('E') {
							goto l1038
						}
						position++
					}
				l1047:
					{
						position1049, tokenIndex1049, depth1049 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1050
						}
						position++
						goto l1049
					l1050:
						position, tokenIndex, depth = position1049, tokenIndex1049, depth1049
						if buffer[position] != rune('S') {
							goto l1038
						}
						position++
					}
				l1049:
					{
						position1051, tokenIndex1051, depth1051 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1052
						}
						position++
						goto l1051
					l1052:
						position, tokenIndex, depth = position1051, tokenIndex1051, depth1051
						if buffer[position] != rune('T') {
							goto l1038
						}
						position++
					}
				l1051:
					{
						position1053, tokenIndex1053, depth1053 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1054
						}
						position++
						goto l1053
					l1054:
						position, tokenIndex, depth = position1053, tokenIndex1053, depth1053
						if buffer[position] != rune('A') {
							goto l1038
						}
						position++
					}
				l1053:
					{
						position1055, tokenIndex1055, depth1055 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1056
						}
						position++
						goto l1055
					l1056:
						position, tokenIndex, depth = position1055, tokenIndex1055, depth1055
						if buffer[position] != rune('M') {
							goto l1038
						}
						position++
					}
				l1055:
					{
						position1057, tokenIndex1057, depth1057 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l1058
						}
						position++
						goto l1057
					l1058:
						position, tokenIndex, depth = position1057, tokenIndex1057, depth1057
						if buffer[position] != rune('P') {
							goto l1038
						}
						position++
					}
				l1057:
					depth--
					add(rulePegText, position1040)
				}
				if !_rules[ruleAction73]() {
					goto l1038
				}
				depth--
				add(ruleTimestamp, position1039)
			}
			return true
		l1038:
			position, tokenIndex, depth = position1038, tokenIndex1038, depth1038
			return false
		},
		/* 97 Array <- <(<(('a' / 'A') ('r' / 'R') ('r' / 'R') ('a' / 'A') ('y' / 'Y'))> Action74)> */
		func() bool {
			position1059, tokenIndex1059, depth1059 := position, tokenIndex, depth
			{
				position1060 := position
				depth++
				{
					position1061 := position
					depth++
					{
						position1062, tokenIndex1062, depth1062 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1063
						}
						position++
						goto l1062
					l1063:
						position, tokenIndex, depth = position1062, tokenIndex1062, depth1062
						if buffer[position] != rune('A') {
							goto l1059
						}
						position++
					}
				l1062:
					{
						position1064, tokenIndex1064, depth1064 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1065
						}
						position++
						goto l1064
					l1065:
						position, tokenIndex, depth = position1064, tokenIndex1064, depth1064
						if buffer[position] != rune('R') {
							goto l1059
						}
						position++
					}
				l1064:
					{
						position1066, tokenIndex1066, depth1066 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1067
						}
						position++
						goto l1066
					l1067:
						position, tokenIndex, depth = position1066, tokenIndex1066, depth1066
						if buffer[position] != rune('R') {
							goto l1059
						}
						position++
					}
				l1066:
					{
						position1068, tokenIndex1068, depth1068 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1069
						}
						position++
						goto l1068
					l1069:
						position, tokenIndex, depth = position1068, tokenIndex1068, depth1068
						if buffer[position] != rune('A') {
							goto l1059
						}
						position++
					}
				l1068:
					{
						position1070, tokenIndex1070, depth1070 := position, tokenIndex, depth
						if buffer[position] != rune('y') {
							goto l1071
						}
						position++
						goto l1070
					l1071:
						position, tokenIndex, depth = position1070, tokenIndex1070, depth1070
						if buffer[position] != rune('Y') {
							goto l1059
						}
						position++
					}
				l1070:
					depth--
					add(rulePegText, position1061)
				}
				if !_rules[ruleAction74]() {
					goto l1059
				}
				depth--
				add(ruleArray, position1060)
			}
			return true
		l1059:
			position, tokenIndex, depth = position1059, tokenIndex1059, depth1059
			return false
		},
		/* 98 Map <- <(<(('m' / 'M') ('a' / 'A') ('p' / 'P'))> Action75)> */
		func() bool {
			position1072, tokenIndex1072, depth1072 := position, tokenIndex, depth
			{
				position1073 := position
				depth++
				{
					position1074 := position
					depth++
					{
						position1075, tokenIndex1075, depth1075 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1076
						}
						position++
						goto l1075
					l1076:
						position, tokenIndex, depth = position1075, tokenIndex1075, depth1075
						if buffer[position] != rune('M') {
							goto l1072
						}
						position++
					}
				l1075:
					{
						position1077, tokenIndex1077, depth1077 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1078
						}
						position++
						goto l1077
					l1078:
						position, tokenIndex, depth = position1077, tokenIndex1077, depth1077
						if buffer[position] != rune('A') {
							goto l1072
						}
						position++
					}
				l1077:
					{
						position1079, tokenIndex1079, depth1079 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l1080
						}
						position++
						goto l1079
					l1080:
						position, tokenIndex, depth = position1079, tokenIndex1079, depth1079
						if buffer[position] != rune('P') {
							goto l1072
						}
						position++
					}
				l1079:
					depth--
					add(rulePegText, position1074)
				}
				if !_rules[ruleAction75]() {
					goto l1072
				}
				depth--
				add(ruleMap, position1073)
			}
			return true
		l1072:
			position, tokenIndex, depth = position1072, tokenIndex1072, depth1072
			return false
		},
		/* 99 Or <- <(<(('o' / 'O') ('r' / 'R'))> Action76)> */
		func() bool {
			position1081, tokenIndex1081, depth1081 := position, tokenIndex, depth
			{
				position1082 := position
				depth++
				{
					position1083 := position
					depth++
					{
						position1084, tokenIndex1084, depth1084 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1085
						}
						position++
						goto l1084
					l1085:
						position, tokenIndex, depth = position1084, tokenIndex1084, depth1084
						if buffer[position] != rune('O') {
							goto l1081
						}
						position++
					}
				l1084:
					{
						position1086, tokenIndex1086, depth1086 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1087
						}
						position++
						goto l1086
					l1087:
						position, tokenIndex, depth = position1086, tokenIndex1086, depth1086
						if buffer[position] != rune('R') {
							goto l1081
						}
						position++
					}
				l1086:
					depth--
					add(rulePegText, position1083)
				}
				if !_rules[ruleAction76]() {
					goto l1081
				}
				depth--
				add(ruleOr, position1082)
			}
			return true
		l1081:
			position, tokenIndex, depth = position1081, tokenIndex1081, depth1081
			return false
		},
		/* 100 And <- <(<(('a' / 'A') ('n' / 'N') ('d' / 'D'))> Action77)> */
		func() bool {
			position1088, tokenIndex1088, depth1088 := position, tokenIndex, depth
			{
				position1089 := position
				depth++
				{
					position1090 := position
					depth++
					{
						position1091, tokenIndex1091, depth1091 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1092
						}
						position++
						goto l1091
					l1092:
						position, tokenIndex, depth = position1091, tokenIndex1091, depth1091
						if buffer[position] != rune('A') {
							goto l1088
						}
						position++
					}
				l1091:
					{
						position1093, tokenIndex1093, depth1093 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1094
						}
						position++
						goto l1093
					l1094:
						position, tokenIndex, depth = position1093, tokenIndex1093, depth1093
						if buffer[position] != rune('N') {
							goto l1088
						}
						position++
					}
				l1093:
					{
						position1095, tokenIndex1095, depth1095 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l1096
						}
						position++
						goto l1095
					l1096:
						position, tokenIndex, depth = position1095, tokenIndex1095, depth1095
						if buffer[position] != rune('D') {
							goto l1088
						}
						position++
					}
				l1095:
					depth--
					add(rulePegText, position1090)
				}
				if !_rules[ruleAction77]() {
					goto l1088
				}
				depth--
				add(ruleAnd, position1089)
			}
			return true
		l1088:
			position, tokenIndex, depth = position1088, tokenIndex1088, depth1088
			return false
		},
		/* 101 Not <- <(<(('n' / 'N') ('o' / 'O') ('t' / 'T'))> Action78)> */
		func() bool {
			position1097, tokenIndex1097, depth1097 := position, tokenIndex, depth
			{
				position1098 := position
				depth++
				{
					position1099 := position
					depth++
					{
						position1100, tokenIndex1100, depth1100 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1101
						}
						position++
						goto l1100
					l1101:
						position, tokenIndex, depth = position1100, tokenIndex1100, depth1100
						if buffer[position] != rune('N') {
							goto l1097
						}
						position++
					}
				l1100:
					{
						position1102, tokenIndex1102, depth1102 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1103
						}
						position++
						goto l1102
					l1103:
						position, tokenIndex, depth = position1102, tokenIndex1102, depth1102
						if buffer[position] != rune('O') {
							goto l1097
						}
						position++
					}
				l1102:
					{
						position1104, tokenIndex1104, depth1104 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1105
						}
						position++
						goto l1104
					l1105:
						position, tokenIndex, depth = position1104, tokenIndex1104, depth1104
						if buffer[position] != rune('T') {
							goto l1097
						}
						position++
					}
				l1104:
					depth--
					add(rulePegText, position1099)
				}
				if !_rules[ruleAction78]() {
					goto l1097
				}
				depth--
				add(ruleNot, position1098)
			}
			return true
		l1097:
			position, tokenIndex, depth = position1097, tokenIndex1097, depth1097
			return false
		},
		/* 102 Equal <- <(<'='> Action79)> */
		func() bool {
			position1106, tokenIndex1106, depth1106 := position, tokenIndex, depth
			{
				position1107 := position
				depth++
				{
					position1108 := position
					depth++
					if buffer[position] != rune('=') {
						goto l1106
					}
					position++
					depth--
					add(rulePegText, position1108)
				}
				if !_rules[ruleAction79]() {
					goto l1106
				}
				depth--
				add(ruleEqual, position1107)
			}
			return true
		l1106:
			position, tokenIndex, depth = position1106, tokenIndex1106, depth1106
			return false
		},
		/* 103 Less <- <(<'<'> Action80)> */
		func() bool {
			position1109, tokenIndex1109, depth1109 := position, tokenIndex, depth
			{
				position1110 := position
				depth++
				{
					position1111 := position
					depth++
					if buffer[position] != rune('<') {
						goto l1109
					}
					position++
					depth--
					add(rulePegText, position1111)
				}
				if !_rules[ruleAction80]() {
					goto l1109
				}
				depth--
				add(ruleLess, position1110)
			}
			return true
		l1109:
			position, tokenIndex, depth = position1109, tokenIndex1109, depth1109
			return false
		},
		/* 104 LessOrEqual <- <(<('<' '=')> Action81)> */
		func() bool {
			position1112, tokenIndex1112, depth1112 := position, tokenIndex, depth
			{
				position1113 := position
				depth++
				{
					position1114 := position
					depth++
					if buffer[position] != rune('<') {
						goto l1112
					}
					position++
					if buffer[position] != rune('=') {
						goto l1112
					}
					position++
					depth--
					add(rulePegText, position1114)
				}
				if !_rules[ruleAction81]() {
					goto l1112
				}
				depth--
				add(ruleLessOrEqual, position1113)
			}
			return true
		l1112:
			position, tokenIndex, depth = position1112, tokenIndex1112, depth1112
			return false
		},
		/* 105 Greater <- <(<'>'> Action82)> */
		func() bool {
			position1115, tokenIndex1115, depth1115 := position, tokenIndex, depth
			{
				position1116 := position
				depth++
				{
					position1117 := position
					depth++
					if buffer[position] != rune('>') {
						goto l1115
					}
					position++
					depth--
					add(rulePegText, position1117)
				}
				if !_rules[ruleAction82]() {
					goto l1115
				}
				depth--
				add(ruleGreater, position1116)
			}
			return true
		l1115:
			position, tokenIndex, depth = position1115, tokenIndex1115, depth1115
			return false
		},
		/* 106 GreaterOrEqual <- <(<('>' '=')> Action83)> */
		func() bool {
			position1118, tokenIndex1118, depth1118 := position, tokenIndex, depth
			{
				position1119 := position
				depth++
				{
					position1120 := position
					depth++
					if buffer[position] != rune('>') {
						goto l1118
					}
					position++
					if buffer[position] != rune('=') {
						goto l1118
					}
					position++
					depth--
					add(rulePegText, position1120)
				}
				if !_rules[ruleAction83]() {
					goto l1118
				}
				depth--
				add(ruleGreaterOrEqual, position1119)
			}
			return true
		l1118:
			position, tokenIndex, depth = position1118, tokenIndex1118, depth1118
			return false
		},
		/* 107 NotEqual <- <(<(('!' '=') / ('<' '>'))> Action84)> */
		func() bool {
			position1121, tokenIndex1121, depth1121 := position, tokenIndex, depth
			{
				position1122 := position
				depth++
				{
					position1123 := position
					depth++
					{
						position1124, tokenIndex1124, depth1124 := position, tokenIndex, depth
						if buffer[position] != rune('!') {
							goto l1125
						}
						position++
						if buffer[position] != rune('=') {
							goto l1125
						}
						position++
						goto l1124
					l1125:
						position, tokenIndex, depth = position1124, tokenIndex1124, depth1124
						if buffer[position] != rune('<') {
							goto l1121
						}
						position++
						if buffer[position] != rune('>') {
							goto l1121
						}
						position++
					}
				l1124:
					depth--
					add(rulePegText, position1123)
				}
				if !_rules[ruleAction84]() {
					goto l1121
				}
				depth--
				add(ruleNotEqual, position1122)
			}
			return true
		l1121:
			position, tokenIndex, depth = position1121, tokenIndex1121, depth1121
			return false
		},
		/* 108 Concat <- <(<('|' '|')> Action85)> */
		func() bool {
			position1126, tokenIndex1126, depth1126 := position, tokenIndex, depth
			{
				position1127 := position
				depth++
				{
					position1128 := position
					depth++
					if buffer[position] != rune('|') {
						goto l1126
					}
					position++
					if buffer[position] != rune('|') {
						goto l1126
					}
					position++
					depth--
					add(rulePegText, position1128)
				}
				if !_rules[ruleAction85]() {
					goto l1126
				}
				depth--
				add(ruleConcat, position1127)
			}
			return true
		l1126:
			position, tokenIndex, depth = position1126, tokenIndex1126, depth1126
			return false
		},
		/* 109 Is <- <(<(('i' / 'I') ('s' / 'S'))> Action86)> */
		func() bool {
			position1129, tokenIndex1129, depth1129 := position, tokenIndex, depth
			{
				position1130 := position
				depth++
				{
					position1131 := position
					depth++
					{
						position1132, tokenIndex1132, depth1132 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1133
						}
						position++
						goto l1132
					l1133:
						position, tokenIndex, depth = position1132, tokenIndex1132, depth1132
						if buffer[position] != rune('I') {
							goto l1129
						}
						position++
					}
				l1132:
					{
						position1134, tokenIndex1134, depth1134 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1135
						}
						position++
						goto l1134
					l1135:
						position, tokenIndex, depth = position1134, tokenIndex1134, depth1134
						if buffer[position] != rune('S') {
							goto l1129
						}
						position++
					}
				l1134:
					depth--
					add(rulePegText, position1131)
				}
				if !_rules[ruleAction86]() {
					goto l1129
				}
				depth--
				add(ruleIs, position1130)
			}
			return true
		l1129:
			position, tokenIndex, depth = position1129, tokenIndex1129, depth1129
			return false
		},
		/* 110 IsNot <- <(<(('i' / 'I') ('s' / 'S') sp (('n' / 'N') ('o' / 'O') ('t' / 'T')))> Action87)> */
		func() bool {
			position1136, tokenIndex1136, depth1136 := position, tokenIndex, depth
			{
				position1137 := position
				depth++
				{
					position1138 := position
					depth++
					{
						position1139, tokenIndex1139, depth1139 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1140
						}
						position++
						goto l1139
					l1140:
						position, tokenIndex, depth = position1139, tokenIndex1139, depth1139
						if buffer[position] != rune('I') {
							goto l1136
						}
						position++
					}
				l1139:
					{
						position1141, tokenIndex1141, depth1141 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1142
						}
						position++
						goto l1141
					l1142:
						position, tokenIndex, depth = position1141, tokenIndex1141, depth1141
						if buffer[position] != rune('S') {
							goto l1136
						}
						position++
					}
				l1141:
					if !_rules[rulesp]() {
						goto l1136
					}
					{
						position1143, tokenIndex1143, depth1143 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1144
						}
						position++
						goto l1143
					l1144:
						position, tokenIndex, depth = position1143, tokenIndex1143, depth1143
						if buffer[position] != rune('N') {
							goto l1136
						}
						position++
					}
				l1143:
					{
						position1145, tokenIndex1145, depth1145 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1146
						}
						position++
						goto l1145
					l1146:
						position, tokenIndex, depth = position1145, tokenIndex1145, depth1145
						if buffer[position] != rune('O') {
							goto l1136
						}
						position++
					}
				l1145:
					{
						position1147, tokenIndex1147, depth1147 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1148
						}
						position++
						goto l1147
					l1148:
						position, tokenIndex, depth = position1147, tokenIndex1147, depth1147
						if buffer[position] != rune('T') {
							goto l1136
						}
						position++
					}
				l1147:
					depth--
					add(rulePegText, position1138)
				}
				if !_rules[ruleAction87]() {
					goto l1136
				}
				depth--
				add(ruleIsNot, position1137)
			}
			return true
		l1136:
			position, tokenIndex, depth = position1136, tokenIndex1136, depth1136
			return false
		},
		/* 111 Plus <- <(<'+'> Action88)> */
		func() bool {
			position1149, tokenIndex1149, depth1149 := position, tokenIndex, depth
			{
				position1150 := position
				depth++
				{
					position1151 := position
					depth++
					if buffer[position] != rune('+') {
						goto l1149
					}
					position++
					depth--
					add(rulePegText, position1151)
				}
				if !_rules[ruleAction88]() {
					goto l1149
				}
				depth--
				add(rulePlus, position1150)
			}
			return true
		l1149:
			position, tokenIndex, depth = position1149, tokenIndex1149, depth1149
			return false
		},
		/* 112 Minus <- <(<'-'> Action89)> */
		func() bool {
			position1152, tokenIndex1152, depth1152 := position, tokenIndex, depth
			{
				position1153 := position
				depth++
				{
					position1154 := position
					depth++
					if buffer[position] != rune('-') {
						goto l1152
					}
					position++
					depth--
					add(rulePegText, position1154)
				}
				if !_rules[ruleAction89]() {
					goto l1152
				}
				depth--
				add(ruleMinus, position1153)
			}
			return true
		l1152:
			position, tokenIndex, depth = position1152, tokenIndex1152, depth1152
			return false
		},
		/* 113 Multiply <- <(<'*'> Action90)> */
		func() bool {
			position1155, tokenIndex1155, depth1155 := position, tokenIndex, depth
			{
				position1156 := position
				depth++
				{
					position1157 := position
					depth++
					if buffer[position] != rune('*') {
						goto l1155
					}
					position++
					depth--
					add(rulePegText, position1157)
				}
				if !_rules[ruleAction90]() {
					goto l1155
				}
				depth--
				add(ruleMultiply, position1156)
			}
			return true
		l1155:
			position, tokenIndex, depth = position1155, tokenIndex1155, depth1155
			return false
		},
		/* 114 Divide <- <(<'/'> Action91)> */
		func() bool {
			position1158, tokenIndex1158, depth1158 := position, tokenIndex, depth
			{
				position1159 := position
				depth++
				{
					position1160 := position
					depth++
					if buffer[position] != rune('/') {
						goto l1158
					}
					position++
					depth--
					add(rulePegText, position1160)
				}
				if !_rules[ruleAction91]() {
					goto l1158
				}
				depth--
				add(ruleDivide, position1159)
			}
			return true
		l1158:
			position, tokenIndex, depth = position1158, tokenIndex1158, depth1158
			return false
		},
		/* 115 Modulo <- <(<'%'> Action92)> */
		func() bool {
			position1161, tokenIndex1161, depth1161 := position, tokenIndex, depth
			{
				position1162 := position
				depth++
				{
					position1163 := position
					depth++
					if buffer[position] != rune('%') {
						goto l1161
					}
					position++
					depth--
					add(rulePegText, position1163)
				}
				if !_rules[ruleAction92]() {
					goto l1161
				}
				depth--
				add(ruleModulo, position1162)
			}
			return true
		l1161:
			position, tokenIndex, depth = position1161, tokenIndex1161, depth1161
			return false
		},
		/* 116 UnaryMinus <- <(<'-'> Action93)> */
		func() bool {
			position1164, tokenIndex1164, depth1164 := position, tokenIndex, depth
			{
				position1165 := position
				depth++
				{
					position1166 := position
					depth++
					if buffer[position] != rune('-') {
						goto l1164
					}
					position++
					depth--
					add(rulePegText, position1166)
				}
				if !_rules[ruleAction93]() {
					goto l1164
				}
				depth--
				add(ruleUnaryMinus, position1165)
			}
			return true
		l1164:
			position, tokenIndex, depth = position1164, tokenIndex1164, depth1164
			return false
		},
		/* 117 Identifier <- <(<ident> Action94)> */
		func() bool {
			position1167, tokenIndex1167, depth1167 := position, tokenIndex, depth
			{
				position1168 := position
				depth++
				{
					position1169 := position
					depth++
					if !_rules[ruleident]() {
						goto l1167
					}
					depth--
					add(rulePegText, position1169)
				}
				if !_rules[ruleAction94]() {
					goto l1167
				}
				depth--
				add(ruleIdentifier, position1168)
			}
			return true
		l1167:
			position, tokenIndex, depth = position1167, tokenIndex1167, depth1167
			return false
		},
		/* 118 TargetIdentifier <- <(<jsonPath> Action95)> */
		func() bool {
			position1170, tokenIndex1170, depth1170 := position, tokenIndex, depth
			{
				position1171 := position
				depth++
				{
					position1172 := position
					depth++
					if !_rules[rulejsonPath]() {
						goto l1170
					}
					depth--
					add(rulePegText, position1172)
				}
				if !_rules[ruleAction95]() {
					goto l1170
				}
				depth--
				add(ruleTargetIdentifier, position1171)
			}
			return true
		l1170:
			position, tokenIndex, depth = position1170, tokenIndex1170, depth1170
			return false
		},
		/* 119 ident <- <(([a-z] / [A-Z]) ([a-z] / [A-Z] / [0-9] / '_')*)> */
		func() bool {
			position1173, tokenIndex1173, depth1173 := position, tokenIndex, depth
			{
				position1174 := position
				depth++
				{
					position1175, tokenIndex1175, depth1175 := position, tokenIndex, depth
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1176
					}
					position++
					goto l1175
				l1176:
					position, tokenIndex, depth = position1175, tokenIndex1175, depth1175
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1173
					}
					position++
				}
			l1175:
			l1177:
				{
					position1178, tokenIndex1178, depth1178 := position, tokenIndex, depth
					{
						position1179, tokenIndex1179, depth1179 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('a') || c > rune('z') {
							goto l1180
						}
						position++
						goto l1179
					l1180:
						position, tokenIndex, depth = position1179, tokenIndex1179, depth1179
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l1181
						}
						position++
						goto l1179
					l1181:
						position, tokenIndex, depth = position1179, tokenIndex1179, depth1179
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1182
						}
						position++
						goto l1179
					l1182:
						position, tokenIndex, depth = position1179, tokenIndex1179, depth1179
						if buffer[position] != rune('_') {
							goto l1178
						}
						position++
					}
				l1179:
					goto l1177
				l1178:
					position, tokenIndex, depth = position1178, tokenIndex1178, depth1178
				}
				depth--
				add(ruleident, position1174)
			}
			return true
		l1173:
			position, tokenIndex, depth = position1173, tokenIndex1173, depth1173
			return false
		},
		/* 120 jsonPath <- <(([a-z] / [A-Z]) ([a-z] / [A-Z] / [0-9] / '_' / '.' / '[' / ']' / '"')*)> */
		func() bool {
			position1183, tokenIndex1183, depth1183 := position, tokenIndex, depth
			{
				position1184 := position
				depth++
				{
					position1185, tokenIndex1185, depth1185 := position, tokenIndex, depth
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1186
					}
					position++
					goto l1185
				l1186:
					position, tokenIndex, depth = position1185, tokenIndex1185, depth1185
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1183
					}
					position++
				}
			l1185:
			l1187:
				{
					position1188, tokenIndex1188, depth1188 := position, tokenIndex, depth
					{
						position1189, tokenIndex1189, depth1189 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('a') || c > rune('z') {
							goto l1190
						}
						position++
						goto l1189
					l1190:
						position, tokenIndex, depth = position1189, tokenIndex1189, depth1189
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l1191
						}
						position++
						goto l1189
					l1191:
						position, tokenIndex, depth = position1189, tokenIndex1189, depth1189
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1192
						}
						position++
						goto l1189
					l1192:
						position, tokenIndex, depth = position1189, tokenIndex1189, depth1189
						if buffer[position] != rune('_') {
							goto l1193
						}
						position++
						goto l1189
					l1193:
						position, tokenIndex, depth = position1189, tokenIndex1189, depth1189
						if buffer[position] != rune('.') {
							goto l1194
						}
						position++
						goto l1189
					l1194:
						position, tokenIndex, depth = position1189, tokenIndex1189, depth1189
						if buffer[position] != rune('[') {
							goto l1195
						}
						position++
						goto l1189
					l1195:
						position, tokenIndex, depth = position1189, tokenIndex1189, depth1189
						if buffer[position] != rune(']') {
							goto l1196
						}
						position++
						goto l1189
					l1196:
						position, tokenIndex, depth = position1189, tokenIndex1189, depth1189
						if buffer[position] != rune('"') {
							goto l1188
						}
						position++
					}
				l1189:
					goto l1187
				l1188:
					position, tokenIndex, depth = position1188, tokenIndex1188, depth1188
				}
				depth--
				add(rulejsonPath, position1184)
			}
			return true
		l1183:
			position, tokenIndex, depth = position1183, tokenIndex1183, depth1183
			return false
		},
		/* 121 sp <- <(' ' / '\t' / '\n' / '\r' / comment)*> */
		func() bool {
			{
				position1198 := position
				depth++
			l1199:
				{
					position1200, tokenIndex1200, depth1200 := position, tokenIndex, depth
					{
						position1201, tokenIndex1201, depth1201 := position, tokenIndex, depth
						if buffer[position] != rune(' ') {
							goto l1202
						}
						position++
						goto l1201
					l1202:
						position, tokenIndex, depth = position1201, tokenIndex1201, depth1201
						if buffer[position] != rune('\t') {
							goto l1203
						}
						position++
						goto l1201
					l1203:
						position, tokenIndex, depth = position1201, tokenIndex1201, depth1201
						if buffer[position] != rune('\n') {
							goto l1204
						}
						position++
						goto l1201
					l1204:
						position, tokenIndex, depth = position1201, tokenIndex1201, depth1201
						if buffer[position] != rune('\r') {
							goto l1205
						}
						position++
						goto l1201
					l1205:
						position, tokenIndex, depth = position1201, tokenIndex1201, depth1201
						if !_rules[rulecomment]() {
							goto l1200
						}
					}
				l1201:
					goto l1199
				l1200:
					position, tokenIndex, depth = position1200, tokenIndex1200, depth1200
				}
				depth--
				add(rulesp, position1198)
			}
			return true
		},
		/* 122 comment <- <('-' '-' (!('\r' / '\n') .)* ('\r' / '\n'))> */
		func() bool {
			position1206, tokenIndex1206, depth1206 := position, tokenIndex, depth
			{
				position1207 := position
				depth++
				if buffer[position] != rune('-') {
					goto l1206
				}
				position++
				if buffer[position] != rune('-') {
					goto l1206
				}
				position++
			l1208:
				{
					position1209, tokenIndex1209, depth1209 := position, tokenIndex, depth
					{
						position1210, tokenIndex1210, depth1210 := position, tokenIndex, depth
						{
							position1211, tokenIndex1211, depth1211 := position, tokenIndex, depth
							if buffer[position] != rune('\r') {
								goto l1212
							}
							position++
							goto l1211
						l1212:
							position, tokenIndex, depth = position1211, tokenIndex1211, depth1211
							if buffer[position] != rune('\n') {
								goto l1210
							}
							position++
						}
					l1211:
						goto l1209
					l1210:
						position, tokenIndex, depth = position1210, tokenIndex1210, depth1210
					}
					if !matchDot() {
						goto l1209
					}
					goto l1208
				l1209:
					position, tokenIndex, depth = position1209, tokenIndex1209, depth1209
				}
				{
					position1213, tokenIndex1213, depth1213 := position, tokenIndex, depth
					if buffer[position] != rune('\r') {
						goto l1214
					}
					position++
					goto l1213
				l1214:
					position, tokenIndex, depth = position1213, tokenIndex1213, depth1213
					if buffer[position] != rune('\n') {
						goto l1206
					}
					position++
				}
			l1213:
				depth--
				add(rulecomment, position1207)
			}
			return true
		l1206:
			position, tokenIndex, depth = position1206, tokenIndex1206, depth1206
			return false
		},
		/* 124 Action0 <- <{
		    p.AssembleSelect()
		}> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
		/* 125 Action1 <- <{
		    p.AssembleCreateStreamAsSelect()
		}> */
		func() bool {
			{
				add(ruleAction1, position)
			}
			return true
		},
		/* 126 Action2 <- <{
		    p.AssembleCreateSource()
		}> */
		func() bool {
			{
				add(ruleAction2, position)
			}
			return true
		},
		/* 127 Action3 <- <{
		    p.AssembleCreateSink()
		}> */
		func() bool {
			{
				add(ruleAction3, position)
			}
			return true
		},
		/* 128 Action4 <- <{
		    p.AssembleCreateState()
		}> */
		func() bool {
			{
				add(ruleAction4, position)
			}
			return true
		},
		/* 129 Action5 <- <{
		    p.AssembleUpdateState()
		}> */
		func() bool {
			{
				add(ruleAction5, position)
			}
			return true
		},
		/* 130 Action6 <- <{
		    p.AssembleUpdateSource()
		}> */
		func() bool {
			{
				add(ruleAction6, position)
			}
			return true
		},
		/* 131 Action7 <- <{
		    p.AssembleUpdateSink()
		}> */
		func() bool {
			{
				add(ruleAction7, position)
			}
			return true
		},
		/* 132 Action8 <- <{
		    p.AssembleInsertIntoSelect()
		}> */
		func() bool {
			{
				add(ruleAction8, position)
			}
			return true
		},
		/* 133 Action9 <- <{
		    p.AssembleInsertIntoFrom()
		}> */
		func() bool {
			{
				add(ruleAction9, position)
			}
			return true
		},
		/* 134 Action10 <- <{
		    p.AssemblePauseSource()
		}> */
		func() bool {
			{
				add(ruleAction10, position)
			}
			return true
		},
		/* 135 Action11 <- <{
		    p.AssembleResumeSource()
		}> */
		func() bool {
			{
				add(ruleAction11, position)
			}
			return true
		},
		/* 136 Action12 <- <{
		    p.AssembleRewindSource()
		}> */
		func() bool {
			{
				add(ruleAction12, position)
			}
			return true
		},
		/* 137 Action13 <- <{
		    p.AssembleDropSource()
		}> */
		func() bool {
			{
				add(ruleAction13, position)
			}
			return true
		},
		/* 138 Action14 <- <{
		    p.AssembleDropStream()
		}> */
		func() bool {
			{
				add(ruleAction14, position)
			}
			return true
		},
		/* 139 Action15 <- <{
		    p.AssembleDropSink()
		}> */
		func() bool {
			{
				add(ruleAction15, position)
			}
			return true
		},
		/* 140 Action16 <- <{
		    p.AssembleDropState()
		}> */
		func() bool {
			{
				add(ruleAction16, position)
			}
			return true
		},
		/* 141 Action17 <- <{
		    p.AssembleEmitter()
		}> */
		func() bool {
			{
				add(ruleAction17, position)
			}
			return true
		},
		nil,
		/* 143 Action18 <- <{
		    p.AssembleProjections(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction18, position)
			}
			return true
		},
		/* 144 Action19 <- <{
		    p.AssembleAlias()
		}> */
		func() bool {
			{
				add(ruleAction19, position)
			}
			return true
		},
		/* 145 Action20 <- <{
		    // This is *always* executed, even if there is no
		    // FROM clause present in the statement.
		    p.AssembleWindowedFrom(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction20, position)
			}
			return true
		},
		/* 146 Action21 <- <{
		    p.AssembleInterval()
		}> */
		func() bool {
			{
				add(ruleAction21, position)
			}
			return true
		},
		/* 147 Action22 <- <{
		    p.AssembleInterval()
		}> */
		func() bool {
			{
				add(ruleAction22, position)
			}
			return true
		},
		/* 148 Action23 <- <{
		    // This is *always* executed, even if there is no
		    // WHERE clause present in the statement.
		    p.AssembleFilter(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction23, position)
			}
			return true
		},
		/* 149 Action24 <- <{
		    // This is *always* executed, even if there is no
		    // GROUP BY clause present in the statement.
		    p.AssembleGrouping(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction24, position)
			}
			return true
		},
		/* 150 Action25 <- <{
		    // This is *always* executed, even if there is no
		    // HAVING clause present in the statement.
		    p.AssembleHaving(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction25, position)
			}
			return true
		},
		/* 151 Action26 <- <{
		    p.EnsureAliasedStreamWindow()
		}> */
		func() bool {
			{
				add(ruleAction26, position)
			}
			return true
		},
		/* 152 Action27 <- <{
		    p.AssembleAliasedStreamWindow()
		}> */
		func() bool {
			{
				add(ruleAction27, position)
			}
			return true
		},
		/* 153 Action28 <- <{
		    p.AssembleStreamWindow()
		}> */
		func() bool {
			{
				add(ruleAction28, position)
			}
			return true
		},
		/* 154 Action29 <- <{
		    p.AssembleUDSFFuncApp()
		}> */
		func() bool {
			{
				add(ruleAction29, position)
			}
			return true
		},
		/* 155 Action30 <- <{
		    p.AssembleSourceSinkSpecs(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction30, position)
			}
			return true
		},
		/* 156 Action31 <- <{
		    p.AssembleSourceSinkSpecs(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction31, position)
			}
			return true
		},
		/* 157 Action32 <- <{
		    p.AssembleSourceSinkParam()
		}> */
		func() bool {
			{
				add(ruleAction32, position)
			}
			return true
		},
		/* 158 Action33 <- <{
		    p.EnsureKeywordPresent(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction33, position)
			}
			return true
		},
		/* 159 Action34 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction34, position)
			}
			return true
		},
		/* 160 Action35 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction35, position)
			}
			return true
		},
		/* 161 Action36 <- <{
		    p.AssembleUnaryPrefixOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction36, position)
			}
			return true
		},
		/* 162 Action37 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction37, position)
			}
			return true
		},
		/* 163 Action38 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction38, position)
			}
			return true
		},
		/* 164 Action39 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction39, position)
			}
			return true
		},
		/* 165 Action40 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction40, position)
			}
			return true
		},
		/* 166 Action41 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction41, position)
			}
			return true
		},
		/* 167 Action42 <- <{
		    p.AssembleUnaryPrefixOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction42, position)
			}
			return true
		},
		/* 168 Action43 <- <{
		    p.AssembleTypeCast(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction43, position)
			}
			return true
		},
		/* 169 Action44 <- <{
		    p.AssembleTypeCast(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction44, position)
			}
			return true
		},
		/* 170 Action45 <- <{
		    p.AssembleFuncApp()
		}> */
		func() bool {
			{
				add(ruleAction45, position)
			}
			return true
		},
		/* 171 Action46 <- <{
		    p.AssembleExpressions(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction46, position)
			}
			return true
		},
		/* 172 Action47 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewStream(substr))
		}> */
		func() bool {
			{
				add(ruleAction47, position)
			}
			return true
		},
		/* 173 Action48 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewRowMeta(substr, TimestampMeta))
		}> */
		func() bool {
			{
				add(ruleAction48, position)
			}
			return true
		},
		/* 174 Action49 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewRowValue(substr))
		}> */
		func() bool {
			{
				add(ruleAction49, position)
			}
			return true
		},
		/* 175 Action50 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewNumericLiteral(substr))
		}> */
		func() bool {
			{
				add(ruleAction50, position)
			}
			return true
		},
		/* 176 Action51 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewFloatLiteral(substr))
		}> */
		func() bool {
			{
				add(ruleAction51, position)
			}
			return true
		},
		/* 177 Action52 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, FuncName(substr))
		}> */
		func() bool {
			{
				add(ruleAction52, position)
			}
			return true
		},
		/* 178 Action53 <- <{
		    p.PushComponent(begin, end, NewNullLiteral())
		}> */
		func() bool {
			{
				add(ruleAction53, position)
			}
			return true
		},
		/* 179 Action54 <- <{
		    p.PushComponent(begin, end, NewBoolLiteral(true))
		}> */
		func() bool {
			{
				add(ruleAction54, position)
			}
			return true
		},
		/* 180 Action55 <- <{
		    p.PushComponent(begin, end, NewBoolLiteral(false))
		}> */
		func() bool {
			{
				add(ruleAction55, position)
			}
			return true
		},
		/* 181 Action56 <- <{
		    p.PushComponent(begin, end, NewWildcard())
		}> */
		func() bool {
			{
				add(ruleAction56, position)
			}
			return true
		},
		/* 182 Action57 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewStringLiteral(substr))
		}> */
		func() bool {
			{
				add(ruleAction57, position)
			}
			return true
		},
		/* 183 Action58 <- <{
		    p.PushComponent(begin, end, Istream)
		}> */
		func() bool {
			{
				add(ruleAction58, position)
			}
			return true
		},
		/* 184 Action59 <- <{
		    p.PushComponent(begin, end, Dstream)
		}> */
		func() bool {
			{
				add(ruleAction59, position)
			}
			return true
		},
		/* 185 Action60 <- <{
		    p.PushComponent(begin, end, Rstream)
		}> */
		func() bool {
			{
				add(ruleAction60, position)
			}
			return true
		},
		/* 186 Action61 <- <{
		    p.PushComponent(begin, end, Tuples)
		}> */
		func() bool {
			{
				add(ruleAction61, position)
			}
			return true
		},
		/* 187 Action62 <- <{
		    p.PushComponent(begin, end, Seconds)
		}> */
		func() bool {
			{
				add(ruleAction62, position)
			}
			return true
		},
		/* 188 Action63 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, StreamIdentifier(substr))
		}> */
		func() bool {
			{
				add(ruleAction63, position)
			}
			return true
		},
		/* 189 Action64 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, SourceSinkType(substr))
		}> */
		func() bool {
			{
				add(ruleAction64, position)
			}
			return true
		},
		/* 190 Action65 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, SourceSinkParamKey(substr))
		}> */
		func() bool {
			{
				add(ruleAction65, position)
			}
			return true
		},
		/* 191 Action66 <- <{
		    p.PushComponent(begin, end, Yes)
		}> */
		func() bool {
			{
				add(ruleAction66, position)
			}
			return true
		},
		/* 192 Action67 <- <{
		    p.PushComponent(begin, end, No)
		}> */
		func() bool {
			{
				add(ruleAction67, position)
			}
			return true
		},
		/* 193 Action68 <- <{
		    p.PushComponent(begin, end, Bool)
		}> */
		func() bool {
			{
				add(ruleAction68, position)
			}
			return true
		},
		/* 194 Action69 <- <{
		    p.PushComponent(begin, end, Int)
		}> */
		func() bool {
			{
				add(ruleAction69, position)
			}
			return true
		},
		/* 195 Action70 <- <{
		    p.PushComponent(begin, end, Float)
		}> */
		func() bool {
			{
				add(ruleAction70, position)
			}
			return true
		},
		/* 196 Action71 <- <{
		    p.PushComponent(begin, end, String)
		}> */
		func() bool {
			{
				add(ruleAction71, position)
			}
			return true
		},
		/* 197 Action72 <- <{
		    p.PushComponent(begin, end, Blob)
		}> */
		func() bool {
			{
				add(ruleAction72, position)
			}
			return true
		},
		/* 198 Action73 <- <{
		    p.PushComponent(begin, end, Timestamp)
		}> */
		func() bool {
			{
				add(ruleAction73, position)
			}
			return true
		},
		/* 199 Action74 <- <{
		    p.PushComponent(begin, end, Array)
		}> */
		func() bool {
			{
				add(ruleAction74, position)
			}
			return true
		},
		/* 200 Action75 <- <{
		    p.PushComponent(begin, end, Map)
		}> */
		func() bool {
			{
				add(ruleAction75, position)
			}
			return true
		},
		/* 201 Action76 <- <{
		    p.PushComponent(begin, end, Or)
		}> */
		func() bool {
			{
				add(ruleAction76, position)
			}
			return true
		},
		/* 202 Action77 <- <{
		    p.PushComponent(begin, end, And)
		}> */
		func() bool {
			{
				add(ruleAction77, position)
			}
			return true
		},
		/* 203 Action78 <- <{
		    p.PushComponent(begin, end, Not)
		}> */
		func() bool {
			{
				add(ruleAction78, position)
			}
			return true
		},
		/* 204 Action79 <- <{
		    p.PushComponent(begin, end, Equal)
		}> */
		func() bool {
			{
				add(ruleAction79, position)
			}
			return true
		},
		/* 205 Action80 <- <{
		    p.PushComponent(begin, end, Less)
		}> */
		func() bool {
			{
				add(ruleAction80, position)
			}
			return true
		},
		/* 206 Action81 <- <{
		    p.PushComponent(begin, end, LessOrEqual)
		}> */
		func() bool {
			{
				add(ruleAction81, position)
			}
			return true
		},
		/* 207 Action82 <- <{
		    p.PushComponent(begin, end, Greater)
		}> */
		func() bool {
			{
				add(ruleAction82, position)
			}
			return true
		},
		/* 208 Action83 <- <{
		    p.PushComponent(begin, end, GreaterOrEqual)
		}> */
		func() bool {
			{
				add(ruleAction83, position)
			}
			return true
		},
		/* 209 Action84 <- <{
		    p.PushComponent(begin, end, NotEqual)
		}> */
		func() bool {
			{
				add(ruleAction84, position)
			}
			return true
		},
		/* 210 Action85 <- <{
		    p.PushComponent(begin, end, Concat)
		}> */
		func() bool {
			{
				add(ruleAction85, position)
			}
			return true
		},
		/* 211 Action86 <- <{
		    p.PushComponent(begin, end, Is)
		}> */
		func() bool {
			{
				add(ruleAction86, position)
			}
			return true
		},
		/* 212 Action87 <- <{
		    p.PushComponent(begin, end, IsNot)
		}> */
		func() bool {
			{
				add(ruleAction87, position)
			}
			return true
		},
		/* 213 Action88 <- <{
		    p.PushComponent(begin, end, Plus)
		}> */
		func() bool {
			{
				add(ruleAction88, position)
			}
			return true
		},
		/* 214 Action89 <- <{
		    p.PushComponent(begin, end, Minus)
		}> */
		func() bool {
			{
				add(ruleAction89, position)
			}
			return true
		},
		/* 215 Action90 <- <{
		    p.PushComponent(begin, end, Multiply)
		}> */
		func() bool {
			{
				add(ruleAction90, position)
			}
			return true
		},
		/* 216 Action91 <- <{
		    p.PushComponent(begin, end, Divide)
		}> */
		func() bool {
			{
				add(ruleAction91, position)
			}
			return true
		},
		/* 217 Action92 <- <{
		    p.PushComponent(begin, end, Modulo)
		}> */
		func() bool {
			{
				add(ruleAction92, position)
			}
			return true
		},
		/* 218 Action93 <- <{
		    p.PushComponent(begin, end, UnaryMinus)
		}> */
		func() bool {
			{
				add(ruleAction93, position)
			}
			return true
		},
		/* 219 Action94 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, Identifier(substr))
		}> */
		func() bool {
			{
				add(ruleAction94, position)
			}
			return true
		},
		/* 220 Action95 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, Identifier(substr))
		}> */
		func() bool {
			{
				add(ruleAction95, position)
			}
			return true
		},
	}
	p.rules = _rules
}
