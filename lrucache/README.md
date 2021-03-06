[LRU Cache](https://leetcode.com/problems/lru-cache/)

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

* LRUCache(int capacity)
* Initialize the LRU cache with positive size capacity.
* int get(int key) Return the value of the key if the key exists, otherwise return -1.
* void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.

Constraints:
* 1 <= capacity <= 3000
* 0 <= key <= 3000
* 0 <= value <= 104
* At most 3 * 104 calls will be made to get and put.