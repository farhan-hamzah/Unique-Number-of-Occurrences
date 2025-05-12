package main
import "fmt"
const NMAX int = 100
func main(){
    var A[NMAX]int
    var n, i, j, x, c int
    var b, b1 bool
    b = false
    b1 = true
    fmt.Scan(&n)
    for i =0; i < n; i++{
        fmt.Scan(&A[i])
    }
    for i = 0; i < n && b1 == true; i++{
        x = A[i]
        for j = i+1; j < n; j++{
            c = 1
            if A[i] == A[j]{
                c++
            }
        }
        if x != c{
            b = true
        }else{
            b = false
        }
        if b == false{
            b1 = false
        }
    }
    fmt.Print(b)
}