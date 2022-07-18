package arrays

//%v placeholder to print the "default" format, which works well for arrays

/* func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
} */

// range lets you iterate over an array. On each iteration, range returns two values - the index and the value. We are choosing to ignore the index value by using _ blank identifier.

func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllArr(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers) // There's a new way to create a slice. make allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through

	for i, numberArr := range numbersToSum {
		sums[i] = SumAll(numberArr)
	}
	return sums
}

/* func SumAllArr(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, SumAll(numbers))
	}

	return sums
}
*/
func SumAllTails(numbersToSum ...[]int) (sum []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			tail := numbers[1:]
			sum = append(sum, SumAll(tail))
		}
	}
	return sum
}
