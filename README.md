# gh-pipelines-hosted-acceptance

Dedicated non-production **Synthetic consumer repository** for `nandohdc/gh-pipelines`.

## Boundary

- Non-production fixtures (`tests/fixtures/{go,node,python,terraform,helm}`), dummy metadata, and hosted `ubuntu-latest` execution only.
- No production cloud, registry, Kubernetes cluster, or Sonar credentials are stored or accessed here.

## Candidate Smoke Contract

- `.github/workflows/candidate-smoke.yml` registers the `workflow_dispatch` inputs (`candidate_sha`, `request_id`).
- `scripts/candidate_smoke.py` in `nandohdc/gh-pipelines` creates an ephemeral `smoke/candidate-<sha12>-<request_id>` branch, stages candidate workflows, nested composite actions, and fixtures, dispatches `candidate-smoke.yml` on that branch, requires a matching `candidate-smoke-result / <candidate_sha>` pass, and deletes the temporary branch on completion.
