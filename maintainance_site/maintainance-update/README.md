# Folder-Specific README for Maintenance Update App

## Overview
This folder contains the **Maintenance Update Admin Panel** for managing maintenance windows, updates, and system components. The admin panel is built using **Next.js** with a modular and reusable component-based design, leveraging **TypeScript** for type safety and **Tailwind CSS** for styling.

This README explains the folder structure, key files, and how to run the project locally (with and without Docker).

---

## Folder Structure

### **Root Folder**
- **`globals.css`**: Global CSS file for the app.
- **`layout.tsx`**: Base layout for the application.
- **`page.tsx`**: Main landing page for the admin panel.
- **`types/`**: Contains reusable TypeScript type definitions.
  - **`form-schemas.ts`**: Schema definitions for forms using `zod`.
  - **`index.ts`**: Exports for all reusable types.

### **Components**
- **`cors-test.tsx`**: A test component for validating CORS configurations.
- **`data-entry-form.tsx`**: Dynamic form wrapper that renders different forms based on selected update types.
- **`forms/`**: Contains specific forms for handling updates.
  - **`maintenance-updates-form.tsx`**: Form for adding/updating maintenance updates.
  - **`maintenance-windows-form.tsx`**: Form for scheduling and managing maintenance windows.
  - **`system-components-form.tsx`**: Form for updating system component statuses.
- **`ui/`**: Modular, reusable UI components like buttons, alerts, tables, etc.

### **Configuration Files**
- **`config.ts`**: Stores application-level configurations, such as API base URL.
- **`next.config.js`**: Configuration file for the Next.js app.
- **`tailwind.config.ts`**: Configuration file for Tailwind CSS.
- **`postcss.config.js`**: Configuration for PostCSS, used alongside Tailwind CSS.
- **`tsconfig.json`**: TypeScript configuration.

### **Hooks**
- **`use-toast.ts`**: Custom React hook for handling toast notifications.

### **Lib**
- **`utils.ts`**: Utility functions for common operations.

### **Package Management**
- **`package.json`**: Manages project dependencies.
- **`pnpm-lock.yaml`**: Lock file for dependency consistency.

---

## Key Components

### **MaintenanceUpdatesForm**
Located at: `components/forms/maintenance-updates-form.tsx`

This form handles updates for ongoing maintenance windows. It includes fields for:
- **Message**: A text input for update details.
- **Maintenance Window ID**: Numeric input for associating updates with a maintenance window.

### **MaintenanceWindowsForm**
Located at: `components/forms/maintenance-windows-form.tsx`

This form is used for creating and managing maintenance windows. It includes:
- **Start Time**: DateTime picker for scheduling the window.
- **Estimated Duration**: Numeric input for specifying the expected downtime.
- **Description**: Text input for additional information.
- **Issue Fixed**: Dropdown to indicate if the issue has been resolved.

### **SystemComponentsForm**
Located at: `components/forms/system-components-form.tsx`

This form manages the status of individual system components. It allows for:
- Adding new components.
- Updating the status of existing components (e.g., Operational, Maintenance).

---

## Running Locally

### **Prerequisites**
- **Node.js** and **npm** or **pnpm** installed.
- **Docker** (if running with containers).

### **With Docker**
1. Build the Docker image:
   ```bash
   docker build -t maintenance-update-app .
   ```
2. Run the container:
   ```bash
   docker run -p 3000:3000 maintenance-update-app
   ```
3. Access the app at [http://localhost:3000](http://localhost:3000).

### **Without Docker**
1. Install dependencies:
   ```bash
   pnpm install
   ```
2. Start the development server:
   ```bash
   pnpm dev
   ```
3. Access the app at [http://localhost:3000](http://localhost:3000).

---

## Key Notes
- **Environment Variables**: Ensure you configure environment variables such as API URLs if necessary.
- **API Endpoints**: The app communicates with a backend API, so ensure the backend is running on the expected port (default: `8080`).
- **Testing CORS**: Use the `cors-test.tsx` component to validate CORS configurations.

---

## Conclusion
This admin panel is designed for flexibility and scalability, making it easy for SRE teams to manage maintenance windows, system statuses, and updates dynamically. Let us know if you need further assistance!