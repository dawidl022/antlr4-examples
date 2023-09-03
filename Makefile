grammar/antlr-4.13.0-complete.jar:
	mkdir -p parser && \
	curl https://www.antlr.org/download/antlr-4.13.0-complete.jar \
	-o grammar/antlr-4.13.0-complete.jar

parser: grammar/antlr-4.13.0-complete.jar grammar/*.g4
	go generate ./...

clean:
	rm -rf parser
.PHONY: clean
