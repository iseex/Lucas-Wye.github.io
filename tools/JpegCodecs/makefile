CXX := g++
main: main.cpp JpegDecoder.o
	$(CXX) -I src/  main.cpp JpegDecoder.o -o main

JpegDecoder.o : src/JpegDecoder.cpp
	$(CXX) -c src/JpegDecoder.cpp
	
clean:
	rm *.o main
