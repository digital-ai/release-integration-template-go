FROM xebialabsunsupported/release-runner:23.3.0-412.113

USER root

RUN apt update && apt install -y curl jq

COPY spin-remote-runner.sh /spin-remote-runner.sh
RUN chmod +x /spin-remote-runner.sh

USER daiuser

ENTRYPOINT ["/spin-remote-runner.sh"]
