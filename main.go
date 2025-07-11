/*
The main.go file serves as the orchestrator and entry point for
your entire compiler application. In this file, you'll handle
command-line arguments to get the input MVL source file, read its
content, and then sequentially initiate each phase of your compiler:
creating a new lexer.Lexer instance, passing it to a new parser.
Parser, calling the parser to build the ast.Program, and finally
invoking the codegen.Generate function to produce the target Go source
code. You'll also include logic to handle and report any errors encountered
during lexing, parsing, or code generation, and to write the final
generated Go code to an output file for execution.
*/