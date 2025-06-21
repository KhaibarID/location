# Clean Architecture Folder Structure

This project follows the Clean Architecture pattern. Below is the initial folder structure:

- `/cmd` - Application entrypoints (main, CLI, etc)
- `/internal`
    - `/domain` - Entities, business rules
    - `/usecase` - Application-specific business rules (interactors)
    - `/repository` - Interfaces for data access
    - `/service` - Implementation of services (external APIs, etc)
    - `/handler` - Delivery mechanisms (HTTP, gRPC, etc)
- `/pkg` - Shared packages (if needed)
- `/configs` - Configuration files
- `/scripts` - Utility scripts
- `/test` - Test data and helpers

You can start building your application by adding code to the respective folders.
