# NPB-GO
Implementation of the NAS Parallel Benchmark in GO language for TECVII class.

## Compilation and Execution

make init --> Use only if there isn't go.mod

make <_benchmark_> CLASS=<_benchmark_class_>

<_benchmark_> = EP or ep

<_benchmark_class_> = S, W, A, B, C, D, E or F

---------------------------------------------

suite.def contains a list of benchmarks and benchmark classes to be used as parameters in the command line.

To compile and run multiple tests:

make suite

## Contributors
- Raissa Nunes Coelho, Computer engineering (student), Universidade Federal de Pelotas(UFPEL)
- Bianca Nunes Coelho, Computer engineering (student), Universidade Federal de Pelotas(UFPEL)

## Benchmarks
- EP
- IS

## Content
|Directory |Files |
| :---|---:|
|config | make.def|
|sys | <ul><li>make.common</li><li>Makefile</li><li>print_header</li><li>print_instructions</li><li>setparams.go</li><li>npbparams.txt</li></ul>|
|EP | <ul><li>Makefile</li><li>EP.go</li><li>npbparams.txt</li></ul> |
|IS | <ul><li>Makefile</li><li>IS.go</li><li>npbparams.txt</li></ul> |
|common | <ul><li>c_randdp.go</li><li>c_print_results.go</li><li>c_file.go</li></ul>  |
|bin | .txt files |

## Notes on the implementation
Implementation was based on:

  -NPB3.4-OMP
      
      NAS Parallel Benchmark Team
  
  -NPB-CPP
  
      Parallel Applications Modelling Group (GMAP) at PUCRS - Brazil.
