
// golang code
%{

package main

import (
	"fmt"
)

%}

// 用来定义一个类型并映射golang的一个数据类型（可以是一个自定义类型）
%union{
	val string
}

// 定义非终结符
%type <val> input

// 定义终结符
%token <val> CHARACTER

// %% 用于分隔，文法定义和文法规则
%%

in : /* empty */
    | in input '\n'
        { fmt.Printf("Read character: %s\n", $2) }
    ;

input : CHARACTER
    | input CHARACTER
        { $$ = $1 + $2 }
    ;

%%
