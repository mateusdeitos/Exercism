defmodule NameBadge do
  def print(id, name, department)
      when id !== nil and department !== nil do
    "[#{id}] - #{name} - #{String.upcase(department)}"
  end

  # new employee
  def print(nil, name, department) when department !== nil do
    "#{name} - #{String.upcase(department)}"
  end

  def print(id, name, nil), do: print_owner(id, name)
  def print(nil, name, nil), do: print_owner(nil, name)

  defp print_owner(id, name) do
    if id === nil do
      "#{name} - OWNER"
    else
      "[#{id}] - #{name} - OWNER"
    end
  end
end
