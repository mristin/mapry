# File automatically generated by mapry. DO NOT EDIT OR APPEND!


"""serializes to JSONable objects."""


import collections
import typing

import some.graph


def serialize_some_graph(
        instance: some.graph.SomeGraph,
        ordered: bool = False
) -> typing.MutableMapping[str, typing.Any]:
    """
    serializes an instance of SomeGraph to a JSONable.

    :param instance: the instance of SomeGraph to be serialized
    :param ordered:
        If set, represents the instance properties as a ``collections.OrderedDict``.
        Otherwise, they are represented as a ``dict``.
    :return: JSONable representation
    """
    if ordered:
        target = (
            collections.OrderedDict()
        )  # type: typing.MutableMapping[str, typing.Any]
    else:
        target = dict()

    ##
    # Serialize array_of_datetimes
    ##

    target_0 = [
        item_0.strftime('%Y-%m-%dT%H:%M:%SZ')
        for item_0 in instance.array_of_datetimes
    ]  # type: typing.List[str]
    target['array_of_datetimes'] = target_0

    return target


# File automatically generated by mapry. DO NOT EDIT OR APPEND!
