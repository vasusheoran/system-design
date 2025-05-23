# Cab Booking Application

## Description

This project implements a simplified cab booking application. The system focuses on core functionalities such as user and driver onboarding, ride searching, booking, and billing, all managed using in-memory data structures.

## Features

* **User Management:** Users can register, update their personal details, and update their current location.
* **Driver Management:** Driving partners can onboard with their vehicle details, update their current location, and change their availability status.
* **Ride Booking:** Users can book rides on a specified route (source and destination).
  * **Ride Search & Selection:** Users can search for available rides based on their source and destination. The system will display drivers nearest to the user (within a maximum distance of 5 units) who are currently available. Users can then select a preferred driver.
* **Billing:** The application calculates the bill for a ride based on the distance between the source and destination.
* **Earnings Tracking:** The system can calculate and display the total earnings for all onboarded drivers. (timestamp or all time)

## API

### User
- AddUser(userID, userDetails)
- UpdateUser(userID, userDetails)
- UpdateLocation(userID, Location)

### Driver
- AddDriver(driverID, driverDetails, vehicleDetails)
- UpdateLocation(driverID, location)
- UpdateStatus(Status)

### Ride
- Search(userID, source, destination) []Drivers
- SelectDriver(userID, driverID) => BookRide(userID, driverID)

### Billing
- CalculateBill(source, destination)
- Earnings(userID, since) => Earnings(driverID, since)