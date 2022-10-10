all:
	go build -o ttt .

run: all
	./ttt

console:
	go run conmain.go ttt.go

clean:
	rm ttt
