# Logs on Redis

Simple test on writing logs to redis.

The core concept is to use the data type "sorted set" to store logs and have efficient retrieval. Redis can also provide an easy publish/subscribe mechanism (although not implemented in this example).

Pros:

* Fast to store.
* Fast to retrieve.
* No dependencies with external services.
* More flexible than external services if you're willing to spend time writing query code.

Cons:

* Logs must fit in memory.
* Requires proper redis configuration (max memory) for safety.
* Requires periodic cleanup of old entries (easy with sorted set).
* Requires custom code to query and display logs (not a given like in external services).


## Setup

Run with:

    make run

You should see a stream of the 10 latest log entries updating on stdout.
