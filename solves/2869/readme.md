달팽이는 정상에 도착한 순간 미끄러지지 않으므로 달팽이가 정확하게 꼭대기 도착시 이동을 수식으로 나타내면 아래와 같다.

$day \times up - (day - 1) \times down = height$

이를 정리하면 아래와 같다. 

$day = (height - down) \div (up - down)$

마지막 이동시 꼭대기를 넘어간다면 day의 값이 소수점과 함께 나눠지므로 코드로 표현하면 아래와 같다.

```go
day := (height - down) / (up - down)
if (height - down) % (up - down) != 0 {
    day++
}
```