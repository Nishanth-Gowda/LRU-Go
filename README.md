# Go LRU Cache

This repository provides a simple and efficient implementation of an LRU (Least Recently Used) cache in Go.

## Features:

  - Thread-safe: Uses a mutex for safe concurrent access from multiple goroutines.
  - Flexible: Accepts keys and values of any data type using an empty interface.
  - Configurable: Allows setting a maximum capacity for the cache.
  - Optional Expiration: Supports storing key-value pairs with optional expiration times.

## Implementation

  1. Data Structure: The LRUCache is implemented using a combination of a hash map (map[interface{}]*list.Element) for fast lookups and a doubly linked list (*list.List) to maintain the order of recently accessed items.

  2. Node: The Node struct represents a key-value pair in the cache along with its expiration time. It has three fields: key for the key, value for the corresponding value, and expiresAt for the expiration time.

  3. LRUCache Struct: The LRUCache struct contains fields to maintain the cache's state. It includes a mutex (sync.Mutex) for concurrent access, capacity to store the maximum number of items allowed in the cache, cache to store key-value pairs, and queue to maintain the order of items based on their access time.

  4. NewLRUCache: The NewLRUCache function creates a new LRUCache instance with a specified capacity. It initializes the cache map and the linked list.

  5. Get: The Get method retrieves the value associated with a given key from the cache. If the key exists and has not expired, it moves the corresponding element to the front of the queue (indicating recent access) and returns the value. If the key does not exist or has expired, it returns nil along with false.

  6. Put: The Put method adds a new key-value pair to the cache or updates the value of an existing key. If the key already exists, it updates the value and expiration time and moves the corresponding element to the front of the queue. If the cache is at full capacity, it removes the least recently used item before adding the new one.

  7. Expiration Handling: The Put method also assigns an expiration time to each item added to the cache. Expired items are removed from the cache automatically during subsequent accesses or when attempting to retrieve them.

  - Helper Methods:

    - removeOldest: Removes the least recently used item from the cache.
    - removeElement: Removes a specific element from the cache.
    - moveToFront: Moves an element to the front of the queue, indicating recent access.
    - isExpired: Checks if a node (key-value pair) has expired based on its expiration time.

## Installation:

  - Clone the repository:
    ```
    git clone https://github.com/your-username/lru-cache.git
    ```
  - Go to the project directory:
    ```
    cd lru-cache
    ```
  - Install Dependencies:
    ```
    go mod tidy
    ```
  - Run the program:
    ```
    go run main.go
    ```
## Contributing:

Feel free to submit pull requests for bug fixes, improvements, or new features.

## License:

This code is distributed under the MIT license. See the LICENSE file for details.
