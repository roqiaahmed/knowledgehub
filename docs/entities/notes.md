# Notes Entity Design (V1)

## Purpose

Represents knowledge stored within a workspace.

Notes are the primary content type in KnowledgeHub and allow teams to document processes, guides, decisions, and project information.

Examples:

- Authentication Guide
- API Standards
- Deployment Checklist
- Team Onboarding Guide

## Ownership

Every note belongs to exactly one workspace.

A workspace may contain multiple notes.

Every note is created by a user.

## Fields

id

- UUID
- Primary Key

workspace_id

- References Workspaces

title

- Required

content

- Required
- Stored as Markdown

created_by

- References Users

created_at

- Timestamp

updated_at

- Timestamp

deleted_at

- Timestamp
- Used for soft deletion

## Constraints

- Every note must belong to a workspace.
- Every note must have a creator.
- Title is required.
- Content is required.

## Content Format

Notes are stored using Markdown.

Examples:

- Headings
- Lists
- Links
- Code blocks

Reason:

Markdown is lightweight, developer-friendly, and easy to render in future versions.

## Deletion Strategy

Soft Delete

When a note is deleted:

- The record is not permanently removed.
- deleted_at is populated.
- The note is hidden from normal queries.

Reason:

Notes contain valuable information and accidental deletion should be recoverable.

## Access Rules

Version 1:

Organization Owner:

- Create notes
- Read notes
- Update notes
- Delete notes

Organization Admin:

- Create notes
- Read notes
- Update notes
- Delete notes

Organization Member:

- Create notes
- Read notes
- Update their own notes

Delete permissions may be restricted based on future business requirements.

## Business Rules

### Note Creation

When a note is created:

1. Workspace must exist.
2. User must belong to the workspace's organization.
3. Creator is recorded in created_by.

### Note Update

Users may update notes they are authorized to access.

updated_at must be updated automatically.

### Note Deletion

Deleted notes are soft deleted.

The original record remains in the database.

## Future Features

The following features are intentionally postponed:

- Comments
- Drafts
- Note version history
- Note sharing
- Attachments
- Rich text editor
- Activity tracking

## Notes

Version 1 focuses on a clean and maintainable note system.

Additional collaboration features will be added only after the core system is stable and deployed.
