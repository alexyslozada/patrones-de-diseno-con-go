# State

```mermaid
flowchart TB
  Initialized --> PendingApprove
  PendingApprove --> Accepted
  PendingApprove --> Rejected
  Rejected --> PendingApprove
  Rejected --> Finish
  Accepted --> Finish
```