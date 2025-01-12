### Understanding Kubernetes Requests & Limits 🚀🔧

Kubernetes **Requests** and **Limits** are mechanisms that help manage resource allocation (CPU and memory) within your cluster. This ensures fair use of resources among all running workloads.

---

### 🏙️ **Requests and Limits Overview**

1. **Requests**:
   - The minimum amount of CPU or memory guaranteed to a pod.
   - The scheduler uses this value to decide which node the pod can run on.
   - Example: If a pod requests `100Mi` of memory, it will be scheduled on a node with at least `100Mi` of free memory.

2. **Limits**:
   - The maximum amount of CPU or memory a pod can use.
   - If a pod exceeds this limit:
     - **For CPU**: It will be throttled.
     - **For Memory**: It will be killed due to an **Out of Memory (OOM)** error.

---

### 🧐 **Why Requests & Limits Are Important?**

1. **Resource Control**:
   - Prevents a single pod from monopolizing cluster resources.
   - For memory, an OOM kill ensures the pod doesn’t consume all memory, potentially crashing the node.

2. **Predictability**:
   - Requests ensure the pod always has the resources it needs to run.
   - Limits act as a safeguard, capping resource usage.

---

### 🔍 **Resource Management in Action**

#### YAML Example 1: Pod Exceeding Available Memory
This pod demonstrates what happens when a container tries to use more memory than the specified **limit**.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: memory-demo-2  # Name of the pod
  namespace: mem-example  # Namespace where this pod will run
spec:
  containers:
  - name: memory-demo-2-ctr  # Name of the container
    image: polinux/stress  # Image that runs a stress test
    resources:  # Resource requests and limits
      requests:
        memory: "50Mi"  # Minimum memory guaranteed
      limits:
        memory: "100Mi"  # Maximum memory allowed
    command: ["stress"]  # Command to run inside the container
    args: ["--vm", "1", "--vm-bytes", "250M", "--vm-hang", "1"]
```

**Explanation**:
- **Requests**: The pod is guaranteed `50Mi` of memory.
- **Limits**: The pod is capped at `100Mi` of memory. 
- **Command**: The `stress` command will attempt to allocate `250M` of memory (`--vm-bytes`, with one worker process `--vm 1`).
- **Outcome**:
  - The pod will attempt to use `250Mi`, exceeding its limit of `100Mi`.
  - Kubernetes will kill the container due to an **OOM (Out of Memory)** error.

---

#### YAML Example 2: Pod Within Limits
This pod demonstrates a well-behaved container that stays within its **limits**.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: memory-demo  # Name of the pod
  namespace: mem-example  # Namespace where this pod will run
spec:
  containers:
  - name: memory-demo-ctr  # Name of the container
    image: polinux/stress  # Image that runs a stress test
    resources:  # Resource requests and limits
      requests:
        memory: "100Mi"  # Minimum memory guaranteed
      limits:
        memory: "200Mi"  # Maximum memory allowed
    command: ["stress"]  # Command to run inside the container
    args: ["--vm", "1", "--vm-bytes", "150M", "--vm-hang", "1"]
```

**Explanation**:
- **Requests**: The pod is guaranteed `100Mi` of memory.
- **Limits**: The pod is capped at `200Mi` of memory.
- **Command**: The `stress` command will allocate `150Mi` of memory (`--vm-bytes`).
- **Outcome**:
  - The pod requests `150Mi`, which is within its limit of `200Mi`.
  - The pod runs successfully without being killed.

---

### Key Commands for Resource Management

1. **Apply a YAML File**:
   ```bash
   kubectl apply -f <filename>.yaml
   ```

2. **Check Pod Status**:
   ```bash
   kubectl get pods -n mem-example
   ```

3. **Describe Pod**:
   ```bash
   kubectl describe pod <pod-name> -n mem-example
   ```
   This shows detailed resource usage and events (e.g., OOM errors).

4. **View Node Resource Utilization**:
   ```bash
   kubectl top nodes
   ```

5. **View Pod Resource Utilization**:
   ```bash
   kubectl top pods -n mem-example
   ```

---

### Common Scenarios

1. **Exceeding Requests**:
   - If a container exceeds its **request** for CPU, it may still use more CPU (up to the limit or node capacity) but will compete with other pods.

2. **Exceeding Limits**:
   - For memory: The container will be killed with an OOM error.
   - For CPU: The container will be throttled.

---

### Best Practices

1. Always set **requests** and **limits** for production workloads to avoid resource contention.
2. Use monitoring tools (e.g., Prometheus, Grafana) to analyze actual resource usage.
3. Test your workloads with different resource settings to find the optimal configuration.

By understanding and properly configuring **Requests** and **Limits**, you can achieve efficient, reliable, and predictable resource management in your Kubernetes cluster. 🚀
