include $(GOROOT)/src/Make.$(GOARCH)

TARG=alloy-d/go140
GOFILES=api.go\
		error.go\

include $(GOROOT)/src/Make.pkg

