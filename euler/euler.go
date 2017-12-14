package main

func main() {
	// sum := 0
	// for i := 0; i < 1000; i++ {
	// 	if i%5 == 0 || i%3 == 0 {
	// 		sum += i
	// 	}
	// }
	// fmt.Println(sum)

	// a := 1
	// b := 2
	// sum := 2

	// for i := a + b; i < 4000000; i = a + b {
	// 	if i%2 == 0 {
	// 		sum += i
	// 	}
	// 	a, b = b, i
	// }

	// fmt.Println(sum)

	// primeFactor := make([]int, 10)
	// n := 600851475143

	// for n%2 == 0 {
	// 	primeFactor = append(primeFactor, 2)
	// 	n /= 2
	// }

	// for i := 3; i < int(math.Sqrt(float64(n))); i += 2 {
	// 	for n%i == 0 {
	// 		primeFactor = append(primeFactor, i)
	// 		n /= i
	// 	}
	// }

	// if n > 2 {
	// 	primeFactor = append(primeFactor, n)
	// }

	// max := primeFactor[0]
	// for _, i := range primeFactor {
	// 	if i > max {
	// 		max = i
	// 	}
	// }

	// fmt.Println(max)

	// currentLargest := 1
	// for i := 1; i < 1000; i++ {
	// 	for j := 1; j < 1000; j++ {
	// 		k := i * j
	// 		if k > currentLargest && isPalindrome(k) {
	// 			currentLargest = k
	// 		}
	// 	}
	// }

	// fmt.Println(currentLargest)

	// i := 20
	// for {
	// 	if isDivBy20(i) {
	// 		break
	// 	} else {
	// 		i += 2
	// 	}
	// }

	// fmt.Println(i)

	// sumOfSquares := 0
	// squareOfSums := 0
	// for i := 0; i <= 100; i++ {
	// 	sumOfSquares += i * i
	// 	squareOfSums += i
	// }

	// squareOfSums = squareOfSums * squareOfSums

	// fmt.Println(squareOfSums - sumOfSquares)

	// count := 2
	// poss := 5
	// for count < 10001 {
	// 	if isPrime(poss) {
	// 		count++
	// 	}
	// 	poss += 2
	// }

	// fmt.Println(poss - 2)

	// str := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"
	// size := len(str)
	// var max uint64
	// for i := 0; i < size-13; i++ {
	// 	var prod uint64 = 1
	// 	for j := 0; j < 13; j++ {
	// 		dig, _ := strconv.Atoi(string(str[i+j]))
	// 		prod *= uint64(dig)
	// 	}
	// 	if prod > max {
	// 		max = prod
	// 	}
	// }
	// fmt.Println(max)

	// for i := 1; i < 1000; i++ {
	// 	for j := 1; j < 1000; j++ {
	// 		for k := 1; k < 1000; k++ {
	// 			if i+j+k == 1000 && ((i*i)+(j*j) == (k * k)) {
	// 				fmt.Println(i * j * k)
	// 				return
	// 			}
	// 		}
	// 	}
	// }

	// var sum uint64 = 5
	// pot := 5
	// for pot < 2000000 {
	// 	if pot%10000 == 0 {
	// 		fmt.Println(pot)
	// 	}
	// 	if isPrime(pot) {
	// 		sum += uint64(pot)
	// 	}
	// 	pot += 2
	// }
	// fmt.Println(sum)

	// inc := 1
	// tri := 0
	// numFactors := 0
	// count := 0
	// for {
	// 	tri += inc
	// 	inc++
	// 	count++
	// 	numFactors = getNumFactors(tri)
	// 	if count%1000 == 0 {
	// 		fmt.Println(tri, " ", numFactors)
	// 	}
	// 	if numFactors > 500 {
	// 		fmt.Println(tri)
	// 		return
	// 	}
	// }

	// num := 1
	// max_len := 1

	// for i := 2; i < 1000000; i++ {
	// 	chain := getChain(i)
	// 	if chain > max_len {
	// 		max_len = chain
	// 		num = i
	// 	}
	// }
	// fmt.Println(num)

}

// func getChain(n int) int {
// 	counter := 0
// 	for {
// 		counter++
// 		if n == 1 {
// 			return counter
// 		}

// 		if n%2 == 0 {
// 			n /= 2
// 		} else {
// 			n = 3*n + 1
// 		}
// 	}
// }

// func getNumFactors(n int) int {
// 	count := 2
// 	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
// 		if n%i == 0 {
// 			count += 2
// 		}
// 	}
// 	return count
// }

// func isPrime(n int) bool {
// 	for i := 3; i < int(math.Sqrt(float64(n)))+1; i += 2 {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func isDivBy20(n int) bool {
// 	for i := 2; i <= 20; i++ {
// 		if n%i != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func isPalindrome(n int) bool {
// 	s := strconv.Itoa(n)
// 	for i := 0; i < len(s)/2; i++ {
// 		if s[i] != s[len(s)-i-1] {
// 			return false
// 		}
// 	}
// 	return true
// }
