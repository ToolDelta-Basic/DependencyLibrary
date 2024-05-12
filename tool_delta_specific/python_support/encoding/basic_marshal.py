from .const import *
from .define import String
from io import BytesIO
from struct import pack


class Writer(BytesIO):
    def String(self, x: String):
        """
        向写入者写入字符串 x 。

        Args:
            x (str): 要写入的字符串

        Raises:
            OSError: 欲写入的字符串长度超过最大上限
        """
        if len(x.x) > MAX_LEN_STRING:
            raise OSError(
                f"Writer(BytesIO): The length of the target string is out of the max limited {MAX_LEN_STRING}"
            )
        # check length
        self.write(pack(">H", len(x.x)))
        self.write(x.x.encode(encoding="utf-8", errors="replace"))
        # write the string with it's length
