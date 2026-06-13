# Workspaces Entity Design (V1)

## Purpose

Represents a collaborative area within an organization.

Workspaces are used to organize knowledge and contain notes.

Examples:

- Backend Team
- Mobile Team
- Marketing Team
- Project Alpha

## Ownership

Every workspace belongs to exactly one organization.

An organization may have multiple workspaces.

## Fields

id

- UUID
- Primary Key

organization_id

- References Organizations

name

- Required

description

- Optional

created_at

- Timestamp

updated_at

- Timestamp

## Constraints

- Every workspace must belong to an organization.
- Workspace names must be unique within the same organization.
- Different organizations may use the same workspace name.

Example:

Organization A

- Backend Team

Organization B

- Backend Team

Allowed.

Organization A

- Backend Team
- Backend Team

Not allowed.

## Access Rules

Version 1:

- All organization members can access all workspaces.
- No workspace-specific permissions.
- No workspace membership table.

Reason:

Keep the system simple while learning core backend concepts.

## Permissions

Owner:

- Create workspaces
- Update workspaces
- Delete workspaces

Admin:

- Create workspaces
- Update workspaces
- Delete workspaces

Member:

- View workspaces

## Business Rules

### Workspace Creation

When a workspace is created:

1. The organization must exist.
2. The creator must be an owner or admin.
3. The workspace is linked to the organization.

### Workspace Deletion

Only owners and admins can delete a workspace.

Deleting a workspace should also remove its notes.

## Future Features

The following features may be added later:

- Private workspaces
- Workspace memberships
- Workspace roles
- Workspace invitations
- Workspace settings
- Workspace activity logs

## Notes

Version 1 focuses on organization-level permissions.

Workspace-level permissions are intentionally postponed to reduce complexity and increase the likelihood of finishing the project.
