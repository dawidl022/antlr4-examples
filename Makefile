parser/antlr-4.13.0-complete.jar:
	mkdir -p parser && \
	curl https://www.antlr.org/download/antlr-4.13.0-complete.jar \
	-o parser/antlr-4.13.0-complete.jar

parser: parser/antlr-4.13.0-complete.jar
	java -Xmx500M -cp "./parser/antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool \
	-Dlanguage=Go -no-visitor -package parser -o parser grammar/*.g4
.PHONY: parser

clean:
	rm -rf parser/grammar
.PHONY: clean
