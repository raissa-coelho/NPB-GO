SHELL=/bin/sh
BENCHMARK=IS
BENCHMARKU=IS

include ../config/make.def

OBJS = IS.o ${COMMON}/c_print_results.o ${COMMON}/c_${RAND}.o

include ../sys/make.common

${PROGRAM}: config ${OBJS}
	${GOLINK} ${GOLINKFLAGS} ${PROGRAM} ${OBJS} ${GO_LIB}

IS.o:	IS.go
	$(UGO) build -o IS IS.go

clean:		
	$(UGO) clean
	rm -f *.o *~
