import * as React from "react";

export function PlantCard({
  name,
}: {
  name: string | null | undefined;
}): JSX.Element {
  return (
    <div className="bg-white mb-3 p-4 rounded-lg shadow-sm">
      <h2 className="font-bold text-2xl">{name}</h2>
    </div>
  );
}
