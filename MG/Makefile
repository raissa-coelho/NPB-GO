SHELL=/bin/sh
BENCHMARK=IS
BENCHMARKU=IS

include ../config/make.def

OBJS = MG.o ${COMMON}/c_print_results.o ${COMMON}/c_${RAND}.o

include ../sys/make.common

${PROGRAM}: config ${OBJS}
	${GOLINK} ${GOLINKFLAGS} ${PROGRAM} ${OBJS} ${GO_LIB}

MG.o:	MG.go
	$(UGO) build -o MG MG.go

clean:		
	$(UGO) clean
	rm -f *.o *~
