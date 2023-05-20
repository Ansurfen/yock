exec({
    redirect = true
}, [[java -jar .\antlr-4.9.1-complete.jar -Dlanguage=Go -visitor -o parser ylang.g4]])