# comp590-finalexam

**Author**: Madison Roberts 
**PID**: 730460151

## Object-Oriented Pillars in Elixir 
### Encapsulation 
Each animal is represented as an isolated process. Internal state (name) is private. It can't be accessed directly, only via message passing ({:get_name, sender}). 
### Abstraction 
The Animal module hides process and message passing logic. Users don't need to know how messages are handled internally. They only know they can :speak or :get_name 
### Inheritance  
While Elixir doesn't support classical inheritance, I simulated it by passing the behavior module (Dog or Cat) into the generic Animal module. This mimics calling super() in base class. 
### Polymorphism 
Each behavior module (Dog, Cat) implements its own version of speak/2. When an Animal receives a :speak message, it delegates to the behavior module allowing runtime polymorphism. 

## How to Run
```bash
elixir animal.exs