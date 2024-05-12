from dataclasses import dataclass
from typing import Any


@dataclass
class GoType:
    """
    GoType refer to the basic data type of the Go language,
    and this encoding library is basing on this basic class.

    The current class carries a payload x, which is the go-lang
    base data type to be recorded.

    It should be noted that you are only allowed to modify the
    payload x which carries by this class.

    For the corresponding go-lang data type to x, you can only
    set its value while init this class.

    Raises:
        AttributeError: Try to modify a value not called "x"
    """

    x: Any

    def __init__(self, x: Any, data_type: type, name_of_data_type: str) -> None:
        """
        Take data x of data type "data_type" and data type
        name "name_of_data_type" as the payload of the current class.

        Args:
            x (Any): The payload to be placed on the current class
            data_type (type): The expected data type of x
            name_of_data_type (str): The name of the go-lang data type corresponding to data_type
        """
        self.__dict__["data_type"] = data_type
        self.__dict__["name_of_data_type"] = name_of_data_type
        self.DataCheck(x)
        self.x: Any = x

    def __setattr__(self, name: str, value: Any) -> None:
        """
        Sets the payload x carried by the current class.

        Args:
            name (str):
                This parameter can only be filled in x,
                because the name of the payload is x.
            value (Any): New value of payload

        Raises:
            AttributeError: Try to modify a value not called "x"
        """
        if name != "x":
            raise AttributeError(
                f"GoType@__setattr__: You can only modify the value of x"
            )
        self.DataCheck(value)
        self.__dict__["x"] = value

    def DataCheck(self, x: Any) -> None:
        """
        Check whether x meets the specified data type which
        recorded in "self.data_type".

        Args:
            x (Any): Data to check

        Raises:
            TypeError: x is not meets self.data_type
        """
        if type(x) != self.__dict__["data_type"]:
            raise TypeError(
                f"DataCheck: Cannot use {x} as {self.__dict__['name_of_data_type']} value in assignment"
            )


class GoInt(GoType):
    """
    GoInt is the base class for all integer data types in Go language.

    Args:
        GoType (class):
            It refers to the base class of GoInt, since GoInt itself is
            a basic data type of the Go language.

    Raises:
        AttributeError:
            The payload x that you want to place on the current class is
            not a int type.
        ValueError:
            The payload (integer) x that you want to place on the current
            class is out of range。
    """

    x: int = 0

    def __init__(
        self, x: int, name_of_data_type: str, int_range: tuple[int, int]
    ) -> None:
        """
        Take data x (integer) and data type name "name_of_data_type" as
        the payload of this GoInt class.

        Args:
            x (int): The payload to be placed on the current class
            name_of_data_type (str): The name of the go-lang data type corresponding to x
            int_range (tuple[int, int]): The range of the integer x
        """
        self.__dict__["int_range"] = int_range
        super().__init__(x, int, name_of_data_type)

    def DataCheck(self, x: int) -> None:
        """
        Check whether x is an integer and whether it meets the range limit.

        Args:
            x (int): Data to check

        Raises:
            TypeError: x is not an integer
            ValueError: x is out of the range
        """
        super().DataCheck(x)
        # check the data type
        int_range: tuple[int, int] = self.__dict__["int_range"]
        if not int_range[0] <= x <= int_range[1]:
            raise ValueError(
                f"DataCheck: Cannot use {x} as {self.__dict__['name_of_data_type']} value in assignment (overflows)"
            )
        # check the range of giving int value


class String(GoType):
    """
    String is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    String is the set of all strings of 8-bit bytes, conventionally but not
    necessarily representing UTF-8-encoded text. A string may be empty, but
    not nil. Values of string type are immutable.
    """

    x: str = ""

    def __init__(self, x: str) -> None:
        super().__init__(x, str, "string")


class Bool(GoType):
    """
    Bool is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    Bool is the set of boolean values, true and false.
    """

    x: bool = False

    def __init__(self, x: bool) -> None:
        super().__init__(x, bool, "bool")


class Uint8(GoInt):
    """
    Uint8 is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    Uint8 is the set of all unsigned 8-bit integers.
    Range: 0 through 255.
    """

    x: int = 0

    def __init__(self, x: int) -> None:
        super().__init__(x, "uint8", (0, 255))


class Int8(GoInt):
    """
    Int8 is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    Int8 is the set of all signed 8-bit integers.
    Range: -128 through 127.
    """

    x: int = 0

    def __init__(self, x: int) -> None:
        super().__init__(x, "int8", (-128, 127))


class Uint16(GoInt):
    """
    Uint16 is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    Uint16 is the set of all unsigned 16-bit integers.
    Range: 0 through 65535.
    """

    x: int = 0

    def __init__(self, x: int) -> None:
        super().__init__(x, "uint16", (0, 65535))


class Int16(GoInt):
    """
    Int16 is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    Int16 is the set of all signed 16-bit integers.
    Range: -32768 through 32767.
    """

    x: int = 0

    def __init__(self, x: int) -> None:
        super().__init__(x, "int16", (-32768, 32767))


class Uint32(GoInt):
    """
    Uint16 is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    Uint32 is the set of all unsigned 32-bit integers.
    Range: 0 through 4294967295.
    """

    x: int = 0

    def __init__(self, x: int) -> None:
        super().__init__(x, "uint32", (0, 4294967295))


class Int32(GoInt):
    """
    Int32 is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    Int32 is the set of all signed 32-bit integers.
    Range: -2147483648 through 2147483647.
    """

    x: int = 0

    def __init__(self, x: int) -> None:
        super().__init__(x, "int32", (-2147483648, 2147483647))


class Uint64(GoInt):
    """
    Uint64 is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    Uint64 is the set of all unsigned 64-bit integers.
    Range: 0 through 18446744073709551615.
    """

    x: int = 0

    def __init__(self, x: int) -> None:
        super().__init__(x, "uint64", (0, 18446744073709551615))


class Int64(GoInt):
    """
    Int64 is one of the basic data types of the Go language.
    The following is an excerpt from the go-lang docs.

    Int64 is the set of all signed 64-bit integers.
    Range: -9223372036854775808 through 9223372036854775807.
    """

    x: int = 0

    def __init__(self, x: int) -> None:
        super().__init__(x, "int64", (-9223372036854775808, 9223372036854775807))
