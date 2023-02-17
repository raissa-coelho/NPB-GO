SHELL=/bin/sh
CLASS=S
SFILE=config/suite.def

default: header
        @ $(SHELL) sys/print_instructions

EP: ep
ep: header
        cd EP; $(MAKE) CLASS=$(CLASS)

cleanall:
        go clean
        -rm -r bin/*

header:
        @ $(SHELL) sys/print_header

help:
        @ $(SHELL) sys/print_instructions
