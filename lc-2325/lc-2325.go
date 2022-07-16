package main

func main() {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	keyMap := make(map[string]int, 0)
	start := 97
	s := ""
	keyMap[" "] = 32
	for i := 0; i < len(key); i++ {
		if _, found := keyMap[string(key[i])]; !found {
			if key[i] > 96 && key[i] < 123 {
				keyMap[string(key[i])] = start
				start++
			}
		}
	}
	// fmt.Println(keyMap)

	for i := 0; i < len(message); i++ {
		s = s + string(keyMap[string(message[i])])
	}
	// fmt.Println(s)
}
