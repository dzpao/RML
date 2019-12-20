package nativelib

import . "../core"

func InitNative(ctx *BindMap){
	var quitToken = Token{
		NATIVE,
		Native{
			"quite",
			1,
			Quit,
			nil,
		},
	}
	ctx.PutNow("quit", &quitToken)
	ctx.PutNow("q", &quitToken)

	var typeOfToken = Token{
		NATIVE,
		Native{
			"type?",
			2,
			TypeOf,
			nil,
		},
	}
	ctx.PutNow("type?", &typeOfToken)

	var doToken = Token{
		NATIVE,
		Native{
			"do",
			2,
			Do,
			nil,
		},
	}
	ctx.PutNow("do", &doToken)

	var copyToken = Token{
		NATIVE,
		Native{
			"copy",
			2,
			Copy,
			nil,
		},
	}
	ctx.PutNow("copy", &copyToken)

	var printToken = Token{
		NATIVE,
		Native{
			"print",
			2,
			Pprint,
			nil,
		},
	}
	ctx.PutNow("print", &printToken)

	var addToken = Token{
		NATIVE,
		Native{
			"add",
			3,
			Add,
			nil,
		},
	}
	ctx.PutNow("add", &addToken)

	var subToken = Token{
		NATIVE,
		Native{
			"sub",
			3,
			Sub,
			nil,
		},
	}
	ctx.PutNow("sub", &subToken)

	var mulToken = Token{
		NATIVE,
		Native{
			"mul",
			3,
			Mul,
			nil,
		},
	}
	ctx.PutNow("mul", &mulToken)

	var divToken = Token{
		NATIVE,
		Native{
			"div",
			3,
			Div,
			nil,
		},
	}
	ctx.PutNow("div", &divToken)

	var modToken = Token{
		NATIVE,
		Native{
			"mod",
			3,
			Mod,
			nil,
		},
	}
	ctx.PutNow("mod", &modToken)

	var ifToken = Token{
		NATIVE,
		Native{
			"if",
			3,
			Iif,
			nil,
		},
	}
	ctx.PutNow("if", &ifToken)

	var eitherToken = Token{
		NATIVE,
		Native{
			"either",
			4,
			Either,
			nil,
		},
	}
	ctx.PutNow("either", &eitherToken)

	var loopToken = Token{
		NATIVE,
		Native{
			"loop",
			3,
			Loop,
			nil,
		},
	}
	ctx.PutNow("loop", &loopToken)

	var repeatToken = Token{
		NATIVE,
		Native{
			"repeat",
			4,
			Repeat,
			[]int{0, 1, 1},
		},
	}
	ctx.PutNow("repeat", &repeatToken)

	var forToken = Token{
		NATIVE,
		Native{
			"for",
			6,
			Ffor,
			[]int{0, 1, 1, 1, 1},
		},
	}
	ctx.PutNow("for", &forToken)

	var whileToken = Token{
		NATIVE,
		Native{
			"while",
			3,
			Wwhile,
			nil,
		},
	}
	ctx.PutNow("while", &whileToken)

	var breakToken = Token{
		NATIVE,
		Native{
			"break",
			1,
			Bbreak,
			nil,
		},
	}
	ctx.PutNow("break", &breakToken)

	var continueToken = Token{
		NATIVE,
		Native{
			"continue",
			1,
			Ccontinue,
			nil,
		},
	}
	ctx.PutNow("continue", &continueToken)

	var returnToken = Token{
		NATIVE,
		Native{
			"return",
			2,
			Rreturn,
			nil,
		},
	}
	ctx.PutNow("return", &returnToken)

	var defFuncToken = Token{
		NATIVE,
		Native{
			"func",
			3,
			DefFunc,
			nil,
		},
	}
	ctx.PutNow("func", &defFuncToken)

	var eqToken = Token{
		NATIVE,
		Native{
			"eq",
			3,
			Eq,
			nil,
		},
	}
	ctx.PutNow("eq", &eqToken)

	var gtToken = Token{
		NATIVE,
		Native{
			"gt",
			3,
			Gt,
			nil,
		},
	}
	ctx.PutNow("gt", &gtToken)

	var ltToken = Token{
		NATIVE,
		Native{
			"lt",
			3,
			Lt,
			nil,
		},
	}
	ctx.PutNow("lt", &ltToken)

	var geToken = Token{
		NATIVE,
		Native{
			"ge",
			3,
			Ge,
			nil,
		},
	}
	ctx.PutNow("ge", &geToken)

	var leToken = Token{
		NATIVE,
		Native{
			"le",
			3,
			Le,
			nil,
		},
	}
	ctx.PutNow("le", &leToken)

	var costToken = Token{
		NATIVE,
		Native{
			"cost",
			2,
			Cost,
			nil,
		},
	}
	ctx.PutNow("cost", &costToken)

}