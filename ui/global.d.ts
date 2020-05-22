declare module '*.woff2';

declare interface NodeModule {
  hot: {
    accept(dependencies?: string | string[], callback?: () => void): void;
  };
}
