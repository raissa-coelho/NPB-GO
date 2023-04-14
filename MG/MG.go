//NPB-GO ~ MG
//Bianca Nunes Coelho
//Raissa NUnes Coelho

package MG

import (
	r "NPB-GO/common"
	"fmt"
	"math"
	//"honnef.co/go/tools/go/ir"
)

// const MAX int
const X float64 = 314159265.0

var (
	M, A                         float64
	NM, NV, NR                   int
	MM                           int
	k, it, n1, n2, n3, lb        int
	is1, is2, is3, ie1, ie2, ie3 int
	t, tinit, mflops             float64
	a, c                         []float64
	rnm2, rnmu, epsilon          float64
	nn, verif, err               float64
	verify                       bool
	class                        string
	nx, ny, nz                   []int
)

func Mg(LM, NDIM1, NDIM2, NDIM3, lt_default, nit, problem_size int) {

	NM = (2 + (1 << LM))
	NV = (1 * (2 + (1 << NDIM1)) * (2 + (1 << NDIM2)) * (2 + (1 << NDIM3)))
	NR = (((NV + NM*NM + 5*NM + 7*(LM) + 6) / 7) * 8)
	MAXLEVEL := lt_default + 1

	M = float64(NM + 1)
	MM = 10
	A = math.Pow(5.0, 13.0)

	debug_vec := make([]int, 8)
	u := make([][][]float64, int(NR))
	v := make([][][]float64, int(NV))
	r := make([][][]float64, int(NR))

	nx[lt_default] = problem_size
	ny[lt_default] = problem_size
	nz[lt_default] = problem_size

	if nx[lt_default] != ny[lt_default] || nx[lt_default] != nz[lt_default] {
		class = "U"
	} else if nx[lt_default] == 32 && nit == 4 {
		class = "S"
	} else if nx[lt_default] == 128 && nit == 4 {
		class = "W"
	} else if nx[lt_default] == 256 && nit == 4 {
		class = "A"
	} else if nx[lt_default] == 256 && nit == 20 {
		class = "B"
	} else if nx[lt_default] == 512 && nit == 20 {
		class = "C"
	} else if nx[lt_default] == 1024 && nit == 50 {
		class = "D"
	} else if nx[lt_default] == 2048 && nit == 50 {
		class = "E"
	} else {
		class = "U"
	}

	a[0] = -8.0 / 3.0
	a[1] = 0.0
	a[2] = 1.0 / 6.0
	a[3] = 1.0 / 12.0

	if class == "A" || class == "S" || class == "W" {
		/* coefficients for the s(a) smoother */
		c[0] = -3.0 / 8.0
		c[1] = +1.0 / 32.0
		c[2] = -1.0 / 64.0
		c[3] = 0.0
	} else {
		/* coefficients for the s(b) smoother */
		c[0] = -3.0 / 17.0
		c[1] = +1.0 / 33.0
		c[2] = -1.0 / 61.0
		c[3] = 0.0
	}

	lb = 1
	k = lt_default

	setup(n1, n2, n3, k, lt_default, MAXLEVEL)
	zero3(u, n1, n2, n3)
	zran3(v, n1, n2, n3, nx[lt_default], ny[lt_default], k, MM)
	norm2u3(v, n1, n2, n3, rnm2, rnmu, nx[lt_default], ny[lt_default], nz[lt_default])

	fmt.Printf("\n\n NAS Parallel Benchmarks 4.1 Parallel GO language version - MG Benchmark\n\n")
	fmt.Printf(" Size: %3dx%3dx%3d (class_npb %1c)\n", nx[lt_default], ny[lt_default], nz[lt_default], class)
	fmt.Printf(" Iterations: %3d\n", nit)

	go resid(debug_vec, u, v, r, n1, n2, n3, a, k)
	go norm2u3(r, n1, n2, n3, rnm2, rnmu, nx[lt_default], ny[lt_default], nz[lt_default])
	go mg3P(debug_vec, lt_default, u, v, r, a, c, n1, n2, n3, k)
	go resid(debug_vec, u, v, r, n1, n2, n3, a, k)

	setup(n1, n2, n3, k, lt_default, MAXLEVEL)
	zero3(u, n1, n2, n3)
	zran3(v, n1, n2, n3, nx[lt_default], ny[lt_default], k, MM)

	//fmt.Printf("Initialization time: %15.3f seconds\n", )

	go resid(debug_vec, u, v, r, n1, n2, n3, a, k)
	go norm2u3(r, n1, n2, n3, rnm2, rnmu, nx[lt_default], ny[lt_default], nz[lt_default])
	for it = 1; it <= nit; it++ {
		if (it == 1) || (it == nit) || ((it % 5) == 0) {
			fmt.Printf("iter %3d\n", it)
		}
		go mg3P(debug_vec, lt_default, u, v, r, a, c, n1, n2, n3, k)
		go resid(debug_vec, u, v, r, n1, n2, n3, a, k)
	}
	go norm2u3(r, n1, n2, n3, rnm2, rnmu, nx[lt_default], ny[lt_default], nz[lt_default])

	var verified bool = false
	var verify_value float64 = 0.0
	fmt.Println("Benchmark Completed!")
	epsilon = 1.0e-8

	if class != "U" {
		if class == "S" {
			verify_value = 0.5307707005734e-04
		} else if class == "W" {
			verify_value = 0.6467329375339e-05
		} else if class == "A" {
			verify_value = 0.2433365309069e-05
		} else if class == "B" {
			verify_value = 0.1800564401355e-05
		} else if class == "C" {
			verify_value = 0.5706732285740e-06
		} else if class == "D" {
			verify_value = 0.1583275060440e-09
		} else if class == "E" {
			verify_value = 0.8157592357404e-10
		}

		err = math.Abs(rnm2-verify_value) / verify_value
		if err <= epsilon {
			verified = true
			fmt.Println("VERIFICATION SUCCESSFUL")
			fmt.Printf("L2 Norm is %20.13e\n", rnm2)
			fmt.Printf("Error is %20.13e\n", err)
		} else {
			verified = false
			fmt.Println("VERIFICATION FAILED")
			fmt.Printf("L2 Norm is %20.13e\n", rnm2)
			fmt.Printf("The correct L2 Norm is %20.13e\n", verify_value)
		}
	} else {
		verified = false
		fmt.Println("Problem size unknown")
		fmt.Println("NO VERIFICATION PERFORMED")
	}

	nn = 1.0 * float64(nx[lt_default]*ny[lt_default]*nz[lt_default])

	if t != 0.0 {
		mflops = 58.0 * float64(nit) * nn * 1.0e-6 / t
	} else {
		mflops = 0.0
	}

}

func setup(n1, n2, n3, k, lt_default, MAXLEVEL int) {
	mi := make([][]int, MAXLEVEL+1)
	for i := 0; i < MAXLEVEL+1; i++ {
		mi[i] = make([]int, 3)
	}

	ng := make([][]int, MAXLEVEL+1)
	for i := 0; i < MAXLEVEL+1; i++ {
		ng[i] = make([]int, 3)
	}

	ng[lt_default][0] = nx[lt_default]
	ng[lt_default][1] = ny[lt_default]
	ng[lt_default][2] = nz[lt_default]

	for ax := 0; ax < 3; ax++ {
		for k := lt_default - 1; k >= 1; k-- {
			ng[k][ax] = ng[k+1][ax] / 2
		}
	}

	for k := lt_default; k >= 1; k-- {
		nx[k] = ng[k][0]
		ny[k] = ng[k][1]
		nz[k] = ng[k][2]
	}

	var m1, m2, m3 []int

	for k := lt_default; k >= 1; k-- {
		for ax := 0; ax < 3; ax++ {
			mi[k][ax] = 2 + ng[k][ax]
		}

		m1[k] = mi[k][0]
		m2[k] = mi[k][1]
		m3[k] = mi[k][2]
	}

	k = lt_default
	is1 = 2 + ng[k][0] - ng[lt_default][0]
	ie1 = 1 + ng[k][0]
	n1 = 3 + ie1 - is1
	is2 = 2 + ng[k][1] - ng[lt_default][1]
	ie2 = 1 + ng[k][1]
	n2 = 3 + ie2 - is2
	is3 = 2 + ng[k][2] - ng[lt_default][2]
	ie3 = 1 + ng[k][2]
	n3 = 3 + ie3 - is3

	var ir []int
	ir[lt_default] = 0

	for j := lt_default - 1; j >= 1; j-- {
		ir[j] = ir[j+1] + 1*m1[j+1]*m2[j+1]*m3[j+1]
	}

}

func zero3(u [][][]float64, n1, n2, n3 int) {

	for i := 0; i < n3; i++ {
		for j := 0; j < n2; j++ {
			for k := 0; k < n1; k++ {
				u[i][j][k] = 0.0
			}
		}
	}
}

func zran3(zz [][][]float64, n1, n2, n3, nx, ny, k, MM int) {

	var z [][][]float64
	z[n2][n1] = zz[n2][n1]

	var i0, m0, m1, i int
	var i1, i2, i3, d1, e1, e2, e3 int
	var xx, x0, x1, a1, a2, ai, best float64

	var ten, j1, j2, j3 [2][]float64
	var jg [2][][4]int

	a1 = math.Pow(A, float64(nx))
	a2 = math.Pow(A, float64(nx*ny))

	zero3(z, n1, n2, n3)

	i = is1 - 2 + nx*(is2-2+ny*(is3-2))

	ai = math.Pow(A, float64(i))
	d1 = ie1 - is1 + 1
	e1 = ie1 - is1 + 2
	e2 = ie2 - is2 + 2
	e3 = ie3 - is3 + 2
	x0 = X

	r.Randlc(&x0, ai)
	for i3 = 1; i3 < e3; i3++ {
		x1 = x0
		for i2 = 1; i2 < e2; i2++ {
			xx = x1
			r.Vrandlc(d1, &xx, A, &(z[i3][i2][1]))
			r.Randlc(&x1, a1)
		}
		r.Randlc(&x0, a2)
	}

	for i = 0; i < MM; i++ {
		ten[1][i] = 0.0
		j1[1][i] = 0
		j2[1][i] = 0
		j3[1][i] = 0
		ten[0][i] = 1.0
		j1[0][i] = 0
		j2[0][i] = 0
		j3[0][i] = 0
	}
	for i3 = 1; i3 < n3-1; i3++ {
		for i2 = 1; i2 < n2-1; i2++ {
			for i1 = 1; i1 < n1-1; i1++ {
				if z[i3][i2][i1] > ten[1][0] {
					ten[1][0] = z[i3][i2][i1]
					j1[1][0] = float64(i1)
					j2[1][0] = float64(i2)
					j3[1][0] = float64(i3)
					bubble(ten, j1, j2, j3, MM, 1)
				}
				if z[i3][i2][i1] < ten[0][0] {
					ten[0][0] = z[i3][i2][i1]
					j1[0][0] = float64(i1)
					j2[0][0] = float64(i2)
					j3[0][0] = float64(i3)
					bubble(ten, j1, j2, j3, MM, 0)
				}
			}
		}
	}

	i1 = MM - 1
	i0 = MM - 1
	for i = MM - 1; i >= 0; i-- {
		best = 0.0
		if best < ten[1][i1] {
			jg[1][i][0] = 0
			jg[1][i][1] = is1 - 2 + int(j1[1][i1])
			jg[1][i][2] = is2 - 2 + int(j2[1][i1])
			jg[1][i][3] = is3 - 2 + int(j3[1][i1])
			i1 = i1 - 1
		} else {
			jg[1][i][0] = 0
			jg[1][i][1] = 0
			jg[1][i][2] = 0
			jg[1][i][3] = 0
		}
		best = 1.0
		if best > ten[0][i0] {
			jg[0][i][0] = 0
			jg[0][i][1] = is1 - 2 + int(j1[0][i0])
			jg[0][i][2] = is2 - 2 + int(j2[0][i0])
			jg[0][i][3] = is3 - 2 + int(j3[0][i0])
			i0 = i0 - 1
		} else {
			jg[0][i][0] = 0
			jg[0][i][1] = 0
			jg[0][i][2] = 0
			jg[0][i][3] = 0
		}
	}
	m1 = 0
	m0 = 0

	for i3 = 0; i3 < n3; i3++ {
		for i2 = 0; i2 < n2; i2++ {
			for i1 = 0; i1 < n1; i1++ {
				z[i3][i2][i1] = 0.0
			}
		}
	}
	for i = MM - 1; i >= m0; i-- {
		z[jg[0][i][3]][jg[0][i][2]][jg[0][i][1]] = -1.0
	}
	for i = MM - 1; i >= m1; i-- {
		z[jg[1][i][3]][jg[1][i][2]][jg[1][i][1]] = +1.0
	}

	go comm3(z, n1, n2, n3, k)
}

func norm2u3(rr [][][]float64, n1, n2, n3 int, rnm2, rnmu float64, nx, ny, nz int) {

	var r [][][]float64
	r[n2][n1] = rr[n2][n1]

	var rnmu_local, s float64
	var a float64
	var i3, i2, i1 int
	var dn float64

	dn = 1.0 * float64(nx*ny*nz)

	s = 0.0
	rnmu_local = 0.0

	for i3 = 1; i3 < n3-1; i3++ {
		for i2 = 1; i2 < n2-1; i2++ {
			for i1 = 1; i1 < n1-1; i1++ {
				s = s + r[i3][i2][i1]*r[i3][i2][i1]
				a = math.Abs(r[i3][i2][i1])
				if a > rnmu_local {
					rnmu_local = a
				}
			}
		}
	}

	rnmu = rnmu_local
	rnm2 = math.Sqrt(s / dn)

}

func mg3P(debug_vec []int, lt int, u, v, r [][][]float64, a, c []float64, n1, n2, n3, k int) {
	var j int

	for k := lt; k >= lb+1; k-- {
		j = k - 1
		rprj3(&r[ir[k]], m1[k], m2[k], m3[k], &r[ir[j]], m1[j], m2[j], m3[j], k)
	}

	k = lb

	zero3(&u[ir[k]], m1[k], m2[k], m3[k], c, k)
	psinv(&r[ir[k]], &u[ir[k]], m1[k], m2[k], m3[k], c, k)

	for k := lb + 1; k <= lt-1; k++ {
		j = k - 1
		zero3(&u[ir[k]], m1[k], m2[k], m3[k])
		interp(&u[ir[j]], m1[j], m2[j], m3[j], &u[ir[k]], m1[k], m2[k], m3[k], k)
		resid(&u[ir[k]], &r[ir[k]], &r[ir[k]], m1[k], m2[k], m3[k], a, k)
		psinv(&r[ir[k]], &u[ir[k]], m1[k], m2[k], m3[k], c, k)
	}

	j = lt - 1
	k = lt

	interp(&u[ir[j]], m1[j], m2[j], m3[j], u, n1, n2, n3, k)
	resid(debug_vec, u, v, r, n1, n2, n3, a, k)
	psinv(debug_vec, r, u, n1, n2, n3, c, k)
}

func rprj3(debug_vec []int, rr [][][]float64, m1k, m2k, m3k int, ss [][][]float64, m1j, m2j, m3j, k int) {

	var j3, j2, j1, i3, i2, i1, d1, d2, d3 int
	var x1, y1 []float64
	var x2, y2 float64
	var r, s [][][]float64
	r[m2k][m1k] = rr[m2k][m1k]
	s[m2j][m1j] = ss[m2j][m1j]

	if m1k == 3 {
		d1 = 2
	} else {
		d1 = 1
	}

	if m2k == 3 {
		d2 = 2
	} else {
		d2 = 1
	}

	if m3k == 3 {
		d3 = 2
	} else {
		d3 = 1
	}

	for j3 = 1; j3 < m3j-1; j3++ {
		i3 := 2*j3 - d3
		for j2 = 1; j2 < m2j-1; j2++ {
			i2 = 2*j2 - d2
			for j1 = 1; j1 < m1j; j1++ {
				i1 = 2*j1 - d1
				x1[i1] = r[i3+1][i2][i1] + r[i3+1][i2+2][i1] + r[i3][i2+1][i1] + r[i3+2][i2+1][i1]
				y1[i1] = r[i3][i2][i1] + r[i3+2][i2][i1] + r[i3][i2+2][i1] + r[i3+2][i2+2][i1]
			}
			for j1 = 1; j1 < m1j-1; j1++ {
				i1 = 2*j1 - d1
				y2 = r[i3][i2][i1+1] + r[i3+2][i2][i1+1] + r[i3][i2+2][i1+1] + r[i3+2][i2+2][i1+1]

				x2 = r[i3+1][i2][i1+1] + r[i3+1][i2+2][i1+1] + r[i3][i2+1][i1+1] + r[i3+2][i2+1][i1+1]

				s[j3][j2][j1] = 0.5*r[i3+1][i2+1][i1+1] + 0.25*(r[i3+1][i2+1][i1]+r[i3+1][i2+1][i1+2]+x2) + 0.125*(x1[i1]+x1[i1+2]+y2) + 0.0625*(y1[i1]+y1[i1+2])
			}
		}
	}

	j := k - 1
	comm3(s, m1j, m2j, m3j, j)

	if debug_vec[0] >= 1 {
		rep_nrm(s, m1j, m2j, m3j, "rprj3", k-1)
	}

	if debug_vec[4] >= k {
		showall(s, m1j, m2j, m3j)
	}

}

func comm3(uu [][][]float64, n1, n2, n3, kk int) {

	var u [][][]float64
	u[n2][n1] = uu[n2][n1]
	var i1, i2, i3 int

	for i3 = 1; i3 < n3-1; i3++ {

		for i2 = 1; i2 < n2-1; i2++ {
			u[i3][i2][0] = u[i3][i2][n1-2]
			u[i3][i2][n1-1] = u[i3][i2][1]
		}

		for i1 = 0; i1 < n1; i1++ {
			u[i3][0][i1] = u[i3][n2-2][i1]
			u[i3][n2-1][i1] = u[i3][1][i1]
		}
	}

	for i2 = 0; i2 < n2; i2++ {
		for i1 = 0; i1 < n1; i1++ {
			u[0][i2][i1] = u[n3-2][i2][i1]
			u[n3-1][i2][i1] = u[1][i2][i1]
		}
	}
}

func power(a float64, n int) float64 {
	var aj, rdummy, power float64
	var nj int

	power = 1.0
	nj = n
	aj = a

	for nj != 0 {
		if (nj % 2) == 1 {
			rdummy = r.Randlc(&power, aj)
		}
		rdummy = r.Randlc(&aj, aj)
		nj = nj / 2
	}

	return power
}

func bubble(ten [][]float64, j1, j2, j3 [][]int, m, ind int) {
	var temp float64
	var j_temp int

	if ind == 1 {
		for i := 0; i < m-1; i++ {
			if ten[ind][i] > ten[ind][i+1] {
				temp = ten[ind][i+1]
				ten[ind][i+1] = ten[ind][i]
				ten[ind][i] = temp

				j_temp = j1[ind][i+1]
				j1[ind][i+1] = j1[ind][i]
				j1[ind][i] = j_temp

				j_temp = j2[ind][i+1]
				j2[ind][i+1] = j2[ind][i]
				j2[ind][i] = j_temp

				j_temp = j3[ind][i+1]
				j3[ind][i+1] = j3[ind][i]
				j3[ind][i] = j_temp
			} else {
				return
			}
		}
	} else {
		for i := 0; i < m-1; i++ {
			if ten[ind][i] < ten[ind][i+1] {
				temp = ten[ind][i+1]
				ten[ind][i+1] = ten[ind][i]
				ten[ind][i] = temp

				j_temp = j1[ind][i+1]
				j1[ind][i+1] = j1[ind][i]
				j1[ind][i] = j_temp

				j_temp = j2[ind][i+1]
				j2[ind][i+1] = j2[ind][i]
				j2[ind][i] = j_temp

				j_temp = j3[ind][i+1]
				j3[ind][i+1] = j3[ind][i]
				j3[ind][i] = j_temp
			} else {
				return
			}
		}
	}
}

func resid(debug_vec []int, uu, vv, rr [][][]float64, n1, n2, n3 int, a []float64, k int) {

	var u, v, r [][][]float64

	u[n2][n1] = uu[n2][n1]
	r[n2][n1] = rr[n2][n1]
	v[n2][n1] = vv[n2][n1]

	var i3, i2, i1 int
	var u1, u2 []float64

	for i3 = 1; i3 < n3-1; i3++ {
		for i2 = 1; i2 < n2-1; i2++ {
			for i2 = 0; i1 < n1; i1++ {
				u1[i1] = u[i3][i2-1][i1] + u[i3][i2+1][i1] + u[i3-1][i2][i1] + u[i3+1][i2][i1]
				u2[i1] = u[i3-1][i2-1][i1] + u[i3-1][i2+1][i1] + u[i3+1][i2-1][i1] + u[i3+1][i2+1][i1]
			}
			for i1 = 1; i1 < n1-1; i1++ {
				r[i3][i2][i1] = v[i3][i2][i1] - a[0]*u[i3][i2][i1] - a[2]*(u2[i1]+u1[i1-1]+u1[i1+1]) - a[3]*(u2[i1-1]+u2[i1+1])
			}
		}
	}

	comm3(r, n1, n2, n3, k)

	if debug_vec[0] >= 1 {
		rep_nrm(r, n1, n2, n3, "resid", k)
	}

	if debug_vec[2] >= k {
		showall(r, n1, n2, n3)
	}
}

func showall(zz [][][]float64, n1, n2, n3 int) {

	var z [][][]float64
	z[n2][n1] = zz[n2][n1]
	var i1, i2, i3, m1, m2, m3 int

	m1 = int(math.Min(float64(n1), float64(18)))
	m2 = int(math.Min(float64(n2), float64(14)))
	m3 = int(math.Min(float64(n3), float64(18)))

	for i3 = 0; i3 < m3; i3++ {
		for i2 = 0; i2 < m2; i2++ {
			for i1 = 0; i1 < m1; i1++ {
				fmt.Printf("%6.3f\n", z[i3][i2][i1])
			}
		}
		fmt.Println("- - - - - - -")
	}
}

func rep_nrm(uu [][][]float64, n1, n2, n3 int, title string, kk int) {
	var rnm2, rnmu float64
	norm2u3(uu, n1, n2, n3, rnm2, rnmu, nx[kk], ny[kk], nz[kk])
	fmt.Printf("Level%2d in %8s: norms =%21.14e%21.14e\n", kk, title, rnm2, rnmu)
}

func psinv(debug_vec []int, rr, uu [][][]float64, n1, n2, n3 int, c []float64, k int) {

	var r, u [][][]float64
	r[n2][n1] = rr[n2][n1]
	u[n2][n1] = uu[n2][n1]

	var i1, i2, i3 int
	var r1, r2 []float64

	for i3 = 1; i3 < n3-1; i3++ {
		for i2 = 1; i2 < n2-1; i2++ {
			for i1 = 0; i1 < n1; i1++ {
				r1[i1] = r[i3][i2-1][i1] + r[i3][i2+1][i1] + r[i3-1][i2][i1] + r[i3+1][i2][i1]
				r2[i1] = r[i3-1][i2-1][i1] + r[i3-1][i2+1][i1] + r[i3+1][i2-1][i1] + r[i3+1][i2+1][i1]
			}
			for i1 = 1; i1 < n1-1; i1++ {
				u[i3][i2][i1] = u[i3][i2][i1] + c[0]*r[i3][i2][i1] + c[1]*(r[i3][i2][i1-1]+r[i3][i2][i1+1]+r1[i1]) + c[2]*(r2[i1]+r1[i1-1]+r1[i1+1])
			}
		}
	}

	comm3(u, n1, n2, n3, k)

	if debug_vec[0] >= 1 {
		rep_nrm(u, n1, n2, n3, "psinv", k)
	}

	if debug_vec[3] >= k {
		showall(u, n1, n2, n3)
	}
}

func interp(debug_vec []int, zz [][][]float64, mm1, mm2, mm3 int, uu [][][]float64, n1, n2, n3, k int) {

	var z, u [][][]float64
	z[mm2][mm1] = zz[mm2][mm1]
	u[n2][n1] = uu[n2][n1]

	var i1, i2, i3, d1, d2, d3, t1, t2, t3 int
	var z1, z2, z3 []float64

	if n1 != 3 && n2 != 3 && n3 != 3 {
		for i3 = 0; i3 < mm3-1; i3++ {
			for i2 = 0; i2 < mm2-1; i2++ {
				for i1 = 0; i1 < mm1; i1++ {
					z1[i1] = z[i3][i2+1][i1] + z[i3][i2][i1]
					z2[i1] = z[i3+1][i2][i1] + z[i3][i2][i1]
					z3[i1] = z[i3+1][i2+1][i1] + z[i3+1][i2][i1] + z1[i1]
				}
				for i1 = 0; i1 < mm1-1; i1++ {
					u[2*i3][2*i2][2*i1] = u[2*i3][2*i2][2*i1] + z[i3][i2][i1]
					u[2*i3][2*i2][2*i1+1] = u[2*i3][2*i2][2*i1+1] + 0.5*(z[i3][i2][i1+1]+z[i3][i2][i1])
				}
				for i1 = 0; i1 < mm1-1; i1++ {
					u[2*i3][2*i2+1][2*i1] = u[2*i3][2*i2+1][2*i1] + 0.5*z1[i1]
					u[2*i3][2*i2+1][2*i1+1] = u[2*i3][2*i2+1][2*i1+1] + 0.25*(z1[i1]+z1[i1+1])
				}
				for i1 = 0; i1 < mm1-1; i1++ {
					u[2*i3+1][2*i2][2*i1] = u[2*i3+1][2*i2][2*i1] + 0.5*z2[i1]
					u[2*i3+1][2*i2][2*i1+1] = u[2*i3+1][2*i2][2*i1+1] + 0.25*(z2[i1]+z2[i1+1])
				}
				for i1 = 0; i1 < mm1-1; i1++ {
					u[2*i3+1][2*i2+1][2*i1] = u[2*i3+1][2*i2+1][2*i1] + 0.25*z3[i1]
					u[2*i3+1][2*i2+1][2*i1+1] = u[2*i3+1][2*i2+1][2*i1+1] + 0.125*(z3[i1]+z3[i1+1])
				}
			}
		}
	} else {
		if n1 == 3 {
			d1 = 2
			t1 = 1
		} else {
			d1 = 1
			t1 = 0
		}
		if n2 == 3 {
			d2 = 2
			t2 = 1
		} else {
			d2 = 1
			t2 = 0
		}
		if n3 == 3 {
			d3 = 2
			t3 = 1
		} else {
			d3 = 1
			t3 = 0
		}

		for i3 = d3; i3 <= mm3-1; i3++ {
			for i2 = d2; i2 <= mm2-1; i2++ {
				for i1 = d1; i1 <= mm1-1; i1++ {
					u[2*i3-d3-1][2*i2-d2-1][2*i1-d1-1] = u[2*i3-d3-1][2*i2-d2-1][2*i1-d1-1] + z[i3-1][i2-1][i1-1]
				}
				for i1 = 1; i1 <= mm1-1; i1++ {
					u[2*i3-d3-1][2*i2-d2-1][2*i1-t1-1] = u[2*i3-d3-1][2*i2-d2-1][2*i1-t1-1] + 0.5*(z[i3-1][i2-1][i1]+z[i3-1][i2-1][i1-1])
				}
			}
			for i2 = 1; i2 <= mm2-1; i2++ {
				for i1 = d1; i1 <= mm1-1; i1++ {
					u[2*i3-d3-1][2*i2-t2-1][2*i1-d1-1] = u[2*i3-d3-1][2*i2-t2-1][2*i1-d1-1] + 0.5*(z[i3-1][i2][i1-1]+z[i3-1][i2-1][i1-1])
				}
				for i1 = 1; i1 <= mm1-1; i1++ {
					u[2*i3-d3-1][2*i2-t2-1][2*i1-t1-1] = u[2*i3-d3-1][2*i2-t2-1][2*i1-t1-1] + 0.25*(z[i3-1][i2][i1]+z[i3-1][i2-1][i1]+z[i3-1][i2][i1-1]+z[i3-1][i2-1][i1-1])
				}
			}
		}

		for i3 = 1; i3 <= mm3-1; i3++ {
			for i2 = d2; i2 <= mm2-1; i2++ {
				for i1 = d1; i1 <= mm1-1; i1++ {
					u[2*i3-t3-1][2*i2-d2-1][2*i1-d1-1] = u[2*i3-t3-1][2*i2-d2-1][2*i1-d1-1] + 0.5*(z[i3][i2-1][i1-1]+z[i3-1][i2-1][i1-1])
				}
				for i1 = 1; i1 <= mm1-1; i1++ {
					u[2*i3-t3-1][2*i2-d2-1][2*i1-t1-1] = u[2*i3-t3-1][2*i2-d2-1][2*i1-t1-1] + 0.25*(z[i3][i2-1][i1]+z[i3][i2-1][i1-1]+z[i3-1][i2-1][i1]+z[i3-1][i2-1][i1-1])
				}
			}
			for i2 = 1; i2 <= mm2-1; i2++ {
				for i1 = d1; i1 <= mm1-1; i1++ {
					u[2*i3-t3-1][2*i2-t2-1][2*i1-d1-1] = u[2*i3-t3-1][2*i2-t2-1][2*i1-d1-1] + 0.25*(z[i3][i2][i1-1]+z[i3][i2-1][i1-1]+z[i3-1][i2][i1-1]+z[i3-1][i2-1][i1-1])
				}
				for i1 = 1; i1 <= mm1-1; i1++ {
					u[2*i3-t3-1][2*i2-t2-1][2*i1-t1-1] = u[2*i3-t3-1][2*i2-t2-1][2*i1-t1-1] + 0.125*(z[i3][i2][i1]+z[i3][i2-1][i1]+z[i3][i2][i1-1]+z[i3][i2-1][i1-1]+z[i3-1][i2][i1]+z[i3-1][i2-1][i1]+z[i3-1][i2][i1-1]+z[i3-1][i2-1][i1-1])
				}
			}
		}
	}

	if debug_vec[0] >= 1 {
		rep_nrm(z, mm1, mm2, mm3, "z: inter", k-1)
		rep_nrm(u, n1, n2, n3, "u: inter", k)
	}
	if debug_vec[5] >= k {
		showall(z, mm1, mm2, mm3)
		showall(u, n1, n2, n3)
	}
}
