func shipWithinDays(weights []int, days int) int {
    // if day is 1 / the smallest -> capacity have to be sum of the weights
    // if day > len(weights) -> capacity have to be the biggest weight value within weights
    // lower limit is max capacity of a weight and upper limit is the sum of the weights
    // if we have a certain capacity, we can check how many days it is
    // feasibility is monotonic, once a smallest weight works (minimum capacity) the larger capacity would also works
    // since we have lower limit, upper limit, and a feasibility, we can go for binary search to find the minimum capacity that satisfy the days
    // from a certain capacity, we can find the days of it

    maxWeight := 0
    sumWeight := 0
    for _, weight := range weights {
        if maxWeight < weight {
            maxWeight = weight
        }
        sumWeight += weight
    }

    leftPointer := maxWeight
    rightPointer := sumWeight
    capacity := 0
    for leftPointer < rightPointer {
        capacity = (leftPointer + rightPointer)/2
        expectedDays := calculateDays(weights,capacity)
        if expectedDays <= days {
            rightPointer = capacity
        } else {
            leftPointer = capacity+1
        }
    }

    return rightPointer
}

func calculateDays(weights []int, capacity int) int {
    days := 1 
    currentCapacity := 0
    for _, weight := range weights {
        currentCapacity += weight
        if currentCapacity > capacity {
            days++
            currentCapacity = weight
        } 
    }

    return days
}