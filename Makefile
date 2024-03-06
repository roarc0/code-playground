SUBDIRS := $(wildcard */)

.PHONY: clean
clean:
	@for dir in $(SUBDIRS); do \
		if [ -f $$dir/Makefile ]; then \
			echo "Cleaning $$dir ..." \
			$(MAKE) -C $$dir clean; \
		fi \
	done