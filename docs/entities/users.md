# Users Table Design (V1)

## Purpose

Represents authenticated users of the platform.

## Authentication

Version 1:

- Email + Password

Future:

- Google OAuth Login

## Fields

id

- UUID
- Primary Key

email

- Unique
- Required

password_hash

- Required
- Never store plain text passwords

full_name

- Required

is_active

- Boolean
- Default: true

created_at

- Timestamp

updated_at

- Timestamp

## Constraints

- Email must be unique.
- Passwords must be hashed.
- A user cannot log in when is_active = false.

## Future Fields

These are not required now but may be added later:

email_verified
google_id
last_login_at
profile_image_url
