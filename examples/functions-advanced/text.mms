// String values and the functions in math.mms, used from commands.

/** The last message built by report, for other functions to read. */
Global last_report: string = ""

/** Lists the primes up to max, joined into one line. */
Function primes_up_to(max: int) {
  let line = "Primes:"
  let count = 0
  for n in 2..max {
    if (!is_prime(n)) {
      continue
    }
    line += " " + n
    count += 1
  }
  last_report = line + " (" + count + " in total)"
  /tellraw @a "${last_report}"
}

/** Says hello in a way that depends on the name. */
Function greet(name: string) {
  let greeting = "Hello, " + name
  if (name == "Notch") {
    greeting += ". Thanks for the game!"
  } else {
    greeting += "!"
  }
  /tellraw @a "${greeting}"
}

/** /function adv:demo shows everything above. */
Function demo {
  greet("Steve")
  /say fib(10) is ${fib(10)}
  /say gcd(84, 36) is ${gcd(84, 36)}
  /say 2^10 is ${pow(2, 10)}
  primes_up_to(30)
}
