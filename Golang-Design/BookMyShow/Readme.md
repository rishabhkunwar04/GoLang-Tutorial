
## Class Diagram
``` go
+----------------------------+
|        Movie              |
+----------------------------+
| - id: String              |
| - title: String           |
| - description: String     |
| - duration: int           |
+----------------------------+

+----------------------------+
|        Theater            |
+----------------------------+
| - id: String              |
| - name: String            |
| - location: String        |
| - shows: List<Show>       |
+----------------------------+

+----------------------------+
|           Show            |
+----------------------------+
| - id: String              |
| - movie: Movie            |
| - theater: Theater        |
| - startTime: LocalDateTime|
| - endTime: LocalDateTime  |
| - seats: Map<String, Seat>|
+----------------------------+

+----------------------------+
|           Seat            |
+----------------------------+
| - id: String              |
| - row: int                |
| - column: int             |
| - type: SeatType          |
| - price: double           |
| - status: SeatStatus      |
+----------------------------+

+----------------------------+
|        Booking            |
+----------------------------+
| - id: String              |
| - user: User              |
| - show: Show              |
| - selectedSeats: List<Seat>|
| - totalPrice: double      |
| - status: BookingStatus   |
+----------------------------+

+----------------------------+
|          User             |
+----------------------------+
| - id: String              |
| - name: String            |
| - email: String           |
+----------------------------+

+----------------------------+
|  MovieTicketBookingSystem |
+----------------------------+
| - movies: Map             |
| - theaters: Map           |
| - shows: Map              |
| - bookings: ConcurrentHashMap |
+----------------------------+
```

## Design pattern we can use here
```go
Pattern	Purpose	Example Use Case
Singleton	One global system instance	MovieTicketBookingSystem
Factory	Create complex objects	Booking, Seat
Strategy	Changeable pricing logic	Normal vs premium pricing
Observer	Notify on status change	Booking confirmation
Command	Encapsulate booking operations	Book, Cancel, Confirm
Builder	Flexible object construction	Show, Theater
Repository	Decouple database access	ShowRepository, UserRepository
Decorator	Dynamic feature extension	Add-ons to bookings or seats
```
