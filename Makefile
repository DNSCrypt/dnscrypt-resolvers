all: clean
	@./utils/format.py
	./utils/subset.py v3/opennic.subset v3/public-resolvers.md > v3/a && mv -f v3/a v3/opennic.md
	./utils/subset.py v3/parental-control.subset v3/public-resolvers.md > v3/a && mv -f v3/a v3/parental-control.md
	@./utils/format.py

clean:
	-@rm -f */*~
