# Cab Booking Application

## Description

This project implements a simplified cab booking application. The system focuses on core functionalities such as user and driver onboarding, ride searching, booking, and billing, all managed using in-memory data structures.

## Features

* **Ride Booking:** Users can book rides on a specified route (source and destination).
* **User Management:** Users can register, update their personal details, and update their current location.
* **Driver Management:** Driving partners can onboard with their vehicle details, update their current location, and change their availability status.
* **Ride Search & Selection:** Users can search for available rides based on their source and destination. The system will display drivers nearest to the user (within a maximum distance of 5 units) who are currently available. Users can then select a preferred driver.
* **Billing:** The application calculates the bill for a ride based on the distance between the source and destination.
* **Earnings Tracking:** The system can calculate and display the total earnings for all onboarded drivers.

## Requirements & API

The application provides the following functionalities:

### User Onboarding

* `add_user(user_detail)`: Adds basic user details to the system.
* `update_user(username, updated_details)`: Allows a user to update their contact details.
* `update_userLocation(username, Location)`: Updates the user's current location using X, Y coordinates, essential for finding the nearest drivers.

### Driver Onboarding

* `add_driver(driver_details, vehicle_details, current_location)`: Registers a new driver and their vehicle, marking their initial current location on the map.
* `update_driverLocation(driver_name)`: Updates the current location of a driver.
* `change_driver_status(driver_name, status)`: Allows a driver to toggle their availability (e.g., `True` for available, `False` for unavailable).

### Ride Management

* `find_ride(Username, Source, destination)`: Returns a list of available drivers for a given user, source, and destination.
    * **Note:** Only drivers within a maximum distance of 5 units from the user's source location will be displayed.
    * **Note:** Only drivers in an "available" state will be considered for booking.
* `choose_ride(Username, driver_name)`: Allows a user to select a specific driver from the list of available rides.

### Billing & Earnings

* `calculateBill(Username)`: Calculates and displays the bill for the completed ride based on the distance between the source and destination.
* `find_total_earning()`: Calculates and displays the total earnings for all drivers in the application.

## Other Notes

* **Demo Class:** A `Driver` class (or similar) will be implemented for demonstration purposes, executing all commands and test cases in one place.
* **Data Storage:** No database or NoSQL store will be used. All data will be managed using in-memory data structures.
* **User Interface:** No graphical user interface (UI) will be developed for this application. It will be a console-based or API-driven application.
* **Prioritization:** The primary focus is on code compilation, execution, and functional completeness.
* **Bonus Features:** Concurrency handling is considered a bonus feature.

## Expectations

* **Working & Demo-able Code:** The final output should be a fully functional and demonstrable application.
* **Functional Correctness:** All features must work as described and produce accurate results.
* **Code Quality (Good to Have):**
    * Proper abstraction and entity modeling.
    * Clear separation of concerns.
    * Modular, readable, and unit-testable code.
    * Code that can easily accommodate new requirements with minimal changes.
    * Robust exception handling.

## Sample Test Cases

```
# Onboard 3 users
add_user("Abhishek", "M", 23)
update_userLocation("Abhishek", (0, 0))

add_user("Rahul", "M", 29)
update_userLocation("Rahul", (10, 0))

add_user("Nandini", "F", 22)
update_userLocation("Nandini", (15, 6))

# Onboard 3 drivers
add_driver("Driver1", "M", 22, "Swift", "KA-01-12345", (10, 1))
add_driver("Driver2", "M", 29, "Swift", "KA-01-12345", (11, 10))
add_driver("Driver3", "M", 24, "Swift", "KA-01-12345", (5, 3))

# User trying to get a ride
print(find_ride("Abhishek", (0, 0), (20, 1)))
# Expected Output: No ride found [Since all the drivers are more than 5 units away from the user]

print(find_ride("Rahul", (10, 0), (15, 3)))
# Expected Output: Driver1 [Available]

print(choose_ride("Rahul", "Driver1"))
# Expected Output: ride Started

print(calculateBill("Rahul"))
# Expected Output: ride Ended bill amount $ 6

# Backend API Calls (simulated)
update_userLocation("Rahul", (15, 3))
update_driverLocation("Driver1", (15, 3))
change_driver_status("Driver1", False)

print(find_ride("Nandini", (15, 6), (20, 4)))
# Expected Output: No ride found [Driver one is set to not available]

# Total earning by drivers
print(find_total_earning())
# Expected Output:
# Driver1 earn $6
# Driver2 earn $0
# Driver3 earn $0
```

## Installation & Usage

*(This section is a placeholder. Detailed instructions will be provided once the code is developed.)*

To run this application:

1.  Clone the repository.
2.  Navigate to the project directory.
3.  Execute the main script (e.g., `python main.py`).

## Technologies Used

*(This section is a placeholder. Will be updated with specific technologies once the code is developed.)*

* Python (for core logic)
* In-memory data structures