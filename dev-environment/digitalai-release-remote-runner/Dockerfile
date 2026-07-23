FROM xebialabs/release-runner:26.1.2-503.130

USER root

RUN apt update && apt install -y curl jq

COPY spin-remote-runner.sh /spin-remote-runner.sh
# Strip Windows CRLF line endings (Git may add them on clone) so the script runs in the Linux container, then make it executable
RUN sed -i 's/\r$//' /spin-remote-runner.sh && chmod +x /spin-remote-runner.sh

ENTRYPOINT ["/spin-remote-runner.sh"]
