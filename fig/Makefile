pdfs := $(patsubst %.svg,%.pdf,$(wildcard *.svg))

%.pdf: %.svg
	inkscape --export-pdf=$@ $<
	pdfcrop $@ output.pdf && mv output.pdf $@

all:	${pdfs}

clean:  
	rm -f *.pdf
