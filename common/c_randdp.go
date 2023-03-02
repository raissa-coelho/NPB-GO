package common

const(
		r23 = (0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5*0.5)
		r46 = r23 * r23
		t23 = (2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0*2.0)
		t46 = t23 * t23
)

func Randlc(x *float64, a float64) float64 {
	
	var t1,t2,t3,t4,a1,a2,x1,x2,z float64

	t1 = r23*a
	a1 = float64(int(t1))
	a2 = a - t23 * a1

	t1 = r23 * (*x)
	x1 = float64(int(t1))
	x2 = (*x) - t23 * x1
	t1 = a1 * x2 + a2 * x1
	t2 = float64(int(r23 * t1))
	z = t1 - t23 * t2
	t3 = t23 * z + a2 * x2
	t4 = float64(int(r46 * t3))
	(*x) = t3 - t46 * t4

	return (r46 * (*x))
}

func Vrandlc(n int, x_speed *float64, a float64, y []float64){

	var x,t1,t2,t3,t4,a1,a2,x1,x2,z float64

	t1 = r23 * a
	a1 = float64(int(t1))
	a2 = a - t23 * a1
	x  = *x_speed

	for i:=0; i < n; i++{
		t1 = r23 * x
		x1 = float64(int(t1))
		x2 = x - t23 * x1
		t1 = a1 * x2 + a2 * x1
		t2 = float64(int(r23 * t1))
		z = t1 - t23 * t2
		t3 = t23 * z + a2 * x2
		t4 = float64(int(r46 * t3))
		x = t3 - t46 * t4
		y[i] = r46 * x
	}

	*x_speed = x
}
