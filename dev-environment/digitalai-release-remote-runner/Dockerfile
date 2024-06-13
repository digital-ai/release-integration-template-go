FROM xebialabs/release-remote-runner:23.3.7

USER root

RUN apt update && apt install -y curl jq

COPY spin-remote-runner.sh /spin-remote-runner.sh
RUN chmod +x /spin-remote-runner.sh

USER daiuser

ENTRYPOINT ["/spin-remote-runner.sh"]
