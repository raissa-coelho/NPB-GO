SHELL=/bin/sh
CLASS=S
SFILE=config/suite.def

default: header
	@ $(SHELL) sys/print_instructions

IS: is
is: header
	cd IS; $(MAKE) CLASS=$(CLASS)
	
EP: ep
ep: header
	cd EP; $(MAKE) CLASS=$(CLASS)

init: 
	go mod init NPB-GO

# Awk script courtesy cmg@cray.com, modified by Haoqiang Jin
suite:
	@ awk -f sys/suite.awk SMAKE=$(MAKE) $(SFILE) | $(SHELL)

clean:
	rm -f core

header:
	@ $(SHELL) sys/print_header

help:
	@ $(SHELL) sys/print_instructions
