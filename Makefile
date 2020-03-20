all: cmd



cmd:
	go  build  -a    -o  gvp  

clean:
	rm  -rf  gvp

doc:
	swag  init