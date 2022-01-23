# Spectator Frontend

This folder contains Spectator's frontend. To see the detail of technologies
being used, please refer to the [Tech section](#tech).

## Development

To develop this app, you will need several things installed on your system.
Prerequisites:
  - Docker
  - Docker Compose
  - Node >=16.13

You can run `npm run dev` to start the dev server and start developing.
Make sure to add tests and test it by running `npm run test`. For the list of
all available commands, refer to the [NPM Scripts
section](#available-npm-scripts).

If you're a bit lost with the folder structure, you can find the details on the
[Folder Structure section](#folder-structure)

You can also try to build the docker image by executing `docker build -t
spectator-frontend` from the root of this folder and run it by executing
`docker run -p 3000:3000 spectator-frontend`

## Tech

- `Typescript` - The language used, I love me some static typings
- `React + Vite` - React from Vite starter, because Vite provides better DX
- `Chakra UI` - Stylings
- `Redux Toolkit` - Manage global state
- `Redux Persist` - Persist state to localStorage
- `React Hook Form + Yup` - Handle form state and validation
- `i18next` - Internationalisation stuff
- `React Testing Library + Vitest` - Bog standard test library, but with bliding ej test runner

## Available NPM Scripts

- `start` - Start the project in production mode
- `dev` - Start the project in development mode
- `build` - Builds the project
- `build:prod` - Builds the project, but disables the type checker. Reason: see [this issue](https://github.com/chakra-ui/chakra-ui/issues/5082#issuecomment-979039787)
- `lint:fix` - Lint and fix the entire project
- `lint:check` - Only checks for linter errors
- `test` - Runs the test and exits
- `test:coverage` - Run test with coverage
- `test:open` - Open Vitest UI, sometime it works, sometime it doesn't
- `test:watch` - Just like dev, but watches for test changes
- `protoc` - Generate typescript typings from proto files

## Folder Structure

- `public/locales` - Translation / Question files
- `src/App.tsx` - Routes list
- `src/components` - Any React components
- `src/data` - Translation / Question files
- `src/events` - Listener for events
- `src/hoc` - High Order Component, mostly for wrapping private route
- `src/hooks` - Our own custom hooks
- `src/images` - Assets for SAM test (they're all SVGs)
- `src/pages` - Any component that serves as a page
- `src/schema` - Schema for form validation
- `src/store` - Redux store
- `src/stub` - Auto-generated typings from proto files
- `src/styles` - Any files related to styling
- `src/tests` - Test files
- `src/utils` - Utilities (ATM it's mostly filled with fake stuff)

