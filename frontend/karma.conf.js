module.exports = function (config) {
  config.set({
    basePath: '',
    frameworks: ['jasmine', '@angular-devkit/build-angular'],
    plugins: [
      require('karma-jasmine'),
      require('karma-chrome-launcher'),
      require('karma-html-reporter'),
      require('@angular-devkit/build-angular/plugins/karma')
    ],
    client: {
      clearContext: false
    },
    reporters: ['progress', 'html'],
    htmlReporter: {
      outputFile: 'tests/units.html',
      pageTitle: 'Unit Tests',
      subPageTitle: 'Angular Component Tests',
      groupSuites: true,
      useCompactStyle: true,
      useLegacyStyle: true
    },
    customLaunchers: {
      ChromeHeadlessCI: {
        base: 'ChromeHeadless',
        flags: [
          '--no-sandbox',
          '--disable-gpu'
        ]
      }
    },
    browsers: ['ChromeHeadlessCI'],
    singleRun: true
  });
};
