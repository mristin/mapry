# File automatically generated by mapry. DO NOT EDIT OR APPEND!


"""parses JSONable objects."""


import collections
import typing

import some.graph
import some.graph.parse


def some_graph_from(
        value: typing.Any,
        ref: str,
        errors: some.graph.parse.Errors
) -> typing.Optional[some.graph.SomeGraph]:
    """
    parses SomeGraph from a JSONable value.

    :param value: JSONable value
    :param ref: reference to the value (e.g., a reference path)
    :param errors: errors encountered during parsing
    :return: parsed SomeGraph, or None if ``errors``
    """
    if errors.full():
        return None

    if not isinstance(value, dict):
        errors.add(
            ref,
            "Expected a dictionary, but got: {}".format(type(value)))
        return None

    graph = some.graph.parse.placeholder_some_graph()

    ##
    # Parse map_of_floats
    ##

    value_0 = value.get(
        'map_of_floats',
        None)

    if value_0 is None:
        errors.add(
            ref,
            'Property is missing: map_of_floats')
    else:
        if not isinstance(value_0, dict):
            errors.add(
                '/'.join((
                    ref, 'map_of_floats')),
                "Expected a dict, but got: {}".format(
                    type(value_0)))
        else:
            if isinstance(value_0, collections.OrderedDict):
                target_1 = (
                    collections.OrderedDict()
                )  # type: typing.MutableMapping[str, float]
            else:
                target_1 = (
                    dict()
                )

            for key_1, value_1 in value_0.items():
                if not isinstance(key_1, str):
                    errors.add(
                        '/'.join((
                            ref, 'map_of_floats')),
                        "Expected the key to be a str, but got: {}".format(
                            type(key_1)))

                    if errors.full():
                        break
                    else:
                        continue

                target_item_1 = (
                    None
                )  # type: typing.Optional[float]
                if not isinstance(value_1, (int, float)):
                    errors.add(
                        '/'.join((
                            ref, 'map_of_floats', repr(key_1))),
                        'Expected a number, but got: {}'.format(
                            type(value_1)))
                else:
                    target_item_1 = float(value_1)

                if target_item_1 is not None:
                    target_1[key_1] = target_item_1

                if errors.full():
                    break

            if target_1 is not None:
                graph.map_of_floats = target_1

    if errors.full():
        return None

    if not errors.empty():
        return None

    return graph
