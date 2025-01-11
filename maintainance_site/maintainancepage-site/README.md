# Site README

## Overview
This folder contains the **customer-facing maintenance site** that provides real-time updates about maintenance windows, system statuses, and ongoing issues. It is built using **React** with **Vite** for a fast development experience and **Tailwind CSS** for styling. This README explains the folder structure, key files, and how to run the site locally (with and without Docker).

---

## Folder Structure

### **Root Folder**
- **`Dockerfile`**: Docker configuration file to containerize the site.
- **`eslint.config.js`**: ESLint configuration for linting the code.
- **`index.html`**: Main HTML file for the application.
- **`package.json`**: Manages project dependencies.
- **`package-lock.json`**: Dependency lock file.
- **`pnpm-lock.yaml`**: Lock file for `pnpm` dependencies.
- **`postcss.config.js`**: Configuration for PostCSS, used alongside Tailwind CSS.
- **`tailwind.config.js`**: Tailwind CSS configuration file.
- **`tsconfig.json`**: TypeScript configuration.
- **`vite.config.ts`**: Configuration file for Vite.

### **`src/`**
Contains the source files for the site.

#### **Root Source Files**
- **`App.tsx`**: Main application component.
- **`main.tsx`**: Entry point of the application.
- **`index.css`**: Global CSS for styling.
- **`vite-env.d.ts`**: TypeScript environment declarations.

#### **`src/components/`**
Modular and reusable UI components for the maintenance site:
- **`ActiveMaintenance.tsx`**: Displays information about currently active maintenance windows.
- **`ContactInfo.tsx`**: Provides contact details for customer support.
- **`CountdownTimer.tsx`**: Shows a countdown timer for ongoing or scheduled maintenance.
- **`MaintenanceIcon.tsx`**: A reusable icon component for maintenance-related visuals.
- **`MaintenancePage.tsx`**: Main page component displaying all maintenance-related information.
- **`StatusInfo.tsx`**: Displays high-level system status.
- **`SystemStatus.tsx`**: Detailed status for individual system components.
- **`UpdatesList.tsx`**: Lists recent updates or announcements about maintenance.

#### **`src/services/`**
- **`api.ts`**: Handles API calls to fetch maintenance and system status data from the backend.

#### **`src/types/`**
- **`index.ts`**: TypeScript type definitions for the application (e.g., API response types).

#### **`src/utils/`**
- **`timeUtils.ts`**: Utility functions for time calculations (e.g., formatting countdowns).

---

## Key Features

### Real-Time Updates
- Displays ongoing and upcoming maintenance windows.
- Countdown timers to show the remaining time until maintenance starts or ends.
- Lists recent updates for transparency.

### System Status
- Displays the overall system status (e.g., Operational, Degraded).
- Shows detailed statuses for individual components.

### Customer Support
- Provides easy access to contact information for customer support.

---

## Running Locally

### **Prerequisites**
- **Node.js** and **npm** or **pnpm** installed.
- **Docker** (if running with containers).

### **With Docker**
1. Build the Docker image:
   ```bash
   docker build -t maintenance-site .
   ```
2. Run the container:
   ```bash
   docker run -p 3000:3000 maintenance-site
   ```
3. Access the site at [http://localhost:3000](http://localhost:3000).

### **Without Docker**
1. Install dependencies:
   ```bash
   pnpm install
   ```
2. Start the development server:
   ```bash
   pnpm dev
   ```
3. Access the site at [http://localhost:3000](http://localhost:3000).

---

## Key Notes

### **Environment Variables**
If the site depends on environment variables for API URLs or configurations, ensure these are set up in a `.env` file or passed during runtime.

### **API Integration**
The site interacts with the backend to fetch data such as system statuses and maintenance updates. Ensure the backend is running and accessible during development.

---

## Conclusion
The maintenance site is designed to be lightweight, fast, and user-friendly, providing customers with real-time insights into system maintenance. Let us know if additional information or support is needed!