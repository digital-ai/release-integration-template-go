FROM xebialabs/release-remote-runner:24.1.2

RUN apk add --no-cache curl jq

COPY spin-remote-runner.sh /spin-remote-runner.sh
RUN chmod +x /spin-remote-runner.sh

ENTRYPOINT ["/spin-remote-runner.sh"]
