#Golang rest api using Chi

##For swagger spect:

- Generate spect with docker: docker run --rm -it -v $(pwd):/swagger -w /swagger quay.io/goswagger/swagger generate spec --scan-models  -o /swagger/swagger.json 

- Validate spect with docker: docker run --rm -it -v $(pwd):/swagger -w /swagger quay.io/goswagger/swagger validate /swagger/swagger.json

- Serve spect with docker: docker run --rm -v $(pwd):/swagger -w /swagger quay.io/goswagger/swagger serve /swagger/swagger.json --no-open  


