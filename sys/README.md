# sys directory
This directory contains utilities and files used by the build process. You should not need to change anything in this directory. 

# Original Files

setparams.go:<br>
Source for the setparams program. This program is used internally in the build process to create the file "npbparams" for each benchmark. npbparams contains GO parameters to build a benchmark for a specific class. The setparams program is never run directly by a user. Its invocation syntax is 

            "setparams benchmark-name class"

make.common<br>
        A makefile segment that is included in each individual benchmark program makefile. It sets up some standard macros (COMPILE, etc) and makes sure everything is configured correctly (npbparams)

Makefile <br>
        Builds  setparams

README<br>
        This file. 

Created files
-------------

setparams<br>
	See descriptions above
