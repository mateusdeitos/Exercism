defmodule LogLevel do
  def to_label(level, legacy?) do
    cond do
      level === 0 and !legacy? -> :trace
      level === 1 -> :debug
      level === 2 -> :info
      level === 3 -> :warning
      level === 4 -> :error
      level === 5 and !legacy? -> :fatal
      true -> :unknown
    end
  end

  def alert_recipient(level, legacy?) when legacy? do
    label = to_label(level, legacy?)
    cond do
      label in [:error, :fatal] -> :ops
      label === :unknown -> :dev1
      true -> false
    end
  end

  def alert_recipient(level, legacy?) do
    label = to_label(level, legacy?)
    cond do
      label in [:error, :fatal] -> :ops
      label === :unknown -> :dev2
      true -> false
    end
  end
end
