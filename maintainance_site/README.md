# Maintenance Page Project

## Overview

This project provides a **Maintenance Page System** designed to ensure smooth communication with customers during application upgrades or maintenance periods. The solution includes both a **customer-facing status page** and an **admin interface for SRE teams** to manage updates and system statuses efficiently.

---

## Features

### 1. **Customer-Facing Maintenance Page**
- A fully deployable maintenance page integrated with your application.
- Redirects customers to a **status page** during maintenance or upgrades.
- Displays:
  - Real-time **system status** of various services.
  - **Recent updates** about ongoing work.
  - **Estimated downtime** and any additional information about the application state.
- Provides an option to contact support with a displayed email address.

### 2. **SRE Admin Panel**
- Accessible interface for SRE engineers to:
  - Manage maintenance windows.
  - Update system status and recent updates dynamically.
- Updates are stored in the database and immediately reflected on the status page for customers.
- Includes fields like:
  - **Update Type**: Maintenance Window, Downtime Notice, etc.
  - **Start Time**: Scheduling the maintenance or downtime.
  - **Estimated Duration**: Time expected for resolution.
  - **Description**: Details of the issue or update.
  - **Issue Fixed**: A toggle to indicate whether the issue is resolved.

### 3. **Database Integration**
- Powered by **PostgreSQL** for robust and scalable data storage.
- Tracks:
  - System component statuses (e.g., API, Database, Cloudflare, etc.).
  - Historical updates for transparency and troubleshooting.
- Database can be deployed alongside the application.
- Provides an admin interface for SRE engineers to make seamless updates.

---

## Technical Details

### **Technology Stack**
- **Frontend**: Built with **Next.js** for modern and efficient UI development.
- **Backend**: Developed using **Go Lang** to ensure high performance and scalability.
- **Database**: **PostgreSQL** to manage and store status updates and system data.

### **Deployment**
- Maintenance Page and PostgreSQL database can be deployed as containers alongside your application.
- Automatically redirects customers to the maintenance page during scheduled updates.

---

## Work in Progress (WIP)
- This project is still in its **early stages** and actively being developed.
- Contributions are encouraged and highly appreciated to help it mature into a reliable solution.

---

## How to Contribute
1. **Clone the repository**:
   ```bash
   git clone <repository_url>
   ```
2. **Set up the environment**:
   - Install dependencies.
   - Set up a PostgreSQL database locally for testing.
3. **Make changes and test**:
   - Ensure new features or bug fixes are functional.
   - Maintain clear and concise code.
4. **Submit a pull request**:
   - Explain the purpose of the changes.
   - Reference any related issues if applicable.

---

## Future Enhancements
- **Role-based access control (RBAC)** for the admin panel.
- Automated reminders for upcoming maintenance windows.
- Integration with monitoring tools for real-time issue detection.
- APIs for updating system statuses programmatically.
- Improved customer-facing UI for better user experience.

---

## Contact
For any issues, suggestions, or contributions, feel free to reach out:
- **Support Email**: 

Together, let’s build a robust system for keeping customers informed and your infrastructure organized!