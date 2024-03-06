SUBDIRS := $(wildcard */)

.PHONY: clean
clean:
	@for dir in $(SUBDIRS); do \
		if [ -f $$dir/Makefile ]; then \
			print "$$dir ...\n" \
			$(MAKE) -C $$dir clean; \
		fi \
	done