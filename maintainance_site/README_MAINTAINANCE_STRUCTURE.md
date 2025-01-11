# Project Structure and Intentions

This project is a robust **Maintenance Page System** designed to handle application maintenance seamlessly. The repository consists of three primary components:

- `maintainance-update`: Contains the admin-facing system for managing maintenance and updates.
- `maintainancepage-backend`: The backend API and database layer.
- `maintainancepage-site`: The customer-facing maintenance page that communicates real-time statuses.

Below is a detailed breakdown of each component and its intentions.

---

## **Project Structure**

### 1. **Root Directory**
- **`README.md`**: Main documentation file for the repository.
- Houses subdirectories for the admin panel, backend services, and customer-facing site.

---

### 2. **`maintainance-update`**
This directory contains the **admin panel** where Site Reliability Engineers (SREs) can manage maintenance windows, system statuses, and updates.

#### Structure:
- **`app`**: The core of the Next.js application.
  - **`globals.css`**: Global CSS for styling.
  - **`layout.tsx`**: Layout structure for the application.
  - **`page.tsx`**: The primary landing page.
  - **`types/`**: Shared TypeScript type definitions.
    - **`form-schemas.ts`**: Schema definitions for forms.
    - **`index.ts`**: Exports all reusable types.
- **`components/`**: Reusable UI components.
  - **`cors-test.tsx`**: Utility for testing CORS functionality.
  - **`data-entry-form.tsx`**: Base component for forms.
  - **`forms/`**: Specific form components.
    - **`maintenance-updates-form.tsx`**: Handles updates on maintenance.
    - **`maintenance-windows-form.tsx`**: Form for scheduling maintenance windows.
    - **`system-components-form.tsx`**: Form for system component updates.
  - **`ui/`**: Modular and reusable UI elements (accordion, button, tooltip, etc.).
- **`components.json`**: Metadata for all reusable components.
- **`config.ts`**: Application configuration.
- **`hooks/`**: Contains custom React hooks.
  - **`use-toast.ts`**: Hook for managing toast notifications.
- **`lib/`**: Utility functions for common operations.
  - **`utils.ts`**: Shared utility methods.
- **Configuration Files**:
  - **`next.config.js`**: Next.js configuration.
  - **`tailwind.config.ts`**: Tailwind CSS configuration.
  - **`postcss.config.js`**: PostCSS configuration.
  - **`package.json`**, **`pnpm-lock.yaml`**: Dependencies and package management.
  - **`tsconfig.json`**: TypeScript configurations.

#### Intentions:
To provide a **user-friendly admin interface** that allows SREs to manage:
- Maintenance windows.
- System statuses.
- Real-time updates visible to customers.

---

### 3. **`maintainancepage-backend`**
This directory contains the **backend API** and database interactions for the system.

#### Structure:
- **`Dockerfile`**: Container configuration for the backend.
- **`cmd/`**: Entry points for backend commands.
  - **`migrate/main.go`**: Handles database migrations.
  - **`seed/main.go`**: Seeds the database with dummy data.
  - **`server/main.go`**: Starts the backend server.
- **`data/`**: SQL files for seeding and initializing the database.
  - **`dummy_data.sql`**: Sample data for development.
- **`internal/`**: Core business logic.
  - **`config/`**: Application configurations (e.g., environment variables).
  - **`cron/cleanup.go`**: Scheduled tasks for cleaning up old data.
  - **`database/database.go`**: Database connection and utilities.
  - **`handlers/`**: API request handlers.
    - **`maintenance.go`**: Endpoints for managing maintenance data.
    - **`system_components.go`**: Endpoints for managing system component statuses.
  - **`middleware/`**: Middleware for security and request validation.
    - **`middleware.go`**: Core middleware implementations.
  - **`models/`**: Data models.
    - **`models.go`**: Defines core models for maintenance updates.
    - **`response.go`**: Standardized response structures.
  - **`router/`**: API routing.
    - **`router.go`**: Defines routes for the API.
- **`go.mod`**, **`go.sum`**: Dependency management for Go modules.

#### Intentions:
To serve as the **backend service** for the system, providing:
- API endpoints for the admin panel and customer-facing site.
- Database operations for storing and retrieving maintenance and system status data.

---

### 4. **`maintainancepage-site`**
This directory contains the **customer-facing site** that displays real-time maintenance updates.

#### Structure:
- **`Dockerfile`**: Container configuration for the site.
- **`index.html`**: Main HTML template for the site.
- **`src/`**: Source files for the React-based frontend.
  - **`App.tsx`**: Main application component.
  - **`components/`**: UI components.
    - **`ActiveMaintenance.tsx`**: Displays active maintenance updates.
    - **`ContactInfo.tsx`**: Provides customer contact details.
    - **`CountdownTimer.tsx`**: Countdown for maintenance windows.
    - **`MaintenancePage.tsx`**: Main maintenance page component.
    - **`StatusInfo.tsx`**: Displays current system status.
    - **`SystemStatus.tsx`**: Shows detailed system component statuses.
    - **`UpdatesList.tsx`**: List of recent maintenance updates.
  - **`services/api.ts`**: Handles API interactions with the backend.
  - **`types/index.ts`**: TypeScript definitions.
  - **`utils/timeUtils.ts`**: Utility functions for time calculations.
- **Configuration Files**:
  - **`tailwind.config.js`**: Tailwind CSS configuration.
  - **`postcss.config.js`**: PostCSS configuration.
  - **`tsconfig.json`**: TypeScript configurations.
  - **`vite.config.ts`**: Vite configuration for building the site.
  - **`package.json`**, **`pnpm-lock.yaml`**: Dependencies and package management.

#### Intentions:
To provide customers with a **real-time status page** that:
- Displays ongoing maintenance or downtime.
- Shares updates about system statuses and resolutions.
- Enhances communication during disruptions.

---

## Summary
This project is designed to streamline the management and communication of system maintenance activities. By leveraging:
- **`maintainance-update`**: Admin-facing panel for SREs.
- **`maintainancepage-backend`**: Robust backend APIs and database operations.
- **`maintainancepage-site`**: Customer-facing site for real-time status updates.

This ensures both internal teams and customers have the tools and information they need during critical maintenance periods.