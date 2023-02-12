# Auth Micro-Service

# Principals

- db modules check integrity, not security/autherization
- service modules check security/authorization, not integrity
- e.g. service checks session user may add another user to the account, then db checks that name is not a duplicate, or service checks user is authorized to manage roles, but db checks user and role belong to same account.

# Done
- Created db and tables and db module, interface in top level and service

# Busy with:
- Implementing default/minimal operations

Todo:
- See if can import one docker-compose into another for auth to import into system compose
- Add GetUser to also return a filtered list of roles like GetAccount, also GetRole with permissions and GetPermission with nothing
- Operations to add role to user and permission to role
- Operation to removed permissions from role and roles from user
- Operation to delete with all references
- Operation to check if has access
- Then add security... tokens that expire and an admin token and user passwords and give admin user permission to add users...
- Operation to renew a token
- Review OAuth mechanisms