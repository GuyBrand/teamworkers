# Teamworkers
A demo that displays how to convert microservices into teamworkers

The code is complete and working
The readme is a WIP, aiming to be completed throughout March 2019 

What's teamworkers?
------------

Team workers are a software development paradigm in which every worker may play a variety of roles dybamically.
In other words a worker is a combination of microservices linked into a singular service that allows 100% microservice behavior OR monolith behavior according to parameters.

Who needs it?
------------

Two very opposite types of developers:
1. Developers who do not have an experience with microservices and discouraged by the overhead needed in terms of orchestration or latency
2. Developers who use microservices, deal with the overhead and would like to find better, easier ways to cope. 


Content
-----

This repo is that includes:
1. A sample of super basic microservices (basket and product) using a share library to initialize their CRUD like service case (micro folder).
2. The same two services after minor modifications, that will allow there to be bundles into a teamworker (workers folder).
3. a bundle that runs them as one (workers/glued).
4. some curl files to test a real external test of the functionality of the services and the worker.

A more detailed breakdown and walkthrough will soon be availalle 

Requirements
------------

GoLang 1.9
curl


Install
-------

```
go get github.com/guybrand/teamworkers
```

Contribute
----------

[![Tags](https://1)](http://2)

.

