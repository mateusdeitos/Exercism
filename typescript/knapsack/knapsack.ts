/* eslint-disable @typescript-eslint/no-shadow */
type Item = {
  weight: number;
  value: number;
};

type Params = {
  maximumWeight: number;
  items: Item[];
};

type Package = {
  itemIndexes: number[];
  value: number;
  weight: number;
};

export function maximumValue({ maximumWeight, ...params }: Params): number {
  const items = params.items.filter((i) => i.weight <= maximumWeight);

  if (!items.length) {
    return 0;
  }

  const packages = getPackagesBelowWeight(maximumWeight, items);

  const result = packages.reduce((result, pack) => {
    if (pack.value > result) {
      return pack.value;
    }

    return result;
  }, 0);

  return result;
}

function getPackagesBelowWeight(
  maximumWeight: number,
  items: Item[]
): Package[] {
  const packages: Package[] = items.map((item, index) => {
    return {
      value: item.value,
      weight: item.weight,
      itemIndexes: [index],
    };
  });

  for (let size = 1; size <= items.length; size++) {
    const additionalPackages: Package[] = [];
    for (const pack of packages) {
      if (pack.itemIndexes.length < size) {
        continue;
      }

      const startIndex = pack.itemIndexes.at(-1)! + 1;

      for (let i = startIndex; i < items.length; i++) {
        const item = items[i];
        if (pack.weight + item.weight > maximumWeight) {
          continue;
        }

        const newPack: Package = {
          value: pack.value + item.value,
          weight: pack.weight + item.weight,
          itemIndexes: [...pack.itemIndexes, i],
        };

        additionalPackages.push(newPack);
      }
    }

    packages.push(...additionalPackages);
  }

  return packages;
}
