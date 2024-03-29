# Numeric

## Float

```go
f1 := 0.1
f2 := 0.2
got := f1 * f2
fmt.Println(got) // 0.020000000000000004
```

```go
f1 := 0.7
f2 := 0.1
got := f1 / f2
fmt.Println(got) // 6.999999999999999
```

## Big rational

```go
import "math/big"
```

```go
r1, _ := new(big.Rat).SetString("0.1")
r2, _ := new(big.Rat).SetString("0.2")
r3 := new(big.Rat).Mul(r1, r2)
fmt.Println(r3.FloatString(2)) // 0.02
```

```go
r1, _ := new(big.Rat).SetString("0.7")
r2, _ := new(big.Rat).SetString("0.1")
r3 := new(big.Rat).Quo(r1, r2)
fmt.Println(r3.FloatString(2)) // 7.00
```

```go
r, _ := new(big.Rat).SetString("13.514")
f, _ := r.Float64()
fmt.Printf("%.2f\n", f) // 13.51
```

```go
r, _ := new(big.Rat).SetString("13.515")
f, _ := r.Float64()
fmt.Printf("%.2f\n", f) // 13.52
```

## Decimal

https://github.com/shopspring/decimal

```go
import "github.com/shopspring/decimal"
```

```go
r1, err := decimal.NewFromString("0.1")
if err != nil {
    panic(err)
}
r2, err := decimal.NewFromString("0.2")
if err != nil {
    panic(err)
}
r3 := r1.Mul(r2)
fmt.Println(r3.StringFixed(2)) // 0.02
```

```go
r1, err := decimal.NewFromString("0.7")
if err != nil {
    panic(err)
}
r2, err := decimal.NewFromString("0.1")
if err != nil {
    panic(err)
}
r3 := r1.Div(r2)
fmt.Println(r3.StringFixed(2)) // 7.00
```

```go
d, err := decimal.NewFromString("13.514")
if err != nil {
    panic(err)
}
fmt.Println(d.StringFixed(2)) // 13.51
```

```go
d, err := decimal.NewFromString("13.515")
if err != nil {
    panic(err)
}
fmt.Println(d.StringFixed(2)) // 13.52
```

## Rounding strategies

| Number | Round Half Up | Bankers' Rounding |
|--------|---------------|-------------------|
| 0.5    | 1             | 0                 |
| 1.5    | 2             | 2                 |
| 2.5    | 3             | 2                 |
| 3.5    | 4             | 4                 |
| 4.5    | 5             | 4                 |

| Number | Round Half Up | Bankers' Rounding |
|--------|---------------|-------------------|
| 0.005  | 0.01          | 0.00              |
| 0.015  | 0.02          | 0.02              |
| 0.025  | 0.03          | 0.02              |
| 0.035  | 0.04          | 0.04              |
| 0.045  | 0.05          | 0.04              |
