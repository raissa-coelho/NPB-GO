package IS

import (
	r "NPB-GO/common"
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
	"sync"
	"golang.org/x/sync/semaphore"
)

var (
	TOTAL_KEYS_LOG_2    int
	MAX_KEY_LOG_2       int
	NUM_BUCKETS_LOG     int
	TOTAL_KEYS          int
	USE_BUCKETS         bool = false
	passed_verification int
	partial_verify_vals [TEST_ARRAY_SIZE]int
	test_index_array    [TEST_ARRAY_SIZE]int
	test_rank_array     [TEST_ARRAY_SIZE]int
	procs               int
)

const (
	MAX_ITERATIONS  int = 10
	TEST_ARRAY_SIZE int = 5
)

func IS(class string) {

	if class == "S" {
		TOTAL_KEYS_LOG_2 = 16
		MAX_KEY_LOG_2 = 11
		NUM_BUCKETS_LOG = 9
	} else if class == "W" {
		TOTAL_KEYS_LOG_2 = 20
		MAX_KEY_LOG_2 = 16
		NUM_BUCKETS_LOG = 10
	} else if class == "A" {
		TOTAL_KEYS_LOG_2 = 23
		MAX_KEY_LOG_2 = 19
		NUM_BUCKETS_LOG = 10
	} else if class == "B" {
		TOTAL_KEYS_LOG_2 = 25
		MAX_KEY_LOG_2 = 21
		NUM_BUCKETS_LOG = 10
	} else if class == "C" {
		TOTAL_KEYS_LOG_2 = 27
		MAX_KEY_LOG_2 = 23
		NUM_BUCKETS_LOG = 10
	} else if class == "D" {
		TOTAL_KEYS_LOG_2 = 31
		MAX_KEY_LOG_2 = 27
		NUM_BUCKETS_LOG = 10
	}

	if class == "D" {
		TOTAL_KEYS = 1 << TOTAL_KEYS_LOG_2
	} else {
		TOTAL_KEYS = 1 << TOTAL_KEYS_LOG_2
	}

	var (
		MAX_KEY     int = 1 << MAX_KEY_LOG_2
		NUM_BUCKETS int = 1 << NUM_BUCKETS_LOG
		NUM_KEYS    int = TOTAL_KEYS
	)

	key_array := make([]int, NUM_KEYS)
	key_buff1 := make([]int, MAX_KEY)
	key_buff2 := make([]int, NUM_KEYS)
	procs = runtime.NumCPU()

	S_test_index_array := [TEST_ARRAY_SIZE]int{48427, 17148, 23627, 62548, 4431}
	S_test_rank_array := [TEST_ARRAY_SIZE]int{0, 18, 346, 64917, 65463}

	W_test_index_array := [TEST_ARRAY_SIZE]int{357773, 934767, 875723, 898999, 404505}
	W_test_rank_array := [TEST_ARRAY_SIZE]int{1249, 11698, 1039987, 1043896, 1048018}

	A_test_index_array := [TEST_ARRAY_SIZE]int{2112377, 662041, 5336171, 3642833, 4250760}
	A_test_rank_array := [TEST_ARRAY_SIZE]int{104, 17523, 123928, 8288932, 8388264}

	B_test_index_array := [TEST_ARRAY_SIZE]int{41869, 812306, 5102857, 18232239, 26860214}
	B_test_rank_array := [TEST_ARRAY_SIZE]int{33422937, 10244, 59149, 33135281, 99}

	C_test_index_array := [TEST_ARRAY_SIZE]int{44172927, 72999161, 74326391, 129606274, 21736814}
	C_test_rank_array := [TEST_ARRAY_SIZE]int{61147, 882988, 266290, 133997595, 133525895}

	D_test_index_array := [TEST_ARRAY_SIZE]int{1317351170, 995930646, 1157283250, 1503301535, 1453734525}
	D_test_rank_array := [TEST_ARRAY_SIZE]int{1, 36538729, 1978098519, 2145192618, 2147425337}

	for i := 0; i < TEST_ARRAY_SIZE; i++ {
		switch class {
		case "S":
			test_index_array[i] = S_test_index_array[i]
			test_rank_array[i] = S_test_rank_array[i]
		case "A":
			test_index_array[i] = A_test_index_array[i]
			test_rank_array[i] = A_test_rank_array[i]
		case "W":
			test_index_array[i] = W_test_index_array[i]
			test_rank_array[i] = W_test_rank_array[i]
		case "B":
			test_index_array[i] = B_test_index_array[i]
			test_rank_array[i] = B_test_rank_array[i]
		case "C":
			test_index_array[i] = C_test_index_array[i]
			test_rank_array[i] = C_test_rank_array[i]
		case "D":
			test_index_array[i] = D_test_index_array[i]
			test_rank_array[i] = D_test_rank_array[i]
		}
	}

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("NAS Parallel Benchmark Parallel GO version - IS Benchmark")
	fmt.Printf("Class: %s\n", class)
	fmt.Printf("Size: %d\n", TOTAL_KEYS)
	fmt.Printf("Iterations: %d\n", MAX_ITERATIONS)

	//rand.Seed(time.Now().UnixNano())
	//myid := rand.Intn(procs)

	key_buff1_aptr := make([][]int, procs)
	for i := range key_buff1_aptr {
		key_buff1_aptr[i] = make([]int, MAX_KEY)
	}

	bucket_size := make([][]int, procs)
	for i := range bucket_size {
		bucket_size[i] = make([]int, NUM_BUCKETS)
	}

	key_buff_ptr_global := make([]int, MAX_KEY)

	bucket_ptrs := make([]int, NUM_BUCKETS)

	fmt.Printf("Myid: %d\n", myid)
	
	// create_seq part
	var groupC_S sync.WaitGroup
	groupC_S.Add(procs)
	for i := 0; i < procs; i++{
		go create_seq(314159265.00, 1220703125.00, i, NUM_KEYS, MAX_KEY, procs, &groupC_S)
	}
	groupC_S.Wait()

	// alloc_key_buff
	alloc_key_buff(NUM_KEYS, key_buff1, key_buff2, key_buff1_aptr)

	passed_verification = 0

	if class != "S" {
		fmt.Println("iteration")
	}

	//var ch chan []int
	ch := make(chan []int, MAX_KEY)
	ch2 := make(chan []int, NUM_KEYS)
	start := time.Now()
	//main iteration
	//CRITICAL
	sem := semaphore.NewWeighted(int64(procs))
	for iteration := 1; iteration <= MAX_ITERATIONS; iteration++ {
		sem.Acquire(context.Background(), 1)
		if class != "S" {
			fmt.Printf("%d\n", iteration)
		}
		ch <- key_buff_ptr_global
		ch2 <- key_array
		go rank(iteration, myid, NUM_KEYS, NUM_BUCKETS, ch2, key_buff1, key_buff2, bucket_ptrs, MAX_KEY, class, key_buff1_aptr, bucket_size, key_buff_ptr_global, ch)
		key_array = <-ch2
		key_buff_ptr_global = <-ch
		sem.Release(1)
	}
	defer close(ch)
	defer close(ch2)
	stop := time.Now()
	t := stop.Sub(start)
	//wg.Wait()

	// This tests that keys are in sequence
	 temp01 := full_verify(myid,procs,NUM_KEYS, MAX_KEY, NUM_BUCKETS, key_buff1, key_buff2, key_array[:], bucket_ptrs, key_buff_ptr_global[:])
	
	passed_verification += temp01
	fmt.Printf("Passed verification: %d\n", passed_verification)
	fmt.Printf("vvv: %d\n", ((5*MAX_ITERATIONS)+1))
	
	var aux bool = true
	if passed_verification != ((5*MAX_ITERATIONS)+1) {
		passed_verification = 0
		aux = false
	}
	
	fmt.Printf("Passed verification: %d\n", passed_verification)
	
	Mops := float64((MAX_ITERATIONS * TOTAL_KEYS)) / t.Seconds() / 1000000.0

	r.C_file_IS(aux, "IS", class, TOTAL_KEYS, MAX_ITERATIONS, Mops, &t)
	r.C_print_results(class, "Keys Ranked", MAX_ITERATIONS, aux, Mops, &t, runtime.NumCPU())
}

func create_seq(seed, a float64, myid, NUM_KEYS, MAX_KEY, procs int, group *sync.WaitGroup) {
	var x, s float64
	var mq, k1, k2, k int
	var an float64 = a

	defer (*group).Done()	
	
	mq = (NUM_KEYS + procs - 1) / procs
	k1 = mq * myid
	k2 = k1 + mq

	if k2 > NUM_KEYS {
		k2 = NUM_KEYS
	}

	s = find_my_seed(myid, procs, int64(4*NUM_KEYS), seed, an)

	k = MAX_KEY / 4

	for i := k1; i < k2; i++ {
		x = r.Orandlc(&s, &an)
		x += r.Orandlc(&s, &an)
		x += r.Orandlc(&s, &an)
		x += r.Orandlc(&s, &an)
		key_array[i] = k * int(x)
	}
}

func find_my_seed(kn, np int, nn int64, s, a float64) float64 {
	var t1, t2 float64
	var mq, nq, kk, ik int64

	if kn == 0 {
		return s
	}
	mq = (nn/4 + int64(np) - 1) / int64(np)
	nq = mq * 4 * int64(kn)
	t1 = s
	t2 = a
	kk = nq

	for kk > 1 {
		ik = kk / 2
		if 2*ik == kk {
			r.Orandlc(&t2, &t2)
			kk = ik
		} else {
			r.Orandlc(&t1, &t2)
			kk = kk - 1
		}
	}
	r.Orandlc(&t1, &t2)
	return t1
}

func alloc_key_buff(NUM_KEYS int, key_buff1, key_buff2 []int, key_buff1_aptr [][]int) {

	if USE_BUCKETS {

		for i := 0; i < NUM_KEYS; i++ {
			key_buff2[i] = 0
		}

	} else {

		key_buff1_aptr[0] = key_buff1

	}
}

func rank(iteration, myid, NUM_KEYS, NUM_BUCKETS int, ch2 chan []int, key_buff1, key_buff2, bucket_ptrs []int, MAX_KEY int, class string, key_buff1_aptr, bucket_size [][]int, key_buff_ptr_global []int, ch chan []int) {
	var k, k1, k2, shift, num_bucket_keys int
	var key_buff_ptr, work_buff []int
	var key_buff_ptr2 []int

	key_array := make([]int, NUM_KEYS)
	key_array = <-ch2

	if USE_BUCKETS {
		shift = MAX_KEY_LOG_2 - NUM_BUCKETS_LOG
		num_bucket_keys = 1 << shift
	}

	key_array[iteration] = iteration
	key_array[iteration+MAX_ITERATIONS] = MAX_KEY - iteration

	for i := 0; i < TEST_ARRAY_SIZE; i++ {
		partial_verify_vals[i] = key_array[test_index_array[i]]
	}

	if USE_BUCKETS {
		key_buff_ptr2 = key_buff2
	} else {
		key_buff_ptr2 = key_array
	}
	key_buff_ptr = key_buff1

	if USE_BUCKETS {

		work_buff = bucket_size[myid]

		for i := 0; i < NUM_BUCKETS; i++ {
			work_buff[i] = 0
		}

		for i := 0; i < NUM_KEYS; i++ {
			work_buff[key_array[i]>>shift]++
		}
		bucket_ptrs[0] = 0

		for k := 0; k < myid; k++ {
			bucket_ptrs[0] += bucket_size[k][0]
		}

		for i := 1; i < NUM_BUCKETS; i++ {
			bucket_ptrs[i] = bucket_ptrs[i-1]
			for k := 0; k < NUM_KEYS; k++ {
				bucket_ptrs[i] += bucket_size[k][i]
			}
			for k := myid; k < procs; k++ {
				bucket_ptrs[i] += bucket_size[k][i-1]
			}
		}

		for i := 0; i < NUM_KEYS; i++ {
			k = key_array[i]
			key_buff2[bucket_ptrs[k>>shift]+1] = k
		}

		if myid < procs-1 {
			for i := 0; i < NUM_BUCKETS; i++ {
				for k := myid + 1; k < procs; k++ {
					bucket_ptrs[i] += bucket_size[k][i]
				}
			}
		}

		for i := 0; i < NUM_BUCKETS; i++ {
			k1 = i * num_bucket_keys
			k2 = k1 + num_bucket_keys

			for j := k1; j < k2; j++ {
				key_buff_ptr[j] = 0
			}

			var m int
			if i > 0 {
				m = bucket_ptrs[i-1]
			} else {
				m = 0
			}

			for k := m; k < bucket_ptrs[i]; k++ {
				key_buff_ptr[key_buff_ptr2[k]] = key_buff_ptr[key_buff_ptr2[k]] + 1
			}

			key_buff_ptr[k1] += m
			for k := k1 + 1; k < k2; k++ {
				key_buff_ptr[k] += key_buff_ptr[k-1]
			}
		}

		//fim USE BUCKET
	} else {

		// cria work_buff com o tamanho da array da key_buff1_aptr[myid]
		work_buff = key_buff1_aptr[myid]

		// init
		for i := 0; i < MAX_KEY; i++ {
			work_buff[i] = 0
		}

		// passa para o work_buff os valores
		for i := 0; i < NUM_KEYS; i++ {
			work_buff[key_buff_ptr2[i]] = work_buff[key_buff_ptr2[i]] + 1
		}
		for i := 0; i < MAX_KEY-1; i++ {
			work_buff[i+1] += work_buff[i]
		}
		for k := 1; k < procs; k++ {
			for i := 0; i < MAX_KEY; i++ {
				key_buff_ptr[i] += key_buff1_aptr[k][i]
			}
		}

	}

	for i := 0; i < TEST_ARRAY_SIZE; i++ {
		k = partial_verify_vals[i]
		if 0 < k && k <= NUM_KEYS-1 {
			var key_rank int = key_buff_ptr[k-1]
			var test_rank int = test_rank_array[i]
			var failed int = 0

			switch class {
			case "S":
				if i <= 2 {
					test_rank += iteration
				} else {
					test_rank -= iteration
				}
			case "W":
				if i < 2 {
					test_rank += iteration - 2
				} else {
					test_rank -= iteration
				}
			case "A":
				if i <= 2 {
					test_rank += iteration - 1
				} else {
					test_rank -= iteration - 1
				}
			case "B":
				if i == 1 || i == 2 || i == 4 {
					test_rank += iteration
				} else {
					test_rank -= iteration
				}
			case "C":
				if i <= 2 {
					test_rank += iteration
				} else {
					test_rank -= iteration
				}
			case "D":
				if i < 2 {
					test_rank += iteration
				} else {
					test_rank -= iteration
				}
			case "E":
				if i < 2 {
					test_rank += iteration - 2
				}else if i == 2{
					test_rank += iteration - 2
					if iteration > 4{
						test_rank -= 2
					}else if iteration > 2 {
						test_rank -= 1
					}
				}else{
					test_rank -= iteration - 2
				}
			}
			
			if key_rank != test_rank {
				failed = 1
			} else {
				passed_verification++
			}

			if failed == 1 {
				fmt.Printf("Failed partial verification: \n")
				fmt.Printf("iteration %d, test key %d\n", iteration, i)
			}

		}
	}

	var t []int

	if iteration == MAX_ITERATIONS {
		t = key_buff_ptr
	}

	ch <- t
}

func full_verify(myid,num_procs,NUM_KEYS, MAX_KEY, NUM_BUCKETS int, key_buff1, key_buff2 []int, key_array []int, bucket_ptrs, key_buff_ptr_global []int) int {
	var k, k1, k2, j int
	//var myid, num_procs int

	//myid = 0
	//num_procs = 1

	if USE_BUCKETS {
		for i := 0; i < NUM_BUCKETS; i++ {
			if i > 0 {
				k1 = bucket_ptrs[i-1]
			} else {
				k1 = 0
			}
			for j := k1; j < bucket_ptrs[i]; j++ {
				k = key_buff_ptr_global[key_buff2[j]] - 1
				key_array[k] = key_buff2[j]
			}
		}
	} else {

		for i := 0; i < NUM_KEYS; i++ {
			key_buff2[i] = key_array[i]
		}

		j = num_procs
		j = (MAX_KEY + j - 1) / j
		k1 = j * myid

		//j = MAX_KEY
		//k1 = 0
		
		k2 = k1 + j

		if k2 > MAX_KEY {
			k2 = MAX_KEY
		}

		for i := 0; i < NUM_KEYS; i++ {
			if key_buff2[i] >= k1 && key_buff2[i] < k2 {
				k = key_buff_ptr_global[key_buff2[i]] - 1
				key_array[k] = key_buff2[i]
			}
		}
	}

	j = 0

	for i := 1; i < NUM_KEYS; i++ {
		if key_array[i-1] > key_array[i] {
			j++
		}
	}

	if j != 0 {
		fmt.Printf("Full_verify: number of keys out of sort: %d\n", j)
	} else {
		passed_verification++
	}
	return passed_verification
}
