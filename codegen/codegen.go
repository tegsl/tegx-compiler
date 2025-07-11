/*
This file contains the code generator, which is the backend of your "jump start"
compiler. Its role is to traverse the AST (built by the parser) and translate
each AST node into equivalent Go source code. You'll implement the main Generate()
function, which sets up the Go program boilerplate, and then create specific
generate functions for each type of AST node (e.g., generateVarStatement,
generatePrintStatement, generateIdentifier, generateIntegerLiteral, generateInfixExpression).
These functions will write the corresponding Go syntax to an output buffer, effectively
transpiling your MVL into runnable Go code.
*/