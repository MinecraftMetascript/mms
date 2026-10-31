// Number helpers with parameters and return values. text.mms calls these.

/** The n-th Fibonacci number, computed recursively. */
Function fib(n: int): int {
  if (n < 2) {
    return n
  }
  return fib(n - 1) + fib(n - 2)
}

/** Greatest common divisor, with a loop. */
Function gcd(a: int, b: int): int {
  while (b != 0) {
    let t = b
    b = a % b
    a = t
  }
  return a
}

/** Whether n is prime: trial division that stops at the first divisor. */
Function is_prime(n: int): bool {
  if (n < 2) {
    return false
  }
  let d = 2
  let prime = true
  while (d * d <= n) {
    if (n % d == 0) {
      prime = false
      break
    }
    d += 1
  }
  return prime
}

/** a to the power b, by squaring. */
Function pow(a: int, b: int): int {
  if (b == 0) {
    return 1
  }
  let half = pow(a, b / 2)
  if (b % 2 == 0) {
    return half * half
  }
  return half * half * a
}
