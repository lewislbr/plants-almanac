export function transformDate(date: Date): string {
  const d = new Date(date)

  return d.toLocaleDateString("en-EN", {
    day: "numeric",
    month: "long",
    year: "numeric",
  })
}
