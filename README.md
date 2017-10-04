# Crypto Accounts Viewer

**Note: At this point, this project is merely an experiment.**

## Build

Install the tools required for building the project. The instructions
below assume you are running macOS:

```
brew install go
npm install -g @juxt/mach
```

Create the Postgres database:

```
createdb crypto-accounts-viewer 
```

## Usage

1. Start the backend: `mach start-backend`
2. Start the frontend: `mach start-frontend` (TODO)
3. Open the frontend: Point your browser to `http://localhost:8081`

## License

Copyright &copy; 2017 Jannis Pohlmann

Licensed under the MIT License.
