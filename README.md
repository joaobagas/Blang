# Introduction

Blang is a Programming language designed by me as a way of understanding programming languages better.
This project will be the compile necessary to transform the language in something the computer can use.

# Requirements

- [ ] Be able to variables;
- [ ] Be able arithmetic;
- [ ] Be able functions;g
- [ ] Be able structs;

# Grammar

program      -> statement*
statement    -> expr ';'
expr         -> term ((ADD | SUB) term)*
term         -> factor ((MUL | DIV) factor)*
factor       -> INT | IDENT | '(' expr ')'
statement    -> (IDENT ASSIGN expr | expr) ';'
