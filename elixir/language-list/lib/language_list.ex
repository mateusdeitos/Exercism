defmodule LanguageList do
  @spec new() :: []
  def new() do
    []
  end

  @spec add(list(), any()) :: [...]
  def add(list, language), do: [language | list]

  @spec remove(list()) :: list()
  def remove([_|tail]), do: tail

  @spec first(list()) :: any()
  def first([first | _]), do: first

  @spec count(list()) :: non_neg_integer()
  def count(list), do: length(list)

  @spec functional_list?(any()) :: boolean()
  def functional_list?(list),
    do: "elixir" in Enum.map(list, fn language -> String.downcase(language) end)
end
