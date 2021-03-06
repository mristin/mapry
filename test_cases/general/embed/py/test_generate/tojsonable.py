# File automatically generated by mapry. DO NOT EDIT OR APPEND!


"""serializes to JSONable objects."""


import collections
import typing

import some.graph


def serialize_empty(
        instance: some.graph.Empty,
        ordered: bool = False
) -> typing.MutableMapping[str, typing.Any]:
    """
    serializes an instance of Empty to a JSONable representation.

    :param instance: the instance of Empty to be serialized
    :param ordered:
        If set, represents the instance as a ``collections.OrderedDict``.
        Otherwise, it is represented as a ``dict``.
    :return: a JSONable
    """
    if ordered:
        target = (
            collections.OrderedDict()
        )  # type: typing.MutableMapping[str, typing.Any]
    else:
        target = dict()

    return target


def serialize_non_empty(
        instance: some.graph.NonEmpty,
        ordered: bool = False
) -> typing.MutableMapping[str, typing.Any]:
    """
    serializes an instance of NonEmpty to a JSONable representation.

    :param instance: the instance of NonEmpty to be serialized
    :param ordered:
        If set, represents the instance as a ``collections.OrderedDict``.
        Otherwise, it is represented as a ``dict``.
    :return: a JSONable
    """
    if ordered:
        target = (
            collections.OrderedDict()
        )  # type: typing.MutableMapping[str, typing.Any]
    else:
        target = dict()

    ##
    # Serialize empty
    ##

    target['empty'] = serialize_empty(instance.empty)

    return target


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
    # Serialize some_embed
    ##

    target['some_embed'] = serialize_non_empty(instance.some_embed)

    return target


# File automatically generated by mapry. DO NOT EDIT OR APPEND!
