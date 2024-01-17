defmodule BirdCount do
  def today([]), do: nil
  def today([today | _]), do: today

  def increment_day_count([]), do: [1]
  def increment_day_count([today | rest]), do: [today + 1 | rest]

  def has_day_without_birds?(list), do:
    list
      |> Enum.any?(fn(dailyCount) -> dailyCount == 0 end)

  def total(list) do
    list
      |> Enum.sum
  end

  def busy_days(list) do
    list
      |> Enum.count(fn(dailyCount) -> dailyCount >= 5 end)
  end
end
