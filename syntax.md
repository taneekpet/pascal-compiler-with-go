Symbol | Expand to
--- | --- 
**PROGRAM** | PROGRAM **VARIABLE** ; **BLOCK** .
**BLOCK** | **DECLARATIONS** **COMPOUND_STATEMENT**
**DECLARATIONS** | VAR (**VARIABLE_DECLARATION** ;)+
| | _empty_
**VARIABLE_DECLARATION** | **$ID** (, **$ID**)* : **TYPE_SPEC**
**TYPE_SPEC** | INTEGER
| | REAL
**COMPOUND_STATEMENT** | BEGIN **STATEMENT_LIST** END
**STATEMENT_LIST** | **STATEMENT**
| | **STATEMENT** : **STATEMENT_LIST**
**STATEMENT** | **COMPOUND_STATEMENT**
| | **ASSIGN**
| | _empty_
**ASSIGN** | **$ID** := **EXPR**
**EXPR** | **TERM** (('+' \| '-') **TERM**)*
**TERM** | **FACTOR** (('\*' \| '/' ) **FACTOR**)*
**FACTOR** | + **FACTOR**
| | - **FACTOR**
| | **$NUMBER_CONST**
| | '(' **EXPR** ')'
| | **$ID**