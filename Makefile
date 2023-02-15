SHELL=/bin/sh
CLASS=S
SFILE=config/suite.def

default: header
        @ $(SHELL) sys/print_instructions

EP: ep
ep: header
        cd EP; $(MAKE) CLASS=$(CLASS)
