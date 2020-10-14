build:
	go build -o bin/pengeluaranku
	./bin/pengeluaranku

serve:
	cd ./frontend && yarn start

.PHONY: build serve