# Overview

When I first began learning Python, I was also learning to code a compiler. I felt like I learned my best lessons as a programmer during this time and if I could recreate said project in any language, I feel I could get a decent grasp on how said language works and learn new ways to do certain things. To start of my learning of Golang, I started with a basic lexer that is able to tokenize any keyword or character I may want in the future. Currently, it does not have the best error messages, but it is a WIP while I learn Golang.

# Lexer

In terms of a compiler, a lexer is the initial start to syntax checking and will tokenize each part of the input for the parser and later stages of the compiler.

# Current Keywords
* keyword
  * "if", "then", "while", "print"
* vartype
  * "int", "str", "bool"
* boolean
  * true", "false"
* assign
  * "="
* operators
  * "+", "-"
* compare
  * "==", "!="
* quote
  * "\""
* space
  * " ", "\t"
* newline"
  * "\n", "\r\n"

# Future Enhancements
* [ ] Take in a file name as an argument
* [ ] Improve error messages
* [ ] Create an actual main package to take in input and pass to lexer package
