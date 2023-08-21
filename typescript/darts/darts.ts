export function score(x: number, y: number): number {
  const RADIUS_OUTER_CIRCLE = 10;
  const RADIUS_MIDDLE_CIRCLE = 5;
  const RADIUS_INNER_CIRCLE = 1;

  const distance = hipotenus(x, y);

  if (distance > RADIUS_OUTER_CIRCLE) {
    return 0;
  }

  if (distance > RADIUS_MIDDLE_CIRCLE) {
    return 1;
  }

  if (distance > RADIUS_INNER_CIRCLE) {
    return 5;
  }

  return 10;
}

const hipotenus = (x: number, y: number): number => Math.sqrt(x * x + y * y);
