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
  // Parse array_of_maps
  ////

  if (not value.isMember("array_of_maps")) {
    errors->add(
      ref,
      "Property is missing: array_of_maps");
  } else {
    const Json::Value& value_0 = value["array_of_maps"];
    if (not value_0.isArray()) {
      constexpr auto expected_but_got(
        "Expected an array, but got: ");

      errors->add(
        std::string(ref)
          .append("/array_of_maps"),
        message(
          expected_but_got,
          strlen(expected_but_got),
          value_type_to_string(
            value_0.type())));
    } else {
      std::vector<std::map<std::string, bool>>& target_0 = target->array_of_maps;
      target_0.resize(value_0.size());
      size_t i_0 = 0;
      for (const Json::Value& item_0 : value_0) {
        if (not item_0.isObject()) {
          constexpr auto expected_but_got(
            "Expected an object, but got: ");

          errors->add(
            std::string(ref)
              .append("/array_of_maps")
              .append("/")
              .append(std::to_string(i_0)),
            message(
              expected_but_got,
              strlen(expected_but_got),
              value_type_to_string(
                item_0.type())));
        } else {
          std::map<std::string, bool>& target_1 = target_0.at(i_0);
          for (Json::ValueConstIterator it_1 = item_0.begin(); it_1 != item_0.end(); ++it_1) {
            const Json::Value& value_2 = *it_1;
            if (not value_2.isBool()) {
              constexpr auto expected_but_got(
                "Expected a bool, but got: ");

              errors->add(
                std::string(ref)
                  .append("/array_of_maps")
                  .append("/")
                  .append(std::to_string(i_0))
                  .append("/")
                  .append(it_1.name()),
                message(
                  expected_but_got,
                  strlen(expected_but_got),
                  value_type_to_string(
                    value_2.type())));
            } else {
              target_1[it_1.name()] = value_2.asBool();
            }

            if (errors->full()) {
              break;
            }
          }
        }
        ++i_0;

        if (errors->full()) {
          break;
        }
      }

    }
  }
  if (errors->full()) {
    return;
  }
}

Json::Value serialize_some_graph(
    const SomeGraph& some_graph) {
  Json::Value some_graph_as_value;

  Json::Value target_0(Json::arrayValue);
  const auto& vector_0 = some_graph.array_of_maps;
  for (int i_0 = 0;
      i_0 < vector_0.size();
      ++i_0) {
    Json::Value target_1(Json::objectValue);
    const auto& map_1 = vector_0[i_0];
    for (const auto& kv_1 : map_1) {
      target_1[kv_1.first] = kv_1.second;
    }
    target_0[i_0] = std::move(target_1);
  }
  some_graph_as_value["array_of_maps"] = std::move(target_0);

  return some_graph_as_value;
}

}  // namespace jsoncpp

}  // namespace graph
}  // namespace some

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
