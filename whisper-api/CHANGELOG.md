# 📜 Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/)
and this project adheres to [Semantic Versioning](https://semver.org/).

---

## [Unreleased]

### Added
- Introduced `UploadBusinessImage` endpoint: `/business/:id/upload`.
- Added `ImageURL` field to `Business` model for storing business image paths.
- Automatically serve uploaded images from `/uploads/business/` via static route.
- Implemented Gin validation for required `Business` fields.
- Added support for **PostgreSQL** via `DATABASE_URL` environment variable.
- Implemented **automatic SQLite fallback** when PostgreSQL is unavailable (useful for local MVP development).
- Configured **Railway deployment compatibility** for seamless Postgres connections.

### Fixed
- Handled incorrect POST payloads creating empty businesses.
- Resolved average rating not updating issue in review controller.
- Improved database connection reliability and startup logs.
- Ensured `uploads/businesses` directory auto-creates on server start.

---

## [v1.0.0] - 2025-10-22
### Added
- Initial project setup with:
  - Gin router
  - SQLite database
  - Business and Review models
  - CRUD endpoints for `/business` and `/review`

---
