all: clean
	@./utils/format.py

clean:
	-@rm -f */*~
