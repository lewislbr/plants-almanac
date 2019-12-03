declare interface GlobalFetch {
  fetch(input: RequestInfo, init?: RequestInit): Promise<Response>;
}

declare module '*.woff2';

declare interface NodeModule {
  hot: {
    accept(dependencies?: string | string[], callback?: () => void): void;
  };
}
