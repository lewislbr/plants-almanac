module.exports = {
  moduleNameMapper: {
    "\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$":
      "test-file-stub",
    "\\.css$": "{}",
  },
  preset: "ts-jest",
  testEnvironment: "node",
  testMatch: ["<rootDir>/src/**/?(*.)test.ts?(x)"],
}
