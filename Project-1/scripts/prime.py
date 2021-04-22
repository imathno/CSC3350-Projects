import sys

remaining = int(sys.argv[1])

counted_primes = 0

for current_num in range(2, remaining + 1):
    for i in range(2, current_num):
        if (current_num % i) == 0:
            break
    else:
        counted_primes += 1

print("Primes counted:", counted_primes)
