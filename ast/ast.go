/*
This file defines the Abstract Syntax Tree (AST), which is the hierarchical,
structured representation of your MVL code after it's parsed. You'll define
interfaces like Node, Statement, and Expression, and then create concrete
struct types for each element of your language's grammar (e.g., Program, VarStatement,
PrintStatement, Identifier, IntegerLiteral, InfixExpression). For each of these
structs, you'll implement the TokenLiteral() and, most importantly for debugging,
the String() method, which provides a human-readable representation of the AST.
*/