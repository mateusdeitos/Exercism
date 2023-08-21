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
  items: BaseComparisonItem[];
  value: number;
  weight: number;
};

type BaseComparisonItem = Item & {
  index: number;
};

const getOptimalPackageOfItems = (
  baseItems: BaseComparisonItem[],
  maximumWeight: number,
  items: Item[],
  cache: { [key: string]: Package } = {}
): Package | undefined => {
  const weight = baseItems.reduce((acc, item) => acc + item.weight, 0);
  const value = baseItems.reduce((acc, item) => acc + item.value, 0);
  const indexesOfBaseItems = baseItems.map((item) => item.index);
  const pack: Package = {
    items: [...baseItems],
    value,
    weight,
  };

  let remainingWeight = maximumWeight - weight;

  if (remainingWeight === 0) {
    return pack;
  }

  if (remainingWeight < 0) {
    return;
  }

  const hash = baseItems.map((item) => item.index).join("-");
  if (cache[hash]) {
    return cache[hash];
  }

  const unfitItems: BaseComparisonItem[] = [];

  items.forEach((item, index) => {
    if (indexesOfBaseItems.includes(index)) {
      return;
    }

    const fit = item.weight <= remainingWeight;

    if (!fit) {
      unfitItems.push({ ...item, index });
      return;
    }

    remainingWeight -= item.weight;
    pack.items.push({ ...item, index });
    pack.value += item.value;
    pack.weight += item.weight;
  });

  const otherPackages = unfitItems.reduce<Package[]>((acc, item) => {
    let _baseItems = [...baseItems, item];
    if (item.weight === maximumWeight) {
      _baseItems = [item];
    }

    let otherCombination = getOptimalPackageOfItems(
      _baseItems,
      maximumWeight,
      items,
      cache
    );

    if (!otherCombination) {
      return acc;
    }

    return [...acc, otherCombination];
  }, []);

  const optimalPackage = [pack, ...otherPackages].reduce(
    (acc, comparisonPackage) => {
      if (comparisonPackage.value > acc.value) {
        return comparisonPackage;
      }

      if (
        comparisonPackage.value === acc.value &&
        comparisonPackage.weight < acc.weight
      ) {
        return comparisonPackage;
      }

      return acc;
    },
    pack
  );

  cache[hash] = optimalPackage;

  return optimalPackage;
};

export function maximumValue({ maximumWeight, ...params }: Params): number {
  const items = params.items.filter((i) => i.weight <= maximumWeight);

  const combination = items.reduce<Package>(
    (currentOptimalPackage, item, index) => {
      if (currentOptimalPackage.items.some((i) => i.index === index)) {
        return currentOptimalPackage;
      }

      const optimalPackage = getOptimalPackageOfItems(
        [{ ...item, index }],
        maximumWeight,
        items.map((i, index) => ({ ...i, index }))
      );

      if (!optimalPackage) {
        return currentOptimalPackage;
      }

      if (optimalPackage.value > currentOptimalPackage.value) {
        return optimalPackage;
      }

      if (
        optimalPackage?.value === currentOptimalPackage.value &&
        optimalPackage?.weight < currentOptimalPackage.weight
      ) {
        return optimalPackage;
      }

      return currentOptimalPackage;
    },
    { value: 0, weight: 0, items: [] }
  );

  return combination.value;
}
