# NSE Thermometer Frontend (Angular)

This project is a web application built with Angular. It allows you to run the app in the browser for development or production. The frontend consumes a backend API to display the thermometer screens, progress bars, and goals.

## Technologies

- **Angular** – Frontend framework for building reactive and component-based UIs.
- **TailwindCSS** – CSS framework for rapid styling and responsive design.
- **RxJS** – For reactive programming and handling intervals.
- **Karma & Jasmine** – For unit testing components.

This combination allows you to develop the app in the browser with fast iteration and hot reload using the Angular development server.

# Install dependencies and run locally

```bash
npm install
ng serve
```

> **Please note:** The backend URLs are currently hardcoded in `app.component.ts`. Consider refactoring them to use environment variables or Angular environment files for more flexibility.

## Available Scripts
Here’s an overview of all the run commands in this project:

| Command               | Description                                                |
| --------------------- | ---------------------------------------------------------- |
| `npm run start`       | Starts the Angular app in the browser with live reload     |
| `npm run build`       | Builds the Angular app for production                      |
| `npm run watch`       | Builds the app in development mode and watches for changes |
| `npm run test`        | Runs unit tests with Karma & Jasmine                       |
| `npm run test:report` | Runs tests headless in Chrome and generates a report       |

## Building for Production
To build the Angular frontend for production:

```bash
ng build --prod
```

The production-ready files will be output to the `dist/` folder. You can then serve them with any web server.

## Testing
The `npm run test:report` method will test the application and create a test report in the `karma_html` folder.
