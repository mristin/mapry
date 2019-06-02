# File automatically generated by mapry. DO NOT EDIT OR APPEND!


"""parses JSONable objects."""


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
    # Parse some_int_gt_0_lt_100
    ##

    value_0 = value.get(
        'some_int_gt_0_lt_100',
        None)

    if value_0 is None:
        errors.add(
            ref,
            'Property is missing: some_int_gt_0_lt_100')
    else:
        if not isinstance(value_0, int):
            errors.add(
                '/'.join((
                    ref, 'some_int_gt_0_lt_100')),
                "Expected an integer, but got: {}".format(
                    type(value_0)))
        else:
            if not (value_0 > 0):
                errors.add(
                    '/'.join((
                        ref, 'some_int_gt_0_lt_100')),
                    'Expected > 0, but got: {}'.format(
                        value_0))
            if not (value_0 < 100):
                errors.add(
                    '/'.join((
                        ref, 'some_int_gt_0_lt_100')),
                    'Expected < 100, but got: {}'.format(
                        value_0))
            else:
                graph.some_int_gt_0_lt_100 = value_0

    if errors.full():
        return None

    ##
    # Parse some_int_ge_0_le_100
    ##

    value_2 = value.get(
        'some_int_ge_0_le_100',
        None)

    if value_2 is None:
        errors.add(
            ref,
            'Property is missing: some_int_ge_0_le_100')
    else:
        if not isinstance(value_2, int):
            errors.add(
                '/'.join((
                    ref, 'some_int_ge_0_le_100')),
                "Expected an integer, but got: {}".format(
                    type(value_2)))
        else:
            if not (value_2 >= 0):
                errors.add(
                    '/'.join((
                        ref, 'some_int_ge_0_le_100')),
                    'Expected >= 0, but got: {}'.format(
                        value_2))
            if not (value_2 <= 100):
                errors.add(
                    '/'.join((
                        ref, 'some_int_ge_0_le_100')),
                    'Expected <= 100, but got: {}'.format(
                        value_2))
            else:
                graph.some_int_ge_0_le_100 = value_2

    if errors.full():
        return None

    ##
    # Parse unconstrained_int
    ##

    value_4 = value.get(
        'unconstrained_int',
        None)

    if value_4 is None:
        errors.add(
            ref,
            'Property is missing: unconstrained_int')
    else:
        if not isinstance(value_4, int):
            errors.add(
                '/'.join((
                    ref, 'unconstrained_int')),
                "Expected an integer, but got: {}".format(
                    type(value_4)))
        else:
            graph.unconstrained_int = value_4

    if errors.full():
        return None

    if not errors.empty():
        return None

    return graph