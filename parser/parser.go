/*
This is where the parser resides, taking the stream of tokens from the lexer
and building the AST defined in ast.go. Your main effort here will be in
implementing the ParseProgram() method, which orchestrates the parsing process.
You'll also write recursive descent parsing functions for each grammar rule
(e.g., parseStatement, parseVarStatement, parseExpression, parseInfixExpression),
and implement helper methods for managing token advancement (nextToken), handling
operator precedence (peekPrecedence, currentPrecedence), and reporting syntax
errors (Errors, peekError, noPrefixParseFnError, expectPeek).
*/