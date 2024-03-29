package IS

import (
	r "NPB-GO/common"
	"fmt"
	"runtime"
	"time"
	"sync"
)

var (
	TOTAL_KEYS_LOG_2    int64
	MAX_KEY_LOG_2       int64
	NUM_BUCKETS_LOG     int64
	TOTAL_KEYS          int64
	USE_BUCKETS         bool
	passed_verification int
	partial_verify_vals [TEST_ARRAY_SIZE]int64
	test_index_array    [TEST_ARRAY_SIZE]int64
	test_rank_array     [TEST_ARRAY_SIZE]int64
	procs               int
	key_array           []int64
	key_buff1 	    []int64
	key_buff2 	    []int64
	key_buff1_aptr      [][]int64
	bucket_size         [][]int64
	key_buff_ptr_global []int64
	bucket_ptrs 	    []int64
	MAX_KEY     	    int
	NUM_BUCKETS         int
	NUM_KEYS            int
	classNPB            string
)

const (
	MAX_ITERATIONS  int = 10
	TEST_ARRAY_SIZE int = 5
)

func IS(class string) {
	
	classNPB = class
	USE_BUCKETS = false
	
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
		TOTAL_KEYS = int64(1) << TOTAL_KEYS_LOG_2
	} else {
		TOTAL_KEYS = 1 << TOTAL_KEYS_LOG_2
	}

	MAX_KEY     = 1 << MAX_KEY_LOG_2
	NUM_BUCKETS = 1 << NUM_BUCKETS_LOG
	NUM_KEYS    = int(TOTAL_KEYS)

	key_array = make([]int64, NUM_KEYS)
	key_buff1 = make([]int64, MAX_KEY)
	key_buff2 = make([]int64, NUM_KEYS)
	procs = runtime.NumCPU()

	S_test_index_array := [TEST_ARRAY_SIZE]int64{48427, 17148, 23627, 62548, 4431}
	S_test_rank_array := [TEST_ARRAY_SIZE]int64{0, 18, 346, 64917, 65463}

	W_test_index_array := [TEST_ARRAY_SIZE]int64{357773, 934767, 875723, 898999, 404505}
	W_test_rank_array := [TEST_ARRAY_SIZE]int64{1249, 11698, 1039987, 1043896, 1048018}

	A_test_index_array := [TEST_ARRAY_SIZE]int64{2112377, 662041, 5336171, 3642833, 4250760}
	A_test_rank_array := [TEST_ARRAY_SIZE]int64{104, 17523, 123928, 8288932, 8388264}

	B_test_index_array := [TEST_ARRAY_SIZE]int64{41869, 812306, 5102857, 18232239, 26860214}
	B_test_rank_array := [TEST_ARRAY_SIZE]int64{33422937, 10244, 59149, 33135281, 99}

	C_test_index_array := [TEST_ARRAY_SIZE]int64{44172927, 72999161, 74326391, 129606274, 21736814}
	C_test_rank_array := [TEST_ARRAY_SIZE]int64{61147, 882988, 266290, 133997595, 133525895}

	D_test_index_array := [TEST_ARRAY_SIZE]int64{1317351170, 995930646, 1157283250, 1503301535, 1453734525}
	D_test_rank_array := [TEST_ARRAY_SIZE]int64{1, 36538729, 1978098519, 2145192618, 2147425337}

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

	key_buff_ptr_global = make([]int64, MAX_KEY)

	bucket_ptrs = make([]int64, NUM_BUCKETS)
	
	// create_seq part
	var groupC_S sync.WaitGroup
	groupC_S.Add(procs)
	for i := 0; i < procs; i++{
		go create_seq(314159265.00, 1220703125.00, i, &groupC_S)
	}
	groupC_S.Wait()

	// alloc_key_buff
	alloc_key_buff()
	
	var tmp sync.WaitGroup
	tmp.Add(1)
	go rank(1, &tmp)
	tmp.Wait()
	
	passed_verification = 0

	if class != "S" {
		fmt.Println("iteration")
	}

	//main iteration
	//CRITICAL
	var groupR sync.WaitGroup
	groupR.Add(procs)
	
	start := time.Now()
	for iteration := 1; iteration <= MAX_ITERATIONS; iteration++ {
		if class != "S" {
			fmt.Printf("%d\n", iteration)
		}
		go rank(iteration, &groupR)
	}
	groupR.Wait()
	stop := time.Now()
	t := stop.Sub(start)

	// This tests that keys are in sequence
	full_verify()

	var aux bool = true
	if passed_verification != ((5*MAX_ITERATIONS)+1) {
		passed_verification = 0
		aux = false
	}
	
	Mops := float64((MAX_ITERATIONS * int(TOTAL_KEYS))) / t.Seconds() / 1000000.0

	r.C_file_IS(aux, "IS", class, int(TOTAL_KEYS), MAX_ITERATIONS, Mops, &t)
	r.C_print_results(class, "Keys Ranked", MAX_ITERATIONS, aux, Mops, &t, runtime.NumCPU())
}

func create_seq(seed, a float64, myid int, group *sync.WaitGroup) {
	var x, s float64
	var mq, k1, k2, k int64
	var an float64 = a

	defer (*group).Done()	
	
	mq = (int64(NUM_KEYS) + int64(procs) - 1) / int64(procs)
	k1 = mq * int64(myid)
	k2 = k1 + mq

	if k2 > int64(NUM_KEYS) {
		k2 = int64(NUM_KEYS)
	}

	s = find_my_seed(myid, procs, int64(4*NUM_KEYS), seed, an)

	k = int64(MAX_KEY) / 4

	for i := k1; i < k2; i++ {
		x = r.Orandlc(&s, &an)
		x += r.Orandlc(&s, &an)
		x += r.Orandlc(&s, &an)
		x += r.Orandlc(&s, &an)
		key_array[i] = k * int64(x)
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

func alloc_key_buff() {
	
	if USE_BUCKETS {

		bucket_size = make([][]int64, procs)
		for i := range bucket_size {
			bucket_size[i] = make([]int64, NUM_BUCKETS)
		}

		for i := 0; i < NUM_KEYS; i++ {
			key_buff2[i] = 0
		}

	} else {
		key_buff1_aptr = make([][]int64, procs)
		key_buff1_aptr[0] = key_buff1
		for i := range key_buff1_aptr {
			key_buff1_aptr[i] = make([]int64, MAX_KEY)
		}
	}
}

func rank(iteration int, groupN *sync.WaitGroup) {
	var k, k1, k2, num_bucket_keys int64
	var shift int
	var key_buff_ptr, work_buff []int64
	var key_buff_ptr2 []int64
	
	defer (*groupN).Done()
	
	myid := iteration
	
	if USE_BUCKETS {
		shift = int(MAX_KEY_LOG_2) -int(NUM_BUCKETS_LOG)
		num_bucket_keys = 1 << shift
	}

	key_array[iteration] = int64(iteration)
	key_array[iteration+MAX_ITERATIONS] = int64(MAX_KEY) - int64(iteration)

	for i := 0; i < TEST_ARRAY_SIZE; i++ {
		partial_verify_vals[i] = key_array[test_index_array[i]]
	}

	if USE_BUCKETS {
		key_buff_ptr2 = key_buff2[:]
	} else {
		key_buff_ptr2 = key_array[:]
	}
	key_buff_ptr = key_buff1[:]

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
			for k := 0; k < myid; k++ {
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
			k1 = int64(i) * num_bucket_keys
			k2 = k1 + num_bucket_keys

			for j := k1; j < k2; j++ {
				key_buff_ptr[j] = 0
			}

			var m int64
			if i > 0 {
				m = bucket_ptrs[i-1]
			} else {
				m = 0
			}

			for k := m; k < bucket_ptrs[i]; k++ {
				key_buff_ptr[key_buff_ptr2[k]]++
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
			work_buff[key_buff_ptr2[i]]++
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
		if 0 < k && k <= int64(NUM_KEYS)-1 {
			var key_rank int64 = key_buff_ptr[k-1]
			var test_rank int64 = test_rank_array[i]
			var failed int = 0

			switch classNPB {
			case "S":
				if i <= 2 {
					test_rank += int64(iteration)
				} else {
					test_rank -= int64(iteration)
				}
			case "W":
				if i < 2 {
					test_rank += int64(iteration) - 2
				} else {
					test_rank -= int64(iteration)
				}
			case "A":
				if i <= 2 {
					test_rank += int64(iteration) - 1
				} else {
					test_rank -= int64(iteration) - 1
				}
			case "B":
				if i == 1 || i == 2 || i == 4 {
					test_rank += int64(iteration)
				} else {
					test_rank -= int64(iteration)
				}
			case "C":
				if i <= 2 {
					test_rank += int64(iteration)
				} else {
					test_rank -= int64(iteration)
				}
			case "D":
				if i < 2 {
					test_rank += int64(iteration)
				} else {
					test_rank -= int64(iteration)
				}
			case "E":
				if i < 2 {
					test_rank += int64(iteration) - 2
				}else if i == 2{
					test_rank += int64(iteration) - 2
					if iteration > 4{
						test_rank -= 2
					}else if iteration > 2 {
						test_rank -= 1
					}
				}else{
					test_rank -= int64(iteration) - 2
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

	if iteration == MAX_ITERATIONS {
		key_buff_ptr_global = key_buff_ptr[:]
	}
}

func full_verify() {
	var k, k1, k2, j int64
	var myid, num_procs int

	myid = 0
	num_procs = 1

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

		j = int64(num_procs)
		j = (int64(MAX_KEY) + j - 1) / j
		k1 = j * int64(myid)
		k2 = k1 + j

		if k2 > int64(MAX_KEY) {
			k2 = int64(MAX_KEY)
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
}
