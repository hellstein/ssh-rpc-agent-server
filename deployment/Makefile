COMP_DIR=$(CURDIR)/compose
ENV_TEMP=$(CURDIR)/temp.env
ENV=$(CURDIR)/.env
SDRUNDIR=$(CURDIR)/sshdata.run
.PHONY: config
config:
	cp -rf $(SSHDATA) $(SDRUNDIR)
	sudo chown -R root:root $(SDRUNDIR)
	sudo chmod -R 500 $(SDRUNDIR)/*
	cp $(ENV_TEMP) $(ENV)
	# add extra env variables
	echo "SSHRUNDATA=$(SDRUNDIR)" >> $(ENV)
	cat $(ENV)
	mkdir -p $(COMP_DIR)
	docker-compose -p $(NAME) -f $(CURDIR)/docker-compose.yml config | tee $(COMP_DIR)/$(NAME).yml
	rm $(ENV)

.PHONY: start
start: $(COMP_DIR)/$(NAME).yml
	docker-compose -p $(NAME) -f $(COMP_DIR)/$(NAME).yml up -d 

.PHONY: stop
stop: $(COMP_DIR)/$(NAME).yml
	docker-compose -p $(NAME) -f $(COMP_DIR)/$(NAME).yml down

.PHONY: restart
restart: $(COMP_DIR)/$(NAME).yml
	docker-compose -p $(NAME) -f $(COMP_DIR)/$(NAME).yml up -d --force-recreate

.PHONY: status
status: $(COMP_DIR)/$(NAME).yml
	docker-compose -p $(NAME) -f $(COMP_DIR)/$(NAME).yml ps
	docker-compose -p $(NAME) -f $(COMP_DIR)/$(NAME).yml logs

.PHONY: list
list: $(COMP_DIR)
	for compose in `ls $(COMP_DIR)`;do name=`echo $$compose|awk -F "." '{print $$1}'`;echo $$name;docker-compose -p $$name -f $(COMP_DIR)/$$compose ps;done

.PHONY: remove
remove: $(COMP_DIR)/$(NAME).yml
	sudo rm -rf $(SDRUNDIR)
	rm $(COMP_DIR)/$(NAME).yml
