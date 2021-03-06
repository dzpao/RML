package nativelib

import . "../core"
import "errors"
import "sync"
// import "fmt"

func Iif(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	var result Token
	if args[1].ToBool(){
		if args[2].Tp == BLOCK {
			return es.Eval(args[2].Tks(), ctx)
		}else if args[2].Tp == STRING {
			return es.EvalStr(args[2].Str(), ctx)
		}
	}else{
		return nil, nil
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil

}

func Either(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	var result Token
	if args[1].ToBool(){
		if args[2].Tp == BLOCK {
			return es.Eval(args[2].Tks(), ctx)
		}else if args[2].Tp == STRING {
			return es.EvalStr(args[2].Str(), ctx)
		}
	}else{
		if args[3].Tp == BLOCK {
			return es.Eval(args[3].Tks(), ctx)
		}else if args[3].Tp == STRING {
			return es.EvalStr(args[3].Str(), ctx)
		}
	}

	result.Tp = ERR
	result.Val = "Type Mismatch"
	return &result, nil
}

func Loop(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if(args[1].Tp == INTEGER && args[2].Tp == BLOCK){
		var rs *Token
		var err error
		for i := 0; i < args[1].Int(); i++ {
			rs, err = es.Eval(args[2].Tks(), ctx) 
			if err != nil {
				if err.Error() == "continue" {
					continue
				}
				if err.Error() == "break" {
					break
				}
				return rs, err
			}
		}

		return rs, nil

	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}


func Repeat(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	// args[2].Echo()

	if(args[1].Tp == WORD && args[2].Tp == INTEGER && args[3].Tp == BLOCK){
		 var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
		 var countToken = Token{INTEGER, 1}
		 
		 c.PutNow(args[1].Str(), &countToken)
		 var rs *Token
		 var err error
		 for countToken.Int() <= args[2].Int() {
			rs, err = es.Eval(args[3].Tks(), &c)
			countToken.Val = countToken.Int() + 1
			if err != nil {
				if err.Error() == "continue" {
					continue
				}
				if err.Error() == "break" {
					break
				}
				return rs, err
			}
			if rs != nil && rs.Tp == ERR {
				return rs, err
			}
		 }
		 return nil, nil

	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}


func Ffor(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if(args[1].Tp == WORD && args[5].Tp == BLOCK && (args[2].Tp == INTEGER || args[2].Tp == DECIMAL) && (args[3].Tp == INTEGER || args[3].Tp == DECIMAL) && (args[4].Tp == INTEGER || args[4].Tp == DECIMAL)){
		var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
		var countToken = args[2].Dup()
		c.PutNow(args[1].Str(), countToken)
		var rs *Token
		var err error

		if(args[2].Tp == INTEGER && args[4].Tp == INTEGER){
			if args[3].Tp == INTEGER {
				for countToken.Int() <= args[3].Int() {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Int() + args[4].Int()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}else{
				for countToken.Int() <= int(args[3].Float()) {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Int() + args[4].Int()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}

		}else if(args[2].Tp == INTEGER && args[4].Tp == DECIMAL) {
			countToken.Tp = DECIMAL
			countToken.Val = float64(countToken.Int())
			if args[3].Tp == INTEGER {
				for countToken.Float() <= float64(args[3].Int()) {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + args[4].Float()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}else{
				for countToken.Float() <= args[3].Float() {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + args[4].Float()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}
		}else if(args[2].Tp == DECIMAL && args[4].Tp == INTEGER) {
			if args[3].Tp == INTEGER {
				for countToken.Float() <= float64(args[3].Int()) {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + float64(args[4].Int())
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}else{
				for countToken.Float() <= args[3].Float() {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + float64(args[4].Int())
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}
		}else if(args[2].Tp == DECIMAL && args[4].Tp == DECIMAL) {
			if args[3].Tp == INTEGER {
				for countToken.Float() <= float64(args[3].Int()) {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + args[4].Float()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}else{
				for countToken.Float() <= args[3].Float() {
					rs, err = es.Eval(args[5].Tks(), &c)
					countToken.Val = countToken.Float() + args[4].Float()
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return rs, err
					}
					if rs != nil && rs.Tp == ERR {
						return rs, err
					}
				}
			}
		}

		return nil, nil
	}

	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}

func Wwhile(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if(args[1].Tp == BLOCK && args[2].Tp == BLOCK){
		var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
		b, e1 := es.Eval(args[1].Tks(), &c)
		if e1 != nil {
			return nil, e1
		}
		var rs *Token
		var err error
		for b.Val.(bool) {
			rs, err = es.Eval(args[2].Tks(), &c)
			if err != nil {
				if err.Error() == "continue" {
					continue
				}
				if err.Error() == "break" {
					break
				}
				return rs, err
			}
			if rs != nil && rs.Tp == ERR {
				return rs, err
			}
			b, err = es.Eval(args[1].Tks(), &c)
			if err != nil {
				return rs, err
			}
			if b.Tp == ERR {
				return rs, err
			}
		}
		return nil, nil
	}
	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}


func Bbreak(es *EvalStack, ctx *BindMap) (*Token, error){
	return nil, errors.New("break")
}

func Ccontinue(es *EvalStack, ctx *BindMap) (*Token, error){
	return nil, errors.New("continue")
}

func Rreturn(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	return args[1], errors.New("return")
}

func Fforeach(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]
	if args[3].Tp != BLOCK && args[3].Tp != STRING {
		var result = Token{ERR, "Type Mismatch"}
		return &result, nil
	}

	if args[1].Tp == WORD {
		if args[2].Tp == BLOCK || args[2].Tp == PAREN || args[2].Tp == PATH {
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for i:=0; i<args[2].List().Len(); i++ {
				c.PutNow(args[1].Str(), args[2].Tks()[i])
				if args[3].Tp == BLOCK {
					temp, err := es.Eval(args[3].Tks(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}

				}else if args[3].Tp == STRING {
					temp, err := es.EvalStr(args[3].Str(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}
				}

			}
			return nil, nil

		}else if args[2].Tp == OBJECT {
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for k, v := range(args[2].Ctx().Table){
				var blk = NewTks(4)
				blk.AddArr([]*Token{&Token{WORD, k}, v})
				c.PutNow(args[1].Str(), &Token{BLOCK, blk})
				if args[3].Tp == BLOCK {
					temp, err := es.Eval(args[3].Tks(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}

				}else if args[3].Tp == STRING {
					temp, err := es.EvalStr(args[3].Str(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}
				}
			}
			return nil, nil
		}else if args[2].Tp == MAP {
			if len(args[2].Table()) == 0 {
				return nil, nil
			}
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for _, v := range(args[2].Table()){
				var blk = NewTks(4)
				blk.AddArr([]*Token{v.Key.CloneDeep(), v.Val})
				c.PutNow(args[1].Str(), &Token{BLOCK, blk})
				if args[3].Tp == BLOCK {
					temp, err := es.Eval(args[3].Tks(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}

				}else if args[3].Tp == STRING {
					temp, err := es.EvalStr(args[3].Str(), &c)
					if err != nil {
						if err.Error() == "continue" {
							continue
						}
						if err.Error() == "break" {
							break
						}
						return temp, err
					}
					if temp != nil && temp.Tp == ERR {
						return temp, err
					}
				}
			}
			return nil, nil
		}

	}else if args[1].Tp == BLOCK {
		for _, item := range(args[1].Tks()) {
			if item.Tp != WORD {
				var result = Token{ERR, "Type Mismatch"}
				return &result, nil
			}
		}

		if args[2].Tp == BLOCK {
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for i:=0; i<args[2].List().Len(); i+=args[1].List().Len() {
				for j:=0; j<args[1].List().Len(); j++ {
					if i+j < args[2].List().Len() {
						c.PutNow(args[1].Tks()[j].Str(), args[2].Tks()[i+j])
					}else{
						c.PutNow(args[1].Tks()[j].Str(), &Token{NONE, "none"})
					}
				}
				temp, err := es.Eval(args[3].Tks(), &c)
				if err != nil {
					if err.Error() == "continue" {
						continue
					}
					if err.Error() == "break" {
						break
					}
					return temp, err
				}
				if temp != nil && temp.Tp == ERR {
					return temp, err
				}
			}
			
			return nil, nil
		}else if args[2].Tp == OBJECT {
			if args[1].List().Len() < 2 || args[1].Tks()[0].Tp != WORD || args[1].Tks()[1].Tp != WORD {
				var result = Token{ERR, "Type Mismatch"}
				return &result, nil
			}
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for k, v := range(args[2].Ctx().Table){
				c.PutNow(args[1].Tks()[0].Str(), &Token{WORD, k})
				c.PutNow(args[1].Tks()[1].Str(), v)
				temp, err := es.Eval(args[3].Tks(), &c)
				if err != nil {
					if err.Error() == "continue" {
						continue
					}
					if err.Error() == "break" {
						break
					}
					return temp, err
				}
				if temp != nil && temp.Tp == ERR {
					return temp, err
				}
			}
			return nil, nil
		}else if args[2].Tp == MAP {
			if args[1].List().Len() < 2 || args[1].Tks()[0].Tp != WORD || args[1].Tks()[1].Tp != WORD {
				var result = Token{ERR, "Type Mismatch"}
				return &result, nil
			}
			if len(args[2].Table()) == 0{
				return nil, nil
			}
			var c = BindMap{make(map[string]*Token, 8), ctx, TMP_CTX, sync.RWMutex{}}
			for _, v := range(args[2].Table()){
				c.PutNow(args[1].Tks()[0].Str(), v.Key.CloneDeep())
				c.PutNow(args[1].Tks()[1].Str(), v.Val)
				temp, err := es.Eval(args[3].Tks(), &c)
				if err != nil {
					if err.Error() == "continue" {
						continue
					}
					if err.Error() == "break" {
						break
					}
					return temp, err
				}
				if temp != nil && temp.Tp == ERR {
					return temp, err
				}
			}

			return nil, nil
		}

	}



	var result = Token{ERR, "Type Mismatch"}
	return &result, nil
}

func Ttry(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if args[1].Tp == BLOCK && args[2].Tp == BLOCK {
		temp, err := es.Eval(args[1].Tks(), ctx)
		if (temp != nil && temp.Tp == ERR) || err != nil {
			if err != nil {
				temp.Val = err.Error()
			}

			var c = BindMap{make(map[string]*Token, 4), ctx, TMP_CTX, sync.RWMutex{}}
			c.PutNow("e", temp)

			return es.Eval(args[2].Tks(), &c)
		}

		return temp, err
	}


	return &Token{ERR, "Type Mismatch"}, nil
}

func Cause(es *EvalStack, ctx *BindMap) (*Token, error){
	var args = es.Line[es.LastStartPos() : es.LastEndPos() + 1]

	if args[1].Tp == STRING {
		return &Token{ERR, args[1].Str()}, nil
	}


	return &Token{ERR, "Type Mismatch"}, nil
}
