include ../config/make.def

# Note that COMPILE is also defined in make.common and should
# be the same. We can't include make.common because it has a lot
# of other garbage. 
GOCOMPILE = $(GCCGO) build $(GO_INC) $(GOFLAGS)

all: setparams 

# setparams creates an npbparam.h file for each benchmark 
# configuration. npbparams.h also contains info about how a benchmark
# was compiled and linked

setparams: setparams.go ../config/make.def
	   $(UGO) mod init NPB-GO
	   $(UGO) build -o setparams setparams.go

clean: 
	go clean
	rm 
