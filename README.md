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

The algorithm is inspired by the [Hashcash](https://achsu3.github.io/client-puzzles-dsn19.pdf) algorithm.

1) The server generates a value and secret numbers. (I choose number to simplify the implementation, but it can be any data type)
2) The server makes a hash of the value and secret and sends it to the client. (sha256 is used)
3) The server server sends the value `value`, first `m` bytes of the hash, and a number `K` to the client.
```   
   v = random_string()
   secret = random_number()
   hash = sha256(v + '$' +  secret)
   
   send(hash, v, hash[:m], K)
```


The client should find `K` solutions that produces a hash with the first `N` bytes equal to the hash sent by the server.
```
   hash = sha256(v + '$' + ki)
   if hash[:m] == hash_from_server:
       bucket += ki
       if len(bucket) == K:
           send(bucket)
```

#### Performance

##### Server

generation = `O(1)`

verification = `O(K)`, where `K` is the number of values that is expected to be received from the client.

##### Client


