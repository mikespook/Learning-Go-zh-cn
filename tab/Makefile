tabs := $(patsubst %.gnumeric,%.tex,$(wildcard *.gnumeric))

%.tex: %.gnumeric
	ssconvert -T Gnumeric_html:latex $< $@ 2>/dev/null

all:	${tabs}

clean:  
