# Spectator Frontend

# Tech

- `Typescript` - The language used, I love me some static typings
- `React + Vite` - React from Vite starter, because Vite provides better DX
- `Chakra UI` - Stylings
- `Redux Toolkit` - Manage global state
- `Redux Persist` - Persist state to localStorage
- `React Hook Form + Yup` - Handle form state and validation
- `i18next` - Internationalisation stuff
- `React Testing Library + Vitest` - Bog standard test library, but with bliding ej test runner

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
