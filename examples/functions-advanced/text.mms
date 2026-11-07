// String values and the functions in math.mms, used from commands.

/** What the reports so far have found, kept in storage adv:mms reports. */
Type Reports {
  /** How many reports have run. */
  count: int = 0
  /** The last line a report built, for other functions to read. */
  last: string = ""
}

/** Read it in the game with /data get storage adv:mms reports. */
Storage reports: Reports

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
  reports.count += 1
  reports.last = line + " (" + count + " in total)"
  tell(@a, "${reports.last}")
}

/** Says hello in a way that depends on the name. */
Function greet(name: string) {
  let greeting = "Hello, " + name
  if (name == "Notch") {
    greeting += ". Thanks for the game!"
  } else {
    greeting += "!"
  }
  tell(@a, "${greeting}")
}

/** /function adv:demo shows everything above. */
Function demo {
  greet("Steve")
  /say fib(10) is ${fib(10)}
  /say gcd(84, 36) is ${gcd(84, 36)}
  /say 2^10 is ${pow(2, 10)}
  primes_up_to(30)
}
