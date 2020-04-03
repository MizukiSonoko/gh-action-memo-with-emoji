
APP=action-app
OPT=CGO_ENABLED=0 GOOS=linux

all:
	${OPT} go build -ldflags="-w -s" -o /app/${APP}

local:
	${OPT} go build -ldflags="-w -s" -o ${APP}

clean:
	-rm /app/${APP} ${APP}

