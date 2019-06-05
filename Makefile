doit:
	@echo $(GOPATH)
	@echo $(PATH)
	PASSWORD ?= $(shell bash -c 'read -s -p "Password: " pwd; echo $$pwd')
	@echo $(PASSWORD)

