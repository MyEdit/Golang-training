package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
    if n <= 0 {
        return 0, fmt.Errorf("input must be positive, got")
    }

    steps := 0
    for ; n != 1; steps++ {
        if n % 2 == 0 {
            n = n / 2
        } else {
            n = n * 3 + 1
        }
    }

    return steps, nil
}
 