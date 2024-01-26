defmodule Username do
  def sanitize([]), do: []

  def sanitize([head | tail]) do
    sanitized =
      case head do
        ?ä -> ~c"ae"
        ?ö -> ~c"oe"
        ?ü -> ~c"ue"
        ?ß -> ~c"ss"
        char when char >= ?a and char <= ?z -> [char]
        ?_ -> ~c"_"
        _ -> []
      end

    sanitized ++ sanitize(tail)
  end
end
