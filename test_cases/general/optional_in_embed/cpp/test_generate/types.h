#pragma once

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

#include <optional.hpp>

#include <string>

namespace some {
namespace graph {

struct SomeGraph;

struct WithOptional;

// defines an embeddable with an optional property.
struct WithOptional {
  // defines an optional property.
  std::experimental::optional<std::string> some_text;
};

// defines some object graph.
struct SomeGraph {
  // tests an embeddable with an optional property.
  WithOptional some_property;
};

}  // namespace graph
}  // namespace some

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
