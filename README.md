# Introduction

Blang is a Programming language designed by me as a way of understanding programming languages better.

This project will be the compile necessary to transform the language in something the computer can use.

This project is inspired in the following [tutorial](https://www.freecodecamp.org/news/the-programming-language-pipeline-91d3f449c919/).

# Requirements

- [ ] Be able use variables;
- [ ] Be able do arithmetic;
- [ ] Be able run functions;
- [ ] Be able create structs;

# Grammar

```
program      -> statement*
statement    -> expr ';'
expr         -> term ((ADD | SUB) term)*
term         -> factor ((MUL | DIV) factor)*
factor       -> INT | IDENT | '(' expr ')'
statement    -> (IDENT ASSIGN expr | expr) ';'
```