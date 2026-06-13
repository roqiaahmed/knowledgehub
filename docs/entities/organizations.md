# Organizations Entity Design (V1)

## Purpose

Represents a company, team, or group that collaborates within KnowledgeHub.

## Fields

id

- UUID
- Primary Key

name

- Required

slug

- Required
- Unique

description

- Optional

created_at

- Timestamp

updated_at

- Timestamp

## Constraints

- Slug must be unique.
- Name is not required to be unique.
- Every organization must have an owner.
- An organization cannot be created without creating an owner membership.

## Business Rules

### Creation

When an organization is created:

1. Organization record is created.
2. Membership record is created.
3. Creator becomes owner.

### Ownership

Every organization must have exactly one owner.

Before ownership is removed:

- Ownership must be transferred to another member.

### Membership

Organizations gain members through memberships.

Users may belong to multiple organizations.

## Future Features

- Invitations
- Organization logo
- Billing
- Organization settings
