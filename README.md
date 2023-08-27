## Task

Design and implement “Word of Wisdom” tcp server.  
• TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.  
• The choice of the POW algorithm should be explained.  
• After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.  
• Docker file should be provided both for the server and for the client that solves the POW challenge

## Design

## Goals

- Protect a TCP resource from DDOS attacks using Prof Of Work at the application layer of the OSI/TCP model.
- After Prof Of Work verification, server should send one of the quotes from "word of wisdom" book or any other collection of quotes.
- After the quote is sent, the connection should be closed.
- If the client fails to solve the POW challenge, the connection should be closed.
- After the connection is closed, the client should be able to reconnect and solve the POW challenge again.

### No targets

- Protection against TCP handshake attacks
- Extend a TCP packet to include a POW challenge in the header
- Filter IP addresses
- Include OS/Kernel level protection (iptables, eBPF, etc)

### Proof of Work algorithm requirements

- The server should be able to generate and verify the POW challenge with less effort than the client needs to solve it.
- The algorithm should be able to generate different challenges for different clients.
- The difficulty of the challenge should be adjustable depending on the load of the server (???).

## Implementation

### Algorithm

The algorithm is inspired by [THE JUELS PUZZLE
SCHEM](https://achsu3.github.io/client-puzzles-dsn19.pdf) paper, page 4.

1) The server generates a value and secret digits. (I choose number to simplify the implementation, but it can be any data type.)
2) The server creates a hash of the `value` and `secret` and calculates hash of them. (sha256 is used)
3) The server sends the value `value`, the first `m` bytes of the hash, and a number `K` to the client.
```   
   v = random_string()
   secret = random_number()
   hash = sha256(value + '$' + secret)
   
   send(value, hash[:m], K)
```


The client should find `K` solutions that produce a hash with the first `m` bytes equal to the hash sent by the server.
```
   for each ki in number_range:
      hash = sha256(value + '$' + ki)
      if hash[:m] == hash_from_server:
          bucket += ki
          if len(bucket) == K:
              send(bucket)
```

### Performance

#### server

generation = `O(1)`.

verification = `O(K)`, where `K` is the number of values expected from the client.

#### Client

`m` (bytes of hash) exponentially affects
`K` (number of expected values) acts linearly

```log
goos: darwin
goarch: arm64
pkg: tcp-pow/challenge
BenchmarkHasherChallenger
BenchmarkHasherChallenger/K_1_m_1
BenchmarkHasherChallenger/K_1_m_1-8         	   41494	     28908 ns/op
BenchmarkHasherChallenger/K_4_m_1
BenchmarkHasherChallenger/K_4_m_1-8         	   10000	    123638 ns/op
BenchmarkHasherChallenger/K_16_m_1
BenchmarkHasherChallenger/K_16_m_1-8        	    2533	    508453 ns/op
BenchmarkHasherChallenger/K_32_m_1
BenchmarkHasherChallenger/K_32_m_1-8        	    1204	    908951 ns/op
BenchmarkHasherChallenger/K_64_m_1
BenchmarkHasherChallenger/K_64_m_1-8        	     639	   1806048 ns/op
BenchmarkHasherChallenger/K_128_m_1
BenchmarkHasherChallenger/K_128_m_1-8       	     333	   3994640 ns/op
BenchmarkHasherChallenger/K_1_m_2
BenchmarkHasherChallenger/K_1_m_2-8         	     128	   8126764 ns/op
BenchmarkHasherChallenger/K_4_m_2
BenchmarkHasherChallenger/K_4_m_2-8         	      87	  30762981 ns/op
BenchmarkHasherChallenger/K_16_m_2
BenchmarkHasherChallenger/K_16_m_2-8        	      10	 118799829 ns/op
BenchmarkHasherChallenger/K_32_m_2
BenchmarkHasherChallenger/K_32_m_2-8        	       5	 229674433 ns/op
BenchmarkHasherChallenger/K_64_m_2
BenchmarkHasherChallenger/K_64_m_2-8        	       3	 522734333 ns/op
BenchmarkHasherChallenger/K_1_m_3
BenchmarkHasherChallenger/K_1_m_3-8         	       1	2432273625 ns/op
BenchmarkHasherChallenger/K_4_m_3
BenchmarkHasherChallenger/K_4_m_3-8         	       1	9757170209 ns/op
BenchmarkHasherChallenger/K_16_m_3
BenchmarkHasherChallenger/K_16_m_3-8        	       1	40964844291 ns/op
```

`This calc could be parallel, so the number of cores can be used to speed up the process*`