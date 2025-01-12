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

### Key Points
1. **Node Affinity**: Targets nodes based on labels.
2. **Pod Affinity**: Ensures pods are co-scheduled based on pod labels.
3. **Pod Anti-Affinity**: Prevents pods from being co-scheduled with certain pods.
4. **Operators**: Used to define matching rules for labels.
5. **Topology Key**: Specifies the scope of affinity rules (e.g., nodes, zones, or regions).

This guide should help you effectively use affinity and anti-affinity rules in Kubernetes!

