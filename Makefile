run:
	go run main.go 

clean:
	rm -rf pb/*

gen:
  
	rm -rf pb/* & protoc -I=proto --go_out=. proto/*.proto 