rl() = [parse(Int64, x) for x in split(readline())]
 
function solve()
    N, M, K = rl()
    A = rl()
    B = rl()
 
    na = 0
    for (i, a) in enumerate(A)
        if K < a
            break
        end
        K -= a
        na += 1
    end
 
    r = na
 
    for (nb, b) in enumerate(B)
        K -= b
        while K < 0 && na > 0
            K += A[na]
            na -= 1
        end
        if K < 0
            break
        end
        r = max(r, na + nb)
    end
    r
end
 
println(solve())