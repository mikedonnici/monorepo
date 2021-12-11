# monorepo structure

Sample monorepo with a Vue frontend, Python (FatsAPI) http service and a Go (gRPC) backend.

Objective is to create a codebase that is:

- easy to comprehend
- easy to work on
- easy to deploy

The application:

- Web front end (Nuxt / Vue) that loads result fetched from a http backend
- Python / FastAPI backend that provides access to the service, and API docs, and communicates with 
  the data layer via gRPC
- Go / gRPC data layer that listens for gRPC requests and responds accordingly 

Ideas:

- Root folders named for the primary code therein.
- Sub folders must comprise deployable projects, where relevant, or common libraries relevant to the language

```
/
├── go
│   ├── hellosrvc      <-- backend service that says hello in a random language
│       ├── cmd        <-- go packages that are executable
│       └── internal
            └── store
│ 
├── proto       <-- proto files accessible to all project that need them
│
├── py
│   ├── apps
│   └── pkg
│
├── testdata
├──  
└── vue          <-- complete frontend apps written in Vue
```
