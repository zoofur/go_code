package main

import "fmt"

// 1640531527 ≈ 2 ^ 32 * (1 - 1 / φ), φ = (√5 + 1) ÷ 2
const HASH_INCREMENT = 0x61c88647

func main() {
	hash(4)
	hash(16)
	hash(32)
}

// 连续生成的哈希码之间的差异(增量值)，转换为几乎最佳分布的乘法哈希值，这些不同的哈希值最终生成一个2的幂次方的哈希表。
func hash(capacity int) {
	keyIndex := 0
	for i := 0; i < capacity; i++ {
		keyIndex = (i + 1) * HASH_INCREMENT & (capacity - 1)
		fmt.Print(keyIndex, " ")
	}
	fmt.Println()
}
