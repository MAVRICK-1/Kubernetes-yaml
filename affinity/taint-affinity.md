# Kubernetes Affinity and Anti-Affinity

Affinity and anti-affinity rules in Kubernetes allow you to control how pods are scheduled onto nodes in a cluster. They enable constraints for co-scheduling or avoiding scheduling of pods on specific nodes or with specific pods.

---

## Types of Affinity
1. **Node Affinity**: Controls which nodes a pod can be scheduled on based on node labels.
2. **Pod Affinity**: Controls scheduling of pods to nodes based on labels of other pods that are already running.
3. **Pod Anti-Affinity**: Prevents scheduling pods on nodes where specific pods are already running.

---

## Node Affinity

### Fields
- `requiredDuringSchedulingIgnoredDuringExecution`: Hard constraints that must be met for scheduling.
- `preferredDuringSchedulingIgnoredDuringExecution`: Soft constraints the scheduler will try to satisfy.

### Example: Node Affinity
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: disktype
            operator: In
            values:
            - ssd
      preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 50
        preference:
          matchExpressions:
          - key: region
            operator: In
            values:
            - us-west
  containers:
  - name: nginx
    image: nginx
```

### Supported Operators for Node Affinity
- **`In`**: Matches any of the specified values.
- **`NotIn`**: Does not match any of the specified values.
- **`Exists`**: Matches if the key exists, regardless of value.
- **`DoesNotExist`**: Matches if the key does not exist.
- **`Gt`**: Matches if the key's value is greater than the specified value.
- **`Lt`**: Matches if the key's value is less than the specified value.

---

## Pod Affinity

### Fields
- `requiredDuringSchedulingIgnoredDuringExecution`: Hard constraints for pod placement.
- `preferredDuringSchedulingIgnoredDuringExecution`: Soft constraints for pod placement.

### Example: Pod Affinity
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: web-pod
spec:
  affinity:
    podAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
      - labelSelector:
          matchExpressions:
          - key: app
            operator: In
            values:
            - backend
        topologyKey: kubernetes.io/hostname
  containers:
  - name: web
    image: nginx
```

### Key Concepts for Pod Affinity
- **`labelSelector`**: Matches pods with specific labels.
- **`topologyKey`**: Defines the topology for co-scheduling, such as nodes or zones.

### Supported Operators
Same as Node Affinity.

---

## Pod Anti-Affinity

### Fields
- `requiredDuringSchedulingIgnoredDuringExecution`: Hard constraints for pod placement.
- `preferredDuringSchedulingIgnoredDuringExecution`: Soft constraints for pod placement.

### Example: Pod Anti-Affinity
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: cache-pod
spec:
  affinity:
    podAntiAffinity:
      preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 100
        podAffinityTerm:
          labelSelector:
            matchExpressions:
            - key: app
              operator: In
              values:
              - cache
          topologyKey: kubernetes.io/hostname
  containers:
  - name: redis
    image: redis
```

### Key Concepts for Pod Anti-Affinity
- **Avoidance**: Ensures pods with similar labels are not co-scheduled.
- **`weight`**: Priority level for the rule when it’s a soft constraint.

---

## Full Example Combining All Types

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: complex-pod
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: disktype
            operator: In
            values:
            - ssd
    podAffinity:
      preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 50
        podAffinityTerm:
          labelSelector:
            matchExpressions:
            - key: app
              operator: In
              values:
              - backend
          topologyKey: kubernetes.io/hostname
    podAntiAffinity:
      preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 100
        podAffinityTerm:
          labelSelector:
            matchExpressions:
            - key: app
              operator: In
              values:
              - frontend
          topologyKey: kubernetes.io/hostname
  containers:
  - name: complex-app
    image: my-app
```

---

## Taints and Affinity Working Together

Taints and tolerations are used to repel pods from certain nodes, while affinity rules can provide additional constraints or preferences for pod scheduling. Together, they allow fine-grained control over pod placement.

### Taints and Tolerations
- **Taints**: Applied to nodes to repel pods unless those pods have matching tolerations.
- **Tolerations**: Added to pods to allow them to be scheduled on tainted nodes.

### Example: Taints and Tolerations
#### Adding a Taint to a Node
```bash
kubectl taint nodes node1 key=value:NoSchedule
```

#### Adding a Toleration to a Pod
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: tolerant-pod
spec:
  tolerations:
  - key: "key"
    operator: "Equal"
    value: "value"
    effect: "NoSchedule"
  containers:
  - name: nginx
    image: nginx
```

---

### Example: Taints and Affinity Together
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: taint-affinity-pod
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: disktype
            operator: In
            values:
            - ssd
  tolerations:
  - key: "dedicated"
    operator: "Equal"
    value: "special"
    effect: "NoSchedule"
  containers:
  - name: nginx
    image: nginx
```

### Diagram: Taints, Tolerations, and Affinity
```plaintext
+----------------------+                +-------------------------+
|     Node 1           |  Taint:       | Pod with toleration    |
| disktype=ssd         |  dedicated=special:NoSchedule           |
| taint applied        |<--------------> Affinity: disktype=ssd  |
+----------------------+                +-------------------------+

  Node 1 accepts the pod because:
  - The pod has a toleration for the taint.
  - The pod's affinity matches the node label.
```

---

### Key Points
1. **Node Affinity**: Targets nodes based on labels.
2. **Pod Affinity**: Ensures pods are co-scheduled based on pod labels.
3. **Pod Anti-Affinity**: Prevents pods from being co-scheduled with certain pods.
4. **Taints and Tolerations**: Control pod scheduling by repelling pods unless tolerations are defined.
5. **Working Together**: Taints repel pods by default, while affinity fine-tunes placement.

