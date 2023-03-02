SHELL=/bin/sh
BENCHMARK=ep
BENCHMARKU=EP

include ../config/make.def

OBJS = ep.o ${COMMON}/c_print_results.o ${COMMON}/c_${RAND}.o \
	${COMMON}/c_timers.o ${COMMON}/c_wtime.o

include ../sys/make.common

${PROGRAM}: config ${OBJS}
	${GOLINK} ${GOLINKFLAGS} ${PROGRAM} ${OBJS} ${GO_LIB}

ep.o:	ep.go npbparams.go
	${GOCOMPILE} ep.go

clean:		
	go clean
	rm -f *.o *~
