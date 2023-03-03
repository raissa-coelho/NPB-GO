SHELL=/bin/sh
CLASS=S
SFILE=config/suite.def

default: header
	@ $(SHELL) sys/print_instructions

EP: ep
ep: header
	cd EP; $(MAKE) CLASS=$(CLASS)

init: 
	go mod init NPB-GO
	go mod tidy -v -x

clean:
	rm -f core

header:
	@ $(SHELL) sys/print_header

help:
	@ $(SHELL) sys/print_instructions
