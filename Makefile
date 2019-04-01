GITBOOK = $(CURDIR)/gitbook
DOCS = $(CURDIR)/docs
IMAGE_ENV = $(CURDIR)/image
DF = $(IMAGE_ENV)/Dockerfile
DEPLOYMENT = $(CURDIR)/deployment
OWNER = hellstein
REPO = ssh-rpc-agent-server
VERSION = test

.PHONY: mk-book clean-book
mk-book: $(GITBOOK)
	gitbook build $(GITBOOK) $(DOCS)

clean-book:
	rm -rf $(DOCS)/*

.PHONY: mk-image clean-image
mk-image:
	go get -v github.com/gorilla/mux 
	go get -v golang.org/x/crypto/ssh
	go get -v github.com/gorilla/websocket
	CGO_ENABLED=0 go build -o image/$(REPO)
	docker run --rm --privileged multiarch/qemu-user-static:register --reset
	docker build -t $(OWNER)/$(REPO):$(VERSION) -f $(DF) $(IMAGE_ENV) 

clean-image:
	rm image/ssh-rpc-agent
	docker rmi $(OWNER)/$(REPO):$(VERSION)


.PHONY: mk-deployment clean-deployment
mk-deployment: $(DEPLOYMENT)
	sed -i s+VERSION=.*+VERSION=$(VERSION)+g $(DEPLOYMENT)/temp.env
	mkdir -p sra-server/
	cp $(DEPLOYMENT)/docker-compose.yml $(DEPLOYMENT)/temp.env $(DEPLOYMENT)/Makefile sra-server/
	zip -r $(REPO)-$(VERSION).zip sra-server
	rm -rf sra-server

clean-deployment: $(REPO)-$(VERSION).zip
	rm $(REPO)-$(VERSION).zip


.PHONY: pushtohub
pushtohub:
	docker login -u $(USER) -p $(PASS)
	docker push $(OWNER)/$(REPO):$(VERSION)
