GITBOOK = $(CURDIR)/gitbook
DOCS = $(CURDIR)/docs
IMAGE_ENV = $(CURDIR)/image
DF = $(IMAGE_ENV)/Dockerfile
DEPLOYMENT = $(CURDIR)/deployment
WSCLIENT = $(CURDIR)/wsclient
OWNER = hellstein
REPO = ssh-rpc-agent

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
	CGO_ENABLED=0 go build -o image/ssh-rpc-agent
	docker run --rm --privileged multiarch/qemu-user-static:register --reset
	docker build -t $(OWNER)/$(REPO)-$(ARCH) -f $(DF)-$(ARCH) $(IMAGE_ENV) 

clean-image:
	rm image/ssh-rpc-agent
	docker rmi $(OWNER)/$(REPO)-$(ARCH)


.PHONY: mk-deployment clean-deployment
mk-deployment: $(DEPLOYMENT) $(WSCLIENT)
	sed -i s+VERSION=.*+VERSION=$(VERSION)+g $(DEPLOYMENT)/temp.env
	mkdir -p agent/imageAPI agent/wsClient
	cp $(DEPLOYMENT)/docker-compose.yml $(DEPLOYMENT)/temp.env $(DEPLOYMENT)/Makefile agent/imageAPI/
	cp $(WSCLIENT)/client.js $(WSCLIENT)/msg.js $(WSCLIENT)/package.json $(WSCLIENT)/package-lock.json agent/wsClient/
	cp -r $(WSCLIENT)/example agent/wsClient/
	zip -r $(REPO)-$(VERSION).zip agent
	rm -rf agent

clean-deployment: $(REPO)-$(VERSION).zip
	rm $(REPO)-$(VERSION).zip


.PHONY: pushtohub
pushtohub:
	docker tag $(OWNER)/$(REPO)-$(ARCH) $(OWNER)/$(REPO)-$(ARCH):$(TAG)
	docker login -u $(USER) -p $(PASS)
	docker push $(OWNER)/$(REPO)-$(ARCH):$(TAG)
