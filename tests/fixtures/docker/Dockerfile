# Minimal fixture image used by test-callers.yml to exercise docker-build.yml
# (build on PRs; build + push + attest-sbom self-verify on push to main).
# Alpine gives Syft a non-empty package set so the SPDX SBOM/attestation is meaningful.
FROM alpine:3.20
RUN echo "gh-pipelines docker build fixture" > /etc/gh-pipelines-fixture
CMD ["true"]
