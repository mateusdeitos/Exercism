defmodule PaintByNumber do
  alias ElixirSense.Core.Bitstring

  def palette_bit_size(color_count, pow \\ 1) do
    case Integer.pow(2, pow) do
      result when result < color_count ->
        palette_bit_size(color_count, pow + 1)

      _ ->
        pow
    end
  end

  def empty_picture(), do: <<>>

  def test_picture(), do: <<0::2, 1::2, 2::2, 3::2>>

  def prepend_pixel(picture, color_count, pixel_color_index) do
    palette_bit_size(color_count)
    |> (fn bit_size -> <<pixel_color_index::size(bit_size)>> end).()
    |> (fn pixel -> <<pixel::bitstring, picture::bitstring>> end).()
  end

  def get_first_pixel(picture, color_count) when bit_size(picture) > 0 do
    palette_bit_size(color_count)
    |> (fn bit_size ->
          <<pixel::size(bit_size), _::bitstring>> = picture
          pixel
        end).()
  end

  def get_first_pixel(<<>>, _), do: nil

  def drop_first_pixel(picture, color_count)
      when bit_size(picture) === 0,
      do: <<>>

  def drop_first_pixel(picture, color_count) do
    palette_bit_size(color_count)
    |> (fn bit_size ->
          <<_::size(bit_size), rest::bitstring>> = picture
          rest
        end).()
  end

  def concat_pictures(picture1, picture2),
    do: <<picture1::bitstring, picture2::bitstring>>
end
