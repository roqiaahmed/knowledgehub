# Database Schema V1

This document defines the PostgreSQL schema for KnowledgeHub Version 1.

Goals:

- Data integrity
- Clear relationships
- Simplicity
- Scalability for future versions

## Users

id UUID PRIMARY KEY DEFAULT gen_random_uuid()
email VARCHAR(255) NOT NULL UNIQUE
full_name VARCHAR(60) NOT NULL
password_hash VARCHAR(255) NOT NULL
is_active BOOLEAN NOT NULL DEFAULT true
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

## Organizations

id UUID PRIMARY KEY DEFAULT gen_random_uuid()
name VARCHAR(120) NOT NULL UNIQUE
slug VARCHAR(120) NOT NULL UNIQUE
description VARCHAR(255)
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

## Memberships

id UUID PRIMARY KEY DEFAULT gen_random_uuid()
user_id UUID NOT NULL
organization_id UUID NOT NULL
role VARCHAR(20) NOT NULL
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

UNIQUE(user_id, organization_id)

## Workspaces

id UUID PRIMARY KEY DEFAULT gen_random_uuid()
organization_id UUID NOT NULL
name VARCHAR(120) NOT NULL
description VARCHAR(255)
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

UNIQUE(organization_id, name)

## Notes

id UUID PRIMARY KEY DEFAULT gen_random_uuid()
workspace_id UUID NOT NULL
created_by UUID NOT NULL
title VARCHAR(255) NOT NULL
content TEXT NOT NULL
created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
deleted_at TIMESTAMP NULL
