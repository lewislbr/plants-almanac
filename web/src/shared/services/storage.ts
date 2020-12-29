export function store(key: string, value: string): void {
  localStorage.setItem(key, value)
}

export function retrieve(key: string): string | null {
  return localStorage.getItem(key)
}
