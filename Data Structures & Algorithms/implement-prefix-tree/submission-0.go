type PrefixTree struct {
	val rune
	children map[rune]*PrefixTree
	isWord bool
}

func Constructor() PrefixTree {
    return PrefixTree{
		children: map[rune]*PrefixTree{},
	}
}

func (this *PrefixTree) Insert(word string) {
	curr := this
	for _, char := range []rune(word){
		_, ok := curr.children[char]
		if !ok {
			curr.children[char] = &PrefixTree{
				children: map[rune]*PrefixTree{},
			}
		}
		curr = curr.children[char]
	}
	curr.isWord=true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this
	for _, char := range []rune(word){
		_, ok := curr.children[char]
		if !ok {
			return false
		}
		curr = curr.children[char]
	}
	return curr.isWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this
	for _, char := range []rune(prefix){
		_, ok := curr.children[char]
		if !ok {
			return false
		}
		curr = curr.children[char]
	}
	return true
}
