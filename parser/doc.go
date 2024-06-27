package parser

//parse implement pratt parser

/*
	-.statement

	letStatement
	let <statement> = <expression>

	return
	return <expression>

	IF
	if (<condition>) <consequence> else <alternative>


*/

/*
	-.Expression


	identifier
	5;
	add(x,y)

	int
	<expression>

	prefix operator
	! and -.
	<prefix operator><expression>;


	infix operator
	5 == 5, 6 >= 5 , etc...
	<expression> <infix operator> <expression>

	boolean (literal)
	true, false


	function
	fn <parameters> <block statement>

	function call
	<expression>(<comma separated expressions>)
*/

/*
	Built-in Data type

	STRING
	"<sequence of characters>"

	Array
	[<expression>, <expression>]

	Array-op index
	<expression>[<expression>]
*/
