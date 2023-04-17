.POSIX:
.SUFFIXES:

PREFIX=/usr/local
BINDIR=bin
MANDIR=share/man
GO=go
RM = rm
INSTALL = install
SCDOC = scdoc
GOBUILD_OPTS=-trimpath

all: build doc

build: # Builds an application binary.
	$(GO) build $(GOBUILD_OPTS)

doc: # Builds the manpage.
	$(SCDOC) <doc/acopw.1.scd >doc/acopw.1

install: # Installs the release binary.
	$(INSTALL) -d \
		$(DESTDIR)$(PREFIX)/$(BINDIR)/ \
		$(DESTDIR)$(PREFIX)/$(MANDIR)/man1/
	$(INSTALL) -pm 0755 acopw $(DESTDIR)$(PREFIX)/$(BINDIR)/
	$(INSTALL) -pm 0644 doc/acopw.1 $(DESTDIR)$(PREFIX)/$(MANDIR)/man1/

clean: # Cleans cache files from tests and deletes any build output.
	$(RM) -f acopw doc/acopw.1

.PHONY: all build doc install clean
