# File automatically generated by mapry. DO NOT EDIT OR APPEND!


"""defines some object graph."""


import typing


class SomeGraph:
    """defines some object graph."""

    def __init__(
            self,
            some_int_gt_0_lt_100: int,
            some_int_ge_0_le_100: int,
            unconstrained_int: int) -> None:
        """
        initializes an instance of SomeGraph with the given values.

        :param some_int_gt_0_lt_100: defines some integer.
        :param some_int_ge_0_le_100: defines some integer.
        :param unconstrained_int: defines some unconstrained integer.

        """
        self.some_int_gt_0_lt_100 = some_int_gt_0_lt_100
        self.some_int_ge_0_le_100 = some_int_ge_0_le_100
        self.unconstrained_int = unconstrained_int


# File automatically generated by mapry. DO NOT EDIT OR APPEND!
