/*
This file houses the lexer (or scanner), which is the first stage of your compiler.
Your primary task here is to implement the NextToken() method.
This method will read characters from your MVL source code one by one,
skip whitespace, identify sequences of characters that form valid tokens
(like numbers, variable names, operators, and keywords), and return them as
token.Token structs. You'll also need to implement helper functions like readChar,
 peekChar, skipWhitespace, readIdentifier, readNumber, isLetter, and isDigit to
 support NextToken()'s logic.
*/