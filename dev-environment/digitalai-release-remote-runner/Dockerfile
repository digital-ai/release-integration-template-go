FROM xebialabsunsupported/release-runner:24.1.0-411.1033

USER root

RUN apt update && apt install -y curl jq

COPY spin-remote-runner.sh /spin-remote-runner.sh
RUN chmod +x /spin-remote-runner.sh

ENTRYPOINT ["/spin-remote-runner.sh"]
