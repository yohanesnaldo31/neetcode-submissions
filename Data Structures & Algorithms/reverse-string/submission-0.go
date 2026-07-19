func reverseString(s []byte)  {
    leftPointer := 0
    rightPointer := len(s)-1

    for leftPointer<rightPointer{
        s[leftPointer],s[rightPointer] = s[rightPointer],s[leftPointer]
        leftPointer++
        rightPointer--
    }
}