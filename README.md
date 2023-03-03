# NPB-GO
Implementation of the NAS Parallel Benchmark in GO language for TECVII class.

## Compilation and Execution

make init --> Use only if there isn't go.mod

make <_benchmark_> CLASS=<_benchmark_type_>

<_benchmark_> = EP or ep

<_benchmark_type_> = S, W, A, B, C, D, E or F

<<<<<<< HEAD
---------------------------------------------
=======
--------------------------------------
>>>>>>> c9d37aada60d3d65da6a37fbe6f300c6e68a328a

suite.def contains a list of benchmarks and benchmark types to be used as parameters in the command line.

To compile and run multiple tests:

make suite
<<<<<<< HEAD
=======

>>>>>>> c9d37aada60d3d65da6a37fbe6f300c6e68a328a
## Contributors
- Raissa Nunes Coelho, Computer engineering (student), Universidade Federal de Pelotas(UFPEL)
- Bianca Nunes Coelho, Computer engineering (student), Universidade Federal de Pelotas(UFPEL)

## Benchmarks
- EP

## Content
|Directory |Files |
| :---|---:|
|config | make.def|
|sys | <ul><li>make.common</li><li>Makefile</li><li>print_header</li><li>print_instructions</li><li>setparams.go</li><li>npbparams.txt</li></ul>|
|EP | <ul><li>Makefile</li><li>EP.go</li><li>npbparams.txt</li></ul> |
|common | <ul><li>c_randdp.go</li><li>c_print_results.go.go</li></ul>  |
|bin | vazio |

## Notes on the implementation
Implementation was based on:

  -NPB3.4-OMP
      
      NAS Parallel Benchmark Team
  
  -NPB-CPP
  
      Parallel Applications Modelling Group (GMAP) at PUCRS - Brazil.
