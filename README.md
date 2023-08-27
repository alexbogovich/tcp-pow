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
- The difficulty of the challenge should be adjustable depending on the load of the server (???)