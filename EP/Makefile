SHELL=/bin/sh
BENCHMARK=ep
BENCHMARKU=EP

include ../config/make.def

OBJS = TODO

include ../sys/make.common

${PROGRAM}: config ${OBJS}
	${GOLINK} ${GOLINKFLAGS} ${PROGRAM} ${OBJS} ${GO_LIB}

ep.o:		TODO

clean:
	- rm -f *.o *~
