defmodule TakeANumber do
  def start(), do: spawn(fn -> listen(0) end)

  defp listen(state) do
    receive do
      {:report_state, sender_pid} ->
        send(sender_pid, state)
        listen(state)

      {:take_a_number, sender_pid} ->
        new_state = state + 1
        send(sender_pid, new_state)
        listen(new_state)

      :stop ->
        :ok

      _ ->
        listen(state)
    end
  end
end
