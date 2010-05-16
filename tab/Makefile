tabs := $(patsubst %.gnumeric,%.tex,$(wildcard *.gnumeric))

%.tex: %.gnumeric
	ssconvert -T Gnumeric_html:latex <$ $@

all:	${tabs}

clean:  
	rm -f *.tex
