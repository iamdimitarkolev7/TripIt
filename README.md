# TripIt
TripIt is a web application which aims to make group travels easier, cheaper and more convenient.

## Technologies
 * Backend: Golang
 * Database: MongoDB
 * Frontend: React

## Entities
 * User: when creating a user we have to choose between two roles(Driver and Passenger); there will be a super user too, who will manage all the data from the server; every user will have an option to keep track of their travels and visited locations; at the end of the travel every user can rate the other users they travelled with;
 * Travel: this entity will keep data for a particular travel(Ex. driver, passengers, date, from-to destination);
 * Car: the driver users will have the option to publish data for their car(Ex. available seats, model, fuel consumption), so the passengers are informed about it and decide easily whether they want to travel with the particular driver or not;
 * Location: this entity will keep the following data: pics from the location, name, description and all the itineraries in this location
 * Itinerary: this entity will be used to inform the users about the distance, difficulty, rating and the name of the itinerary in a particular location.

## Functionalities
 * User authentication
 * Drivers will publish their travel suggestions
 * Passengers will apply for a particular travel
 * Displaying users visited locations and travelled itineraries
