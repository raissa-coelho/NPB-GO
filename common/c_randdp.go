package common

func randlc(*x, a float64){
	var t1,t2,t3,t4,a1,a2,x1,x2,z float64

	t1 = r23*a
	a1 = int64(t1)
	a2 = a - t23 * a1

	t1 = r23 * (*x)
	x1 = int64(t1)
	x2 = (*x) - t23 * x1
	t1 = a1 * x2 + a2 * x1
	t2 = int64(r23 * t1)
	z = t1 - t23 * t2
	t3 = t23 * z + a2 * x2
	t4 = int64(r46 * t3)
	(*x) = t3 - t46 * t4

	return (r46 * (*x))
}