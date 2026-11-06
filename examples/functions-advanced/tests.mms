// Tests of math.mms and text.mms. Run them with `mms test`, in the web
// editor's Tests tab, or in VS Code's Testing view. They are never
// written to the pack.

Test fibonacci {
  expect fib(0) == 0
  expect fib(1) == 1
  expect fib(10) == 55
}

Test greatest_common_divisor {
  expect gcd(84, 36) == 12
  expect gcd(17, 5) == 1
}

Test primes {
  expect is_prime(2)
  expect is_prime(97)
  expect !is_prime(1)
  expect !is_prime(91)
}

Test powers {
  expect pow(2, 10) == 1024
  expect pow(7, 0) == 1
}

Test primes_report {
  setup {
    reports.count = 4
  }
  run primes_up_to(10)
  expect reports.count == 5
  expect reports.last == "Primes: 2 3 5 7 (4 in total)"
}
