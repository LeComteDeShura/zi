package lab1

import "fmt"
import "math"
import "time"
import "math/rand"


func rand64() uint64 {
    return uint64(rand.Uint32() % 1000);
}

func isPrime(p uint64) bool {
    if p <= 1 {
        return false;
    }

    var b uint64 = uint64(math.Pow(float64(p), 0.5));

    var i uint64;
    for i = 2; i <= b; i++ {
        if (p % i == 0) {
            return false;
        }
    }
    //
    return true;
}

func modular_pow(a, e, m uint64) uint64 {
    var r uint64= 1;
    for e != 0 {
         if e % 2 == 1 {
             r = r * a % m;
         }
         e = e / 2;
         a = a * a % m;
     }
     return r;
}

func ea(a, b int64) [3]int64{
    if a < b {
        c := a;
        a = b;
        b = c;
    }

    var t [3]int64;
    u := [3]int64{a, 1, 0};
    v := [3]int64{b, 0, 1};
    var q int64;

    for v[0] > 0 {
        q = u[0] / v[0];
        // t := [3]int64{u[0] % v[0], u[1] - q * v[1], u[2] - q * v[2]}
        t[0] = u[0] % v[0];
        t[1] = u[1] - q * v[1];
        t[2] = u[2] - q * v[2];

        u = v;
        v = t;
    }
   return u;
}

func genPG() (uint64, uint64) {
    var q uint64=0;
    var p uint64=0;
    for !isPrime(q) || !isPrime(p) {
        q = uint64(rand.Intn(1000000000))
        p = 2*q+1
        // fmt.Println(q, p)
    }

    var g uint64 = uint64(rand.Intn(int(p-1)));
    for modular_pow(g, q, p) == 1 {
        g  = uint64(rand.Intn(int(p-1)));
    }

    return p, g
}

func bsgs(a, p, y uint64) (uint64, bool) {
    m := uint64(math.Sqrt(float64(p))) + 1
    k := uint64(math.Sqrt(float64(p))) + 1

    var g = make(map[uint64]uint64)
    var i uint64;
    for i = 1; i <= k; i++ {
        g[modular_pow(a, i*m, p)] = (i)
    }

    fmt.Println()
    var j uint64;
    for j = 0; j < m; j++ {
        b := ((y % p)*(modular_pow(a, j, p))) % p
        if v, found := g[b]; found {
            return v*m-j, true
        }
    }

    return 0, false
}


// func main() {
//     print(isPrime(5), "\n\n")
//
//
//     rand.Seed(time.Now().UnixNano())
//     /////////////////////////////////////////////////////////
//     fmt.Println("Алгоритм быстрого возведения числа в степень по модулю")
//     a1 := rand64()
//     x := rand64()
//     p := rand64()
//     y1 := modular_pow(a1, x, p)
//
//     fmt.Printf("%d^%d mod %d = %d\n\n\n", a1, x, p, y1)
//     /////////////////////////////////////////////////////////
//
//
//     /////////////////////////////////////////////////////////
//     fmt.Println("Обобщённый алгоритм Евклида")
//     a2 := int64(rand64())
//     b2 := int64(rand64())
//     fmt.Printf("a = %d b = %d", a2, b2)
//     u := ea(a2, b2)
//     fmt.Printf("\n%d %d %d\n", u[0], u[1], u[2]);
//     if a2 < b2 {
//         c2 := a2;
//         a2 = b2;
//         b2 = c2;
//     }
//     fmt.Print(u[1]*a2+u[2]*b2, u[0], "\n\n\n");
//     /////////////////////////////////////////////////////////
//
//
//     /////////////////////////////////////////////////////////
//     fmt.Println("Диффи-Хеллман")
//     p, g := genPG()
//     fmt.Printf("общие данные: p = %d, g = %d\n", p, g)
//
//     var xa uint64 = rand64();
//     var xb uint64 = rand64();
//     fmt.Printf("закрытые ключи: Xa = %d, Xb = %d\n", xa, xb)
//
//     ya := modular_pow(g, xa, p)
//     yb := modular_pow(g, xb, p)
//     fmt.Printf("открытые ключи: Ya = %d, Yb = %d\n", ya, yb)
//
//     za := modular_pow(yb, xa, p)
//     zb := modular_pow(ya, xb, p)
//     if za == zb {
//         fmt.Printf("общий ключ: Z = %d\n", za)
//     } else {
//         fmt.Printf("error");
//     }
//     /////////////////////////////////////////////////////////
//
//
//     /////////////////////////////////////////////////////////
//     fmt.Println("\n\nШаг младенца, шаг великана")
//     // p = rand64()
//     a := rand64() % p
//     y := rand64() % p
//
//     fmt.Printf("%d^x mod %d = %d", a, p, y)
//
//     x, found := bsgs(a, p, y)
//
//     if found {
//         fmt.Println("x =", x)
//         fmt.Printf("%d^%d mod %d = %d\n", a, x, p, modular_pow(a, x, p))
//     } else {
//         fmt.Println("not found")
//     }
//     /////////////////////////////////////////////////////////
// }
