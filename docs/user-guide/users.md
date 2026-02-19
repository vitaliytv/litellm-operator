# Users

Users represent individual accounts in your LiteLLM system. They provide the foundation for authentication, authorization, and budget management.

## Overview

User resources in the LiteLLM Operator provide:

- **Identity Management** - Define user accounts and roles
- **Budget Control** - Set spending limits per user
- **Role-Based Access** - Assign different permission levels
- **Usage Tracking** - Monitor individual user consumption
- **Automatic Key Creation** - Generate virtual keys automatically

## Creating Users

### Basic User

```yaml
apiVersion: auth.litellm.ai/v1alpha1
kind: User
metadata:
  name: alice
spec:
  userEmail: "alice@example.com"
  userAlias: "alice"
  userRole: "internal_user_viewer"
  keyAlias: "alice-key"
  autoCreateKey: true
  models:
    - "gpt-4o"
  maxBudget: "10"
  budgetDuration: 1h
  connectionRef:
    instanceRef:
      name: litellm-example
      namespace: litellm
```

### Admin User

```yaml
apiVersion: auth.litellm.ai/v1alpha1
kind: User
metadata:
  name: admin-user
spec:
  userEmail: "admin@example.com"
  userAlias: "admin"
  userRole: "admin"
  keyAlias: "admin-key"
  autoCreateKey: true
  models:
    - "gpt-4o"
    - "claude-3-sonnet"
  maxBudget: "1000"
  budgetDuration: 30d
  connectionRef:
    instanceRef:
      name: litellm-example
      namespace: litellm
```

### User Without Auto-Created Key

When `autoCreateKey: false`, no virtual key is created for the user. The user account is created in LiteLLM only; you can attach a key later via a separate [VirtualKey](virtual-keys.md) resource. The `keyAlias` field is ignored when `autoCreateKey` is false.

```yaml
apiVersion: auth.litellm.ai/v1alpha1
kind: User
metadata:
  name: external-user
spec:
  userEmail: "user@example.com"
  userAlias: "external-user"
  userRole: "internal_user_viewer"
  autoCreateKey: false
  models:
    - "gpt-4o"
  maxBudget: "50"
  budgetDuration: 7d
  connectionRef:
    instanceRef:
      name: litellm-example
      namespace: litellm
```

## Specification Reference

| Field | Type | Description | Required |
|-------|------|-------------|----------|
| `userEmail` | string | User's email address | Yes |
| `userAlias` | string | User alias/username | Yes |
| `userRole` | string | User role (one of "proxy_admin", "proxy_admin_viewer", "internal_user", "internal_user_viewer") | Yes |
| `keyAlias` | string | Alias for the virtual key when `autoCreateKey` is true; ignored when `autoCreateKey` is false | No |
| `autoCreateKey` | boolean | When true, a virtual key is created automatically with the given `keyAlias`. When false, no key is created. | No (default: false) |
| `models` | []string | Allowed models for this user | No |
| `maxBudget` | string | Maximum spend limit in dollars | Yes |
| `budgetDuration` | string | Budget duration (e.g., "1h", "30d") | Yes |
| `connectionRef` | object | Reference to LiteLLM instance | Yes |

## Managing Users

### List Users

```bash
kubectl get users
```

### Get User Details

```bash
kubectl describe user alice
```

### Update User Budget

```bash
kubectl patch user alice --type='merge' -p='{"spec":{"maxBudget":"200"}}'
```

### Delete a User

```bash
kubectl delete user alice
```

## Best Practices

- Set appropriate budget limits based on usage patterns
- Use meaningful key aliases for easy identification
- Use `autoCreateKey: true` with a `keyAlias` for seamless user onboarding when a key is needed; use `autoCreateKey: false` when the user will use a separately managed VirtualKey
- Set reasonable budget durations to prevent overspending
- Regularly review and update user permissions

## Next Steps

- Learn about [Teams](teams.md) and [Team Member Associations](team-member-associations.md)
- Create [Virtual Keys](virtual-keys.md) for users
- Set up user monitoring and alerts 
