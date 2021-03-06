// File automatically generated by mapry. DO NOT EDIT OR APPEND!

#include "jsoncpp.h"
#include "parse.h"
#include "types.h"

#include <cstring>
#include <memory>
#include <sstream>
#include <stdexcept>
#include <string>
#include <utility>

namespace some {
namespace graph {

namespace jsoncpp {

/**
 * generates an error message.
 *
 * @param cc char array as the description part of the message
 * @param cc_size size of the char array
 * @param s string as the detail part of the message
 * @return concatenated string
 */
std::string message(const char* cc, size_t cc_size, std::string s) {
  std::string result;
  result.reserve(cc_size + s.size());
  result.append(cc, cc_size);
  result.append(s);
  return result;
}

/**
 * converts a JSON value type to a human-readable string representation.
 *
 * @param value_type to be converted
 * @return string representation of the JSON value type
 */
std::string value_type_to_string(Json::ValueType value_type) {
  switch (value_type) {
    case Json::ValueType::nullValue: return "null";
    case Json::ValueType::intValue: return "int";
    case Json::ValueType::uintValue: return "uint";
    case Json::ValueType::realValue: return "real";
    case Json::ValueType::stringValue: return "string";
    case Json::ValueType::booleanValue: return "bool";
    case Json::ValueType::arrayValue: return "array";
    case Json::ValueType::objectValue: return "object";
    default:
      std::stringstream ss;
      ss << "Unhandled value type in value_to_string: "
        << value_type;
      throw std::domain_error(ss.str());
  }
}

void some_graph_from(
    const Json::Value& value,
    std::string ref,
    SomeGraph* target,
    parse::Errors* errors) {
  if (errors == nullptr) {
    throw std::invalid_argument("Unexpected null errors");
  }

  if (not errors->empty()) {
    throw std::invalid_argument("Unexpected non-empty errors");
  }

  if (not value.isObject()) {
    constexpr auto expected_but_got(
      "Expected an object, but got: ");

    errors->add(
      ref,
      message(
        expected_but_got,
        strlen(expected_but_got),
        value_type_to_string(
          value.type())));
    return;
  }

  ////
  // Parse some_float_gt_0_lt_100
  ////

  if (not value.isMember("some_float_gt_0_lt_100")) {
    errors->add(
      ref,
      "Property is missing: some_float_gt_0_lt_100");
  } else {
    const Json::Value& value_0 = value["some_float_gt_0_lt_100"];
    if (not value_0.isDouble()) {
      constexpr auto expected_but_got(
        "Expected a double, but got: ");

      errors->add(
        std::string(ref)
          .append("/some_float_gt_0_lt_100"),
        message(
          expected_but_got,
          strlen(expected_but_got),
          value_type_to_string(
            value_0.type())));
    } else {
      const auto cast_0 = value_0.asDouble();
      bool ok_0 = true;

      if (not (cast_0 > 0)) {
        constexpr auto expected_but_got(
          "Expected "
          "> 0.000000"
          ", but got: ");

        errors->add(
          std::string(ref)
            .append("/some_float_gt_0_lt_100"),
          message(
            expected_but_got,
            strlen(expected_but_got),
            std::to_string(cast_0)));
        ok_0 = false;
      }

      if (not (cast_0 < 100)) {
        constexpr auto expected_but_got(
          "Expected "
          "< 100.000000"
          ", but got: ");

        errors->add(
          std::string(ref)
            .append("/some_float_gt_0_lt_100"),
          message(
            expected_but_got,
            strlen(expected_but_got),
            std::to_string(cast_0)));
        ok_0 = false;
      }

      if (ok_0) {
        target->some_float_gt_0_lt_100 = cast_0;
      }
    }
  }
  if (errors->full()) {
    return;
  }

  ////
  // Parse some_float_ge_0_le_100
  ////

  if (not value.isMember("some_float_ge_0_le_100")) {
    errors->add(
      ref,
      "Property is missing: some_float_ge_0_le_100");
  } else {
    const Json::Value& value_1 = value["some_float_ge_0_le_100"];
    if (not value_1.isDouble()) {
      constexpr auto expected_but_got(
        "Expected a double, but got: ");

      errors->add(
        std::string(ref)
          .append("/some_float_ge_0_le_100"),
        message(
          expected_but_got,
          strlen(expected_but_got),
          value_type_to_string(
            value_1.type())));
    } else {
      const auto cast_1 = value_1.asDouble();
      bool ok_1 = true;

      if (not (cast_1 >= 0)) {
        constexpr auto expected_but_got(
          "Expected "
          ">= 0.000000"
          ", but got: ");

        errors->add(
          std::string(ref)
            .append("/some_float_ge_0_le_100"),
          message(
            expected_but_got,
            strlen(expected_but_got),
            std::to_string(cast_1)));
        ok_1 = false;
      }

      if (not (cast_1 <= 100)) {
        constexpr auto expected_but_got(
          "Expected "
          "<= 100.000000"
          ", but got: ");

        errors->add(
          std::string(ref)
            .append("/some_float_ge_0_le_100"),
          message(
            expected_but_got,
            strlen(expected_but_got),
            std::to_string(cast_1)));
        ok_1 = false;
      }

      if (ok_1) {
        target->some_float_ge_0_le_100 = cast_1;
      }
    }
  }
  if (errors->full()) {
    return;
  }

  ////
  // Parse unconstrained_float
  ////

  if (not value.isMember("unconstrained_float")) {
    errors->add(
      ref,
      "Property is missing: unconstrained_float");
  } else {
    const Json::Value& value_2 = value["unconstrained_float"];
    if (not value_2.isDouble()) {
      constexpr auto expected_but_got(
        "Expected a double, but got: ");

      errors->add(
        std::string(ref)
          .append("/unconstrained_float"),
        message(
          expected_but_got,
          strlen(expected_but_got),
          value_type_to_string(
            value_2.type())));
    } else {
      target->unconstrained_float = value_2.asDouble();
    }
  }
  if (errors->full()) {
    return;
  }
}

Json::Value serialize_some_graph(
    const SomeGraph& some_graph) {
  Json::Value some_graph_as_value;

  some_graph_as_value["some_float_gt_0_lt_100"] = some_graph.some_float_gt_0_lt_100;

  some_graph_as_value["some_float_ge_0_le_100"] = some_graph.some_float_ge_0_le_100;

  some_graph_as_value["unconstrained_float"] = some_graph.unconstrained_float;

  return some_graph_as_value;
}

}  // namespace jsoncpp

}  // namespace graph
}  // namespace some

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
