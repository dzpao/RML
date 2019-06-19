proc iff*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    if args[2].tp == TypeEnum.list:
        case args[1].tp
        of TypeEnum.logic:
            if args[1].val.logic:
                var unit = newEvalUnit(cont)
                result = unit.eval(args[2].val.list)
                freeEvalUnit(unit)
                return result
            else:
                return nil
        of TypeEnum.integer:
            if args[1].val.integer != 0:
                var unit = newEvalUnit(cont)
                result = unit.eval(args[2].val.list)
                freeEvalUnit(unit)
                return result
            else:
                return nil
        of TypeEnum.decimal:
            if args[1].val.decimal != 0.0:
                var unit = newEvalUnit(cont)
                result = unit.eval(args[2].val.list)
                freeEvalUnit(unit)
                return result
            else:
                return nil
        of TypeEnum.string:
            if args[1].val.string != "":
                var unit = newEvalUnit(cont)
                result = unit.eval(args[2].val.list)
                freeEvalUnit(unit)
                return result
            else:
                return nil
        of TypeEnum.none:
            return nil
        else:
            var unit = newEvalUnit(cont)
            result = unit.eval(args[2].val.list)
            freeEvalUnit(unit)
            return result
    else:
        result = newToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
        return result


proc either*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    if args[2].tp == TypeEnum.list and args[3].tp == TypeEnum.list:
        var unit = newEvalUnit(cont)
        case args[1].tp
        of TypeEnum.logic:
            if args[1].val.logic:
                result = unit.eval(args[2].val.list)
            else:
                result = unit.eval(args[3].val.list)
        of TypeEnum.integer:
            if args[1].val.integer != 0:
                result = unit.eval(args[2].val.list)
            else:
                result = unit.eval(args[3].val.list)
        of TypeEnum.decimal:
            if args[1].val.decimal != 0.0:
                result = unit.eval(args[2].val.list)
            else:
                result = unit.eval(args[3].val.list)
        of TypeEnum.string:
            if args[1].val.string != "":
                result = unit.eval(args[2].val.list)
            else:
                result = unit.eval(args[3].val.list)
        of TypeEnum.none:
            result = unit.eval(args[3].val.list)
        else:
            result = unit.eval(args[2].val.list)
        freeEvalUnit(unit)
    else:
        result = newToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result


proc loop*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    if args[1].tp == TypeEnum.integer and args[2].tp == TypeEnum.list:
        var unit = newEvalUnit(cont)
        for i in 1..args[1].val.integer:
            result = unit.eval(args[2].val.list)
        freeEvalUnit(unit)
    else:
        result = newToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result

proc repeat*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    if args[1].tp == TypeEnum.word and args[2].tp == TypeEnum.integer and args[3].tp == TypeEnum.list:
        var unit = newEvalUnit(cont)
        var countToken = newToken(TypeEnum.integer)
        countToken.val.integer = 1
        unit.mainCtx[args[1].val.string] = countToken 
        while unit.mainCtx[args[1].val.string].val.integer <= args[2].val.integer:
            result = unit.eval(args[3].val.list)
            # echo(repr(unit.mainCtx[args[1].val.string]))
            unit.mainCtx[args[1].val.string].val.integer += 1
        freeEvalUnit(unit)
    else:
        result = newToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result


proc ffor*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    if args[1].tp == TypeEnum.word and args[5].tp == TypeEnum.list and (args[2].tp == TypeEnum.integer or args[2].tp == TypeEnum.decimal) and (args[3].tp == TypeEnum.integer or args[3].tp == TypeEnum.decimal) and (args[4].tp == TypeEnum.integer or args[4].tp == TypeEnum.decimal) :
        if(args[2].tp == TypeEnum.integer and args[3].tp == TypeEnum.integer and args[4].tp == TypeEnum.integer):
            var unit = newEvalUnit(cont)
            var count = newToken(TypeEnum.integer)
            count.val.integer = args[2].val.integer
            unit.mainCtx[args[1].val.string] = count
            while unit.mainCtx[args[1].val.string].val.integer <= args[3].val.integer:
                result = unit.eval(args[5].val.list)
                unit.mainCtx[args[1].val.string].val.integer += args[4].val.integer
            freeEvalUnit(unit)
        else:
            var unit = newEvalUnit(cont)
            var count = newToken(TypeEnum.decimal)
            if args[2].tp == TypeEnum.integer:
                count.val.decimal = float64(args[2].val.integer)
            else:
                count.val.decimal = args[2].val.decimal
            unit.mainCtx[args[1].val.string] = count

            if args[3].tp == TypeEnum.integer:
                while unit.mainCtx[args[1].val.string].val.decimal <= args[3].val.integer.float64:
                    result = unit.eval(args[5].val.list)
                    if args[4].tp == TypeEnum.integer:
                        unit.mainCtx[args[1].val.string].val.decimal += args[4].val.integer.float64
                    else:
                        unit.mainCtx[args[1].val.string].val.decimal += args[4].val.decimal
            else:
                while unit.mainCtx[args[1].val.string].val.decimal <= args[3].val.decimal:
                    result = unit.eval(args[5].val.list)
                    if args[4].tp == TypeEnum.integer:
                        unit.mainCtx[args[1].val.string].val.decimal += args[4].val.integer.float64
                    else:
                        unit.mainCtx[args[1].val.string].val.decimal += args[4].val.decimal
            freeEvalUnit(unit)
    else:
        result = newToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result


proc wwhile*(args: ptr List[ptr Token], cont: ptr BindMap[ptr Token] = nil):ptr Token=
    if args[1].tp == TypeEnum.list and args[2].tp == TypeEnum.list:
        var condUnit = newEvalUnitWithMap(cont)
        var bodyUnit = newEvalUnitWithMap(cont)
        var b = condUnit.eval(args[1].val.list)
        while b.tp == TypeEnum.logic and b.val.logic:
            result = bodyUnit.eval(args[2].val.list)
            b = condUnit.eval(args[1].val.list)

        if not (b.tp == TypeEnum.logic):
            result = newToken(TypeEnum.err)
            result.val.string = "Bad Logic Expression "
        freeEvalUnitWithoutMap(condUnit)
        freeEvalUnitWithoutMap(bodyUnit)
    else:
        result = newToken(TypeEnum.err)
        result.val.string = "Type Mismatch"
    return result