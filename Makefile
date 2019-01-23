# Variables --------------------------------------------------------------------
export GO111MODULE=on

GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD) fmt
BINARY_NAME=stopandsearch
FRONTEND=$(CURDIR)/web
VERSION=0.0.0

# Targets for app --------------------------------------------------------------
.PHONY: all
all: deps backend frontend;

# Targets for docker  ----------------------------------------------------------

.PHONY: docker
docker: clean ; $(info $(M) building docker image...)
	$Q docker build -t stopandsearch/$(BINARY_NAME):$(VERSION) .
# Targets for vue app ----------------------------------------------------------

.PHONY: frontend
frontend: ; $(info $(M) building web frontend ui...)	     
	$Q npm install --prefix $(FRONTEND) \
	   && npm rebuild node-sass \
	   && npm run build --prefix $(FRONTEND)

# Targets for go app -----------------------------------------------------------
.PHONY: backend												
backend: ; $(info $(M) building backend binary...)
	$(GOBUILD) cmd/$(BINARY_NAME)/$(BINARY_NAME).go

.PHONY: test
test:
	$(GOTEST) -v ./...

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	# $(GOCLEAN) #go clean is having a problem with modules currently.

.PHONY: format
 format:
	$(GOFMT) ./...

.PHONY: deps
deps:
	$(GOBUILD) -v ./...

.PHONY: upgrade
upgrade:
	$(GOGET) -u



