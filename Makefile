build: 
	@sh -c "./commands/build.sh"

download:
	@sh -c "./commands/download.sh $(track) $(exercise)"

submit:
	@sh -c "./commands/submit.sh $(track) $(exercise)"
