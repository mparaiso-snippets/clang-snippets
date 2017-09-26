%.o: %.c
	gcc -Wall -c $<

test_stack: stack.o test_stack.o
	gcc $^ -o $@
compute: stack.o compute.o

clean: *.o *.exe
	rm $^