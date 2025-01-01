# Go Blockchain Project
### Video Demo:  
<URL HERE>

### Overview
This project is a blockchain implementation developed in Go, inspired by concepts from Harvard's CS50 course. It serves as an educational tool to understand blockchain technology from a coding perspective, emphasizing efficient memory management through the use of pointers and Go modules for dependency management.

### Features
Blockchain Implementation: A simplified blockchain that includes block creation, validation, and chain verification, providing insights into the mechanics of blockchain technology.

Go Modules: Utilizes Go modules for efficient dependency management, ensuring a modular and maintainable codebase.

Memory Optimization with Pointers: Implements pointers to manage memory effectively, enhancing performance and demonstrating Go's capabilities in handling data structures.

### Project Structure
The project comprises several key components:

Block Structure: 
Defines the fundamental unit of the blockchain, including attributes like index, timestamp, data, previous hash, and hash.

Blockchain Management: 
Functions to create the genesis block, add new blocks, validate the chain, and reset the blockchain.

Persistence: 
Implements saving and loading of the blockchain to and from a JSON file, ensuring data persistence across sessions.

User Interface: 
A command-line interface that allows users to interact with the blockchain, including adding new blocks, viewing the blockchain, verifying its integrity, and resetting it.

### Installation and Usage
Clone the Repository:
git clone https://github.com/keenans1/go_blockchain.git
cd go-blockchain

Run the Application:
go run main.go
Follow the on-screen instructions to interact with the blockchain.

### Learning Outcomes
This project provided valuable insights into:

Blockchain Mechanics: 
Understanding how blocks are linked, how data integrity is maintained, and how consensus can be achieved in a decentralized system.

Go Programming: 
Gaining proficiency in Go, particularly in using pointers for memory management and Go modules for dependency handling.

CS50 Concepts: 
Applying theoretical knowledge from CS50 to a practical, real-world application, reinforcing learning through hands-on experience.

### Future Enhancements
Potential improvements for this project include:

Networking: Implementing peer-to-peer communication to simulate a decentralized network of nodes.

Consensus Algorithms: Introducing consensus mechanisms like Proof of Work (PoW) or Proof of Stake (PoS) to enhance the blockchain's robustness.

Advanced Data Structures: Utilizing more complex data structures, such as Merkle trees, to improve data integrity and verification processes.

### Conclusion
Developing this Go-based blockchain project has been an enlightening experience, bridging theoretical concepts from CS50 with practical application. It has deepened my understanding of both blockchain technology and the Go programming language, highlighting the importance of efficient memory management and modular code design in building scalable and maintainable systems.

For more information and to explore the codebase, visit the GitHub repository.