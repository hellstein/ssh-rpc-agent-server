GITBOOK = $(CURDIR)/gitbook
DOCS = $(CURDIR)/docs
IMAGE_ENV = $(CURDIR)/image
DF = $(IMAGE_ENV)/Dockerfile
DEPLOYMENT = $(CURDIR)/deployment
OWNER = hellstein
REPO = ssh-rpc-agent

.PHONY: mk-book clean-book
mk-book: $(GITBOOK)
	gitbook build $(GITBOOK) $(DOCS)

clean-book:
	rm -rf $(DOCS)/*

.PHONY: mk-image clean-image
mk-image:
	CGO_ENABLED=0 go build -o image/ssh-rpc-agent
	docker run --rm --privileged multiarch/qemu-user-static:register --reset
	docker build -t $(OWNER)/$(REPO)-$(ARCH) -f $(DF)-$(ARCH) $(IMAGE_ENV) 

clean-image:
	rm image/ssh-rpc-agent
	docker rmi $(OWNER)/$(REPO)-$(ARCH)


.PHONY: mk-deployment clean-deployment
mk-deployment: $(DEPLOYMENT)
	sed -i s+VERSION=.*+VERSION=$(VERSION)+g $(DEPLOYMENT)/temp.env
	mkdir imageAPI 
	cp $(DEPLOYMENT)/docker-compose.yml $(DEPLOYMENT)/temp.env $(DEPLOYMENT)/Makefile imageAPI/
	zip -r $(REPO)-$(VERSION).zip imageAPI
	rm -rf imageAPI

clean-deployment: $(REPO)-$(VERSION).zip
	rm $(REPO)-$(VERSION).zip


.PHONY: pushtohub
pushtohub:
	docker tag $(OWNER)/$(REPO)-$(ARCH) $(OWNER)/$(REPO)-$(ARCH):$(TAG)
	docker login -u $(USER) -p $(PASS)
	docker push $(OWNER)/$(REPO)-$(ARCH):$(TAG)
