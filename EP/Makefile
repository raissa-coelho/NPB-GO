SHELL=/bin/sh
BENCHMARK=EP
BENCHMARKU=EP

include ../config/make.def

OBJS = EP.o ${COMMON}/c_print_results.o ${COMMON}/c_${RAND}.o

include ../sys/make.common

${PROGRAM}: config ${OBJS}
	${GOLINK} ${GOLINKFLAGS} ${PROGRAM} ${OBJS} ${GO_LIB}

EP.o:	EP.go
	$(UGO) build -o EP EP.go

clean:		
	go clean
	rm 
