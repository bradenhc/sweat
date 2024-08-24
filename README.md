# Sweat - A Simple Workout Tracker

This application is meant to be run as a stand-alone microservice.

## Concepts

The top-level element is a **workout**. Workouts group together multiple
**exercises**. An exercise could be "Running", "Squats", "Yoga", or "Dumbell
Curls".

Sweat makes no distinction on the type of exercise. You can tag exercises to
make them easier to organize, but the tagging does not have any effect on how
sweat treats the exercise. Instead, each exercise has associated **metrics**
used for collecting data. For example, you could create an exercise called
"Cycling Intervals" and apply the metrics "sets" and "time" to create an
interval training exercise while you are using a bike.

Every time you record data for an exercise, the data is associated with a
**session**. Sessions keep track of related sets of data and are used to create
simple time-series analytics of your workout trends.

Sweat can generate very rudimentary **reports** that provide tables and graphs
of your workout habits over time.

## API

| Endpoint         | Method | Action                                |
| :--------------- | :----- | :------------------------------------ |
| `/workouts`      | GET    | Fetch a list of current workouts      |
| `/workouts`      | POST   | Add a new workout                     |
| `/workouts/:id`  | DELETE | Remove an existing workout            |
| `/exercises`     | GET    | Fetch a list of exercises             |
| `/exercises`     | POST   | Add a new exercise                    |
| `/exercises/:id` | GET    | Fetch a single exercise by ID         |
| `/metrics`       | GET    | Fetch a list of all available metrics |
| `/sessions`      | GET    | Fetch a list of all known sessions    |

## State

All state in sweat is kept in-memory and is managed by a single Goroutine. All
requests to read from/write to the state are handled by the state manager. This
design synchronizes access to the state from multiple Goroutines in a way that
clearly communicates responsibility.

## State Persistence

Rather than connect to an external database, sweat is designed to keep its own
file-based state log. Each sweat instance runs isolated from the rest, and each
instance owns a specific state log. From a user's perspective, if multiple users
want to use sweat, each one gets its own sweat instance. This prevents the need
to lock a sweat state log so it can be shared by multiple instances.

The StateManager is responsible for committing requests to modify state to the
log before the state is applied internally.

At a configurable limit, the current state of the sweat instance is written to a
snapshot file. This prevents the state log from growing too large and causing a
sweat instance to drop in performance while it manages that state.

By default the state log is a JSONL file and the snapshot is a JSON file. You
can also configure sweat to use a binary format for increased performance.

## State Structure

The following JSON structure describes the organization of state in Sweat. The
organization is somewhat normalized. A JSON snapshot file would have the same
structure.

```json
{
    "workouts": {
        "<wid>": {
            "id": "<wid>",
            "name": "",
            "exercises": ["<eid>"],
            "created": "<timestamp>",
            "updated": "<timestamp>",
            "archived": false
        }
    },
    "exercises": {
        "<eid>": {
            "id": "<eid>",
            "name": "",
            "metrics": ["sets", "reps", "time", "distance", "weight"],
            "created": "<timestamp>",
            "updated": "<timestamp>",
            "archived": false
        }
    },
    "sessions": [
        {
            "id": "<sid>",
            "duration": 0,
            "workout": "<wid>",
            "metrics": [
                {
                    "exercise": "<eid>",
                    "set": 0,
                    "reps": 0,
                    "weight": 0,
                    "time": 0,
                    "distance": 0
                }
            ],
            "notes": "",
            "created": "<timestamp>",
            "updated": "<timestamp>"
        }
    ]
}
```
