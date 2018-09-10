package processor

// ConditionalSum performs the sum of the elements of the input slice of integers considering only elements greater than 128.
func ConditionalSum(data []int, iterations int) int {
	var sum int
	for i := 0; i < iterations; i++ {
		for c := 0; c < len(data); c++ {
			if data[c] >= 128 {
				sum += data[c]
			}
		}
	}
	return sum
}

/***************************************

         ,_---~~~~~----._
  _,,_,*^____      _____``*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \`/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |

***************************************/

// ConditionalSumNoBranching performs the sum of the elements of the input slice of integers considering only elements greater than 128, using bitwise operations in order to avoid branching.
func ConditionalSumNoBranching(data []int, iterations int) int {
	var sum int
	for i := 0; i < iterations; i++ {
		for c := 0; c < len(data); c++ {
			sum += ^((data[c] - 128) >> 31) & data[c]
			// t := (data[c] - 128) >> 31
			// sum += ^t & data[c]
			//fmt.Printf("data[c] = %d\tdata[c] - 128 = %d [%032b]\t\tt = %032b\t^t = %032b\tadding %d\n", data[c], data[c]-128, data[c]-128, t, ^t, ^t&data[c])
		}
	}
	return sum
}
