# Memberships Entity Design (V1)

## Purpose

Represents a user's membership inside an organization.

A membership connects:

- User
- Organization
- Role

## Fields

id

- UUID
- Primary Key

user_id

- References Users

organization_id

- References Organizations

role

- owner
- admin
- member

created_at

- Timestamp

updated_at

- Timestamp

## Constraints

A user may belong to multiple organizations.

An organization may have multiple users.

A user may only have one membership per organization.

Exactly one owner must exist per organization.

Multiple admins are allowed.

## Business Rules

### Organization Creation

When an organization is created:

1. Organization is created.
2. Membership is created.
3. Creator becomes owner.

### Ownership Transfer

Before ownership is removed:

- another member must become owner.

### Membership Removal

Removing a membership removes the user's access to the organization.

## Future Features

- Invitations
- Pending memberships
- Custom roles
- Role history
