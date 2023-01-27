Introduction

[backend]

Before writing any services, you must at least create a 1 entity/model for a specific
resources (e.g. Comments, Reacts, Users ...), implement its data sources by writing efficient queries
of any form (mysql, sqlite, ...). These implementations must pass a series of written less strict,
regularly updated unit tests, after that we proceed at writing a series of microservices that
will use these entities for a basic CRUD operation. Of course these steps are iterative until the 
blog looks finally refurbished.

[frontend]

For the user interface that will interact with the backend, we will provide a variety of methods
to which a user can interact with the blogs. For simplicity we will start working with the web
and when everything looks fine we will now proceed with terminal and mobile devices.

[Phases]

Phase 1, will deal with writing the entities and model.                           (iterative)
Phase 2, will focus on making the microservices via GRPC service.                 (iterative)
Phase 3, will heavily focus on frontend development.                              (iterative)


