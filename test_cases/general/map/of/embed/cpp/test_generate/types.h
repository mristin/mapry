#pragma once

// File automatically generated by mapry. DO NOT EDIT OR APPEND!

#include <map>

namespace some {
namespace graph {

struct SomeGraph;

struct SomeEmbed;

// defines an embeddable structure.
struct SomeEmbed {
  // defines a boolean property.
  bool some_property = false;
};

// defines some object graph.
struct SomeGraph {
  // tests a map of embeddable structures.
  std::map<std::string, SomeEmbed> map_of_embeds;
};

}  // namespace graph
}  // namespace some

// File automatically generated by mapry. DO NOT EDIT OR APPEND!
