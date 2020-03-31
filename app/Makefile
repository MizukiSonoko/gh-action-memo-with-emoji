
APP=gen-memo
OPT=CGO_ENABLED=0 GOOS=linux

all:
	${OPT} go build -ldflags="-w -s" -o /app/${APP}

clean:
	-rm /app/${APP}

