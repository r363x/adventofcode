CC=clang
CFLAGS=-Wall -I --std=gnu11
ODIR=obj

poly1:	poly1.c
	$(CC) $(CFLAGS) $@.c -o $@

poly2:	poly2.c
	$(CC) $(CFLAGS) $@.c -o $@

poly1-debug: poly1.c
	     $(CC) $(CFLAGS) poly1.c -o $@

poly2-debug: poly2.c
	     $(CC) $(CFLAGS) poly1.c -o $@
clean:
	rm -rf poly1 poly2 poly1.dSYM poly2.dSYM

.PHONY: clean

